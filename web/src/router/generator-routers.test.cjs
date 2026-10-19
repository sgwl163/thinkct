const assert = require('node:assert/strict');
const fs = require('node:fs/promises');
const os = require('node:os');
const path = require('node:path');
const { pathToFileURL } = require('node:url');
const vm = require('node:vm');
const ts = require('typescript');
const test = require('node:test');

// 使用真实 Vite 模块图与文件事件，不触碰项目视图目录或真实生成接口。
test('生成多个页面只热更新组件映射，现有路由引用也能找到新增页面', async () => {
  const { createServer } = await import('vite');
  const root = await fs.mkdtemp(path.join(os.tmpdir(), 'hotgo-router-hmr-'));
  let server;
  try {
    const write = async (name, content) => {
      const file = path.join(root, name);
      await fs.mkdir(path.dirname(file), { recursive: true });
      await fs.writeFile(file, content);
      return file;
    };
    await write(
      'src/router/generator-routers.ts',
      await fs.readFile(path.join(__dirname, 'generator-routers.ts'), 'utf8')
    );
    await write(
      'src/main.ts',
      "import { asyncImportRoute } from './router/generator-routers'; asyncImportRoute([]);"
    );
    await write('src/api/system/menu.ts', 'export const adminMenus = async () => ({ list: [] });');
    await write('src/router/router-icons.ts', 'export const constantRouterIcon = {};');
    await write(
      'src/router/constant.ts',
      'export const Layout = () => {}; export const ParentLayout = () => {};'
    );
    await write('src/views/iframe/index.vue', '<template><div /></template>');
    server = await createServer({
      root,
      configFile: false,
      logLevel: 'silent',
      resolve: { alias: { '@': path.join(root, 'src') } },
      server: { middlewareMode: true, watch: null },
      optimizeDeps: { noDiscovery: true, include: [] },
    });
    const messages = [];
    server.ws.send = (message) => messages.push(message);
    await server.transformRequest('/src/main.ts');
    const hotData = {};
    const hot = { data: hotData, accept() {} };
    function execute(code) {
      const context = {
        exports: {},
        console,
        meta: { hot, url: pathToFileURL(path.join(root, 'src/router/generator-routers.ts')).href },
        require: () => ({
          createHotContext: () => hot,
          Layout: () => {},
          ParentLayout: () => {},
          constantRouterIcon: {},
        }),
      };
      const compiled = ts.transpileModule(code.replaceAll('import.meta', 'meta'), {
        compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
      }).outputText;
      vm.runInNewContext(compiled, context);
      return context.exports;
    }
    async function load() {
      return execute((await server.transformRequest('/src/router/generator-routers.ts')).code);
    }
    const original = await load();
    // 原有调用方继续持有旧模块导出的函数，HMR 后也必须读取最新映射。
    original.asyncImportRoute([]);
    async function fileEvent(event, file) {
      messages.length = 0;
      server.watcher.emit(event, file);
      const deadline = Date.now() + 3000;
      while (messages.length === 0 && Date.now() < deadline) {
        await new Promise((resolve) => setTimeout(resolve, 10));
      }
      assert.ok(messages.length > 0, 'Vite 未处理文件事件');
      assert.equal(
        messages.some((message) => message.type === 'full-reload'),
        false,
        `新增/删除视图导致整页刷新：${JSON.stringify(messages)}`
      );
      assert.ok(messages.some((message) => message.type === 'update'));
      await load();
    }
    for (const name of ['index', 'edit', 'view']) {
      const file = await write(
        `src/views/generated/${name}.vue`,
        '<template><div>generated</div></template>'
      );
      await fileEvent('add', file);
      const routes = [{ component: `/generated/${name}` }];
      original.asyncImportRoute(routes);
      assert.equal(typeof routes[0].component, 'function', `旧路由函数找不到新页面 ${name}`);
    }
    const removed = path.join(root, 'src/views/generated/view.vue');
    await fs.unlink(removed);
    await fileEvent('unlink', removed);
    const routes = [{ component: '/generated/view' }];
    original.asyncImportRoute(routes);
    assert.equal(routes[0].component, undefined, '被删除页面仍残留在组件映射');
  } finally {
    await server?.close();
    await fs.rm(root, { recursive: true, force: true });
  }
});

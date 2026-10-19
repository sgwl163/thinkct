const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const test = require('node:test');
const vm = require('node:vm');
const ts = require('typescript');

// 执行页面实际的 script setup，隔离网络、弹窗和定时器，不连接真实后端。
function loadPage({ build, view, selects = async () => ({ status: [] }) }) {
  const source = fs.readFileSync(path.join(__dirname, 'deploy.vue'), 'utf8');
  const script = source.match(/<script lang="ts" setup>([\s\S]*?)<\/script>/)[1];
  const calls = { success: [], warning: [], error: [], notifications: [] };
  let confirm;
  let unmount;
  let mount;
  let now = 0;
  const modules = {
    vue: {
      ref: (value) => ({ value }),
      onMounted: (callback) => (mount = callback),
      onBeforeUnmount: (callback) => (unmount = callback),
      watch() {},
    },
    'vue-router': {
      useRouter: () => ({ currentRoute: { value: { params: { id: 52 } } } }),
    },
    'naive-ui': {
      useDialog: () => ({ info: (options) => (confirm = options.onPositiveClick) }),
      useMessage: () => ({
        success: (text) => calls.success.push(text),
        warning: (text) => calls.warning.push(text),
        error: (text) => calls.error.push(text),
      }),
      useNotification: () => ({ success: (options) => calls.notifications.push(options) }),
    },
    '@/api/develop/code': { Build: build, View: view, Selects: selects },
    '@/enums/httpEnum': { ResultEnum: { SUCCESS: 0, ERROR: -1, TIMEOUT: 61 } },
    '@/views/develop/code/components/model': {
      newState: () => ({ id: 52, status: 1, options: {} }),
      selectListObj: { status: [{ value: 1, label: '后台成功标签' }] },
      formatColumns: (columns) => columns,
    },
    '@/utils/is': { isJsonString: () => false },
  };
  const context = {
    exports: {},
    require(name) {
      if (name.endsWith('.vue')) return {};
      assert.ok(modules[name], `unexpected module: ${name}`);
      return modules[name];
    },
    defineProps: () => ({}),
    withDefaults: (props, defaults) => ({ ...defaults, ...props }),
    window: {
      setTimeout: (callback, delay) => {
        now += delay;
        callback();
      },
    },
    Date: { now: () => now },
    console,
  };
  const compiled = ts.transpileModule(script, {
    compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
  }).outputText;
  vm.runInNewContext(
    compiled +
      '\nexports.page = { submitBuild, waitForBuildResult, buildMessage, genInfo, formBtnLoading, show };',
    context
  );
  return {
    calls,
    page: context.exports.page,
    mount: () => mount(),
    unmount: () => unmount(),
    confirm() {
      context.exports.page.submitBuild();
      return confirm();
    },
    async submit() {
      context.exports.page.submitBuild();
      confirm();
      // 确认回调同步返回后，等待独立执行的生成和状态核对任务结束。
      await new Promise((resolve) => setImmediate(resolve));
    },
  };
}

test('生成接口和详情核对选项传递给项目 HTTP 封装', () => {
  const requests = [];
  const source = fs.readFileSync(path.join(__dirname, '../../../api/develop/code.ts'), 'utf8');
  const compiled = ts.transpileModule(source, {
    compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
  }).outputText;
  const context = {
    exports: {},
    require: () => ({ http: { request: (...args) => requests.push(args) } }),
  };
  vm.runInNewContext(compiled, context);
  context.exports.Build({ id: 54 });
  context.exports.View({ id: 54 }, { isShowErrorMessage: false });
  context.exports.View({ id: 54 });
  context.exports.Selects({}, { isShowErrorMessage: false });
  context.exports.TableSelect({}, { isShowErrorMessage: false });
  context.exports.ColumnSelect({}, { isShowErrorMessage: false });
  assert.equal(requests[0][0].timeout, 10 * 60 * 1000);
  assert.equal(requests[0][1].isTransformResponse, false);
  assert.equal(requests[0][1].isShowErrorMessage, false);
  assert.equal(requests[1][1].isShowErrorMessage, false);
  assert.equal(requests[2][1], undefined, '普通详情请求被强制改为静默');
  for (const request of requests.slice(3)) {
    assert.equal(request[1].isShowErrorMessage, false, '选项加载未传递静默错误选项');
  }
});

test('页面重载遇到后端重启时自动重试配置详情，不留下空 ID 和永久遮罩', async () => {
  let reads = 0;
  const { mount, page } = loadPage({
    build: async () => ({}),
    view: async () => {
      if (++reads < 3) throw new Error('后端正在重启');
      return completed();
    },
  });
  await mount();
  assert.equal(reads, 3);
  assert.equal(page.genInfo.value.id, 52);
  assert.equal(page.show.value, false);
});

test('配置初始化持续失败时释放遮罩并显示重新加载提示', async () => {
  const { mount, page } = loadPage({
    build: async () => ({}),
    view: async () => {
      throw new Error('服务不可用');
    },
  });
  await mount();
  assert.equal(page.show.value, false);
  assert.ok(page.buildMessage.value.includes('加载失败'));
});

const completed = () => ({
  id: 52,
  status: 1,
  building: false,
  buildSucceeded: true,
  buildFiles: 11,
  options: {},
  masterColumns: [],
});

test('确认提交立即关闭确认框，不等待生成请求完成', async () => {
  let finish;
  const pending = new Promise((resolve) => (finish = resolve));
  const { confirm } = loadPage({ build: () => pending, view: async () => completed() });
  const result = confirm();
  assert.equal(
    result,
    undefined,
    '确认回调返回了生成 Promise，Naive UI 会一直等到请求结束才关闭弹窗'
  );
  finish({ code: 0, message: '操作成功' });
  await new Promise((resolve) => setImmediate(resolve));
});

test('明确的生成失败不能被旧的成功记录覆盖，错误详情保留', async () => {
  const detail = '生成未完成，可能已写入部分文件：服务接口文件被占用';
  const { submit, calls, page } = loadPage({
    build: async () => ({ code: -1, message: detail }),
    view: async () => completed(),
  });
  await submit();
  assert.equal(calls.notifications.length, 0, '失败响应被显示为生成成功');
  assert.equal(calls.success.length, 0);
  assert.ok(page.buildMessage.value.includes(detail), '实际生成错误被刷新覆盖');
});

test('请求中断后只能确认后台记录，不能把旧成功当成本次成功', async () => {
  const { submit, calls, page } = loadPage({
    build: async () => {
      throw new Error('连接已中断');
    },
    view: async () => completed(),
  });
  await submit();
  assert.equal(calls.notifications.length, 0);
  assert.equal(calls.success.length, 0, '旧成功被误报为本次生成成功');
  assert.ok(page.buildMessage.value.includes('核对'));
  assert.equal(page.genInfo.value.buildFiles, 11, '核对结果没有同步回页面');
});

test('观察到正在生成后再完成，页面同步最终结果', async () => {
  let reads = 0;
  const { calls, page } = loadPage({
    build: async () => ({ code: 0, message: '操作成功' }),
    view: async () =>
      ++reads === 1 ? { ...completed(), building: true, buildSucceeded: false } : completed(),
  });
  await page.waitForBuildResult();
  assert.equal(reads, 2);
  assert.equal(calls.success.length, 1);
  assert.equal(page.genInfo.value.buildFiles, 11);
});

test('生成已成功但热重启使详情查询短暂失败时，自动重试并清除旧未完成提示', async () => {
  let reads = 0;
  const { submit, calls, page } = loadPage({
    build: async () => ({ code: 0, message: '操作成功' }),
    view: async (_params, options) => {
      assert.equal(options.isShowErrorMessage, false, '状态核对仍会弹出全局网络错误');
      if (++reads < 3) throw new Error('后端正在重启');
      return completed();
    },
  });
  page.genInfo.value.buildMessage = '上次生成未完整完成，可能已有部分文件';
  await submit();
  assert.equal(reads, 3, '生成成功后详情查询失败便停止，页面仍保留旧结果');
  assert.equal(calls.notifications.length, 1, '同一任务重复提示成功');
  assert.equal(page.genInfo.value.buildFiles, 11);
  assert.equal(page.buildMessage.value, '');
});

test('生成期间禁止重复提交，确认框关闭后按钮持续加载到结果核对完成', async () => {
  let finish;
  let finishView;
  let builds = 0;
  const pending = new Promise((resolve) => (finish = resolve));
  const pendingView = new Promise((resolve) => (finishView = resolve));
  const { confirm, page } = loadPage({
    build: () => {
      builds++;
      return pending;
    },
    view: () => pendingView,
  });
  confirm();
  assert.equal(page.formBtnLoading.value, true, '确认框关闭后生成按钮未保持加载状态');
  confirm();
  assert.equal(builds, 1, '确认框重复点击触发了重复生成请求');
  finish({ code: 0, message: '操作成功' });
  await new Promise((resolve) => setImmediate(resolve));
  assert.equal(page.formBtnLoading.value, true, '详情尚未核对就允许重复提交');
  confirm();
  assert.equal(builds, 1, '详情核对期间仍能触发重复生成');
  finishView(completed());
  await new Promise((resolve) => setImmediate(resolve));
  assert.equal(page.formBtnLoading.value, false, '结果核对结束后按钮仍无法操作');
});

test('已确认生成成功但服务持续不可用时，不把详情同步超时显示成生成失败', async () => {
  const { submit, calls, page } = loadPage({
    build: async () => ({ code: 0, message: '操作成功' }),
    view: async () => {
      throw new Error('服务不可用');
    },
  });
  await submit();
  assert.equal(calls.notifications.length, 1);
  assert.equal(calls.error.length, 0);
  assert.ok(page.buildMessage.value.startsWith('代码已生成'));
  assert.ok(page.buildMessage.value.includes('无需重复提交'));
  assert.equal(page.formBtnLoading.value, false);
});

test('离开页面后终止结果核对，避免继续发出详情请求', async () => {
  const { page, unmount } = loadPage({
    build: async () => ({ code: 0, message: '操作成功' }),
    view: async () => assert.fail('页面卸载后仍请求详情'),
  });
  unmount();
  await page.waitForBuildResult();
  assert.equal(page.formBtnLoading.value, false);
});

test('生成请求确认成功后显示通知并刷新文件信息', async () => {
  const { submit, calls, page } = loadPage({
    build: async () => ({ code: 0, message: '操作成功' }),
    view: async () => completed(),
  });
  await submit();
  assert.equal(calls.notifications.length, 1);
  assert.equal(page.genInfo.value.buildFiles, 11);
});

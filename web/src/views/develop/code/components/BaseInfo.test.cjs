const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const test = require('node:test');
const vm = require('node:vm');
const ts = require('typescript');
const vue = require('vue');

// 执行真实组件脚本，以可控网络和时钟复现热重启期间的初始化。
function loadComponent({ column, table, menu, id = 55 } = {}) {
  const source = fs.readFileSync(path.join(__dirname, 'BaseInfo.vue'), 'utf8');
  const script = source.match(/<script lang="ts" setup>([\s\S]*?)<\/script>/)[1];
  const props = vue.reactive({
    value: { id, dbName: 'default', tableName: 'hg_app_products', options: { join: [] } },
    selectList: {},
  });
  const timers = new Map();
  const errors = [];
  let timerId = 0;
  let mounted;
  let unmount;
  const modules = {
    vue: { ...vue, onMounted: (fn) => (mounted = fn), onBeforeUnmount: (fn) => (unmount = fn) },
    'naive-ui': {
      useMessage: () => ({ error: (text) => errors.push(text) }),
      useDialog: () => ({}),
    },
    './model': { newState: () => props.value, selectListObj: {} },
    '@/api/develop/code': {
      ColumnSelect: column || (async () => []),
      TableSelect: table || (async () => []),
    },
    '@/api/system/menu': { getMenuList: menu || (async () => ({ list: [] })) },
    'lodash-es': { cloneDeep: (value) => structuredClone(value) },
    '@/utils/charset': {},
    '@/utils/is': {},
    '@vicons/antd': {},
  };
  const schedule = (fn) => {
    timers.set(++timerId, fn);
    return timerId;
  };
  const context = {
    exports: {},
    console,
    require: (name) => (name.endsWith('.vue') ? {} : modules[name]),
    defineProps: () => props,
    withDefaults: (value) => value,
    defineEmits: () => () => {},
    setInterval: schedule,
    setTimeout: schedule,
    clearTimeout: (id) => timers.delete(id),
    clearInterval: (id) => timers.delete(id),
    window: { setTimeout: schedule, clearTimeout: (id) => timers.delete(id) },
  };
  const compiled = ts.transpileModule(script, {
    compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
  }).outputText;
  vm.runInNewContext(
    compiled +
      '\nexports.state = { bodyShow, tablesLoading, columnsLoading, linkColumnsLoading }; exports.retry = () => initialize();',
    context
  );
  mounted?.();
  async function flush() {
    await new Promise((resolve) => setImmediate(resolve));
  }
  async function settle() {
    await flush();
    for (let i = 0; i < 30 && timers.size; i++) {
      const [id, fn] = timers.entries().next().value;
      timers.delete(id);
      // 旧实现的异步定时回调会抛错；仍检查用户可见的 loading 状态。
      try {
        await fn();
      } catch (error) {
        errors.push(error.message);
      }
      await flush();
    }
  }
  return {
    state: context.exports.state,
    retry: context.exports.retry,
    props,
    timers,
    errors,
    settle,
    flush,
    unmount: () => unmount?.(),
  };
}

for (const api of ['column', 'table', 'menu']) {
  test(`${api} 初始化请求失败后结束请稍候并允许重新加载`, async () => {
    let available = false;
    const component = loadComponent({
      [api]: async () => {
        if (!available) throw new Error('服务正在重启');
        return api === 'menu' ? { list: [] } : [];
      },
    });
    await component.settle();
    assert.equal(component.state.bodyShow.value, false, '初始化失败后永久停留在请稍候');
    assert.equal(component.state.tablesLoading.value, false);
    assert.equal(component.state.columnsLoading.value, false);
    available = true;
    await component.retry();
    assert.equal(component.state.bodyShow.value, false);
    component.unmount();
  });
}

test('详情稍后返回有效 ID 时自动初始化，不使用永久轮询', async () => {
  let reads = 0;
  const component = loadComponent({
    id: 0,
    column: async () => {
      reads++;
      return [];
    },
  });
  await component.flush();
  assert.equal(reads, 0);
  component.props.value = { ...component.props.value, id: 55 };
  await component.settle();
  assert.equal(reads, 1);
  assert.equal(component.state.bodyShow.value, false);
  assert.equal(component.timers.size, 0);
  component.unmount();
});

test('等待详情时卸载组件会停止所有定时任务', async () => {
  const component = loadComponent({ id: 0 });
  component.unmount();
  assert.equal(component.timers.size, 0, '卸载后仍然每 30ms 轮询详情 ID');
});

test('卸载时终止重试，未完成的字段请求返回后不再请求表和菜单', async () => {
  let finish;
  let tables = 0;
  let menus = 0;
  const component = loadComponent({
    column: () => new Promise((resolve) => (finish = resolve)),
    table: async () => {
      tables++;
      return [];
    },
    menu: async () => {
      menus++;
      return { list: [] };
    },
  });
  component.unmount();
  finish([]);
  await component.settle();
  assert.equal(tables, 0, '页面卸载后仍请求表选项');
  assert.equal(menus, 0, '页面卸载后仍请求菜单');
  assert.equal(component.timers.size, 0);
});

test('重试等待期间卸载组件会清理定时器', async () => {
  let reads = 0;
  const component = loadComponent({
    column: async () => {
      reads++;
      throw new Error('502');
    },
  });
  await component.flush();
  assert.equal(component.timers.size, 1);
  component.unmount();
  await component.settle();
  assert.equal(reads, 1);
  assert.equal(component.timers.size, 0);
});

test('初始化请求短暂失败后自动恢复', async () => {
  let reads = 0;
  const component = loadComponent({
    column: async () => {
      if (++reads < 3) throw new Error('502');
      return [];
    },
  });
  await component.settle();
  assert.equal(reads, 3);
  assert.equal(component.state.bodyShow.value, false);
  component.unmount();
});

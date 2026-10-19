const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const test = require('node:test');
const vm = require('node:vm');
const ts = require('typescript');

// 使用真实 HTTP 响应处理逻辑，模拟请求上下文和消息组件，不发送网络请求。
function loadTransform() {
  let transform;
  const calls = { error: [], dialog: [], status: [] };
  const message = { error: (text) => calls.error.push(text) };
  const dialog = {
    info: (options) => calls.dialog.push(options),
    warning: (options) => calls.dialog.push(options),
  };
  const modules = {
    './Axios': {
      VAxios: class {
        constructor(options) {
          transform = options.transform;
        }
      },
    },
    axios: { default: { isCancel: () => false } },
    './checkStatus': { checkStatus: (...args) => calls.status.push(args) },
    '@/enums/httpEnum': {
      ContentTypeEnum: { JSON: 'application/json' },
      ResultEnum: { SUCCESS: 0, ERROR: -1, TIMEOUT: 61 },
    },
    '@/enums/pageEnum': { PageEnum: { BASE_LOGIN_NAME: 'Login', BASE_LOGIN: '/login' } },
    '@/hooks/setting': { useGlobSetting: () => ({}) },
    '@/utils': { deepMerge: (target, source) => ({ ...target, ...source }) },
    '@/router': { default: { currentRoute: { value: { name: 'generator' } } } },
  };
  const script = ts.transpileModule(fs.readFileSync(path.join(__dirname, 'index.ts'), 'utf8'), {
    compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
  }).outputText;
  vm.runInNewContext(script, {
    exports: {},
    require(name) {
      return modules[name] || {};
    },
    window: { $message: message, $dialog: dialog },
    console,
  });
  return { transform, calls };
}

test('生成结果核对显式关闭错误提示时，网络故障只向调用方拒绝，不弹出全局错误框', async () => {
  const { transform, calls } = loadTransform();
  for (const scenario of ['network', 'timeout', 'http']) {
    const error = new Error(scenario === 'network' ? 'Network Error' : 'timeout');
    error.config = { requestOptions: { isShowErrorMessage: false } };
    if (scenario === 'timeout') error.code = 'ECONNABORTED';
    if (scenario === 'http') error.response = { status: 502, data: { message: '服务重启中' } };
    await assert.rejects(
      Promise.resolve(transform.responseInterceptorsCatch(error)),
      (caught) => caught === error
    );
  }
  assert.equal(calls.error.length, 0);
  assert.equal(calls.dialog.length, 0);
  assert.equal(calls.status.length, 0);
});

test('详情核对中的业务失败仍抛出错误，但遵守显式关闭错误提示选项', () => {
  const { transform, calls } = loadTransform();
  assert.throws(
    () =>
      transform.transformRequestData(
        { data: { code: -1, message: '状态暂时不可用' } },
        { isTransformResponse: true, isShowErrorMessage: false }
      ),
    /状态暂时不可用/
  );
  assert.equal(calls.error.length, 0);
});

test('普通请求保留全局网络错误提示，登录失效仍显示登录提示', async () => {
  const { transform, calls } = loadTransform();
  const error = new Error('Network Error');
  await assert.rejects(transform.responseInterceptorsCatch(error));
  assert.equal(calls.dialog.length, 1);
  assert.throws(
    () =>
      transform.transformRequestData(
        { data: { code: 61, message: '登录失效' } },
        { isTransformResponse: true, isShowErrorMessage: false }
      ),
    /登录失效/
  );
  assert.equal(calls.dialog.length, 2, '核对请求屏蔽了登录失效提醒');
});

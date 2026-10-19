const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const test = require('node:test');
const vm = require('node:vm');
const ts = require('typescript');
const vue = require('vue');
const { NTag } = require('naive-ui');
const { parse, compileTemplate } = require('@vue/compiler-sfc');

function execute(source, modules, extra = {}) {
  const context = {
    exports: {},
    require: (name) => {
      if (!(name in modules)) throw new Error(`Unexpected import: ${name}`);
      return modules[name];
    },
    ...extra,
  };
  vm.runInNewContext(
    ts.transpileModule(source, {
      compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
    }).outputText,
    context
  );
  return context.exports;
}

function loadModel() {
  const options = vue.reactive({});
  const findOption = (key, value) =>
    options[key]?.find((item) => String(item.value) === String(value));
  const dict = {
    getOptionUnRef: (key) => options[key] || [],
    getLabel: (key, value) => findOption(key, value)?.label || '',
    getType: (key, value) => findOption(key, value)?.listClass || 'default',
    getExtra: (key, value) => findOption(key, value)?.extra,
    loadOptions: (keys) => keys.forEach((key) => (options[key] ??= [])),
  };
  const model = execute(fs.readFileSync(path.join(__dirname, 'model.ts'), 'utf8'), {
    vue,
    'naive-ui': { NTag },
    'lodash-es': { cloneDeep: (value) => structuredClone(value) },
    '@/utils/dateUtil': {
      defRangeShortcuts: () => ({}),
      formatToDateTime: (value) => `formatted:${value}`,
    },
    '@/utils/validateUtil': { validate: { qq: () => true } },
    '@/utils': {
      renderOptionTag: (key, value) =>
        vue.h(
          NTag,
          { type: dict.getType(key, value) },
          { default: () => dict.getLabel(key, value) }
        ),
    },
    '@/store/modules/dict': { useDictStore: () => dict },
  });
  return { model, options };
}

function tagText(node) {
  assert.equal(node.type, NTag);
  return node.children.default();
}

test('列表保留日常授权字段和创建时间，ID标题统一，敏感辅助字段放到详情', () => {
  const { model } = loadModel();
  assert.deepEqual(
    Array.from(model.columns, (column) => column.key),
    [
      'id',
      'name',
      'productId',
      'userId',
      'host',
      'versionId',
      'authMode',
      'authDomain',
      'status',
      'expiredAt',
      'createdAt',
    ]
  );
  assert.equal(model.columns[0].title, 'ID');
  assert.equal(model.schemas.value[0].label, 'ID');
  const createdAt = model.columns.find((column) => column.key === 'createdAt');
  assert.equal(
    createdAt.render({ createdAt: '2026-10-19 15:00:00' }),
    'formatted:2026-10-19 15:00:00'
  );
  assert.equal(createdAt.render({ createdAt: '' }), '-');
});

test('产品名称与验证策略使用后端字典，异步加载后保持响应式', () => {
  const { model, options } = loadModel();
  assert.equal(model.productOptions.value.length, 0);
  assert.equal(model.getProductLabel(7), '产品 #7');
  options.AppAuthProductOptions = [{ value: 7, label: 'product-name' }];
  options.AppAuthModeOptions = [
    { value: '0', label: '单域名验证', listClass: 'info' },
    { value: '1', label: 'IP双重验证', listClass: 'success', extra: { requireIp: true } },
  ];
  options.AppAuthDomainOptions = [
    { value: '0', label: '当前单域名', listClass: 'primary' },
    { value: '1', label: '顶级泛域名', listClass: 'warning' },
  ];
  assert.equal(model.productOptions.value[0].value, 7);
  assert.equal(model.getProductLabel(7), 'product-name');
  assert.equal(model.getProductLabel(null), '-');
  assert.equal(model.requiresIp('0'), false);
  assert.equal(model.requiresIp('1'), true);
  for (const [key, value, text, color] of [
    ['authMode', '0', '单域名验证', 'info'],
    ['authMode', '1', 'IP双重验证', 'success'],
    ['authDomain', '0', '当前单域名', 'primary'],
    ['authDomain', '1', '顶级泛域名', 'warning'],
  ]) {
    const node = model.columns.find((column) => column.key === key).render({ [key]: value });
    assert.equal(tagText(node), text);
    assert.equal(node.props.type, color);
  }
  const label = model.renderDictOption(options.AppAuthDomainOptions[1]);
  assert.equal(tagText(label), '顶级泛域名');
  assert.equal(label.props.type, 'warning');
});

test('用户选项加载后更新标签，用户与版本使用Naive UI标签而非原始ID', () => {
  const { model, options } = loadModel();
  const userColumn = model.columns.find((column) => column.key === 'userId');
  assert.equal(model.userOptions.value.length, 0);
  assert.equal(tagText(userColumn.render({ userId: 12 })), '用户 #12');
  model.loadOptions();
  assert.ok('AppAuthUserOptions' in options);
  options.AppAuthUserOptions = [{ value: 12, label: 'demo（演示用户） #12', listClass: 'primary' }];
  assert.equal(model.userOptions.value[0].value, 12);
  assert.equal(tagText(userColumn.render({ userId: 12 })), 'demo（演示用户） #12');
  assert.equal(tagText(model.renderUserOption(model.userOptions.value[0])), 'demo（演示用户） #12');
  assert.equal(model.getUserLabel(null), '-');
  assert.equal(model.renderUserTag(null), '-');
  const version = model.columns
    .find((column) => column.key === 'versionId')
    .render({ versionId: '1.2.3' });
  assert.equal(tagText(version), '1.2.3');
  assert.equal(version.props.size, 'small');
  assert.equal(model.renderVersionTag(''), '-');
});

test('搜索仅使用后端支持的字段，产品ID与版本号保持不同的数据类型', () => {
  const { model } = loadModel();
  assert.deepEqual(
    Array.from(model.schemas.value, (schema) => schema.field),
    ['id', 'status', 'createdAt']
  );
  const state = model.newState(null);
  assert.equal(state.productId, null);
  assert.equal(state.userId, null);
  assert.equal(state.versionId, '1.0.0');
  assert.equal(model.rules.productId.type, 'number');
  assert.equal(model.rules.userId.message, '请选择所属用户');
  assert.equal(model.rules.authMode.type, 'string');
});

test('编辑回填保留隐藏字段和用户ID，打开新增时清空上次授权', async () => {
  const { model } = loadModel();
  const source = fs.readFileSync(path.join(__dirname, 'edit.vue'), 'utf8');
  const { descriptor } = parse(source);
  let resolveView;
  let exposed;
  const component = execute(
    descriptor.scriptSetup.content +
      '\nexports.state = { formValue, loading, isFormValid, formRules };',
    {
      vue,
      'naive-ui': { useMessage: () => ({ success() {}, error() {} }) },
      './model': model,
      '@/api/addons/appAuth/appAuth': {
        View: () => new Promise((resolve) => (resolveView = resolve)),
      },
      '@/components/DatePicker/datePicker.vue': {},
      '@/store/modules/projectSetting': { useProjectSettingStore: () => ({ isMobile: false }) },
      '@/utils/hotgo': { adaModalWidth: () => '840px' },
    },
    { defineEmits: () => () => {}, defineExpose: (value) => (exposed = value) }
  );
  exposed.openModal(model.newState({ id: 12 }));
  assert.equal(component.state.isFormValid.value, false);
  resolveView({ id: 12, productId: 7, userId: 12, views: 88, qq: '1234567', ip: '192.0.2.1' });
  await new Promise((resolve) => setImmediate(resolve));
  assert.equal(component.state.formValue.value.views, 88);
  assert.equal(component.state.formValue.value.ip, '192.0.2.1');
  assert.equal(component.state.formValue.value.productId, 7);
  assert.equal(component.state.formValue.value.userId, 12);
  assert.equal(component.state.isFormValid.value, true);
  exposed.openModal(null);
  assert.equal(component.state.formValue.value.id, 0);
  assert.equal(component.state.formValue.value.productId, null);
  assert.equal(component.state.formValue.value.userId, null);
});

test('编辑字段匹配选择组件，详情使用标签，保留原有布局且模板可编译', () => {
  for (const filename of ['index.vue', 'edit.vue', 'view.vue']) {
    const source = fs.readFileSync(path.join(__dirname, filename), 'utf8');
    const { descriptor, errors } = parse(source);
    assert.deepEqual(errors, [], filename);
    const result = compileTemplate({
      source: descriptor.template.content,
      filename,
      id: 'app-auth-test',
    });
    assert.deepEqual(result.errors, [], filename);
    assert.ok(
      descriptor.styles.every((style) => style.content.trim() === ''),
      filename
    );
  }
  const edit = fs.readFileSync(path.join(__dirname, 'edit.vue'), 'utf8');
  const { descriptor } = parse(edit);
  const selectFields = [];
  const visit = (node) => {
    if (node.tag === 'n-select') {
      const model = node.props.find((prop) => prop.name === 'model');
      if (model) selectFields.push(model.exp.content);
    }
    for (const child of node.children || []) visit(child);
  };
  visit(descriptor.template.ast);
  assert.deepEqual(
    selectFields.sort(),
    ['productId', 'userId', 'authMode', 'authDomain', 'status']
      .map((field) => `formValue.${field}`)
      .sort()
  );
  assert.doesNotMatch(edit, /v-model:value="formValue\.(views|createdAt|updatedAt)"/);
  assert.match(edit, /:options="userOptions"/);
  assert.match(edit, /:render-label="renderUserOption"/);
  assert.match(edit, /show-password-on="click"/);
  assert.match(edit, /DatePicker v-model:formValue="formValue\.expiredAt"/);
  assert.match(edit, /cols="1 s:1 m:1 l:1 xl:1 2xl:1"/);
  const view = fs.readFileSync(path.join(__dirname, 'view.vue'), 'utf8');
  assert.match(view, /label="ID"/);
  assert.match(view, /getUserLabel\(formValue.userId\)/);
  assert.match(view, /dict.getType\('AppAuthModeOptions'/);
  assert.match(view, /dict.getType\('AppAuthDomainOptions'/);
});

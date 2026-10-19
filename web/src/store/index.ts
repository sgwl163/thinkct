import type { App } from 'vue';
import { createPinia } from 'pinia';
import { get, set } from 'lodash-es';
import { storage } from '@/utils/Storage';
import { DESIGN_CFG_KEY, LIST_CFG_KEY, PROJ_CFG_KEY } from '@/enums/cacheEnum';
import { animates } from '@/settings/animateSetting';

type SettingValidator = (value: unknown) => boolean;

const store = createPinia();
const isBoolean: SettingValidator = (value) => typeof value === 'boolean';
const isPositiveNumber: SettingValidator = (value) =>
  typeof value === 'number' && Number.isFinite(value) && value > 0;
const isHexColor: SettingValidator = (value) =>
  typeof value === 'string' && /^#(?:[\da-f]{3}|[\da-f]{6})$/i.test(value);

const settingsCache: Record<string, { key: string; fields: Record<string, SettingValidator> }> = {
  'app-project-setting': {
    key: PROJ_CFG_KEY,
    fields: {
      navMode: (value) =>
        typeof value === 'string' && ['vertical', 'horizontal', 'horizontal-mix'].includes(value),
      navTheme: (value) =>
        typeof value === 'string' && ['dark', 'light', 'header-dark'].includes(value),
      'headerSetting.bgColor': isHexColor,
      'headerSetting.fixed': isBoolean,
      'headerSetting.isReload': isBoolean,
      showFooter: isBoolean,
      'menuSetting.minMenuWidth': isPositiveNumber,
      'menuSetting.menuWidth': isPositiveNumber,
      'menuSetting.fixed': isBoolean,
      'menuSetting.mixMenu': isBoolean,
      'menuSetting.mobileWidth': isPositiveNumber,
      'menuSetting.collapsed': isBoolean,
      'multiTabsSetting.bgColor': isHexColor,
      'multiTabsSetting.show': isBoolean,
      'multiTabsSetting.fixed': isBoolean,
      'crumbsSetting.show': isBoolean,
      'crumbsSetting.showIcon': isBoolean,
      isPageAnimate: isBoolean,
      pageAnimateType: (value) => animates.some((item) => item.value === value),
    },
  },
  'app-list-setting': {
    key: LIST_CFG_KEY,
    fields: {
      searchCollapsed: isBoolean,
      striped: isBoolean,
      size: (value) => typeof value === 'string' && ['small', 'medium', 'large'].includes(value),
      pageSize: (value) => typeof value === 'number' && Number.isSafeInteger(value) && value > 0,
    },
  },
  'app-design-setting': {
    key: DESIGN_CFG_KEY,
    fields: {
      darkTheme: isBoolean,
      // 主题色还用于生成亮色阶，只接受完整的六位十六进制颜色。
      appTheme: (value) => typeof value === 'string' && /^#[\da-f]{6}$/i.test(value),
    },
  },
};

function getSettingsState(state: unknown, fields: Record<string, SettingValidator>) {
  const settings = {};
  // 按叶子字段校验，防止错误类型或未知的嵌套字段覆盖默认配置。
  for (const [path, validate] of Object.entries(fields)) {
    const value = get(state, path);
    if (validate(value)) {
      set(settings, path, value);
    }
  }
  return settings;
}

// 仅持久化界面偏好，移动端状态、权限模式和内置主题列表仍使用运行时或默认配置。
store.use(({ store }) => {
  const config = settingsCache[store.$id];
  if (!config) {
    return;
  }

  try {
    const saved = storage.get(config.key);
    // $patch 深度合并，旧缓存缺少或校验不通过的字段继续使用当前默认值。
    store.$patch(getSettingsState(saved, config.fields));
  } catch (error) {
    console.warn('读取界面配置失败，已使用默认配置。', error);
  }

  if (store.$id === 'app-project-setting') {
    if (store.navMode !== 'horizontal-mix') {
      store.menuSetting.mixMenu = false;
    }
    if (store.navMode === 'horizontal' && store.navTheme === 'light') {
      store.navTheme = 'dark';
    }
  }

  store.$subscribe(
    (_mutation, state) => {
      try {
        storage.set(config.key, getSettingsState(state, config.fields), null);
      } catch (error) {
        // 浏览器禁用存储或空间不足时，仍允许当前页面正常切换配置。
        console.warn('保存界面配置失败，本次修改仅在当前页面生效。', error);
      }
    },
    { detached: true, flush: 'sync' }
  );
});

export function setupStore(app: App<Element>) {
  app.use(store);
}

export { store };

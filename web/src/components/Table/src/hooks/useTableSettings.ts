import { ref, computed, unref, getCurrentInstance, type ComputedRef } from 'vue';
import { useListSettingStore, type TableDisplaySettings } from '@/store/modules/listSetting';
import { useRoute } from 'vue-router';
import { storage } from '@/utils/Storage';
import { TABLE_CFG_KEY } from '@/enums/cacheEnum';
import type { BasicTableProps } from '../types/table';

export interface TableColumnSetting {
  key: string | number;
  visible: boolean;
  fixed?: 'left' | 'right' | false;
}

export interface TableSettings extends TableDisplaySettings {
  selection?: boolean;
  columns?: TableColumnSetting[];
}

function validateSettings(value: unknown): TableSettings {
  if (!value || typeof value !== 'object' || Array.isArray(value)) return {};
  const saved = value as TableSettings;
  const settings: TableSettings = {};
  if (typeof saved.striped === 'boolean') settings.striped = saved.striped;
  if (['small', 'medium', 'large'].includes(saved.size!)) settings.size = saved.size;
  if (Number.isSafeInteger(saved.pageSize) && saved.pageSize! > 0) {
    settings.pageSize = saved.pageSize;
  }
  if (typeof saved.selection === 'boolean') settings.selection = saved.selection;
  if (Array.isArray(saved.columns)) {
    const keys = new Set<string | number>();
    settings.columns = saved.columns
      .filter((column) => {
        if (!column || typeof column !== 'object') return false;
        if (typeof column.key !== 'string' && typeof column.key !== 'number') return false;
        if (typeof column.visible !== 'boolean' || keys.has(column.key)) return false;
        if (column.fixed !== undefined && !['left', 'right', false].includes(column.fixed)) {
          return false;
        }
        keys.add(column.key);
        return true;
      })
      .map(({ key, visible, fixed }) => ({ key, visible, fixed }));
  }
  return settings;
}

export function useTableSettings(propsRef: ComputedRef<BasicTableProps>) {
  const route = useRoute();
  const props = unref(propsRef);
  // 捕获所属页面，避免 keep-alive 页面在路由切换后写入其他列表的配置。
  // 同页多表格默认按所属组件和 ref 区分，也可通过 cacheKey 指定稳定标识。
  const instance = getCurrentInstance();
  const owners: string[] = [];
  let parent = instance?.parent;
  while (parent) {
    const type = parent.type as { name?: string; __name?: string };
    owners.push(type.name || type.__name || 'component');
    parent = parent.parent;
  }
  const tableRef = (instance?.vnode.ref as { r?: unknown } | null)?.r;
  const identity =
    props.cacheKey || JSON.stringify([owners, typeof tableRef === 'string' ? tableRef : 'default']);
  const key = `${TABLE_CFG_KEY}${JSON.stringify([route?.path || '', identity])}`;
  const localSettings = ref<TableSettings>({});
  const listSetting = useListSettingStore();
  try {
    const saved = validateSettings(storage.get(key));
    // 兼容旧缓存中的列设置，旧版按页面保存的通用样式不再使用。
    localSettings.value = { columns: saved.columns, selection: saved.selection };
  } catch (error) {
    console.warn('读取表格配置失败，已使用默认配置。', error);
  }
  const settings = computed<TableSettings>(() => ({
    ...localSettings.value,
    striped: listSetting.striped,
    size: listSetting.size,
    pageSize: listSetting.pageSize,
  }));

  function saveSettings(value: Partial<TableSettings>) {
    const validated = validateSettings(value);
    const shared: TableDisplaySettings = {};
    for (const field of ['striped', 'size', 'pageSize'] as const) {
      if (validated[field] !== undefined) Object.assign(shared, { [field]: validated[field] });
    }
    if (Object.keys(shared).length) listSetting.setTableSettings(shared);
    if (!('columns' in value) && !('selection' in value)) return;

    localSettings.value = validateSettings({ ...localSettings.value, ...value });
    const { columns, selection } = localSettings.value;
    localSettings.value = { columns, selection };
    try {
      // 每个列表只保存列展示偏好，不保存表格数据、搜索条件或渲染函数。
      storage.set(key, localSettings.value, null);
    } catch (error) {
      console.warn('保存表格配置失败，本次修改仅在当前页面生效。', error);
    }
  }

  return { settings, saveSettings };
}

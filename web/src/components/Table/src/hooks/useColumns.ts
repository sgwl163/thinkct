import { ref, Ref, ComputedRef, unref, computed, watch, toRaw, h } from 'vue';
import type { BasicColumn, BasicTableProps } from '../types/table';
import { cloneDeep } from 'lodash-es';
import { isArray, isBoolean, isFunction } from '@/utils/is';
import type { useTableSettings, TableColumnSetting } from './useTableSettings';
import { usePermission } from '@/hooks/web/usePermission';
import { ActionItem } from '@/components/Table';
import { renderEditCell } from '../components/editable';
import { NTooltip, NIcon } from 'naive-ui';
import { FormOutlined } from '@vicons/antd';

export function useColumns(
  propsRef: ComputedRef<BasicTableProps>,
  preferences: ReturnType<typeof useTableSettings>
) {
  const { settings, saveSettings } = preferences;
  const columnsRef = ref<BasicColumn[]>([]) as Ref<BasicColumn[]>;
  const cacheColumns = ref<BasicColumn[]>([]) as Ref<BasicColumn[]>;
  const selection = ref(settings.value.selection ?? !!unref(propsRef).openChecked);

  function restoreColumns() {
    const defaults = cloneDeep(unref(propsRef).columns);
    const saved = settings.value.columns || [];
    const known = new Set(saved.map((column) => column.key));
    const ordered = saved.flatMap((setting) => {
      const column = defaults.find((item) => item.key === setting.key);
      return column
        ? [
            {
              ...column,
              fixed: setting.fixed === undefined ? column.fixed : setting.fixed || undefined,
            },
          ]
        : [];
    });
    ordered.push(...defaults.filter((column) => !known.has(column.key)));
    cacheColumns.value = ordered;
    columnsRef.value = ordered.filter((column) => {
      const setting = saved.find((item) => item.key === column.key);
      return !setting || setting.visible;
    });
  }

  restoreColumns();

  function saveColumns() {
    const visible = new Set(columnsRef.value.map((column) => column.key));
    const columns: TableColumnSetting[] = cacheColumns.value
      .filter((column) => typeof column.key === 'string' || typeof column.key === 'number')
      .map((column) => ({
        key: column.key,
        visible: visible.has(column.key),
        fixed: column.fixed || false,
      }));
    saveSettings({ columns, selection: selection.value });
  }

  const getColumnsRef = computed(() => {
    const columns = cloneDeep(unref(columnsRef));
    if (selection.value && !columns.some((column) => (column as any).type === 'selection')) {
      columns.unshift({ type: 'selection', key: 'selection' } as any);
    }
    handleActionColumn(propsRef, columns);
    return columns;
  });

  const { hasPermission } = usePermission();

  function isIfShow(action: ActionItem): boolean {
    const ifShow = action.ifShow;

    let isIfShow = true;

    if (isBoolean(ifShow)) {
      isIfShow = ifShow;
    }
    if (isFunction(ifShow)) {
      isIfShow = ifShow(action);
    }
    return isIfShow;
  }

  const renderTooltip = (trigger, content) => {
    return h(NTooltip, null, {
      trigger: () => trigger,
      default: () => content,
    });
  };

  const getPageColumns = computed(() => {
    const pageColumns = unref(getColumnsRef);
    const columns = cloneDeep(pageColumns);
    return columns
      .filter((column) => {
        return hasPermission(column.auth as string[]) && isIfShow(column);
      })
      .map((column) => {
        //默认 ellipsis 为true
        column.ellipsis = typeof column.ellipsis === 'undefined' ? { tooltip: true } : false;
        const { edit } = column;
        if (edit) {
          column.render = renderEditCell(column);
          if (edit) {
            const title: any = column.title;
            column.title = () => {
              return renderTooltip(
                h('span', {}, [
                  h('span', { style: { 'margin-right': '5px' } }, title),
                  h(
                    NIcon,
                    {
                      size: 14,
                    },
                    {
                      default: () => h(FormOutlined),
                    }
                  ),
                ]),
                '该列可编辑'
              );
            };
          }
        }
        return column;
      });
  });

  watch(() => unref(propsRef).columns, restoreColumns);

  function handleActionColumn(propsRef: ComputedRef<BasicTableProps>, columns: BasicColumn[]) {
    const { actionColumn } = unref(propsRef);
    if (!actionColumn) return;
    if (!columns.find((col) => col.key === 'action')) {
      if ((actionColumn as any).fixed && (actionColumn as any).fixed === 'left') {
        columns.unshift({
          ...(actionColumn as any),
        });
      } else {
        columns.push({
          ...(actionColumn as any),
        });
      }
    }
  }

  // 更新可见列，同时保留隐藏列的位置和固定设置。
  function setColumns(columnList: BasicColumn[] | (string | number)[]) {
    if (!isArray(columnList)) return;
    const visible = new Map<string | number, BasicColumn>();
    for (const value of columnList) {
      const key = typeof value === 'object' ? value.key : value;
      const column = cacheColumns.value.find((item) => item.key === key);
      if (typeof value === 'object') {
        const updated = { ...column, ...cloneDeep(value) };
        if (column) Object.assign(column, updated);
        else cacheColumns.value.push(updated);
        visible.set(key, updated);
      } else if (column) {
        visible.set(key, column);
      }
    }
    const ordered = [...visible.values()];
    let index = 0;
    cacheColumns.value = cacheColumns.value.map((column) =>
      visible.has(column.key) ? ordered[index++] : column
    );
    columnsRef.value = ordered;
    saveColumns();
  }

  function setColumnOrder(keys: (string | number)[]) {
    const columns = cacheColumns.value;
    cacheColumns.value = [
      ...[...new Set(keys)].flatMap((key) => columns.filter((column) => column.key === key)),
      ...columns.filter((column) => !keys.includes(column.key)),
    ];
    const visible = new Set(columnsRef.value.map((column) => column.key));
    columnsRef.value = cacheColumns.value.filter((column) => visible.has(column.key));
    saveColumns();
  }

  function setSelection(value: boolean) {
    selection.value = value;
    saveColumns();
  }

  function resetColumns() {
    selection.value = !!unref(propsRef).openChecked;
    saveSettings({ columns: undefined, selection: undefined });
    restoreColumns();
  }

  //获取
  function getColumns(): BasicColumn[] {
    const columns = toRaw(unref(getColumnsRef));
    return columns.map((item) => {
      return { ...item, title: item.title, key: item.key, fixed: item.fixed || undefined };
    });
  }

  //获取原始
  function getCacheColumns(isKey?: boolean): any[] {
    return isKey ? cacheColumns.value.map((item) => item.key) : cloneDeep(cacheColumns.value);
  }

  // 按 key 更新，避免隐藏列或勾选列改变下标后固定到错误的列。
  function setCacheColumnsField(key: string | number | undefined, value: Partial<BasicColumn>) {
    for (const columns of [cacheColumns.value, columnsRef.value]) {
      const column = columns.find((item) => item.key === key);
      if (column) Object.assign(column, value);
    }
    saveColumns();
  }

  return {
    getColumnsRef,
    getCacheColumns,
    setCacheColumnsField,
    setColumns,
    setColumnOrder,
    setSelection,
    resetColumns,
    getColumns,
    getPageColumns,
  };
}

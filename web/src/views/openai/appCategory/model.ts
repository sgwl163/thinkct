import { h, ref } from 'vue';
import { NSwitch } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Status } from '@/api/openai/appCategory';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0; // 分类ID
  public name = ''; // 分类名称
  public sort = 0; // 排序
  public status = 1; // 状态(1:启用, 2:禁用)
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入分类名称',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '分类名称',
    componentProps: {
      placeholder: '请输入分类名称',
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
      options: dict.getOption('sys_normal_disable'),
      placeholder: '请选择状态',
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
    },
  },
]);

import { computed } from 'vue';

// 表格列
export const columns = [
  {
    title: '分类ID',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '分类名称',
    key: 'name',
    align: 'left',
    width: 200,
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: 100,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 150,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        checked: '启用',
        unchecked: '禁用',
        disabled: !hasPermission(['/openaiAppCategory/status']),
        onUpdateValue: function (e) {
          row.status = e ? 1 : 2;
          Status({ id: row.id, status: row.status }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: 180,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

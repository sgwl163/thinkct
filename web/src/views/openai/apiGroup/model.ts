import { h, ref } from 'vue';
import { NSwitch, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Status } from '@/api/openai/apiGroup';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0; // ID
  public name = ''; // 组名称
  public baseUrl = ''; // 接口地址
  public apiKey = ''; // API Key
  public status = 1; // 状态
  public sort = 0; // 排序
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间

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
    message: '请输入组名称',
  },
  baseUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入接口地址',
  },
  apiKey: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入API Key',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '组名称',
    componentProps: {
      placeholder: '请输入组名称',
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: dict.getOption('sys_normal_disable'),
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

// 表格列
export const columns = [
  {
    title: '组名称',
    key: 'name',
    align: 'left',
    width: 150,
  },
  {
    title: '接口地址',
    key: 'baseUrl',
    align: 'left',
    width: 220,
    ellipsis: { tooltip: true },
  },
  {
    title: 'API Key',
    key: 'apiKey',
    align: 'left',
    width: 180,
    render(row: State) {
      if (!row.apiKey) return '';
      // 脱敏显示
      const key = row.apiKey;
      return key.length > 10 ? key.slice(0, 6) + '****' + key.slice(-4) : '****';
    },
  },
  {
    title: '排序',
    key: 'sort',
    align: 'center',
    width: 70,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        disabled: !hasPermission(['/apiGroup/status']),
        onUpdateValue: (val: boolean) => {
          row.status = val ? 1 : 2;
          Status({ id: row.id, status: row.status }).then(() => {
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
    width: 160,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

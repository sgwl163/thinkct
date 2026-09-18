import { h, ref } from 'vue';
import { NSwitch } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage } from '@/utils';
import { Status } from '@/api/openai/app';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0; // 应用ID
  public categoryId = 0; // 分类ID
  public title = ''; // 应用名称
  public icon = ''; // 应用图标
  public description = ''; // 应用描述
  public prompt = ''; // 系统提示词
  public useCount = 0; // 使用次数
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
  categoryId: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    message: '请选择应用分类',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入应用名称',
  },
  prompt: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入系统提示词',
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
    field: 'title',
    component: 'NInput',
    label: '应用名称',
    componentProps: {
      placeholder: '请输入应用名称',
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
    title: '应用名称',
    key: 'title',
    align: 'left',
    width: 160,
  },
  {
    title: '应用图标',
    key: 'icon',
    align: 'center',
    width: 80,
    render(row: State) {
      return renderImage(row.icon);
    },
  },
  {
    title: '应用描述',
    key: 'description',
    align: 'left',
    width: 220,
    ellipsis: { tooltip: true },
  },
  {
    title: '使用次数',
    key: 'useCount',
    align: 'center',
    width: 90,
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
        checkedValue: true,
        uncheckedValue: false,
        disabled: !hasPermission(['/app/status']),
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

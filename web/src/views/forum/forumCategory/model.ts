import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // 板块ID
  public pid = 0; // 上级板块ID
  public name = ''; // 板块名称
  public description = ''; // 板块简介
  public image = ''; // 板块图标
  public sort = 0; // 排序
  public status = 1; // 状态:1正常,2禁用
  public topicCount = 0; // 主题数量
  public updatedAt = ''; // 修改时间
  public createdAt = ''; // 创建时间

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
  pid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入上级板块ID',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入板块名称',
  },
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入板块简介',
  },
  image: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入板块图标',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态:1正常,2禁用',
  },
  topicCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入主题数量',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '板块ID',
    componentProps: {
      placeholder: '请输入板块ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态:1正常,2禁用',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态:1正常,2禁用',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
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
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '板块ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '上级板块ID',
    key: 'pid',
    align: 'left',
    width: -1,
  },
  {
    title: '板块名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '板块图标',
    key: 'image',
    align: 'left',
    width: -1,
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:1正常,2禁用',
    key: 'status',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '主题数量',
    key: 'topicCount',
    align: 'left',
    width: -1,
  },
  {
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
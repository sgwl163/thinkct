import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 举报ID
  public targetType = ''; // 举报对象类型:topic,comment
  public targetId = 0; // 举报对象ID
  public userId = 0; // 举报用户ID
  public reason = ''; // 举报原因
  public description = ''; // 补充说明
  public status = 'pending'; // 处理状态:pending,processed,rejected
  public handledBy = 0; // 处理管理员ID
  public handledAt = ''; // 处理时间
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
  targetType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入举报对象类型:topic,comment',
  },
  targetId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入举报对象ID',
  },
  userId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入举报用户ID',
  },
  reason: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入举报原因',
  },
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入补充说明',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入处理状态:pending,processed,rejected',
  },
  handledBy: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入处理管理员ID',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '举报ID',
    componentProps: {
      placeholder: '请输入举报ID',
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
    title: '举报ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '举报对象类型:topic,comment',
    key: 'targetType',
    align: 'left',
    width: -1,
  },
  {
    title: '举报对象ID',
    key: 'targetId',
    align: 'left',
    width: -1,
  },
  {
    title: '举报用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '举报原因',
    key: 'reason',
    align: 'left',
    width: -1,
  },
  {
    title: '处理状态:pending,processed,rejected',
    key: 'status',
    align: 'left',
    width: -1,
  },
  {
    title: '处理管理员ID',
    key: 'handledBy',
    align: 'left',
    width: -1,
  },
  {
    title: '处理时间',
    key: 'handledAt',
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
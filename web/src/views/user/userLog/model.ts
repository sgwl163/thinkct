import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // 日志ID
  public userId = 0; // 用户ID
  public username = ''; // 用户名
  public logType = 'login'; // 记录类型:login=登录,operation=操作
  public action = 'login'; // 操作标识
  public actionName = '用户登录'; // 操作名称
  public eventAt = ''; // 事件时间
  public ip = ''; // IP地址
  public userAgent = ''; // UA信息
  public message = ''; // 结果消息
  public extra = ''; // 扩展数据
  public status = 1; // 状态
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
  userId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入用户ID',
  },
  username: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入用户名',
  },
  logType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入记录类型:login=登录,operation=操作',
  },
  action: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入操作标识',
  },
  actionName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入操作名称',
  },
  ip: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入IP地址',
  },
  userAgent: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入UA信息',
  },
  message: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入结果消息',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '日志ID',
    componentProps: {
      placeholder: '请输入日志ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
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
    title: '日志ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '用户名',
    key: 'username',
    align: 'left',
    width: -1,
  },
  {
    title: '记录类型:login=登录,operation=操作',
    key: 'logType',
    align: 'left',
    width: -1,
  },
  {
    title: '操作标识',
    key: 'action',
    align: 'left',
    width: -1,
  },
  {
    title: '操作名称',
    key: 'actionName',
    align: 'left',
    width: -1,
  },
  {
    title: '事件时间',
    key: 'eventAt',
    align: 'left',
    width: -1,
  },
  {
    title: 'IP地址',
    key: 'ip',
    align: 'left',
    width: -1,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
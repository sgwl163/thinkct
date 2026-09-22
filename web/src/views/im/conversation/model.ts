import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 会话ID
  public userId = 0; // 归属用户ID
  public chatType = 1; // 会话类型: 1=单聊, 2=群聊, 3=系统通知

  public targetId = 0; // 目标ID(好友ID或群组ID)
  public lastMsgId = 0; // 最后一条消息ID
  public lastMsgContent = ''; // 最后消息内容摘要(用于列表展示)
  public lastMsgTime = ''; // 最后消息时间
  public unreadCount = 0; // 未读消息数
  public isTop = 0; // 是否置顶: 0=否, 1=是
  public isDisturb = 0; // 消息免打扰: 0=否, 1=是
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
    message: '请输入归属用户ID',
  },
  chatType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入会话类型: 1=单聊, 2=群聊, 3=系统通知',
  },
  targetId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入目标ID(好友ID或群组ID)',
  },
  unreadCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入未读消息数',
  },
  isTop: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否置顶: 0=否, 1=是',
  },
  isDisturb: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入消息免打扰: 0=否, 1=是',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '会话ID',
    componentProps: {
      placeholder: '请输入会话ID',
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
    title: '会话ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '归属用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '会话类型: 1=单聊, 2=群聊, 3=系统通知',
    key: 'chatType',
    align: 'left',
    width: -1,
  },
  {
    title: '目标ID(好友ID或群组ID)',
    key: 'targetId',
    align: 'left',
    width: -1,
  },
  {
    title: '最后一条消息ID',
    key: 'lastMsgId',
    align: 'left',
    width: -1,
  },
  {
    title: '最后消息时间',
    key: 'lastMsgTime',
    align: 'left',
    width: -1,
  },
  {
    title: '未读消息数',
    key: 'unreadCount',
    align: 'left',
    width: -1,
  },
  {
    title: '是否置顶: 0=否, 1=是',
    key: 'isTop',
    align: 'left',
    width: -1,
  },
  {
    title: '消息免打扰: 0=否, 1=是',
    key: 'isDisturb',
    align: 'left',
    width: -1,
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
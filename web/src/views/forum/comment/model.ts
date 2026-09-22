import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 评论ID
  public topicId = 0; // 主题ID
  public userId = 0; // 评论用户ID
  public parentId = 0; // 父评论ID
  public replyUserId = 0; // 被回复用户ID
  public content = ''; // 评论内容
  public status = 'normal'; // 状态:normal,pending,rejected,hidden
  public likeCount = 0; // 点赞数量
  public updatedAt = ''; // 修改时间
  public createdAt = ''; // 创建时间
  public deletedAt = ''; // 删除时间

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
  topicId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入主题ID',
  },
  userId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入评论用户ID',
  },
  parentId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入父评论ID',
  },
  replyUserId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入被回复用户ID',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入评论内容',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入状态:normal,pending,rejected,hidden',
  },
  likeCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入点赞数量',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '评论ID',
    componentProps: {
      placeholder: '请输入评论ID',
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
    title: '评论ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '主题ID',
    key: 'topicId',
    align: 'left',
    width: -1,
  },
  {
    title: '评论用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '父评论ID',
    key: 'parentId',
    align: 'left',
    width: -1,
  },
  {
    title: '被回复用户ID',
    key: 'replyUserId',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:normal,pending,rejected,hidden',
    key: 'status',
    align: 'left',
    width: -1,
  },
  {
    title: '点赞数量',
    key: 'likeCount',
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
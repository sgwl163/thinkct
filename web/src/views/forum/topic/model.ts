import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 主题ID
  public categoryId = 0; // 板块ID
  public userId = 0; // 发布用户ID
  public title = ''; // 主题标题
  public keywords = ''; // 主题关键词，英文逗号分隔
  public description = ''; // 主题摘要
  public content = ''; // 主题内容
  public images = ''; // 主题图片，英文逗号分隔
  public status = 'normal'; // 状态:normal,pending,rejected,hidden,closed
  public isTop = 0; // 是否置顶:0否,1是
  public isEssence = 0; // 是否精华:0否,1是
  public commentCount = 0; // 评论数量
  public likeCount = 0; // 点赞数量
  public favoriteCount = 0; // 收藏数量
  public viewCount = 0; // 浏览数量
  public lastReplyUserId = 0; // 最后回复用户ID
  public lastReplyAt = ''; // 最后回复时间
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
  categoryId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入板块ID',
  },
  userId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入发布用户ID',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入主题标题',
  },
  keywords: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入主题关键词，英文逗号分隔',
  },
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入主题摘要',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入主题内容',
  },
  images: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入主题图片，英文逗号分隔',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入状态:normal,pending,rejected,hidden,closed',
  },
  isTop: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否置顶:0否,1是',
  },
  isEssence: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否精华:0否,1是',
  },
  commentCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入评论数量',
  },
  likeCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入点赞数量',
  },
  favoriteCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入收藏数量',
  },
  viewCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入浏览数量',
  },
  lastReplyUserId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入最后回复用户ID',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '主题ID',
    componentProps: {
      placeholder: '请输入主题ID',
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
    title: '主题ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '板块ID',
    key: 'categoryId',
    align: 'left',
    width: -1,
  },
  {
    title: '发布用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '主题标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:normal,pending,rejected,hidden,closed',
    key: 'status',
    align: 'left',
    width: -1,
  },
  {
    title: '是否置顶:0否,1是',
    key: 'isTop',
    align: 'left',
    width: -1,
  },
  {
    title: '是否精华:0否,1是',
    key: 'isEssence',
    align: 'left',
    width: -1,
  },
  {
    title: '评论数量',
    key: 'commentCount',
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
    title: '收藏数量',
    key: 'favoriteCount',
    align: 'left',
    width: -1,
  },
  {
    title: '浏览数量',
    key: 'viewCount',
    align: 'left',
    width: -1,
  },
  {
    title: '最后回复用户ID',
    key: 'lastReplyUserId',
    align: 'left',
    width: -1,
  },
  {
    title: '最后回复时间',
    key: 'lastReplyAt',
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
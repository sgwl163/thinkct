import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // ID
  public name = ''; // 组名
  public rules = ''; // 权限节点
  public status = 1; // 状态:0=禁用,1=启用
  public updateTime = 0; // 更新时间
  public createTime = 0; // 创建时间

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
    message: '请输入组名',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入状态:0=禁用,1=启用',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '组名',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:0=禁用,1=启用',
    key: 'status',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updateTime',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createTime',
    align: 'left',
    width: -1,
  },
];
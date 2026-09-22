import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // ID
  public productId = 0; // 产品ID
  public host = ''; // 域名
  public ip = '0.0.0.0'; // IP
  public version = '0.0.0'; // 版本号
  public code = ''; // 执行代码
  public data = ''; // 数据
  public views = 1; // 浏览量
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
  version: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入版本号',
  },
  views: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入浏览量',
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
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '产品ID',
    key: 'productId',
    align: 'left',
    width: -1,
  },
  {
    title: '域名',
    key: 'host',
    align: 'left',
    width: -1,
  },
  {
    title: 'IP',
    key: 'ip',
    align: 'left',
    width: -1,
  },
  {
    title: '版本号',
    key: 'version',
    align: 'left',
    width: -1,
  },
  {
    title: '执行代码',
    key: 'code',
    align: 'left',
    width: -1,
  },
  {
    title: '浏览量',
    key: 'views',
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
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];
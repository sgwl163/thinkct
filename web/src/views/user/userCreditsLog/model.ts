import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // 流水ID
  public userId = 0; // 用户ID
  public creditType = ''; // 资产类型:money=余额,score=积分
  public beforeNum = 0; // 变动前数量
  public num = 0; // 变动数量
  public afterNum = 0; // 变动后数量
  public remark = ''; // 备注
  public mapId = 0; // 关联ID
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
  creditType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入资产类型:money=余额,score=积分',
  },
  beforeNum: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入变动前数量',
  },
  num: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入变动数量',
  },
  afterNum: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入变动后数量',
  },
  remark: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入备注',
  },
  mapId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入关联ID',
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
    label: '流水ID',
    componentProps: {
      placeholder: '请输入流水ID',
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
    title: '流水ID',
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
    title: '资产类型:money=余额,score=积分',
    key: 'creditType',
    align: 'left',
    width: -1,
  },
  {
    title: '变动前数量',
    key: 'beforeNum',
    align: 'left',
    width: -1,
  },
  {
    title: '变动数量',
    key: 'num',
    align: 'left',
    width: -1,
  },
  {
    title: '变动后数量',
    key: 'afterNum',
    align: 'left',
    width: -1,
  },
  {
    title: '备注',
    key: 'remark',
    align: 'left',
    width: -1,
  },
  {
    title: '关联ID',
    key: 'mapId',
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
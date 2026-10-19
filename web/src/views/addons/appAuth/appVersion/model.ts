import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public productId = 0; // 产品ID
  public cover = ''; // 封面
  public title = ''; // 标题
  public content = ''; // 详情
  public version = ''; // 版本号
  public file = ''; // 文件
  public status = 1; // 状态:1=启用,0=禁用
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
  productId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入产品ID',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态:1=启用,0=禁用',
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
    field: 'status',
    component: 'NSelect',
    label: '状态:1=启用,0=禁用',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态:1=启用,0=禁用',
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
    title: '封面',
    key: 'cover',
    align: 'left',
    width: -1,
  },
  {
    title: '标题',
    key: 'title',
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
    title: '文件',
    key: 'file',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:1=启用,0=禁用',
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
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
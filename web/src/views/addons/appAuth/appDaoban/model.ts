import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';
import { productOptions, renderProductTag } from '../appAuth/model';

export { productOptions };

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public productId: number | null = null; // 产品ID
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
  productId: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    min: 1,
    message: '请选择所属产品',
  },
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
      min: 1,
      precision: 0,
      clearable: true,
      showButton: false,
    },
  },
  {
    field: 'productId',
    component: 'NSelect',
    label: '所属产品',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择所属产品',
      filterable: true,
      clearable: true,
      options: productOptions,
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
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '所属产品',
    key: 'productId',
    align: 'left',
    width: 160,
    render(row: State) {
      return renderProductTag(row.productId);
    },
  },
  {
    title: '域名',
    key: 'host',
    align: 'left',
    width: 180,
  },
  {
    title: 'IP',
    key: 'ip',
    align: 'left',
    width: 150,
  },
  {
    title: '版本号',
    key: 'version',
    align: 'left',
    width: 120,
  },
  {
    title: '执行代码',
    key: 'code',
    align: 'left',
    width: 180,
  },
  {
    title: '浏览量',
    key: 'views',
    align: 'left',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: 180,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['AppAuthProductOptions']);
}

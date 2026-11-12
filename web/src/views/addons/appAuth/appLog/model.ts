import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';
import {
  productOptions,
  userOptions,
  renderProductTag,
  renderUserTag,
} from '../appAuth/model';

export { productOptions, userOptions };

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public userId: number | null = null; // 用户ID
  public productId: number | null = null; // 产品ID
  public host = ''; // 域名
  public content = ''; // 内容
  public ip = ''; // IP
  public version = '0.0.0'; // 版本号
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
  userId: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    min: 1,
    message: '请选择所属用户',
  },
  productId: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    min: 1,
    message: '请选择所属产品',
  },
  host: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入域名',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入内容',
  },
  ip: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入IP',
  },
  version: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入版本号',
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
    field: 'userId',
    component: 'NSelect',
    label: '所属用户',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择用户，可搜索昵称',
      filterable: true,
      clearable: true,
      options: userOptions,
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
    title: '所属用户',
    key: 'userId',
    align: 'left',
    width: 180,
    render(row: State) {
      return renderUserTag(row.userId);
    },
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
  dict.loadOptions(['AppAuthProductOptions', 'AppAuthUserOptions']);
}

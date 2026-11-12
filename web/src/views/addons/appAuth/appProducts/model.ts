import { computed, ref } from 'vue';
import type { FormRules } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import type { FormSchema } from '@/components/Form';
import type { BasicColumn } from '@/components/Table';
import { defRangeShortcuts, formatToDateTime } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { renderImage, renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public cover = ''; // 封面
  public title = ''; // 标题
  public keywords = ''; // 关键词
  public description = ''; // 描述
  public name = ''; // 产品名称
  public alias = ''; // 别名
  public content = ''; // 内容
  public price = 0; // 价格
  public ysweb = ''; // 演示站
  public unique = ''; // 特点
  public weigh = 0; // 权重
  public status = 1; // 状态，使用系统启用/禁用字典
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

export const statusOptions = computed(() => dict.getOptionUnRef('sys_normal_disable'));

export function formatPrice(value: number | null | undefined): string {
  return typeof value === 'number' && Number.isFinite(value) ? value.toFixed(2) : '-';
}

export function formatDateTime(value: string): string {
  return value ? formatToDateTime(value) : '-';
}

export const rules: FormRules = {
  name: {
    required: true,
    whitespace: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入产品名称',
  },
  cover: {
    required: true,
    whitespace: true,
    trigger: ['blur', 'input', 'change'],
    type: 'string',
    message: '请上传封面或输入图片地址',
  },
  price: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    // 通用金额校验将数字0视为空，产品允许价格为0，与后端保持一致。
    validator: (rule, value, callback) =>
      validate.amount(rule, value === 0 ? '0' : value, callback),
  },
  unique: {
    required: true,
    whitespace: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入特点',
  },
  status: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    message: '请选择状态',
  },
};

// 搜索条件由列表和导出接口共同处理。
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
    field: 'name',
    component: 'NInput',
    label: '产品名称',
    componentProps: {
      placeholder: '请输入产品名称',
      clearable: true,
    },
  },
  {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      clearable: true,
    },
  },
  {
    field: 'alias',
    component: 'NInput',
    label: '别名',
    componentProps: {
      placeholder: '请输入别名',
      clearable: true,
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      clearable: true,
      options: statusOptions,
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

// 列表展示常用字段，关键词、描述、特点和完整内容在详情中保留。
export const columns: BasicColumn[] = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '封面',
    key: 'cover',
    align: 'left',
    width: 80,
    render(row: State) {
      return row.cover ? renderImage(row.cover) : '-';
    },
  },
  {
    title: '产品名称',
    key: 'name',
    align: 'left',
    width: 180,
    ellipsis: { tooltip: true },
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 180,
    ellipsis: { tooltip: true },
  },
  {
    title: '别名',
    key: 'alias',
    align: 'left',
    width: 140,
    ellipsis: { tooltip: true },
  },
  {
    title: '价格',
    key: 'price',
    align: 'right',
    width: 120,
    render(row: State) {
      return formatPrice(row.price);
    },
  },
  {
    title: '演示站',
    key: 'ysweb',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
    render(row: State) {
      return row.ysweb || '-';
    },
  },
  {
    title: '权重',
    key: 'weigh',
    align: 'left',
    width: 90,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 90,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
    render(row: State) {
      return formatDateTime(row.createdAt);
    },
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: 180,
    render(row: State) {
      return formatDateTime(row.updatedAt);
    },
  },
];

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

import { computed, h, ref } from 'vue';
import { NTag } from 'naive-ui';
import type { FormRules } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import type { FormSchema } from '@/components/Form';
import type { BasicColumn } from '@/components/Table';
import { defRangeShortcuts, formatToDateTime } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public userId: number | null = null; // 用户ID
  public productId: number | null = null; // 产品ID
  public name = ''; // 授权名称
  public host = ''; // 授权域名
  public ip = '127.0.0.1'; // IP
  public qq = ''; // QQ
  public phone = ''; // 联系电话
  public syskey = ''; // 授权码
  public versionId = '1.0.0'; // 版本号（字符串，不是版本记录ID）
  public authMode = '0'; // 验证方式
  public authDomain = '0'; // 域名范围
  public views = 1; // 浏览量，仅在详情展示
  public status = 1; // 状态，使用系统启用/禁用字典
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public expiredAt = ''; // 到期时间

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

export const productOptions = computed(() => dict.getOptionUnRef('AppAuthProductOptions'));
export const userOptions = computed(() => dict.getOptionUnRef('AppAuthUserOptions'));
export const authModeOptions = computed(() => dict.getOptionUnRef('AppAuthModeOptions'));
export const authDomainOptions = computed(() => dict.getOptionUnRef('AppAuthDomainOptions'));
export const statusOptions = computed(() => dict.getOptionUnRef('sys_normal_disable'));

export function requiresIp(authMode: string): boolean {
  return dict.getExtra('AppAuthModeOptions', authMode)?.requireIp === true;
}

export function getProductLabel(productId: number | null): string {
  return (
    dict.getLabel('AppAuthProductOptions', productId) || (productId ? `产品 #${productId}` : '-')
  );
}

export function getUserLabel(userId: number | null): string {
  return dict.getLabel('AppAuthUserOptions', userId) || '-';
}

// 名称、域名保留文本显示，产品、用户、分类和版本在列表与详情中使用标签。
export function renderProductTag(productId: number | null) {
  if (!productId) return '-';
  return h(
    NTag,
    { type: dict.getType('AppAuthProductOptions', productId), bordered: false },
    { default: () => getProductLabel(productId) }
  );
}

export function renderUserTag(userId: number | null) {
  if (!userId) return '-';
  return h(
    NTag,
    { type: dict.getType('AppAuthUserOptions', userId), bordered: false },
    { default: () => getUserLabel(userId) }
  );
}

export function renderVersionTag(version: string) {
  if (!version) return '-';
  return h(NTag, { type: 'default', bordered: false }, { default: () => version });
}

export function formatDateTime(value: string): string {
  return value ? formatToDateTime(value) : '-';
}

// 保持字段与后端校验一致，关联ID不能以0提交。
export const rules: FormRules = {
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
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入授权名称',
  },
  qq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.qq,
  },
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入联系电话',
  },
  versionId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入版本号',
  },
  authMode: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'string',
    message: '请选择验证方式',
  },
  authDomain: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'string',
    message: '请选择域名范围',
  },
  status: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    message: '请选择状态',
  },
};

// 搜索条件与后端列表、导出接口保持一致，选择项统一使用后端字典。
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
    label: '授权名称',
    componentProps: {
      placeholder: '请输入授权名称',
      clearable: true,
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
    field: 'host',
    component: 'NInput',
    label: '授权域名',
    componentProps: {
      placeholder: '请输入授权域名',
      clearable: true,
    },
  },
  {
    field: 'authMode',
    component: 'NSelect',
    label: '验证方式',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择验证方式',
      clearable: true,
      options: authModeOptions,
    },
  },
  {
    field: 'authDomain',
    component: 'NSelect',
    label: '域名范围',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择域名范围',
      clearable: true,
      options: authDomainOptions,
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
  {
    field: 'expiredAt',
    component: 'NDatePicker',
    label: '到期时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
    },
  },
]);

// 列表保留授权识别、验证策略、创建/更新及到期时间，敏感和辅助字段放到详情。
export const columns: BasicColumn[] = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '授权名称',
    key: 'name',
    align: 'left',
    width: 160,
    ellipsis: { tooltip: true },
  },
  {
    title: '所属产品',
    key: 'productId',
    align: 'left',
    width: 160,
    ellipsis: { tooltip: true },
    render(row: State) {
      return renderProductTag(row.productId);
    },
  },
  {
    title: '所属用户',
    key: 'userId',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
    render(row: State) {
      return renderUserTag(row.userId);
    },
  },
  {
    title: '授权域名',
    key: 'host',
    align: 'left',
    width: 180,
    ellipsis: { tooltip: true },
    render(row: State) {
      return row.host || '-';
    },
  },
  {
    title: '版本号',
    key: 'versionId',
    align: 'left',
    width: 100,
    render(row: State) {
      return renderVersionTag(row.versionId);
    },
  },
  {
    title: '验证方式',
    key: 'authMode',
    align: 'left',
    width: 130,
    render(row: State) {
      return renderOptionTag('AppAuthModeOptions', row.authMode);
    },
  },
  {
    title: '域名范围',
    key: 'authDomain',
    align: 'left',
    width: 130,
    render(row: State) {
      return renderOptionTag('AppAuthDomainOptions', row.authDomain);
    },
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
  {
    title: '到期时间',
    key: 'expiredAt',
    align: 'left',
    width: 180,
    render(row: State) {
      return formatDateTime(row.expiredAt);
    },
  },
];

export function loadOptions() {
  dict.loadOptions([
    'sys_normal_disable',
    'AppAuthProductOptions',
    'AppAuthUserOptions',
    'AppAuthModeOptions',
    'AppAuthDomainOptions',
  ]);
}

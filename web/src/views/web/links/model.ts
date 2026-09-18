import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { renderImage } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export const linksStatusOptions = [
  { label: '待审核', value: 'pending_trial', type: 'warning' },
  { label: '已启用', value: 'enable', type: 'success' },
  { label: '已禁用', value: 'disable', type: 'default' },
];

export const targetOptions = [
  { label: '新标签页', value: '_blank' },
  { label: '当前标签页', value: '_self' },
];

export function getStatusLabel(v: string) {
  return linksStatusOptions.find((o) => o.value === v)?.label ?? v;
}
export function getStatusType(v: string): any {
  return (linksStatusOptions.find((o) => o.value === v) as any)?.type ?? 'default';
}

export class State {
  public id = 0;
  public userId = 0;
  public title = '';
  public link = '';
  public target = '_blank';
  public logo = '';
  public weigh = 0;
  public status = 'pending_trial';
  public remark = '';

  constructor(state?: Partial<State>) {
    if (state) { Object.assign(this, state); }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) { return cloneDeep(state); }
    return new State(state);
  }
  return new State();
}

export const rules = {
  title: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入标题' },
  link: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入链接' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: { placeholder: '请输入标题' },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择状态', options: linksStatusOptions },
  },
]);

export const columns = [
  {
    title: 'LOGO',
    key: 'logo',
    align: 'center',
    width: 70,
    render(row: State) { return renderImage(row.logo); },
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 150,
  },
  {
    title: '链接',
    key: 'link',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
  },
  {
    title: '跳转方式',
    key: 'target',
    align: 'center',
    width: 90,
    render(row: State) {
      const label = targetOptions.find((o) => o.value === row.target)?.label ?? row.target;
      return h(NTag, { size: 'small' }, { default: () => label });
    },
  },
  {
    title: '权重',
    key: 'weigh',
    align: 'center',
    width: 70,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 90,
    render(row: State) {
      return h(NTag, { type: getStatusType(row.status), size: 'small' }, { default: () => getStatusLabel(row.status) });
    },
  },
];

export function loadOptions() {
  dict.loadOptions([]);
}

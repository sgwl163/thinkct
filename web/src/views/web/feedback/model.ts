import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, timestampToTime } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export const feedbackStatusOptions = [
  { label: '待处理', value: 0, type: 'warning' },
  { label: '已回复', value: 1, type: 'success' },
  { label: '已关闭', value: 2, type: 'error' },
];

export function getStatusLabel(v: number) {
  return feedbackStatusOptions.find((o) => o.value === v)?.label ?? String(v);
}
export function getStatusType(v: number): any {
  return (feedbackStatusOptions.find((o) => o.value === v) as any)?.type ?? 'default';
}

export class State {
  public id = 0;
  public name = '';
  public email = '';
  public phone = '';
  public company = '';
  public content = '';
  public answer = '';
  public ip = '';
  public status = 0;
  public createTime = 0;
  public updateTime = 0;

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

export const rules = {
  name: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入姓名' },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入反馈内容',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '联系人',
    componentProps: { placeholder: '请输入联系人姓名' },
  },
  {
    field: 'phone',
    component: 'NInput',
    label: '手机号',
    componentProps: { placeholder: '请输入手机号' },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '处理状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择处理状态', options: feedbackStatusOptions },
  },
  {
    field: 'createTime',
    component: 'NDatePicker',
    label: '提交时间',
    componentProps: { type: 'datetimerange', clearable: true, shortcuts: defRangeShortcuts() },
  },
]);

export const columns = [
  {
    title: '联系人',
    key: 'name',
    align: 'left',
    width: 130,
    ellipsis: { tooltip: true },
  },
  {
    title: '联系方式',
    key: 'phone',
    align: 'left',
    width: 140,
    ellipsis: { tooltip: true },
  },
  {
    title: '公司',
    key: 'company',
    align: 'left',
    width: 160,
    ellipsis: { tooltip: true },
  },
  {
    title: '反馈内容',
    key: 'content',
    align: 'left',
    width: 260,
    ellipsis: { tooltip: true },
  },
  {
    title: '回复内容',
    key: 'answer',
    align: 'left',
    width: 220,
    ellipsis: { tooltip: true },
  },
  {
    title: '处理状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(
        NTag,
        { type: getStatusType(row.status), size: 'small' },
        { default: () => getStatusLabel(row.status) }
      );
    },
  },
  {
    title: '提交时间',
    key: 'createTime',
    align: 'center',
    width: 170,
    render(row: State) {
      return row.createTime > 0 ? timestampToTime(row.createTime) : '-';
    },
  },
];

export function loadOptions() {
  dict.loadOptions([]);
}

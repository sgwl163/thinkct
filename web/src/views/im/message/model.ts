import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export const chatTypeOptions = [
  { label: '私聊', value: 1 },
  { label: '群聊', value: 2 },
];

export const msgTypeOptions = [
  { label: '文本', value: 'text' },
  { label: '图片', value: 'image' },
  { label: '语音', value: 'audio' },
  { label: '视频', value: 'video' },
  { label: '文件', value: 'file' },
];

export class State {
  public id = 0;
  public chatType = 1;
  public fromId = 0;
  public toId = 0;
  public msgType = 'text';
  public content = '';
  public status = 1;
  public createdAt = '';
  public updatedAt = '';

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
  fromId: { required: true, trigger: ['blur', 'input'], type: 'number', min: 1, message: '请输入发送者ID' },
  toId: { required: true, trigger: ['blur', 'input'], type: 'number', min: 1, message: '请输入接收者ID' },
  content: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入消息内容' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'fromId',
    component: 'NInputNumber',
    label: '发送者ID',
    componentProps: { placeholder: '请输入发送者ID' },
  },
  {
    field: 'toId',
    component: 'NInputNumber',
    label: '接收者ID',
    componentProps: { placeholder: '请输入接收者ID' },
  },
  {
    field: 'chatType',
    component: 'NSelect',
    label: '聊天类型',
    defaultValue: null,
    componentProps: { placeholder: '请选择聊天类型', options: chatTypeOptions },
  },
  {
    field: 'msgType',
    component: 'NSelect',
    label: '消息类型',
    defaultValue: null,
    componentProps: { placeholder: '请选择消息类型', options: msgTypeOptions },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择状态', options: dict.getOption('sys_normal_disable') },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '发送时间',
    componentProps: { type: 'datetimerange', clearable: true, shortcuts: defRangeShortcuts() },
  },
]);

export const columns = [
  {
    title: '发送者',
    key: 'fromId',
    align: 'center',
    width: 80,
  },
  {
    title: '接收者',
    key: 'toId',
    align: 'center',
    width: 80,
  },
  {
    title: '类型',
    key: 'chatType',
    align: 'center',
    width: 70,
    render(row: State) {
      const label = chatTypeOptions.find((o) => o.value === row.chatType)?.label ?? String(row.chatType);
      return h(NTag, { size: 'small' }, { default: () => label });
    },
  },
  {
    title: '消息类型',
    key: 'msgType',
    align: 'center',
    width: 80,
    render(row: State) {
      const label = msgTypeOptions.find((o) => o.value === row.msgType)?.label ?? row.msgType;
      return h(NTag, { size: 'small' }, { default: () => label });
    },
  },
  {
    title: '消息内容',
    key: 'content',
    align: 'left',
    width: 260,
    ellipsis: { tooltip: true },
  },
  {
    title: '发送时间',
    key: 'createdAt',
    align: 'left',
    width: 160,
  },
];

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

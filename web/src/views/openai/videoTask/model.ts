import { h, ref, computed } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

// 任务状态
export const taskStatusOptions = [
  { label: '进行中', value: 0, type: 'info' },
  { label: '成功', value: 1, type: 'success' },
  { label: '失败', value: 2, type: 'error' },
];

export function getStatusLabel(v: number) {
  return taskStatusOptions.find((o) => o.value === v)?.label ?? String(v);
}
export function getStatusType(v: number): any {
  return taskStatusOptions.find((o) => o.value === v)?.type ?? 'default';
}

// 视频模型选项
export const videoModelOptions = [
  { label: 'Seedance 1.5 Pro（全能）', value: 'doubao-seedance-1-5-pro-251215' },
  { label: 'Seedance 1.0 Pro（全能）', value: 'doubao-seedance-1-0-pro-250528' },
  { label: 'Seedance 1.0 Pro Fast（快速）', value: 'doubao-seedance-1-0-pro-fast-251015' },
  { label: 'Seedance 1.0 Lite 文生视频', value: 'doubao-seedance-1-0-lite-t2v-250428' },
  { label: 'Seedance 1.0 Lite 首尾帧', value: 'doubao-seedance-1-0-lite-i2v-250428' },
];

export class State {
  public id = 0;
  public userId = 0;
  public prompt = '';
  public referenceImage = '';
  public lastFrame = '';
  public voiceDubUrl = '';
  public bgMusicUrl = '';
  public generateAudio = 1;
  public model = '';
  public videoUrl = '';
  public status = 0;
  public failReason = '';
  public remoteTaskId = '';
  public createdAt = '';
  public updatedAt = '';

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
    min: 1,
    message: '请输入用户ID',
  },
  prompt: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入视频提示词',
  },
  model: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'string',
    message: '请选择使用的模型',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'userId',
    component: 'NInputNumber',
    label: '用户ID',
    componentProps: { placeholder: '请输入用户ID' },
  },
  {
    field: 'model',
    component: 'NSelect',
    label: '模型',
    componentProps: {
      placeholder: '请选择模型',
      options: videoModelOptions,
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
      options: taskStatusOptions,
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
    title: '用户ID',
    key: 'userId',
    align: 'center',
    width: 80,
  },
  {
    title: '参考图',
    key: 'referenceImage',
    align: 'center',
    width: 70,
    render(row: State) {
      return renderImage(row.referenceImage);
    },
  },
  {
    title: '模型',
    key: 'model',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
    render(row: State) {
      return videoModelOptions.find((o) => o.value === row.model)?.label ?? row.model;
    },
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
  {
    title: '失败原因',
    key: 'failReason',
    align: 'left',
    width: 180,
    ellipsis: { tooltip: true },
    render(row: State) {
      return row.failReason || '-';
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 160,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions([]);
}

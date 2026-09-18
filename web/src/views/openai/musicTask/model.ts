import { h, ref } from 'vue';
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

// 创作模式
export const musicTypeOptions = [
  { label: '灵感模式', value: 1 },
  { label: '定制模式', value: 2 },
];

export class State {
  public id = 0; // 任务ID
  public userId = 0; // 用户ID
  public prompt = ''; // 提示词(灵感或歌词)
  public type = 1; // 创作模式: 1=灵感模式, 2=定制模式
  public makeInstrumental = 0; // 纯音乐: 0=否 1=是
  public title = ''; // 歌名
  public tags = ''; // 风格标签
  public model = ''; // 使用的模型(mv)
  public audioUrl = ''; // 生成的音频URL
  public imageUrl = ''; // 封面图URL
  public videoUrl = ''; // 生成的视频URL
  public duration = 0; // 时长(秒)
  public status = 0; // 状态: 0=进行中, 1=成功, 2=失败
  public failReason = ''; // 失败原因
  public remoteTaskId = ''; // 远端任务ID
  public remoteClipId = ''; // 远端歌曲ID
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
    trigger: ['blur', 'input'],
    type: 'number',
    min: 1,
    message: '请输入用户ID',
  },
  prompt: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入提示词',
  },
  model: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入使用的模型',
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
    field: 'title',
    component: 'NInput',
    label: '歌名',
    componentProps: { placeholder: '请输入歌名' },
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
    title: '封面',
    key: 'imageUrl',
    align: 'center',
    width: 70,
    render(row: State) {
      return renderImage(row.imageUrl);
    },
  },
  {
    title: '歌名',
    key: 'title',
    align: 'left',
    width: 150,
    ellipsis: { tooltip: true },
  },
  {
    title: '创作模式',
    key: 'type',
    align: 'center',
    width: 90,
    render(row: State) {
      const label = musicTypeOptions.find((o) => o.value === row.type)?.label ?? String(row.type);
      return h(NTag, { size: 'small' }, { default: () => label });
    },
  },
  {
    title: '纯音乐',
    key: 'makeInstrumental',
    align: 'center',
    width: 75,
    render(row: State) {
      return h(NTag, { type: row.makeInstrumental === 1 ? 'success' : 'default', size: 'small' }, { default: () => (row.makeInstrumental === 1 ? '是' : '否') });
    },
  },
  {
    title: '模型',
    key: 'model',
    align: 'left',
    width: 120,
  },
  {
    title: '时长(s)',
    key: 'duration',
    align: 'center',
    width: 80,
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

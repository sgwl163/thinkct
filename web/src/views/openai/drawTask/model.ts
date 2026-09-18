import { h, ref } from 'vue';
import { NTag, NProgress } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

// 绘画任务状态选项
export const drawTaskStatusOptions = [
  { label: '进行中', value: 0, listClass: 'info' },
  { label: '成功', value: 1, listClass: 'success' },
  { label: '失败', value: 2, listClass: 'error' },
];

export function getStatusLabel(status: number) {
  return drawTaskStatusOptions.find((o) => o.value === status)?.label ?? String(status);
}

export function getStatusType(status: number) {
  return (drawTaskStatusOptions.find((o) => o.value === status)?.listClass as any) ?? 'default';
}

export class State {
  public id = 0; // 任务ID
  public userId = 0; // 用户ID
  public prompt = ''; // 绘画提示词
  public negativePrompt = ''; // 负面提示词
  public referenceImage = ''; // 参考图URL
  public model = ''; // 使用的模型
  public version = ''; // 模型版本
  public ratio = '1:1'; // 图片比例
  public raw = 0; // 是否开启RAW模式(0否,1是)
  public imageUrl = ''; // 生成的图片URL
  public status = 0; // 状态: 0=进行中, 1=成功, 2=失败
  public failReason = ''; // 失败原因
  public stylize = 0; // 风格化程度(0-1000)
  public chaos = 0; // 多样性(0-100)
  public seed = 0; // 随机种子
  public mjTaskId = ''; // MJ远端任务ID
  public buttons = ''; // MJ操作按钮JSON
  public progress = 0; // 进度(0-100)
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
  prompt: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入绘画提示词',
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
    componentProps: {
      placeholder: '请输入用户ID',
    },
  },
  {
    field: 'model',
    component: 'NInput',
    label: '模型',
    componentProps: {
      placeholder: '请输入模型名称',
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: drawTaskStatusOptions,
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
    title: '任务ID',
    key: 'id',
    align: 'center',
    width: 80,
  },
  {
    title: '用户ID',
    key: 'userId',
    align: 'center',
    width: 80,
  },
  {
    title: '结果图',
    key: 'imageUrl',
    align: 'center',
    width: 70,
    render(row: State) {
      return renderImage(row.imageUrl);
    },
  },
  {
    title: '模型',
    key: 'model',
    align: 'left',
    width: 140,
  },
  {
    title: '比例',
    key: 'ratio',
    align: 'center',
    width: 70,
  },
  {
    title: '进度',
    key: 'progress',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NProgress, {
        type: 'line',
        percentage: row.progress,
        status: row.status === 2 ? 'error' : row.status === 1 ? 'success' : 'default',
        showIndicator: true,
        style: 'width: 80px',
      });
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

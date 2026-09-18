import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public memberId = 0; // 用户ID
  public folderId = 0; // 文件夹ID
  public title = ''; // 会话标题
  public prompt = ''; // 系统提示词
  public model = ''; // 使用的模型
  public status = 1; // 状态
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
  memberId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    min: 1,
    message: '请输入用户ID',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入会话标题',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'memberId',
    component: 'NInputNumber',
    label: '用户ID',
    componentProps: {
      placeholder: '请输入用户ID',
    },
  },
  {
    field: 'title',
    component: 'NInput',
    label: '会话标题',
    componentProps: {
      placeholder: '请输入会话标题',
    },
  },
  {
    field: 'model',
    component: 'NInput',
    label: '模型',
    componentProps: {
      placeholder: '请输入模型',
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: dict.getOption('sys_normal_disable'),
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
    key: 'memberId',
    align: 'center',
    width: 80,
  },
  {
    title: '会话标题',
    key: 'title',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
  },
  {
    title: '使用的模型',
    key: 'model',
    align: 'left',
    width: 160,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 90,
    render(row: State) {
      const type = row.status === 1 ? 'success' : 'default';
      const label = dict.getLabel('sys_normal_disable', row.status) || (row.status === 1 ? '启用' : '禁用');
      return h(NTag, { type, size: 'small' }, { default: () => label });
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
  dict.loadOptions(['sys_normal_disable']);
}

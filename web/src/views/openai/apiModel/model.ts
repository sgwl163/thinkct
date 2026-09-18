import { h, ref } from 'vue';
import { NSwitch, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Status } from '@/api/openai/apiModel';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0; // ID
  public groupId = 0; // 分组ID
  public name = ''; // 模型别名
  public model = ''; // 模型标识
  public isVision = 0; // 是否支持识图(0否,1是)
  public isFile = 0; // 是否支持上传文件(0否,1是)
  public isSearch = 0; // 是否支持联网搜索(0否,1是)
  public status = 1; // 状态
  public sort = 0; // 排序
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
  groupId: {
    required: true,
    trigger: ['blur', 'change'],
    type: 'number',
    min: 1,
    message: '请选择所属分组',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入模型别名',
  },
  model: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入模型标识',
  },
};

function boolTag(val: number) {
  return h(NTag, { type: val === 1 ? 'success' : 'default', size: 'small' }, { default: () => (val === 1 ? '支持' : '不支持') });
}

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '模型别名',
    componentProps: {
      placeholder: '请输入模型别名',
    },
  },
  {
    field: 'model',
    component: 'NInput',
    label: '模型标识',
    componentProps: {
      placeholder: '请输入模型标识',
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
    title: '模型别名',
    key: 'name',
    align: 'left',
    width: 140,
  },
  {
    title: '模型标识',
    key: 'model',
    align: 'left',
    width: 180,
    ellipsis: { tooltip: true },
  },
  {
    title: '识图',
    key: 'isVision',
    align: 'center',
    width: 70,
    render(row: State) { return boolTag(row.isVision); },
  },
  {
    title: '文件',
    key: 'isFile',
    align: 'center',
    width: 70,
    render(row: State) { return boolTag(row.isFile); },
  },
  {
    title: '联网',
    key: 'isSearch',
    align: 'center',
    width: 70,
    render(row: State) { return boolTag(row.isSearch); },
  },
  {
    title: '排序',
    key: 'sort',
    align: 'center',
    width: 70,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        disabled: !hasPermission(['/apiModel/status']),
        onUpdateValue: (val: boolean) => {
          row.status = val ? 1 : 2;
          Status({ id: row.id, status: row.status }).then(() => {
            $message.success('操作成功');
          });
        },
      });
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

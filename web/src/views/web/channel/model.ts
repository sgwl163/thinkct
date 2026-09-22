import { h, ref } from 'vue';
import { NSwitch, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Status, TreeOption } from '@/api/web/channel';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export const channelTypeOptions = [
  { label: '封面频道', value: 'cover' },
  { label: '列表频道', value: 'list' },
  { label: '跳转链接', value: 'link' },
];

export const targetOptions = [
  { label: '当前标签页', value: '_self' },
  { label: '新标签页', value: '_blank' },
  { label: '当前窗体', value: '_top' },
  { label: '父窗口', value: '_parent' },
];

export const visitGroupOptions = [
  { label: '全部', value: 'all' },
  { label: '仅会员', value: 'user' },
];

export const modelOptions = [
  { label: '文章模型', value: 'article' },
  { label: '产品模型', value: 'products' },
];

export class State {
  public id = 0;
  public pid = 0;
  public level = 0;
  public tree = '';
  public model = '';
  public type = 'list';
  public name = '';
  public diyname = '';
  public template = '';
  public url = '';
  public target = '_self';
  public seotitle = '';
  public keywords = '';
  public description = '';
  public allowVisitGroups = 'all';
  public weigh = 0;
  public status = 1;

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
  name: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入频道名称' },
  type: { required: true, trigger: ['blur', 'change'], type: 'string', message: '请选择类型' },
  model: { required: true, trigger: ['blur', 'change'], type: 'string', message: '请选择模型' },
  template: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入模板地址',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '频道名称',
    componentProps: { placeholder: '请输入频道名称' },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择状态', options: dict.getOption('sys_normal_disable') },
  },
]);

export const columns = [
  {
    title: '栏目名称',
    key: 'name',
    align: 'left',
    width: 220,
    ellipsis: { tooltip: true },
  },
  {
    title: '自定义名称',
    key: 'diyname',
    align: 'left',
    width: 160,
    ellipsis: { tooltip: true },
  },
  {
    title: '模型',
    key: 'model',
    align: 'center',
    width: 110,
    render(row: State) {
      const label = modelOptions.find((option) => option.value === row.model)?.label ?? row.model;
      return h(NTag, { size: 'small', type: 'info' }, { default: () => label });
    },
  },
  {
    title: '类型',
    key: 'type',
    align: 'center',
    width: 110,
    render(row: State) {
      const label =
        channelTypeOptions.find((option) => option.value === row.type)?.label ?? row.type;
      const type = row.type === 'cover' ? 'success' : row.type === 'link' ? 'warning' : 'info';
      return h(NTag, { size: 'small', type }, { default: () => label });
    },
  },
  {
    title: '访问策略',
    key: 'allowVisitGroups',
    align: 'center',
    width: 100,
    render(row: State) {
      const label =
        visitGroupOptions.find((option) => option.value === row.allowVisitGroups)?.label ??
        row.allowVisitGroups;
      return h(
        NTag,
        { size: 'small', type: row.allowVisitGroups === 'user' ? 'warning' : 'info' },
        { default: () => label }
      );
    },
  },
  {
    title: '权重',
    key: 'weigh',
    align: 'center',
    width: 80,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        disabled: !hasPermission(['/channel/status']),
        onUpdateValue: (val: boolean) => {
          row.status = val ? 1 : 2;
          Status({ id: row.id, status: row.status }).then(() => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
];

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

export const treeOption = ref([]);

export function loadTreeOption() {
  TreeOption().then((res) => {
    treeOption.value = res;
  });
}

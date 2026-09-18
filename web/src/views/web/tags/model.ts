import { h, ref } from 'vue';
import { NSwitch, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Status } from '@/api/web/tags';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export const tagTypeOptions = [
  { label: '默认', value: 'default' },
  { label: '成功', value: 'success' },
  { label: '信息', value: 'info' },
  { label: '警告', value: 'warning' },
  { label: '危险', value: 'danger' },
];

export class State {
  public id = 0;
  public name = '';
  public type = 'default';
  public documentCount = 0;
  public seotitle = '';
  public keywords = '';
  public description = '';
  public innerChain = 0;
  public weigh = 0;
  public status = 1;

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
  name: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入标签名称' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '名称',
    componentProps: { placeholder: '请输入名称' },
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
    title: '名称',
    key: 'name',
    align: 'left',
    width: 160,
  },
  {
    title: '类型',
    key: 'type',
    align: 'center',
    width: 80,
    render(row: State) {
      const label = tagTypeOptions.find((o) => o.value === row.type)?.label ?? row.type;
      const typeMap: Record<string, any> = { success: 'success', info: 'info', warning: 'warning', danger: 'error' };
      return h(NTag, { type: typeMap[row.type] ?? 'default', size: 'small' }, { default: () => label });
    },
  },
  {
    title: '文档数',
    key: 'documentCount',
    align: 'center',
    width: 80,
  },
  {
    title: '自动内链',
    key: 'innerChain',
    align: 'center',
    width: 90,
    render(row: State) {
      return h(NTag, { type: row.innerChain === 1 ? 'success' : 'default', size: 'small' }, { default: () => (row.innerChain === 1 ? '开启' : '关闭') });
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
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        disabled: !hasPermission(['/tags/status']),
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

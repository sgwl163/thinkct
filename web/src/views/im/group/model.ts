import { h, ref } from 'vue';
import { NSwitch, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage } from '@/utils';
import { Status } from '@/api/im/group';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0;
  public name = '';
  public avatar = '';
  public ownerId = 0;
  public notice = '';
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
  name: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入群名称' },
  ownerId: { required: true, trigger: ['blur', 'input'], type: 'number', min: 1, message: '请输入群主ID' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '群名称',
    componentProps: { placeholder: '请输入群名称' },
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
    label: '创建时间',
    componentProps: { type: 'datetimerange', clearable: true, shortcuts: defRangeShortcuts() },
  },
]);

export const columns = [
  {
    title: '群头像',
    key: 'avatar',
    align: 'center',
    width: 70,
    render(row: State) { return renderImage(row.avatar); },
  },
  {
    title: '群名称',
    key: 'name',
    align: 'left',
    width: 160,
  },
  {
    title: '群主ID',
    key: 'ownerId',
    align: 'center',
    width: 80,
  },
  {
    title: '群公告',
    key: 'notice',
    align: 'left',
    width: 200,
    ellipsis: { tooltip: true },
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 1,
        disabled: !hasPermission(['/group/status']),
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

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

import { h, ref } from 'vue';
import { NSwitch } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDateTime } from '@/utils/dateUtil';
import { renderImage, renderOptionTag } from '@/utils';
import { validate } from '@/utils/validateUtil';
import { Status } from '@/api/user/user';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export const userStatusOptions = [
  { label: '启用', value: 'enable', type: 'success' },
  { label: '禁用', value: 'disable', type: 'default' },
];

export const genderOptions = [
  { label: '未知', value: 0 },
  { label: '男', value: 1 },
  { label: '女', value: 2 },
];

export class State {
  public id = 0;
  public groupId = 0;
  public username = '';
  public nickname = '';
  public email = '';
  public mobile = '';
  public avatar = '';
  public gender = 1;
  public birthday = '';
  public money = 0;
  public score = 0;
  public createdAt = '';
  public lastLoginAt = '';
  public lastLoginIp = '';
  public loginFailure = 0;
  public joinIp = '';
  public joinedAt = '';
  public motto = '';
  public password = '';
  public status = 'enable';

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
  username: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入用户名',
  },
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
  mobile: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.phone,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'username',
    component: 'NInput',
    label: '用户名',
    componentProps: { placeholder: '请输入用户名' },
  },
  {
    field: 'mobile',
    component: 'NInput',
    label: '手机号',
    componentProps: { placeholder: '请输入手机号' },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择状态', options: userStatusOptions },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '注册时间',
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
    key: 'id',
    align: 'center',
    width: 80,
  },
  {
    title: '头像',
    key: 'avatar',
    align: 'center',
    width: 65,
    render(row: State) {
      return renderImage(row.avatar);
    },
  },
  {
    title: '用户名',
    key: 'username',
    align: 'left',
    width: 130,
  },
  {
    title: '昵称',
    key: 'nickname',
    align: 'left',
    width: 120,
  },
  {
    title: '手机',
    key: 'mobile',
    align: 'left',
    width: 120,
  },
  {
    title: '邮箱',
    key: 'email',
    align: 'left',
    width: 160,
    ellipsis: { tooltip: true },
  },
  {
    title: '性别',
    key: 'gender',
    align: 'center',
    width: 70,
    render(row: State) {
      return renderOptionTag('sys_user_sex', row.gender);
    },
  },
  {
    title: '余额',
    key: 'money',
    align: 'right',
    width: 90,
    render(row: State) {
      return row.money > 0 ? `¥${row.money}` : '¥0';
    },
  },
  {
    title: '积分',
    key: 'score',
    align: 'center',
    width: 80,
  },
  {
    title: '注册时间',
    key: 'createdAt',
    align: 'left',
    width: 160,
    render(row: State) {
      return row.createdAt ? formatToDateTime(row.createdAt) : '-';
    },
  },
  {
    title: '上次登录IP',
    key: 'lastLoginIp',
    align: 'left',
    width: 130,
  },
  {
    title: '上次登录时间',
    key: 'lastLoginAt',
    align: 'left',
    width: 160,
    render(row: State) {
      return row.lastLoginAt ? formatToDateTime(row.lastLoginAt) : '-';
    },
  },
  {
    title: '登录失败次数',
    key: 'loginFailure',
    align: 'center',
    width: 120,
  },
  {
    title: '状态',
    key: 'status',
    align: 'center',
    width: 100,
    render(row: State) {
      return h(NSwitch, {
        value: row.status === 'enable',
        disabled: !hasPermission(['/user/status']),
        onUpdateValue: (val: boolean) => {
          row.status = val ? 'enable' : 'disable';
          Status({ id: row.id, status: row.status }).then(() => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_user_sex']);
}

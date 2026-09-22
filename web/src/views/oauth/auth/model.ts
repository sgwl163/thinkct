import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { validate } from '@/utils/validateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // ID
  public userId = 0; // 用户ID
  public productId = 0; // 产品ID
  public name = ''; // 名称
  public host = ''; // 域名
  public ip = '127.0.0.1'; // IP
  public qq = ''; // QQ
  public phone = ''; // 电话
  public syskey = ''; // 授权码
  public versionId = '1.0.0'; // 版本
  public authMode = '0'; // 验证方式:0=单域名验证,1=IP双重验证
  public authDomain = '0'; // 验证域名:0=当前单域名方式,1=顶级泛域名方式
  public views = 1; // 浏览量
  public status = 1; // 状态:1=启用,0=禁用
  public createTime = 0; // 创建时间
  public updateTime = 0; // 修改时间
  public dueTime = 0; // 到期时间

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
    message: '请输入用户ID',
  },
  productId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入产品ID',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入名称',
  },
  qq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.qq,
  },
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入电话',
  },
  versionId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入版本',
  },
  authMode: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入验证方式:0=单域名验证,1=IP双重验证',
  },
  authDomain: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入验证域名:0=当前单域名方式,1=顶级泛域名方式',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态:1=启用,0=禁用',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态:1=启用,0=禁用',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态:1=启用,0=禁用',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '用户ID',
    key: 'userId',
    align: 'left',
    width: -1,
  },
  {
    title: '产品ID',
    key: 'productId',
    align: 'left',
    width: -1,
  },
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '域名',
    key: 'host',
    align: 'left',
    width: -1,
  },
  {
    title: 'IP',
    key: 'ip',
    align: 'left',
    width: -1,
  },
  {
    title: 'QQ',
    key: 'qq',
    align: 'left',
    width: -1,
  },
  {
    title: '电话',
    key: 'phone',
    align: 'left',
    width: -1,
  },
  {
    title: '授权码',
    key: 'syskey',
    align: 'left',
    width: -1,
  },
  {
    title: '版本',
    key: 'versionId',
    align: 'left',
    width: -1,
  },
  {
    title: '验证方式:0=单域名验证,1=IP双重验证',
    key: 'authMode',
    align: 'left',
    width: -1,
  },
  {
    title: '验证域名:0=当前单域名方式,1=顶级泛域名方式',
    key: 'authDomain',
    align: 'left',
    width: -1,
  },
  {
    title: '浏览量',
    key: 'views',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:1=启用,0=禁用',
    key: 'status',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '创建时间',
    key: 'createTime',
    align: 'left',
    width: -1,
  },
  {
    title: '修改时间',
    key: 'updateTime',
    align: 'left',
    width: -1,
  },
  {
    title: '到期时间',
    key: 'dueTime',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
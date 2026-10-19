import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // ID
  public pid = 0; // 上级菜单
  public type = 'menu'; // 类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
  public title = ''; // 标题
  public name = ''; // 规则名称
  public path = ''; // 路由路径
  public icon = ''; // 图标
  public menuType = 'tab'; // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
  public url = ''; // Url
  public component = ''; // 组件路径
  public noLoginValid = 0; // 未登录有效:0=否,1=是
  public extend = 'none'; // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
  public remark = ''; // 备注
  public weigh = 0; // 权重
  public status = '1'; // 状态:0=禁用,1=启用
  public createdAt = ''; // 更新时间
  public updatedAt = ''; // 创建时间

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
  pid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入上级菜单',
  },
  type: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入规则名称',
  },
  path: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入路由路径',
  },
  icon: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入图标',
  },
  menuType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入菜单类型:tab=选项卡,link=链接,iframe=Iframe',
  },
  url: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入Url',
  },
  component: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入组件路径',
  },
  noLoginValid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入未登录有效:0=否,1=是',
  },
  extend: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单',
  },
  remark: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入备注',
  },
  weigh: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入权重',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入状态:0=禁用,1=启用',
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
    field: 'createdAt',
    component: 'NDatePicker',
    label: '更新时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
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
    title: '上级菜单',
    key: 'pid',
    align: 'left',
    width: -1,
  },
  {
    title: '类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮',
    key: 'type',
    align: 'left',
    width: -1,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '规则名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '路由路径',
    key: 'path',
    align: 'left',
    width: -1,
  },
  {
    title: '图标',
    key: 'icon',
    align: 'left',
    width: -1,
  },
  {
    title: '菜单类型:tab=选项卡,link=链接,iframe=Iframe',
    key: 'menuType',
    align: 'left',
    width: -1,
  },
  {
    title: 'Url',
    key: 'url',
    align: 'left',
    width: -1,
  },
  {
    title: '组件路径',
    key: 'component',
    align: 'left',
    width: -1,
  },
  {
    title: '未登录有效:0=否,1=是',
    key: 'noLoginValid',
    align: 'left',
    width: -1,
  },
  {
    title: '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单',
    key: 'extend',
    align: 'left',
    width: -1,
  },
  {
    title: '备注',
    key: 'remark',
    align: 'left',
    width: -1,
  },
  {
    title: '权重',
    key: 'weigh',
    align: 'left',
    width: -1,
  },
  {
    title: '状态:0=禁用,1=启用',
    key: 'status',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];
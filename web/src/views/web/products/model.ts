import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, timestampToTime } from '@/utils/dateUtil';
import { renderImage } from '@/utils';
import { useDictStore } from '@/store/modules/dict';
import { TreeOption } from '@/api/web/channel/index';

const dict = useDictStore();

export const channelTreeOption = ref<any[]>([]);
export function loadChannelTreeOption() {
  TreeOption({ model: 'products' }).then((res) => {
    channelTreeOption.value = res;
  });
}

export const articleFlagOptions = [
  { label: '最新', value: 'new' },
  { label: '置顶', value: 'top' },
  { label: '热门', value: 'hot' },
  { label: '推荐', value: 'recommend' },
];

export const articleStatusOptions = [
  { label: '正常', value: 'normal', type: 'success' },
  { label: '待审核', value: 'unaudited', type: 'warning' },
  { label: '已拒绝', value: 'refused', type: 'error' },
  { label: '已下线', value: 'offline', type: 'default' },
];

export const targetOptions = [
  { label: '新标签页', value: '_blank' },
  { label: '当前标签页', value: '_self' },
  { label: '当前窗体', value: '_top' },
  { label: '父窗口', value: '_parent' },
];

export function getStatusLabel(v: string) {
  return articleStatusOptions.find((o) => o.value === v)?.label ?? v;
}
export function getStatusType(v: string): any {
  return (articleStatusOptions.find((o) => o.value === v) as any)?.type ?? 'default';
}

export class State {
  public id = 0;
  public userId = 0;
  public adminId = 0;
  public channelId = 0;
  public channelIds = '';
  public diyname = '';
  public title = '';
  public titleStyle = '';
  public flag = '';
  public image = '';
  public images = '';
  public content = '';
  public seotitle = '';
  public keywords = '';
  public description = '';
  public tags = '';
  public url = '';
  public target = '_blank';
  public views = 0;
  public comments = 0;
  public likes = 0;
  public dislikes = 0;
  public weigh = 0;
  public price = 0;
  public status = 'normal';
  public publishTime = 0;
  public name = '';
  public updateTime = 0;
  public createTime = 0;

  constructor(state?: Partial<State>) {
    if (state) { Object.assign(this, state); }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) { return cloneDeep(state); }
    const obj = new State(state);
    if (state.publishTime) {
      obj.publishTime = state.publishTime * 1000;
    } else {
      obj.publishTime = 0;
    }
    return obj;
  }
  return new State();
}

export const rules = {
  title: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入标题' },
  // name: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入型号名称' },
  // diyname: { required: true, trigger: ['blur', 'input'], type: 'string', message: '请输入自定义名称' },
  channelId: { required: true, trigger: ['blur', 'change'], type: 'number', min: 1, message: '请输入频道ID' },
};

export const schemas = ref<FormSchema[]>([
    {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: { placeholder: '请输入标题' },
  },
  {
    field: 'publishTime',
    component: 'NDatePicker',
    label: '发布时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      valueFormat: 'yyyy-MM-dd HH:mm:ss',
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: { placeholder: '请选择状态', options: articleStatusOptions, clearable: true },
  },
  {
    field: 'channelId',
    component: 'NTreeSelect',
    label: '栏目',
    componentProps: {
      options: channelTreeOption,
      keyField: 'id',
      labelField: 'name',
      childrenField: 'children',
      placeholder: '请选择栏目',
      clearable: true,
      defaultExpandAll: true,
      cascade: true,
      checkStrategy: 'all',
    },
  },
  {
    field: 'channelIds',
    component: 'NTreeSelect',
    label: '副栏目',
    componentProps: {
      options: channelTreeOption,
      keyField: 'id',
      labelField: 'name',
      childrenField: 'children',
      placeholder: '请选择副栏目',
      clearable: true,
      multiple: true,
      defaultExpandAll: true,
    },
  },
  {
    field: 'id',
    component: 'NInput',
    label: 'ID',
    componentProps: { placeholder: '请输入Id' },
  },
  {
    field: 'flag',
    component: 'NSelect',
    label: '标志',
    defaultValue: null,
    componentProps: { placeholder: '请选择标志', options: articleFlagOptions, clearable: true },
  },
  {
    field: 'weigh',
    component: 'NInputNumber',
    label: '权重',
    componentProps: { placeholder: '请输入权重' },
  },
  {
    field: 'price',
    component: 'NInputNumber',
    label: '价格',
    componentProps: { placeholder: '请输入价格' },
  },
  {
    field: 'createTime',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      valueFormat: 'yyyy-MM-dd HH:mm:ss',
    },
  },
  {
    field: 'updateTime',
    component: 'NDatePicker',
    label: '更新时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      valueFormat: 'yyyy-MM-dd HH:mm:ss',
    },
  },
]);

// 树选项辅助函数
function getChannelName(channels: any[], channelId: number): string {
  if (!channels || !Array.isArray(channels)) return '';
  for (const item of channels) {
    if (String(item.value) === String(channelId) || String(item.id) === String(channelId) || String(item.key) === String(channelId)) {
      return item.name || item.label || item.title || String(channelId);
    }
    if (item.children && item.children.length > 0) {
      const found = getChannelName(item.children, channelId);
      if (found) return found;
    }
  }
  return '';
}

export const columns = [
  {
    title: '栏目',
    key: 'channelId',
    align: 'left',
    width: 120,
    render(row: State) {
      const channelName = channelTreeOption.value.length === 0 
        ? String(row.channelId) 
        : (getChannelName(channelTreeOption.value, row.channelId) || String(row.channelId));
      return h(NTag, { type: 'info', size: 'small' }, { default: () => channelName });
    },
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 220,
    ellipsis: { tooltip: true },
    render(row: State) {
      const style = row.titleStyle || '';
      const parts = style.split(',');
      const isBold = parts[0] === 'bold';
      const color = parts[1] || '';
      return h('span', { style: { fontWeight: isBold ? 'bold' : 'normal', color: color || 'inherit' } }, { default: () => row.title });
    }
  },
  {
    title: '封面',
    key: 'image',
    align: 'center',
    width: 70,
    render(row: State) {
      if (row.image) {
        return renderImage(row.image);
      } else if (row.images && row.images.length > 0) {
        const imagesList = row.images.split(',');
        return renderImage(imagesList[0]);
      }
      return renderImage('暂无封面');
    },
  },
  // {
  //   title: '自定义名称',
  //   key: 'diyname',
  //   align: 'left',
  //   width: 120,
  //   ellipsis: { tooltip: true },
  // },
  // {
  //   title: '型号名称',
  //   key: 'name',
  //   align: 'left',
  //   width: 120,
  // },
  {
    title: '价格',
    key: 'price',
    align: 'center',
    width: 80,
    //sorter: 'custom',
    render(row: State) {
      if (row.price && row.price > 0) {
        return `${Number(row.price).toFixed(2)}`;
      }
      return '-';
    },
  },
  {
    title: '浏览量',
    key: 'views',
    align: 'center',
    width: 80,
    //sorter: 'custom',
  },
  // {
  //   title: '标志',
  //   key: 'flag',
  //   align: 'center',
  //   width: 80,
  //   render(row: State) {
  //     const label = articleFlagOptions.find((o) => o.value === row.flag)?.label ?? row.flag;
  //     return h(NTag, { size: 'small' }, { default: () => label });
  //   },
  // },
  {
    title: '权重',
    key: 'weigh',
    align: 'center',
    width: 80,
    //sorter: 'custom',
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
    title: '发布时间',
    key: 'publishTime',
    align: 'left',
    width: 160,
    //sorter: 'custom',
    render(row: State) {
      return row.publishTime ? timestampToTime(row.publishTime) : '-';
    },
  },
  {
    title: '更新时间',
    key: 'updateTime',
    align: 'left',
    width: 160,
    //sorter: 'custom',
    render(row: State) {
      return (row as any).updateTime ? timestampToTime((row as any).updateTime) : '-';
    },
  },
];

export function loadOptions() {
  dict.loadOptions([]);
}

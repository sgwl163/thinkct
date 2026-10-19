import { defineStore } from 'pinia';

export interface TableDisplaySettings {
  striped?: boolean;
  size?: 'small' | 'medium' | 'large';
  pageSize?: number;
}

interface ListSettingState extends TableDisplaySettings {
  searchCollapsed: boolean;
}

// 列表通用展示偏好由 Pinia 共享，列设置仍由每个表格独立维护。
export const useListSettingStore = defineStore({
  id: 'app-list-setting',
  state: (): ListSettingState => ({
    searchCollapsed: true,
    striped: undefined,
    size: undefined,
    pageSize: undefined,
  }),
  actions: {
    setSearchCollapsed(collapsed: boolean) {
      this.searchCollapsed = collapsed;
    },
    setTableSettings(settings: TableDisplaySettings) {
      this.$patch(settings);
    },
  },
});

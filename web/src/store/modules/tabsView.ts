import { defineStore } from 'pinia';
import { RouteLocationNormalized } from 'vue-router';

// 不需要出现在标签页中的路由
const whiteList = ['Redirect', 'login', 'ErrorPage'];
const redirectPath = '/redirect';

export function isExcludedTabRoute(route: Partial<RouteLocationNormalized>): boolean {
  return Boolean(
    whiteList.includes(route.name as string) ||
      route.path === redirectPath ||
      route.path?.startsWith(`${redirectPath}/`)
  );
}

export type RouteItem = Partial<RouteLocationNormalized> & {
  fullPath: string;
  path: string;
  name: string;
  hash: string;
  meta: object;
  params: object;
  query: object;
};

export type ITabsViewState = {
  tabsList: RouteItem[]; // 标签页
  signal: string; // 关闭信号
};

//保留固定路由
function retainAffixRoute(list: any[]) {
  return list.filter((item) => item?.meta?.affix ?? false);
}

export const useTabsViewStore = defineStore({
  id: 'app-tabs-view',
  state: (): ITabsViewState => ({
    tabsList: [],
    signal: '',
  }),
  getters: {},
  actions: {
    initTabs(routes) {
      // 初始化标签页
      this.tabsList = routes;
    },
    addTabs(route): boolean {
      // 添加标签页
      if (isExcludedTabRoute(route)) return false;
      const isExists = this.tabsList.some((item) => item.fullPath == route.fullPath);
      if (!isExists) {
        this.tabsList.push(route);
      }
      return true;
    },
    replaceTabs(fullPath: string, route: RouteItem) {
      if (isExcludedTabRoute(route)) return;
      const index = this.tabsList.findIndex((item) => item.fullPath === fullPath);
      if (index === -1) return;
      // 参数替换仍属于当前标签页，移除路由监听可能已添加的目标标签。
      this.tabsList[index] = route;
      this.tabsList = this.tabsList.filter(
        (item, itemIndex) => itemIndex === index || item.fullPath !== route.fullPath
      );
    },
    closeLeftTabs(route) {
      // 关闭左侧
      const index = this.tabsList.findIndex((item) => item.fullPath == route.fullPath);
      this.tabsList = this.tabsList.filter((item, i) => i >= index || (item?.meta?.affix ?? false));
    },
    closeRightTabs(route) {
      // 关闭右侧
      const index = this.tabsList.findIndex((item) => item.fullPath == route.fullPath);
      this.tabsList = this.tabsList.filter((item, i) => i <= index || (item?.meta?.affix ?? false));
    },
    closeOtherTabs(route) {
      // 关闭其他
      this.tabsList = this.tabsList.filter(
        (item) => item.fullPath == route.fullPath || (item?.meta?.affix ?? false)
      );
    },
    closeCurrentTab(route) {
      // 关闭当前页
      const index = this.tabsList.findIndex((item) => item.fullPath == route.fullPath);
      this.tabsList.splice(index, 1);
    },
    closeAllTabs() {
      // 关闭全部
      console.log(retainAffixRoute(this.tabsList));
      this.tabsList = retainAffixRoute(this.tabsList);
    },
    closeSignal(signal: string) {
      // 设置tab关闭信号， 1刷新 2关闭当前 3关闭其他 4关闭全部
      this.signal = signal;
    },
  },
});

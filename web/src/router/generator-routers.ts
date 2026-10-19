import { adminMenus } from '@/api/system/menu';
import { constantRouterIcon } from './router-icons';
import { RouteRecordRaw } from 'vue-router';
import { Layout, ParentLayout } from '@/router/constant';
import type { AppRouteRecordRaw } from '@/router/types';

const Iframe = () => import('@/views/iframe/index.vue');
const LayoutMap = new Map<string, () => Promise<typeof import('*.vue')>>();

LayoutMap.set('LAYOUT', Layout);
LayoutMap.set('IFRAME', Iframe);

/** 规范化内部路由路径，保留查询参数、锚点和外链地址。 */
function normalizeRoutePath(path: string): string {
  if (/^https?:\/\//i.test(path)) return path;
  const suffixIndex = path.search(/[?#]/);
  const pathname = suffixIndex === -1 ? path : path.slice(0, suffixIndex);
  const suffix = suffixIndex === -1 ? '' : path.slice(suffixIndex);
  return (pathname.replace(/\/+/g, '/').replace(/\/$/, '') || '/') + suffix;
}

/**
 * 格式化 后端 结构信息并递归生成层级路由表
 * @param routerMap
 * @param parent
 * @returns {*}
 */
export const routerGenerator = (routerMap, parent?): any[] => {
  return routerMap.map((item) => {
    const path = normalizeRoutePath(`${(parent && parent.path) || ''}/${item.path}`);
    const currentRouter: any = {
      // 子菜单沿用后端的层级拼接规则，统一清理重复斜杠。
      path,
      // 路由名称，建议唯一
      name: item.name || '',
      // 该路由对应页面的 组件
      component: item.component,
      // meta: 页面标题, 菜单图标, 页面权限(供指令权限用，可去掉)
      meta: {
        ...item.meta,
        label: item.meta.title,
        icon: constantRouterIcon[item.meta.icon] || null,
        permissions: item.meta.permissions || null,
      },
    };

    // 显式配置的目录跳转也需与路由路径使用相同的斜杠规则。
    if (item.meta.type === 1 && item.redirect) {
      currentRouter.redirect = normalizeRoutePath(item.redirect);
    }
    // 先生成完整子路由，再使用实际地址作为目录的默认跳转。
    if (item.children && item.children.length > 0) {
      currentRouter.children = routerGenerator(item.children, currentRouter);
      if (item.meta.type === 1 && !item.redirect) {
        currentRouter.redirect = currentRouter.children[0].path;
      }
    }

    return currentRouter;
  });
};

/**
 * 动态生成菜单
 */
export const generatorDynamicRouter = (): Promise<RouteRecordRaw[]> => {
  return new Promise((resolve, reject) => {
    adminMenus()
      .then((result) => {
        const routeList = routerGenerator(result.list);
        asyncImportRoute(routeList);

        resolve(routeList);
      })
      .catch((err) => {
        reject(err);
      });
  });
};

/**
 * 查找views中对应的组件文件
 * */
const currentViewModules = import.meta.glob('../views/**/*.{vue,tsx}');
// 保留映射对象引用，让既有路由调用方在热更新后也能找到新生成的页面。
const viewsModules: Record<string, () => Promise<Recordable>> =
  import.meta.hot?.data.viewsModules || {};
for (const key of Object.keys(viewsModules)) delete viewsModules[key];
Object.assign(viewsModules, currentViewModules);
if (import.meta.hot) {
  import.meta.hot.data.viewsModules = viewsModules;
  // 新增或删除页面只更新组件映射，避免生成期间反复刷新整页并中断请求。
  import.meta.hot.accept();
}
export const asyncImportRoute = (routes: AppRouteRecordRaw[] | undefined): void => {
  if (!routes) return;
  routes.forEach((item) => {
    if (!item.component && item.meta?.frameSrc) {
      item.component = 'IFRAME';
    }
    const { component, name } = item;
    const { children } = item;
    if (component) {
      const layoutFound = LayoutMap.get(component as string);
      if (layoutFound) {
        item.component = layoutFound;
      } else {
        item.component = dynamicImport(viewsModules, component as string);
      }
    } else if (name) {
      item.component = ParentLayout;
    }
    if (children) asyncImportRoute(children);
  });
};

/**
 * 动态导入
 * */
export const dynamicImport = (
  viewsModules: Record<string, () => Promise<Recordable>>,
  component: string
) => {
  const keys = Object.keys(viewsModules);
  const matchKeys = keys.filter((key) => {
    let k = key.replace('../views', '');
    const lastIndex = k.lastIndexOf('.');
    k = k.substring(0, lastIndex);
    return k === component;
  });
  if (matchKeys?.length === 1) {
    const matchKey = matchKeys[0];
    return viewsModules[matchKey];
  }
  if (matchKeys?.length > 1) {
    console.warn(
      'Please do not create `.vue` and `.TSX` files with the same file name in the same hierarchical directory under the views folder. This will cause dynamic introduction failure'
    );
    return;
  }
};

/**
 * 移除隐藏的菜单
 * @param menus
 */
export const removeHiddenMenus = (menus: any[]) => {
  const arr: any[] = [];
  for (let j = 0; j < menus.length; j++) {
    if (menus[j].meta?.type === 3) {
      continue;
    }
    if (menus[j].meta?.hidden === true) {
      continue;
    }

    if (menus[j].children?.length > 0) {
      menus[j].children = removeHiddenMenus(menus[j].children);
      if (menus[j].children?.length === 0) {
        delete menus[j].children;
      }
    }
    arr.push(menus[j]);
  }
  return arr;
};

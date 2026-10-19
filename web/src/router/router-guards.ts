import type { RouteRecordRaw } from 'vue-router';
import { Router } from 'vue-router';
import { useUserStoreWidthOut } from '@/store/modules/user';
import { useAsyncRouteStoreWidthOut } from '@/store/modules/asyncRoute';
import { ACCESS_TOKEN } from '@/store/mutation-types';
import { storage } from '@/utils/Storage';
import { PageEnum } from '@/enums/pageEnum';
import { ResultEnum } from '@/enums/httpEnum';
import { ErrorPageRoute } from '@/router/base';
import { jump } from '@/utils/http/axios';
import { getNowUrl } from '@/utils/urlUtils';

const LOGIN_PATH = PageEnum.BASE_LOGIN;
const whitePathList = [LOGIN_PATH]; // no redirect whitelist

export function createRouterGuards(router: Router) {
  const userStore = useUserStoreWidthOut();
  const asyncRouteStore = useAsyncRouteStoreWidthOut();
  let errorDialogOpen = false;

  router.beforeEach(async (to) => {
    window['$loading']?.start();

    // Whitelist can be directly entered
    if (whitePathList.includes(to.path as PageEnum)) {
      await userStore.LoadLoginConfig();
      return true;
    }

    const token = storage.get(ACCESS_TOKEN);

    if (!token) {
      // You can access without permissions. You need to set the routing meta.ignoreAuth to true
      if (to.meta.ignoreAuth) {
        return true;
      }

      return {
        path: LOGIN_PATH,
        replace: true,
        query: { redirect: to.fullPath },
      };
    }

    if (asyncRouteStore.getIsDynamicAddedRoute) {
      return true;
    }

    const userInfo = await userStore.GetInfo();
    await userStore.LoadLoginConfig();

    // 是否允许获取微信openid，保留查询参数和页面锚点。
    if (userStore.allowWxOpenId()) {
      const URI = getNowUrl() + '#' + to.fullPath;
      jump('/wechat/authorize', { type: 'openId', syncRedirect: URI });
      return false;
    }

    await userStore.GetConfig();
    const routes = await asyncRouteStore.generateRoutes(userInfo);

    // 动态添加可访问路由表
    routes.forEach((item) => {
      router.addRoute(item as unknown as RouteRecordRaw);
    });

    // 添加404
    if (!router.hasRoute(ErrorPageRoute.name!)) {
      router.addRoute(ErrorPageRoute as unknown as RouteRecordRaw);
    }

    asyncRouteStore.setDynamicAddedRoute(true);
    // 新增路由后重新匹配目标页面，不再从上一页读取或解码回跳地址。
    return { path: to.path, query: to.query, hash: to.hash, replace: true };
  });

  router.afterEach((to, _, failure) => {
    if (!failure) {
      document.title = (to.meta.title as string) || document.title;
    }
    window['$loading']?.finish();
  });

  router.onError((error, to) => {
    window['$loading']?.finish();
    console.error('页面路由加载失败。', error);
    const dialog = window['$dialog'];
    if (!dialog || errorDialogOpen || error?.code === ResultEnum.TIMEOUT) {
      return;
    }

    errorDialogOpen = true;
    const closeDialog = () => {
      errorDialogOpen = false;
    };
    dialog.warning({
      title: '页面加载失败',
      content: '页面、用户信息或菜单加载失败，请检查网络后重试；若仍然失败，请联系管理员。',
      positiveText: '重试',
      negativeText: '取消',
      onPositiveClick: async () => {
        try {
          await router.replace(to.fullPath);
          return true;
        } catch {
          // 保留当前对话框，网络恢复后可继续重试，避免叠加提示。
          return false;
        }
      },
      onAfterLeave: closeDialog,
    });
  });
}

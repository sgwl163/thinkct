import { RouteLocationRaw } from 'vue-router';
import router from '@/router';
import { PageEnum } from '@/enums/pageEnum';
import { useTabsViewStore } from '@/store/modules/tabsView';

/**
 * 将对象添加当作参数拼接到URL上面
 * @param baseUrl 需要拼接的url
 * @param obj 参数对象
 * @returns {string} 拼接后的对象
 * 例子:
 *  let obj = {a: '3', b: '4'}
 *  setObjToUrlParams('www.baidu.com', obj)
 *  ==>www.baidu.com?a=3&b=4
 */
export function setObjToUrlParams(baseUrl: string, obj: object): string {
  let parameters = '';
  let url = '';
  for (const key in obj) {
    parameters += key + '=' + encodeURIComponent(obj[key]) + '&';
  }
  parameters = parameters.replace(/&$/, '');
  if (/\?$/.test(baseUrl)) {
    url = baseUrl + parameters;
  } else {
    url = baseUrl.replace(/\/?$/, '?') + parameters;
  }
  return url;
}

export function encodeParams(obj) {
  const arr = [];
  for (const p in obj) {
    // @ts-ignore
    arr.push(encodeURIComponent(p) + '=' + encodeURIComponent(obj[p]));
  }
  return arr.join('&');
}

/**
 * 获取文件后缀
 */
export function getFileExt(fileName: string) {
  if (fileName === undefined || fileName === '') {
    return ``;
  }
  return fileName.substring(fileName.lastIndexOf('.') + 1);
}

/**
 * 获取当访问的url，不含参数
 */
export function getNowUrl(): string {
  const w = window.location;
  return w.protocol + '//' + w.host + w.pathname;
}

// 返回上一页
export function goBackOrToPage(to: RouteLocationRaw): void {
  if (router.currentRoute.value.matched.length > 0) {
    router.go(-1);
  } else {
    router.push(to).catch((e) => {
      console.log('返回上一页失败:', e);
    });
  }
}

/** 校验登录回跳地址。query 已由路由解码，保留原始参数编码，避免二次解码。 */
export function getLoginRedirect(value: unknown): string {
  if (typeof value !== 'string' || !value.startsWith('/') || value.startsWith('//')) {
    return PageEnum.BASE_HOME;
  }

  try {
    const path = decodeURIComponent(value.split(/[?#]/)[0]);
    if (
      path.startsWith('//') ||
      /[\\\s]/.test(path) ||
      /[\u0000-\u001f\u007f]/.test(path + value) ||
      path.replace(/\/+$/, '').toLowerCase() === PageEnum.BASE_LOGIN
    ) {
      return PageEnum.BASE_HOME;
    }
  } catch {
    return PageEnum.BASE_HOME;
  }

  return value;
}

/** 更新当前页面查询参数，同时更新当前标签页，不增加后退记录。 */
export async function setRouteQuery(key: string, value: string | number) {
  const route = router.currentRoute.value;
  const target = {
    path: route.path,
    query: { ...route.query, [key]: String(value) },
    hash: route.hash,
  };
  const failure = await router.replace(target);
  if (!failure && router.currentRoute.value.fullPath === router.resolve(target).fullPath) {
    const { fullPath, hash, meta, name, params, path, query } = router.currentRoute.value;
    useTabsViewStore().replaceTabs(route.fullPath, {
      fullPath,
      hash,
      meta,
      name: name as string,
      params,
      path,
      query,
    });
  }
  return failure;
}

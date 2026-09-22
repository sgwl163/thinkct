import { http, jumpExport } from '@/utils/http/axios';

// 获取升级记录列表
export function List(params) {
  return http.request({
    url: '/oauthLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除升级记录
export function Delete(params) {
  return http.request({
    url: '/oauthLog/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑升级记录
export function Edit(params) {
  return http.request({
    url: '/oauthLog/edit',
    method: 'POST',
    params,
  });
}

// 获取升级记录指定详情
export function View(params) {
  return http.request({
    url: '/oauthLog/view',
    method: 'GET',
    params,
  });
}

// 导出升级记录
export function Export(params) {
  jumpExport('/oauthLog/export', params);
}
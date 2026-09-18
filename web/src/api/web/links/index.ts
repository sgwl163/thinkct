import { http, jumpExport } from '@/utils/http/axios';

// 获取友情链接列表
export function List(params) {
  return http.request({
    url: '/links/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除友情链接
export function Delete(params) {
  return http.request({
    url: '/links/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑友情链接
export function Edit(params) {
  return http.request({
    url: '/links/edit',
    method: 'POST',
    params,
  });
}

// 修改友情链接状态
export function Status(params) {
  return http.request({
    url: '/links/status',
    method: 'POST',
    params,
  });
}

// 获取友情链接指定详情
export function View(params) {
  return http.request({
    url: '/links/view',
    method: 'GET',
    params,
  });
}

// 导出友情链接
export function Export(params) {
  jumpExport('/links/export', params);
}
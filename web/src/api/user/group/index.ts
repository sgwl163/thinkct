import { http, jumpExport } from '@/utils/http/axios';

// 获取分组管理列表
export function List(params) {
  return http.request({
    url: '/group/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除分组管理
export function Delete(params) {
  return http.request({
    url: '/group/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑分组管理
export function Edit(params) {
  return http.request({
    url: '/group/edit',
    method: 'POST',
    params,
  });
}

// 修改分组管理状态
export function Status(params) {
  return http.request({
    url: '/group/status',
    method: 'POST',
    params,
  });
}

// 获取分组管理指定详情
export function View(params) {
  return http.request({
    url: '/group/view',
    method: 'GET',
    params,
  });
}

// 导出分组管理
export function Export(params) {
  jumpExport('/group/export', params);
}
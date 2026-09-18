import { http, jumpExport } from '@/utils/http/axios';

// 获取产品管理列表
export function List(params) {
  return http.request({
    url: '/products/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除产品管理
export function Delete(params) {
  return http.request({
    url: '/products/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑产品管理
export function Edit(params) {
  return http.request({
    url: '/products/edit',
    method: 'POST',
    params,
  });
}

// 修改产品管理状态
export function Status(params) {
  return http.request({
    url: '/products/status',
    method: 'POST',
    params,
  });
}

// 获取产品管理指定详情
export function View(params) {
  return http.request({
    url: '/products/view',
    method: 'GET',
    params,
  });
}

// 导出产品管理
export function Export(params) {
  jumpExport('/products/export', params);
}
import { http, jumpExport } from '@/utils/http/axios';

// 获取盗版管理列表
export function List(params) {
  return http.request({
    url: '/daoban/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除盗版管理
export function Delete(params) {
  return http.request({
    url: '/daoban/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑盗版管理
export function Edit(params) {
  return http.request({
    url: '/daoban/edit',
    method: 'POST',
    params,
  });
}

// 获取盗版管理指定详情
export function View(params) {
  return http.request({
    url: '/daoban/view',
    method: 'GET',
    params,
  });
}

// 导出盗版管理
export function Export(params) {
  jumpExport('/daoban/export', params);
}
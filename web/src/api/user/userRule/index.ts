import { http, jumpExport } from '@/utils/http/axios';

// 获取权限规则列表
export function List(params) {
  return http.request({
    url: '/userRule/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除权限规则
export function Delete(params) {
  return http.request({
    url: '/userRule/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑权限规则
export function Edit(params) {
  return http.request({
    url: '/userRule/edit',
    method: 'POST',
    params,
  });
}

// 获取权限规则指定详情
export function View(params) {
  return http.request({
    url: '/userRule/view',
    method: 'GET',
    params,
  });
}

// 导出权限规则
export function Export(params) {
  jumpExport('/userRule/export', params);
}
import { http, jumpExport } from '@/utils/http/axios';

// 获取用户分组列表
export function List(params) {
  return http.request({
    url: '/userGroup/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除用户分组
export function Delete(params) {
  return http.request({
    url: '/userGroup/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑用户分组
export function Edit(params) {
  return http.request({
    url: '/userGroup/edit',
    method: 'POST',
    params,
  });
}

// 获取用户分组指定详情
export function View(params) {
  return http.request({
    url: '/userGroup/view',
    method: 'GET',
    params,
  });
}

// 导出用户分组
export function Export(params) {
  jumpExport('/userGroup/export', params);
}
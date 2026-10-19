import { http, jumpExport } from '@/utils/http/axios';

// 获取授权管理列表
export function List(params) {
  return http.request({
    url: '/appAuth/appAuth/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除授权管理
export function Delete(params) {
  return http.request({
    url: '/appAuth/appAuth/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑授权管理
export function Edit(params) {
  return http.request({
    url: '/appAuth/appAuth/edit',
    method: 'POST',
    params,
  });
}

// 修改授权管理状态
export function Status(params) {
  return http.request({
    url: '/appAuth/appAuth/status',
    method: 'POST',
    params,
  });
}

// 获取授权管理指定详情
export function View(params) {
  return http.request({
    url: '/appAuth/appAuth/view',
    method: 'GET',
    params,
  });
}

// 导出授权管理
export function Export(params) {
  jumpExport('/appAuth/appAuth/export', params);
}
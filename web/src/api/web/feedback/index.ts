import { http, jumpExport } from '@/utils/http/axios';

// 获取留言管理列表
export function List(params) {
  return http.request({
    url: '/feedback/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除留言管理
export function Delete(params) {
  return http.request({
    url: '/feedback/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑留言管理
export function Edit(params) {
  return http.request({
    url: '/feedback/edit',
    method: 'POST',
    params,
  });
}

// 修改留言管理状态
export function Status(params) {
  return http.request({
    url: '/feedback/status',
    method: 'POST',
    params,
  });
}

// 获取留言管理指定详情
export function View(params) {
  return http.request({
    url: '/feedback/view',
    method: 'GET',
    params,
  });
}

// 导出留言管理
export function Export(params) {
  jumpExport('/feedback/export', params);
}
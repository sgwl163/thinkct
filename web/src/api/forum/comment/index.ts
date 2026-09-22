import { http, jumpExport } from '@/utils/http/axios';

// 获取评论管理列表
export function List(params) {
  return http.request({
    url: '/comment/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除评论管理
export function Delete(params) {
  return http.request({
    url: '/comment/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑评论管理
export function Edit(params) {
  return http.request({
    url: '/comment/edit',
    method: 'POST',
    params,
  });
}

// 获取评论管理指定详情
export function View(params) {
  return http.request({
    url: '/comment/view',
    method: 'GET',
    params,
  });
}

// 导出评论管理
export function Export(params) {
  jumpExport('/comment/export', params);
}
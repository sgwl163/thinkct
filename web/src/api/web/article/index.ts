import { http, jumpExport } from '@/utils/http/axios';

// 获取文章管理列表
export function List(params) {
  return http.request({
    url: '/article/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除文章管理
export function Delete(params) {
  return http.request({
    url: '/article/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑文章管理
export function Edit(params) {
  return http.request({
    url: '/article/edit',
    method: 'POST',
    params,
  });
}

// 修改文章管理状态
export function Status(params) {
  return http.request({
    url: '/article/status',
    method: 'POST',
    params,
  });
}

// 获取文章管理指定详情
export function View(params) {
  return http.request({
    url: '/article/view',
    method: 'GET',
    params,
  });
}

// 导出文章管理
export function Export(params) {
  jumpExport('/article/export', params);
}
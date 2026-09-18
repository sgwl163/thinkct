import { http, jumpExport } from '@/utils/http/axios';

// 获取标签管理列表
export function List(params) {
  return http.request({
    url: '/tags/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除标签管理
export function Delete(params) {
  return http.request({
    url: '/tags/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑标签管理
export function Edit(params) {
  return http.request({
    url: '/tags/edit',
    method: 'POST',
    params,
  });
}

// 修改标签管理状态
export function Status(params) {
  return http.request({
    url: '/tags/status',
    method: 'POST',
    params,
  });
}

// 获取标签管理指定详情
export function View(params) {
  return http.request({
    url: '/tags/view',
    method: 'GET',
    params,
  });
}

// 导出标签管理
export function Export(params) {
  jumpExport('/tags/export', params);
}
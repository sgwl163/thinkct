import { http, jumpExport } from '@/utils/http/axios';

// 获取会话列表管理列表
export function List(params) {
  return http.request({
    url: '/conversation/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除会话列表管理
export function Delete(params) {
  return http.request({
    url: '/conversation/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会话列表管理
export function Edit(params) {
  return http.request({
    url: '/conversation/edit',
    method: 'POST',
    params,
  });
}

// 获取会话列表管理指定详情
export function View(params) {
  return http.request({
    url: '/conversation/view',
    method: 'GET',
    params,
  });
}

// 导出会话列表管理
export function Export(params) {
  jumpExport('/conversation/export', params);
}
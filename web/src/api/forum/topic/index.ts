import { http, jumpExport } from '@/utils/http/axios';

// 获取社区主题列表
export function List(params) {
  return http.request({
    url: '/topic/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除社区主题
export function Delete(params) {
  return http.request({
    url: '/topic/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑社区主题
export function Edit(params) {
  return http.request({
    url: '/topic/edit',
    method: 'POST',
    params,
  });
}

// 获取社区主题指定详情
export function View(params) {
  return http.request({
    url: '/topic/view',
    method: 'GET',
    params,
  });
}

// 导出社区主题
export function Export(params) {
  jumpExport('/topic/export', params);
}
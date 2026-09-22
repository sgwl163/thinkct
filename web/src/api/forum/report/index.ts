import { http, jumpExport } from '@/utils/http/axios';

// 获取社区论坛举报表列表
export function List(params) {
  return http.request({
    url: '/report/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除社区论坛举报表
export function Delete(params) {
  return http.request({
    url: '/report/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑社区论坛举报表
export function Edit(params) {
  return http.request({
    url: '/report/edit',
    method: 'POST',
    params,
  });
}

// 获取社区论坛举报表指定详情
export function View(params) {
  return http.request({
    url: '/report/view',
    method: 'GET',
    params,
  });
}

// 导出社区论坛举报表
export function Export(params) {
  jumpExport('/report/export', params);
}
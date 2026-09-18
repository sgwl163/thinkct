import { http, jumpExport } from '@/utils/http/axios';

// 获取AI视频任务列表
export function List(params) {
  return http.request({
    url: '/videoTask/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI视频任务
export function Delete(params) {
  return http.request({
    url: '/videoTask/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI视频任务
export function Edit(params) {
  return http.request({
    url: '/videoTask/edit',
    method: 'POST',
    params,
  });
}

// 修改AI视频任务状态
export function Status(params) {
  return http.request({
    url: '/videoTask/status',
    method: 'POST',
    params,
  });
}

// 获取AI视频任务指定详情
export function View(params) {
  return http.request({
    url: '/videoTask/view',
    method: 'GET',
    params,
  });
}

// 导出AI视频任务
export function Export(params) {
  jumpExport('/videoTask/export', params);
}
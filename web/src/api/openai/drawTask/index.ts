import { http, jumpExport } from '@/utils/http/axios';

// 获取AI绘画任务列表
export function List(params) {
  return http.request({
    url: '/drawTask/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI绘画任务
export function Delete(params) {
  return http.request({
    url: '/drawTask/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI绘画任务
export function Edit(params) {
  return http.request({
    url: '/drawTask/edit',
    method: 'POST',
    params,
  });
}

// 修改AI绘画任务状态
export function Status(params) {
  return http.request({
    url: '/drawTask/status',
    method: 'POST',
    params,
  });
}

// 获取AI绘画任务指定详情
export function View(params) {
  return http.request({
    url: '/drawTask/view',
    method: 'GET',
    params,
  });
}

// 导出AI绘画任务
export function Export(params) {
  jumpExport('/drawTask/export', params);
}
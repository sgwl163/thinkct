import { http, jumpExport } from '@/utils/http/axios';

// 获取AI音乐任务列表
export function List(params) {
  return http.request({
    url: '/musicTask/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI音乐任务
export function Delete(params) {
  return http.request({
    url: '/musicTask/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI音乐任务
export function Edit(params) {
  return http.request({
    url: '/musicTask/edit',
    method: 'POST',
    params,
  });
}

// 修改AI音乐任务状态
export function Status(params) {
  return http.request({
    url: '/musicTask/status',
    method: 'POST',
    params,
  });
}

// 获取AI音乐任务指定详情
export function View(params) {
  return http.request({
    url: '/musicTask/view',
    method: 'GET',
    params,
  });
}

// 导出AI音乐任务
export function Export(params) {
  jumpExport('/musicTask/export', params);
}
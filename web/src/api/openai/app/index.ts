import { http, jumpExport } from '@/utils/http/axios';

// 获取AI应用列表
export function List(params) {
  return http.request({
    url: '/app/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI应用
export function Delete(params) {
  return http.request({
    url: '/app/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI应用
export function Edit(params) {
  return http.request({
    url: '/app/edit',
    method: 'POST',
    params,
  });
}

// 修改AI应用状态
export function Status(params) {
  return http.request({
    url: '/app/status',
    method: 'POST',
    params,
  });
}

// 获取AI应用指定详情
export function View(params) {
  return http.request({
    url: '/app/view',
    method: 'GET',
    params,
  });
}

// 获取AI应用最大排序
export function MaxSort() {
  return http.request({
    url: '/app/maxSort',
    method: 'GET',
  });
}

// 导出AI应用
export function Export(params) {
  jumpExport('/app/export', params);
}
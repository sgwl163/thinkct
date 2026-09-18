import { http, jumpExport } from '@/utils/http/axios';

// 获取AI应用分类列表
export function List(params) {
  return http.request({
    url: '/openaiAppCategory/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI应用分类
export function Delete(params) {
  return http.request({
    url: '/openaiAppCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI应用分类
export function Edit(params) {
  return http.request({
    url: '/openaiAppCategory/edit',
    method: 'POST',
    params,
  });
}

// 修改AI应用分类状态
export function Status(params) {
  return http.request({
    url: '/openaiAppCategory/status',
    method: 'POST',
    params,
  });
}

// 获取AI应用分类指定详情
export function View(params) {
  return http.request({
    url: '/openaiAppCategory/view',
    method: 'GET',
    params,
  });
}

// 获取AI应用分类最大排序
export function MaxSort() {
  return http.request({
    url: '/openaiAppCategory/maxSort',
    method: 'GET',
  });
}

// 导出AI应用分类
export function Export(params) {
  jumpExport('/openaiAppCategory/export', params);
}
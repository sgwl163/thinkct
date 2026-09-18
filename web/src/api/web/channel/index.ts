import { http, jumpExport } from '@/utils/http/axios';

// 获取栏目管理列表
export function List(params) {
  return http.request({
    url: '/channel/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除栏目管理
export function Delete(params) {
  return http.request({
    url: '/channel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑栏目管理
export function Edit(params) {
  return http.request({
    url: '/channel/edit',
    method: 'POST',
    params,
  });
}

// 修改栏目管理状态
export function Status(params) {
  return http.request({
    url: '/channel/status',
    method: 'POST',
    params,
  });
}

// 获取栏目管理指定详情
export function View(params) {
  return http.request({
    url: '/channel/view',
    method: 'GET',
    params,
  });
}

// 导出栏目管理
export function Export(params) {
  jumpExport('/channel/export', params);
}

// 获取栏目管理关系树选项
export function TreeOption(params?) {
  return http.request({
    url: '/channel/treeOption',
    method: 'GET',
    params,
  });
}
import { http, jumpExport } from '@/utils/http/axios';

// 获取规则管理列表
export function List(params) {
  return http.request({
    url: '/userRule/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除规则管理
export function Delete(params) {
  return http.request({
    url: '/userRule/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑规则管理
export function Edit(params) {
  return http.request({
    url: '/userRule/edit',
    method: 'POST',
    params,
  });
}

// 修改规则管理状态
export function Status(params) {
  return http.request({
    url: '/userRule/status',
    method: 'POST',
    params,
  });
}

// 获取规则管理指定详情
export function View(params) {
  return http.request({
    url: '/userRule/view',
    method: 'GET',
    params,
  });
}

// 导出规则管理
export function Export(params) {
  jumpExport('/userRule/export', params);
}
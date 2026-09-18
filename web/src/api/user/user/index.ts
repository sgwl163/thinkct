import { http, jumpExport } from '@/utils/http/axios';

// 获取会员管理列表
export function List(params) {
  return http.request({
    url: '/user/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除会员管理
export function Delete(params) {
  return http.request({
    url: '/user/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑会员管理
export function Edit(params) {
  return http.request({
    url: '/user/edit',
    method: 'POST',
    params,
  });
}

// 修改会员管理状态
export function Status(params) {
  return http.request({
    url: '/user/status',
    method: 'POST',
    params,
  });
}

// 操作会员管理开关
export function Switch(params) {
  return http.request({
    url: '/user/switch',
    method: 'POST',
    params,
  });
}

// 获取会员管理指定详情
export function View(params) {
  return http.request({
    url: '/user/view',
    method: 'GET',
    params,
  });
}

// 导出会员管理
export function Export(params) {
  jumpExport('/user/export', params);
}
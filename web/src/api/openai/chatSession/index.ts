import { http, jumpExport } from '@/utils/http/axios';

// 获取聊天会话列表
export function List(params) {
  return http.request({
    url: '/chatSession/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除聊天会话
export function Delete(params) {
  return http.request({
    url: '/chatSession/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑聊天会话
export function Edit(params) {
  return http.request({
    url: '/chatSession/edit',
    method: 'POST',
    params,
  });
}

// 修改聊天会话状态
export function Status(params) {
  return http.request({
    url: '/chatSession/status',
    method: 'POST',
    params,
  });
}

// 获取聊天会话指定详情
export function View(params) {
  return http.request({
    url: '/chatSession/view',
    method: 'GET',
    params,
  });
}

// 导出聊天会话
export function Export(params) {
  jumpExport('/chatSession/export', params);
}
import { http, jumpExport } from '@/utils/http/axios';

// 获取IM-聊天记录表列表
export function List(params) {
  return http.request({
    url: '/message/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除IM-聊天记录表
export function Delete(params) {
  return http.request({
    url: '/message/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑IM-聊天记录表
export function Edit(params) {
  return http.request({
    url: '/message/edit',
    method: 'POST',
    params,
  });
}

// 修改IM-聊天记录表状态
export function Status(params) {
  return http.request({
    url: '/message/status',
    method: 'POST',
    params,
  });
}

// 获取IM-聊天记录表指定详情
export function View(params) {
  return http.request({
    url: '/message/view',
    method: 'GET',
    params,
  });
}

// 导出IM-聊天记录表
export function Export(params) {
  jumpExport('/message/export', params);
}
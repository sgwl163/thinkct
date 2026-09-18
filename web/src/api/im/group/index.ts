import { http, jumpExport } from '@/utils/http/axios';

// 获取IM-群组表列表
export function List(params) {
  return http.request({
    url: '/group/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除IM-群组表
export function Delete(params) {
  return http.request({
    url: '/group/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑IM-群组表
export function Edit(params) {
  return http.request({
    url: '/group/edit',
    method: 'POST',
    params,
  });
}

// 修改IM-群组表状态
export function Status(params) {
  return http.request({
    url: '/group/status',
    method: 'POST',
    params,
  });
}

// 获取IM-群组表指定详情
export function View(params) {
  return http.request({
    url: '/group/view',
    method: 'GET',
    params,
  });
}

// 导出IM-群组表
export function Export(params) {
  jumpExport('/group/export', params);
}
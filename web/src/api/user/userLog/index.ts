import { http, jumpExport } from '@/utils/http/axios';

// 获取操作记录列表
export function List(params) {
  return http.request({
    url: '/userLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除操作记录
export function Delete(params) {
  return http.request({
    url: '/userLog/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑操作记录
export function Edit(params) {
  return http.request({
    url: '/userLog/edit',
    method: 'POST',
    params,
  });
}

// 修改操作记录状态
export function Status(params) {
  return http.request({
    url: '/userLog/status',
    method: 'POST',
    params,
  });
}

// 获取操作记录指定详情
export function View(params) {
  return http.request({
    url: '/userLog/view',
    method: 'GET',
    params,
  });
}

// 导出操作记录
export function Export(params) {
  jumpExport('/userLog/export', params);
}
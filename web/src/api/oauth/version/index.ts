import { http, jumpExport } from '@/utils/http/axios';

// 获取版本更新列表
export function List(params) {
  return http.request({
    url: '/version/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除版本更新
export function Delete(params) {
  return http.request({
    url: '/version/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑版本更新
export function Edit(params) {
  return http.request({
    url: '/version/edit',
    method: 'POST',
    params,
  });
}

// 修改版本更新状态
export function Status(params) {
  return http.request({
    url: '/version/status',
    method: 'POST',
    params,
  });
}

// 获取版本更新指定详情
export function View(params) {
  return http.request({
    url: '/version/view',
    method: 'GET',
    params,
  });
}

// 导出版本更新
export function Export(params) {
  jumpExport('/version/export', params);
}
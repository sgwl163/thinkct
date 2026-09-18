import { http, jumpExport } from '@/utils/http/axios';

// 获取分组列表
export function List(params) {
  return http.request({
    url: '/group/list',
    method: 'get',
    params,
  });
}

// 删除分组
export function Delete(params) {
  return http.request({
    url: '/group/delete',
    method: 'POST',
    params,
  });
}

// 编辑/新增分组
export function Edit(params) {
  return http.request({
    url: '/group/edit',
    method: 'POST',
    params,
  });
}

// 获取指定分组信息
export function View(params) {
  return http.request({
    url: '/group/view',
    method: 'GET',
    params,
  });
}

// 导出分组列表
export function Export(params) {
  jumpExport('/group/export', params);
}

// 修改分组状态
export function Status(params) {
  return http.request({
    url: '/group/status',
    method: 'POST',
    params,
  });
}

// 获取最大排序值
export function MaxSort(params) {
  return http.request({
    url: '/group/maxSort',
    method: 'GET',
    params,
  });
}

import { http, jumpExport } from '@/utils/http/axios';

// 获取AI接口分组列表
export function List(params) {
  return http.request({
    url: '/apiGroup/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI接口分组
export function Delete(params) {
  return http.request({
    url: '/apiGroup/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI接口分组
export function Edit(params) {
  return http.request({
    url: '/apiGroup/edit',
    method: 'POST',
    params,
  });
}

// 修改AI接口分组状态
export function Status(params) {
  return http.request({
    url: '/apiGroup/status',
    method: 'POST',
    params,
  });
}

// 获取AI接口分组指定详情
export function View(params) {
  return http.request({
    url: '/apiGroup/view',
    method: 'GET',
    params,
  });
}

// 获取AI接口分组最大排序
export function MaxSort() {
  return http.request({
    url: '/apiGroup/maxSort',
    method: 'GET',
  });
}

// 导出AI接口分组
export function Export(params) {
  jumpExport('/apiGroup/export', params);
}
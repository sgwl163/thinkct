import { http, jumpExport } from '@/utils/http/axios';

// 获取AI模型列表列表
export function List(params) {
  return http.request({
    url: '/apiModel/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AI模型列表
export function Delete(params) {
  return http.request({
    url: '/apiModel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AI模型列表
export function Edit(params) {
  return http.request({
    url: '/apiModel/edit',
    method: 'POST',
    params,
  });
}

// 修改AI模型列表状态
export function Status(params) {
  return http.request({
    url: '/apiModel/status',
    method: 'POST',
    params,
  });
}

// 获取AI模型列表指定详情
export function View(params) {
  return http.request({
    url: '/apiModel/view',
    method: 'GET',
    params,
  });
}

// 获取AI模型列表最大排序
export function MaxSort() {
  return http.request({
    url: '/apiModel/maxSort',
    method: 'GET',
  });
}

// 导出AI模型列表
export function Export(params) {
  jumpExport('/apiModel/export', params);
}
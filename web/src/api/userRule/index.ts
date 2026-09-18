import { http, jumpExport } from '@/utils/http/axios';

// 获取规则列表
export function List(params) {
  return http.request({
    url: '/userRule/list',
    method: 'get',
    params,
  });
}

// 删除规则
export function Delete(params) {
  return http.request({
    url: '/userRule/delete',
    method: 'POST',
    params,
  });
}

// 编辑/新增规则
export function Edit(params) {
  return http.request({
    url: '/userRule/edit',
    method: 'POST',
    params,
  });
}

// 获取指定规则信息
export function View(params) {
  return http.request({
    url: '/userRule/view',
    method: 'GET',
    params,
  });
}

// 导出规则列表
export function Export(params) {
  jumpExport('/userRule/export', params);
}

// 修改规则状态
export function Status(params) {
  return http.request({
    url: '/userRule/status',
    method: 'POST',
    params,
  });
}

// 获取最大排序值
export function MaxSort(params) {
  return http.request({
    url: '/userRule/maxSort',
    method: 'GET',
    params,
  });
}

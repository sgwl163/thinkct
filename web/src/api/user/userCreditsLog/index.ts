import { http, jumpExport } from '@/utils/http/axios';

// 获取资产明细列表
export function List(params) {
  return http.request({
    url: '/userCreditsLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除资产明细
export function Delete(params) {
  return http.request({
    url: '/userCreditsLog/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑资产明细
export function Edit(params) {
  return http.request({
    url: '/userCreditsLog/edit',
    method: 'POST',
    params,
  });
}

// 修改资产明细状态
export function Status(params) {
  return http.request({
    url: '/userCreditsLog/status',
    method: 'POST',
    params,
  });
}

// 获取资产明细指定详情
export function View(params) {
  return http.request({
    url: '/userCreditsLog/view',
    method: 'GET',
    params,
  });
}

// 导出资产明细
export function Export(params) {
  jumpExport('/userCreditsLog/export', params);
}
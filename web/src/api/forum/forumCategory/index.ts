import { http, jumpExport } from '@/utils/http/axios';

// 获取社区板块列表
export function List(params) {
  return http.request({
    url: '/forumCategory/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除社区板块
export function Delete(params) {
  return http.request({
    url: '/forumCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑社区板块
export function Edit(params) {
  return http.request({
    url: '/forumCategory/edit',
    method: 'POST',
    params,
  });
}

// 修改社区板块状态
export function Status(params) {
  return http.request({
    url: '/forumCategory/status',
    method: 'POST',
    params,
  });
}

// 获取社区板块指定详情
export function View(params) {
  return http.request({
    url: '/forumCategory/view',
    method: 'GET',
    params,
  });
}

// 获取社区板块最大排序
export function MaxSort() {
  return http.request({
    url: '/forumCategory/maxSort',
    method: 'GET',
  });
}

// 导出社区板块
export function Export(params) {
  jumpExport('/forumCategory/export', params);
}
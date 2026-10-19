import type { RequestOptions } from '@/utils/http/axios/types';
import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/genCodes/list',
    method: 'get',
    params,
  });
}
export function Delete(params) {
  return http.request({
    url: '/genCodes/delete',
    method: 'POST',
    params,
  });
}

export function DeleteBuild(params) {
  return http.request({
    url: '/genCodes/deleteBuild',
    method: 'POST',
    params,
  });
}

export function View(params, options?: RequestOptions) {
  return http.request(
    {
      url: '/genCodes/view',
      method: 'GET',
      params,
    },
    options
  );
}

export function Edit(params) {
  return http.request({
    url: '/genCodes/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/genCodes/status',
    method: 'POST',
    params,
  });
}
export function Selects(params, options?: RequestOptions) {
  return http.request({ url: '/genCodes/selects', method: 'get', params }, options);
}

export function TableSelect(params, options?: RequestOptions) {
  return http.request({ url: '/genCodes/tableSelect', method: 'get', params }, options);
}

export function ColumnSelect(params, options?: RequestOptions) {
  return http.request({ url: '/genCodes/columnSelect', method: 'get', params }, options);
}

export function ColumnList(params) {
  return http.request({
    url: '/genCodes/columnList',
    method: 'get',
    params,
  });
}

export function Preview(params) {
  return http.request({
    url: '/genCodes/preview',
    method: 'post',
    params,
  });
}

export function Build(params) {
  return http.request(
    {
      url: '/genCodes/build',
      method: 'post',
      // DAO 和服务接口生成可能超过全局 30 秒超时，等待真实生成结果。
      timeout: 10 * 60 * 1000,
      params,
    },
    // 页面区分后台明确失败与网络结果未知，避免将旧成功状态误当成本次成功。
    { isTransformResponse: false, isShowErrorMessage: false }
  );
}

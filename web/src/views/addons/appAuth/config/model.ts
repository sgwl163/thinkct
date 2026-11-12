import type { FormItemRule, FormRules } from 'naive-ui';
import { validate } from '@/utils/validateUtil';

export class State {
  public unauthorizedMessage = '当前域名未获得产品授权，请联系服务商';
  public expiredMessage = '产品授权已到期，请联系服务商续费';
  public disabledMessage = '产品授权已被禁用，请联系服务商';
  public purchaseUrl = '';
  public serviceName = '';
  public serviceQQ = '';
  public servicePhone = '';
  public supportUrl = '';
  public authQueryCount: number | null = 0;
}

const textFields = [
  'unauthorizedMessage',
  'expiredMessage',
  'disabledMessage',
  'purchaseUrl',
  'serviceName',
  'serviceQQ',
  'servicePhone',
  'supportUrl',
] as const;

export function newState(source: Record<string, unknown> | Partial<State> | null = null): State {
  const state = new State();
  if (!source) return state;
  for (const key of textFields) {
    if (typeof source[key] === 'string') state[key] = source[key];
  }
  const rawCount = source.authQueryCount;
  const count =
    typeof rawCount === 'string' && /^[0-9]+$/.test(rawCount) ? Number(rawCount) : rawCount;
  if (typeof count === 'number' && Number.isSafeInteger(count) && count >= 0) {
    state.authQueryCount = count;
  }
  return state;
}

// 仅提交声明的配置字段，兼容旧数据但不继续提交历史测试变量。
export function toConfigList(state: State): State {
  const list = new State();
  for (const key of textFields) list[key] = state[key].trim();
  list.authQueryCount = state.authQueryCount;
  return list;
}

function textRule(label: string, max: number, required = false): FormItemRule {
  return {
    required,
    trigger: ['blur', 'input'],
    validator: (_rule, value: unknown) => {
      if (typeof value !== 'string') return new Error(`请填写${label}`);
      const text = value.trim();
      if (required && !text) return new Error(`请填写${label}`);
      if (Array.from(text).length > max) return new Error(`${label}不能超过${max}个字符`);
      return true;
    },
  };
}

function urlRule(label: string): FormItemRule {
  return {
    trigger: ['blur', 'input'],
    validator: (_rule, value: string) => {
      const text = value.trim();
      if (!text) return true;
      try {
        const url = new URL(text);
        if (/^https?:\/\//i.test(text) && ['http:', 'https:'].includes(url.protocol) && url.host) {
          return true;
        }
      } catch {
        // 无效地址通过表单校验消息反馈。
      }
      return new Error(`${label}必须是有效的HTTP或HTTPS地址`);
    },
  };
}

export const rules: FormRules = {
  unauthorizedMessage: textRule('未授权提示', 500, true),
  expiredMessage: textRule('到期提示', 500, true),
  disabledMessage: textRule('禁用提示', 500, true),
  purchaseUrl: [textRule('购买地址', 500), urlRule('购买地址')],
  serviceName: textRule('服务商名称', 100),
  serviceQQ: [
    textRule('客服QQ', 13),
    {
      trigger: ['blur', 'input'],
      validator: (rule, value: string) => {
        let error: Error | undefined;
        validate.qq(rule, value.trim(), (result?: Error) => (error = result));
        return error || true;
      },
    },
  ],
  servicePhone: textRule('联系电话', 50),
  supportUrl: [textRule('支持地址', 500), urlRule('支持地址')],
  authQueryCount: {
    required: true,
    trigger: ['blur', 'change'],
    validator: (_rule, value: unknown) =>
      typeof value === 'number' && Number.isSafeInteger(value) && value >= 0
        ? true
        : new Error('查询总次数必须是0到9007199254740991之间的整数'),
  },
};

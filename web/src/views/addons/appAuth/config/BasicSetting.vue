<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-form
        :label-width="80"
        :model="formValue"
        :rules="rules"
        :disabled="!loaded || loading || saving"
        ref="formRef"
      >
        <n-form-item label="未授权提示" path="unauthorizedMessage">
          <n-input
            v-model:value="formValue.unauthorizedMessage"
            type="textarea"
            placeholder="请输入未授权时的提示信息"
          />
        </n-form-item>

        <n-form-item label="到期提示" path="expiredMessage">
          <n-input
            v-model:value="formValue.expiredMessage"
            type="textarea"
            placeholder="请输入授权到期时的提示信息"
          />
        </n-form-item>

        <n-form-item label="禁用提示" path="disabledMessage">
          <n-input
            v-model:value="formValue.disabledMessage"
            type="textarea"
            placeholder="请输入授权被禁用时的提示信息"
          />
        </n-form-item>

        <n-form-item label="购买地址" path="purchaseUrl">
          <n-input
            v-model:value="formValue.purchaseUrl"
            placeholder="请输入购买或续费地址（选填）"
          />
        </n-form-item>

        <n-form-item label="服务商名称" path="serviceName">
          <n-input v-model:value="formValue.serviceName" placeholder="请输入服务商名称（选填）" />
        </n-form-item>

        <n-form-item label="客服QQ" path="serviceQQ">
          <n-input v-model:value="formValue.serviceQQ" placeholder="请输入客服QQ号码（选填）" />
        </n-form-item>

        <n-form-item label="联系电话" path="servicePhone">
          <n-input
            v-model:value="formValue.servicePhone"
            placeholder="请输入客服电话或座机（选填）"
          />
        </n-form-item>

        <n-form-item label="支持地址" path="supportUrl">
          <n-input v-model:value="formValue.supportUrl" placeholder="请输入技术支持地址（选填）" />
        </n-form-item>

        <n-form-item label="查询总次数" path="authQueryCount">
          <n-input-number
            v-model:value="formValue.authQueryCount"
            :min="0"
            :max="Number.MAX_SAFE_INTEGER"
            :step="1"
            placeholder="请输入累计授权查询次数"
          />
          <template #feedback>可手动调整累计次数，每次完成授权查询后自动加 1，不是查询次数上限。</template>
        </n-form-item>

        <div>
          <n-space>
            <n-button
              v-if="hasPermission(['/appAuth/config/update'])"
              type="primary"
              :loading="saving"
              :disabled="!loaded || loading || saving"
              @click="formSubmit"
            >
              保存更新
            </n-button>
            <n-button v-if="loadFailed" :disabled="loading" @click="load">重新加载</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import type { FormInst } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/addons/appAuth/config';
  import { usePermission } from '@/hooks/web/usePermission';
  import { newState, rules, toConfigList } from './model';

  const group = 'basic';
  const loading = ref(false);
  const saving = ref(false);
  const loaded = ref(false);
  const loadFailed = ref(false);
  const formRef = ref<FormInst | null>(null);
  const formValue = ref(newState());
  const loadedQueryCount = ref<number | null>(0);
  const message = useMessage();
  const { hasPermission } = usePermission();

  async function formSubmit() {
    if (
      !loaded.value ||
      loading.value ||
      saving.value ||
      !formRef.value ||
      !hasPermission(['/appAuth/config/update'])
    )
      return;
    saving.value = true;
    try {
      try {
        await formRef.value.validate();
      } catch {
        message.error('验证失败，请检查配置内容');
        return;
      }
      const list = toConfigList(formValue.value);
      const changes: Partial<typeof list> = { ...list };
      // 未主动修改次数时不提交旧值，避免覆盖页面打开后发生的查询累加。
      if (list.authQueryCount === loadedQueryCount.value) delete changes.authQueryCount;
      await updateConfig({ group, list: changes });
      formValue.value = newState(list);
      loadedQueryCount.value = list.authQueryCount;
      message.success('更新成功');
    } catch {
      message.error('保存授权配置失败，请稍后重试');
    } finally {
      saving.value = false;
    }
  }

  async function load() {
    if (loading.value || saving.value) return;
    loading.value = true;
    loaded.value = false;
    loadFailed.value = false;
    try {
      const res = await getConfig({ group });
      formValue.value = newState(res?.list);
      loadedQueryCount.value = formValue.value.authQueryCount;
      formRef.value?.restoreValidation();
      loaded.value = true;
    } catch {
      loadFailed.value = true;
      message.error('加载授权配置失败，请重新加载');
    } finally {
      loading.value = false;
    }
  }

  onMounted(load);
</script>

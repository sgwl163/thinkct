<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="API Key" path="openaiApiKey">
          <n-input
            v-model:value="formValue.openaiApiKey"
            type="password"
            show-password-on="click"
            placeholder="请输入模型服务 API Key"
          >
            <template #password-visible-icon>
              <n-icon :size="16" :component="GlassesOutline" />
            </template>
            <template #password-invisible-icon>
              <n-icon :size="16" :component="Glasses" />
            </template>
          </n-input>
          <template #feedback>用于访问 OpenAI 兼容模型服务的密钥</template>
        </n-form-item>

        <n-form-item label="接口地址" path="openaiBaseUrl">
          <n-input
            v-model:value="formValue.openaiBaseUrl"
            placeholder="例如：https://api.openai.com/v1"
          />
          <template #feedback>请填写包含版本路径的 OpenAI 兼容接口地址</template>
        </n-form-item>

        <n-form-item label="默认模型" path="openaiDefaultModel">
          <n-input
            v-model:value="formValue.openaiDefaultModel"
            placeholder="例如：gpt-3.5-turbo"
          />
          <template #feedback>聊天会话未指定模型时使用，也可填写服务商支持的其他模型名称</template>
        </n-form-item>

        <n-space>
          <n-button type="primary" @click="formSubmit">保存更新</n-button>
        </n-space>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { Glasses, GlassesOutline } from '@vicons/ionicons5';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('openai');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    openaiApiKey: '',
    openaiBaseUrl: 'https://api.gpt.ge/v1',
    openaiDefaultModel: 'gpt-3.5-turbo',
  });

  const rules = {
    openaiApiKey: {
      required: true,
      message: '请输入模型服务 API Key',
      trigger: 'blur',
    },
    openaiBaseUrl: {
      required: true,
      message: '请输入模型服务接口地址',
      trigger: 'blur',
    },
    openaiDefaultModel: {
      required: true,
      message: '请输入默认模型名称',
      trigger: 'blur',
    },
  };

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (errors) {
        message.error('验证失败，请填写完整信息');
        return;
      }
      updateConfig({ group: group.value, list: formValue.value }).then(() => {
        message.success('更新成功');
        load();
      });
    });
  }

  function load() {
    show.value = true;
    getConfig({ group: group.value })
      .then((res) => {
        formValue.value = {
          ...formValue.value,
          ...res.list,
        };
      })
      .finally(() => {
        show.value = false;
      });
  }

  onMounted(load);
</script>

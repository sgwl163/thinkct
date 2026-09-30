<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="120" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="站点名称" path="siteName">
          <n-input v-model:value="formValue.siteName" placeholder="请输入站点名称" />
        </n-form-item>

        <n-form-item label="站点标题" path="siteTitle">
          <n-input v-model:value="formValue.siteTitle" placeholder="请输入站点标题" />
        </n-form-item>

        <n-form-item label="站点关键词" path="siteKeywords">
          <n-input
            v-model:value="formValue.siteKeywords"
            placeholder="请输入站点关键词(逗号分隔)"
          />
        </n-form-item>

        <n-form-item label="站点描述" path="siteDescription">
          <n-input
            v-model:value="formValue.siteDescription"
            type="textarea"
            placeholder="请输入站点描述"
          />
        </n-form-item>

        <n-form-item label="版权所有" path="siteCopyRight">
          <n-input v-model:value="formValue.siteCopyRight" placeholder="请输入版权信息" />
        </n-form-item>

        <n-form-item label="ICP备案号" path="siteIcp">
          <n-input v-model:value="formValue.siteIcp" placeholder="请输入ICP备案号" />
        </n-form-item>

        <n-form-item label="公安备案号" path="sitePoliceRecord">
          <n-input v-model:value="formValue.sitePoliceRecord" placeholder="请输入公安备案号" />
        </n-form-item>

        <n-form-item label="邮箱地址" path="email">
          <n-input v-model:value="formValue.email" placeholder="请输入邮箱地址" />
        </n-form-item>

        <n-form-item label="联系电话" path="phone">
          <n-input v-model:value="formValue.phone" placeholder="请输入联系电话" />
        </n-form-item>

        <n-form-item label="微信二维码" path="wechatQrCode">
          <FileChooser v-model:value="formValue.wechatQrCode" file-type="image" />
        </n-form-item>

        <n-form-item label="公众号二维码" path="mpQrCode">
          <FileChooser v-model:value="formValue.mpQrCode" file-type="image" />
        </n-form-item>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import FileChooser from '@/components/FileChooser/index.vue';

  const group = ref('site');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    siteName: '',
    siteTitle: '',
    siteKeywords: '',
    siteDescription: '',
    siteCopyRight: '',
    siteIcp: '',
    sitePoliceRecord: '',
    email: '',
    phone: '',
    wechatQrCode: '',
    mpQrCode: '',
  });

  const rules = {
    siteName: {
      required: true,
      message: '请输入站点名称',
      trigger: 'blur',
    },
  };

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          message.success('更新成功');
          load();
        });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if (res.list) {
            formValue.value = Object.assign(formValue.value, res.list);
          }
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>

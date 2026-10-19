<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑授权管理 #' + formValue.id : '添加授权管理'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="formRules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="授权名称" path="name">
                  <n-input placeholder="请输入授权名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="所属产品" path="productId">
                  <n-select
                    v-model:value="formValue.productId"
                    :options="productOptions"
                    placeholder="请选择所属产品"
                    filterable
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="选择用户" path="userId">
                  <n-select
                    v-model:value="formValue.userId"
                    :options="userOptions"
                    placeholder="请选择用户，可搜索昵称"
                    filterable
                    clearable
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="授权域名" path="host">
                  <n-input
                    placeholder="请输入域名，如 example.com"
                    v-model:value="formValue.host"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="验证方式" path="authMode">
                  <n-select
                    v-model:value="formValue.authMode"
                    :options="authModeOptions"
                    placeholder="请选择验证方式"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="requiresIp(formValue.authMode)">
                <n-form-item label="授权IP" path="ip">
                  <n-input placeholder="请输入授权服务器IP" v-model:value="formValue.ip" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="域名范围" path="authDomain">
                  <n-select
                    v-model:value="formValue.authDomain"
                    :options="authDomainOptions"
                    placeholder="请选择域名范围"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="版本号" path="versionId">
                  <n-input
                    placeholder="请输入版本号，如 1.0.0"
                    v-model:value="formValue.versionId"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="授权码" path="syskey">
                  <n-input
                    placeholder="请输入授权码"
                    v-model:value="formValue.syskey"
                    type="password"
                    show-password-on="click"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="QQ" path="qq">
                  <n-input placeholder="请输入QQ号码" v-model:value="formValue.qq" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系电话" path="phone">
                  <n-input placeholder="请输入联系电话" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select
                    v-model:value="formValue.status"
                    :options="statusOptions"
                    placeholder="请选择状态"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="到期时间" path="expiredAt">
                  <DatePicker v-model:formValue="formValue.expiredAt" type="datetime" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button
            type="info"
            :loading="formBtnLoading"
            :disabled="!isFormValid"
            @click="confirmForm"
          >
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import type { FormInst, FormRules } from 'naive-ui';
  import { Edit, View } from '@/api/addons/appAuth/appAuth';
  import {
    State,
    newState,
    rules,
    productOptions,
    userOptions,
    authModeOptions,
    authDomainOptions,
    statusOptions,
    requiresIp,
  } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<FormInst | null>(null);
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = computed(() => !loading.value && !formBtnLoading.value);
  const formRules = computed<FormRules>(() => ({
    ...rules,
    ip: {
      required: requiresIp(formValue.value.authMode),
      trigger: ['blur', 'input'],
      type: 'string',
      message: '请输入授权服务器IP',
    },
  }));

  // 提交表单
  function confirmForm(e: MouseEvent) {
    e.preventDefault();
    if (!isFormValid.value) return;
    formRef.value?.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then(() => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State | null) {
    formValue.value = newState(state);
    showModal.value = true;

    if (!state || state.id < 1) {
      return;
    }

    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = newState(res);
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>

<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      :closable="!formBtnLoading"
      :close-on-esc="!formBtnLoading"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑产品 #' + formValue.id : '添加产品'"
      class="app-products-edit-modal"
    >
      <n-scrollbar class="app-products-edit-scrollbar pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="产品名称" path="name">
                  <n-input placeholder="请输入产品名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="封面" path="cover">
                  <n-space vertical class="w-full">
                    <UploadImage :maxNumber="1" v-model:value="formValue.cover" />
                    <n-input
                      placeholder="也可直接输入封面图片地址"
                      v-model:value="formValue.cover"
                      clearable
                    />
                  </n-space>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="标题" path="title">
                  <n-input placeholder="请输入标题" v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="别名" path="alias">
                  <n-input placeholder="请输入别名" v-model:value="formValue.alias" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="价格" path="price">
                  <n-input-number
                    placeholder="请输入价格，可为0"
                    v-model:value="formValue.price"
                    :min="0"
                    :max="9999999999.99"
                    :precision="2"
                    class="w-full"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="演示站" path="ysweb">
                  <n-input placeholder="请输入演示站地址" v-model:value="formValue.ysweb" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="特点" path="unique">
                  <n-input
                    type="textarea"
                    placeholder="请输入特点"
                    v-model:value="formValue.unique"
                    :autosize="{ minRows: 2, maxRows: 5 }"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="关键词" path="keywords">
                  <n-input placeholder="请输入关键词" v-model:value="formValue.keywords" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="描述" path="description">
                  <n-input
                    type="textarea"
                    placeholder="请输入描述"
                    v-model:value="formValue.description"
                    :autosize="{ minRows: 3, maxRows: 6 }"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="权重" path="weigh">
                  <n-input-number
                    placeholder="请输入权重"
                    v-model:value="formValue.weigh"
                    :precision="0"
                    class="w-full"
                  />
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
                <n-form-item label="内容" path="content">
                  <Editor
                    v-if="formReady"
                    :key="editorKey"
                    id="app-products-content"
                    v-model:value="formValue.content"
                    class="app-products-editor"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button :disabled="formBtnLoading" @click="closeForm">取消</n-button>
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
  import { computed, ref, watch } from 'vue';
  import { useMessage } from 'naive-ui';
  import type { FormInst } from 'naive-ui';
  import { Edit, View } from '@/api/addons/appAuth/appProducts';
  import { State, newState, rules, statusOptions } from './model';
  import Editor from '@/components/Editor/editor.vue';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<FormInst | null>(null);
  const formBtnLoading = ref(false);
  const formReady = ref(false);
  const editorKey = ref(0);
  let requestId = 0;
  const isFormValid = computed(() => formReady.value && !loading.value && !formBtnLoading.value);

  watch(
    showModal,
    (show) => {
      if (!show) {
        requestId++;
        formReady.value = false;
      }
    },
    { flush: 'sync' }
  );

  function confirmForm(e: MouseEvent) {
    e.preventDefault();
    if (!isFormValid.value) return;
    const currentRequest = requestId;
    formRef.value?.validate((errors) => {
      if (currentRequest !== requestId || !isFormValid.value) return;
      if (errors) {
        message.error('请填写完整信息');
        return;
      }
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
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  function openModal(state: State | null) {
    const currentRequest = ++requestId;
    editorKey.value++;
    formValue.value = newState(state);
    formRef.value?.restoreValidation();
    formReady.value = !state || state.id < 1;
    loading.value = !formReady.value;
    showModal.value = true;
    if (!state || state.id < 1) return;

    View({ id: state.id })
      .then((res) => {
        if (currentRequest !== requestId || !showModal.value) return;
        if (!res?.id) {
          message.error('产品信息不存在，请刷新列表后重试');
          return;
        }
        formValue.value = newState(res);
        formReady.value = true;
      })
      .catch(() => {
        // 请求错误由项目HTTP封装提示，加载失败时保持禁止提交。
      })
      .finally(() => {
        if (currentRequest === requestId) loading.value = false;
      });
  }

  defineExpose({ openModal });
</script>

<style lang="less">
  .n-dialog.n-modal.app-products-edit-modal {
    width: 840px;
    max-width: 100%;
  }

  .app-products-edit-scrollbar {
    max-height: 87vh;
  }

  .app-products-editor {
    height: 450px;
  }
</style>

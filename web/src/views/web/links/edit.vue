<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑友情链接 #' + formValue.id : '添加友情链接'"
      :style="{ width: dialogWidth }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="90"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:2 l:2 xl:2 2xl:2" responsive="screen" :x-gap="12">
              <n-gi span="2">
                <n-form-item label="标题" path="title">
                  <n-input placeholder="请输入标题" v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="链接" path="link">
                  <n-input placeholder="请输入完整URL，如 https://example.com" v-model:value="formValue.link" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="LOGO" path="logo">
                  <UploadImage v-model:value="formValue.logo" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="跳转方式" path="target">
                  <n-select v-model:value="formValue.target" :options="targetOptions" placeholder="请选择跳转方式" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="linksStatusOptions" placeholder="请选择状态" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="权重" path="weigh">
                  <n-input-number v-model:value="formValue.weigh" :min="0" style="width:100%" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="用户ID" path="userId">
                  <n-input-number v-model:value="formValue.userId" :min="0" style="width:100%" placeholder="留空则为匿名" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="备注" path="remark">
                  <n-input type="textarea" placeholder="请输入备注" v-model:value="formValue.remark" :rows="2" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { Edit, View } from '@/api/web/links';
  import { State, newState, rules, linksStatusOptions, targetOptions } from './model';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => adaModalWidth(600));

  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => { formBtnLoading.value = false; });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  function openModal(state: State) {
    showModal.value = true;
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      return;
    }
    loading.value = true;
    View({ id: state.id })
      .then((res) => { formValue.value = newState(res); })
      .finally(() => { loading.value = false; });
  }

  defineExpose({ openModal });
</script>

<style lang="less"></style>

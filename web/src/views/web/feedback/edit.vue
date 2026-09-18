<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '处理反馈 #' + formValue.id : '添加反馈'"
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
              <n-gi span="1">
                <n-form-item label="姓名" path="name">
                  <n-input placeholder="请输入姓名" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手机号" path="phone">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="email">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.email" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="公司名称" path="company">
                  <n-input placeholder="请输入公司名称" v-model:value="formValue.company" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="反馈内容" path="content">
                  <n-input type="textarea" :rows="4" v-model:value="formValue.content" :disabled="formValue.id > 0" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="回复内容" path="answer">
                  <n-input type="textarea" placeholder="请输入回复内容" :rows="4" v-model:value="formValue.answer" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="feedbackStatusOptions" placeholder="请选择状态" />
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
  import { Edit, View } from '@/api/web/feedback';
  import { State, newState, rules, feedbackStatusOptions } from './model';
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
  const dialogWidth = computed(() => adaModalWidth(640));

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

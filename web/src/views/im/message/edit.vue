<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑消息 #' + formValue.id : '添加消息'"
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
                <n-form-item label="聊天类型" path="chatType">
                  <n-select v-model:value="formValue.chatType" :options="chatTypeOptions" placeholder="请选择聊天类型" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="消息类型" path="msgType">
                  <n-select v-model:value="formValue.msgType" :options="msgTypeOptions" placeholder="请选择消息类型" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="发送者ID" path="fromId">
                  <n-input-number v-model:value="formValue.fromId" :min="1" style="width:100%" placeholder="请输入发送者ID" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="接收者ID" path="toId">
                  <n-input-number v-model:value="formValue.toId" :min="1" style="width:100%" placeholder="请输入接收者ID" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="消息内容" path="content">
                  <n-input type="textarea" placeholder="请输入消息内容" v-model:value="formValue.content" :rows="4" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" placeholder="请选择状态" />
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
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/im/message';
  import { State, newState, rules, chatTypeOptions, msgTypeOptions } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => adaModalWidth(620));

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

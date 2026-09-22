<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑会话列表管理 #' + formValue.id : '添加会话列表管理'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid
              cols="1 s:1 m:1 l:1 xl:1 2xl:1"
              responsive="screen"
            >
              <n-gi span="1">
                <n-form-item label="归属用户ID" path="userId">
                  <n-input-number
                    placeholder="请输入归属用户ID"
                    v-model:value="formValue.userId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="会话类型: 1=单聊, 2=群聊, 3=系统通知" path="chatType">
                  <n-input-number
                    placeholder="请输入会话类型: 1=单聊, 2=群聊, 3=系统通知"
                    v-model:value="formValue.chatType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="目标ID(好友ID或群组ID)" path="targetId">
                  <n-input-number
                    placeholder="请输入目标ID(好友ID或群组ID)"
                    v-model:value="formValue.targetId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最后一条消息ID" path="lastMsgId">
                  <n-input-number
                    placeholder="请输入最后一条消息ID"
                    v-model:value="formValue.lastMsgId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最后消息内容摘要(用于列表展示)" path="lastMsgContent">
                  <n-input
                    type="textarea"
                    placeholder="最后消息内容摘要(用于列表展示)"
                    v-model:value="formValue.lastMsgContent"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最后消息时间" path="lastMsgTime">
                  <DatePicker v-model:formValue="formValue.lastMsgTime" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="未读消息数" path="unreadCount">
                  <n-input-number
                    placeholder="请输入未读消息数"
                    v-model:value="formValue.unreadCount"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否置顶: 0=否, 1=是" path="isTop">
                  <n-input-number
                    placeholder="请输入是否置顶: 0=否, 1=是"
                    v-model:value="formValue.isTop"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="消息免打扰: 0=否, 1=是" path="isDisturb">
                  <n-input-number
                    placeholder="请输入消息免打扰: 0=否, 1=是"
                    v-model:value="formValue.isDisturb"
                    />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" :disabled="!isFormValid" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { Edit, View } from '@/api/im/conversation';
  import { State, newState, rules } from './model';
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
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = ref(true);

  // 提交表单
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
  function openModal(state: State) {
    showModal.value = true;
    
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      
      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
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
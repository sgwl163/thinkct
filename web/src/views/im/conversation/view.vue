<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="会话列表管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                归属用户ID
              </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                会话类型: 1=单聊, 2=群聊, 3=系统通知
              </template>
              {{ formValue.chatType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                目标ID(好友ID或群组ID)
              </template>
              {{ formValue.targetId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最后一条消息ID
              </template>
              {{ formValue.lastMsgId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最后消息内容摘要(用于列表展示)
              </template>
              <span v-html="formValue.lastMsgContent"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最后消息时间
              </template>
              {{ formValue.lastMsgTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                未读消息数
              </template>
              {{ formValue.unreadCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否置顶: 0=否, 1=是
              </template>
              {{ formValue.isTop }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                消息免打扰: 0=否, 1=是
              </template>
              {{ formValue.isDisturb }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/im/conversation';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const message = useMessage();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
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

<style lang="less" scoped></style>
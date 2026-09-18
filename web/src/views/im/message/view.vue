<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="IM-聊天记录表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                聊天类型: 1=单聊, 2=群聊, 3=系统消息
              </template>
              {{ formValue.chatType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                发送者ID
              </template>
              {{ formValue.fromId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                接收者ID(目标用户ID或群组ID)
              </template>
              {{ formValue.toId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包
              </template>
              {{ formValue.msgType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                消息内容(文本内容或文件URL/JSON)
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item label="状态: 1=正常, 2=已撤回, 3=已删除">
              <n-tag
            :type="dict.getType('sys_normal_disable', formValue.status)"
            size="small"
            class="min-left-space"
          >
                {{ dict.getLabel('sys_normal_disable', formValue.status) }}
              </n-tag>
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/im/message';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';

  const message = useMessage();
  const dict = useDictStore();
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

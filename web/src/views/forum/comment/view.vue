<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="评论管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                主题ID
              </template>
              {{ formValue.topicId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                评论用户ID
              </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                父评论ID
              </template>
              {{ formValue.parentId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被回复用户ID
              </template>
              {{ formValue.replyUserId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                评论内容
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                状态:normal,pending,rejected,hidden
              </template>
              {{ formValue.status }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点赞数量
              </template>
              {{ formValue.likeCount }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/forum/comment';
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
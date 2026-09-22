<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="社区论坛举报表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                举报对象类型:topic,comment
              </template>
              {{ formValue.targetType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                举报对象ID
              </template>
              {{ formValue.targetId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                举报用户ID
              </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                举报原因
              </template>
              {{ formValue.reason }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                补充说明
              </template>
              <span v-html="formValue.description"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                处理状态:pending,processed,rejected
              </template>
              {{ formValue.status }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                处理管理员ID
              </template>
              {{ formValue.handledBy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                处理时间
              </template>
              {{ formValue.handledAt }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/forum/report';
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
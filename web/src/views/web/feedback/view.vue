<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="留言管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label> 名称 </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 邮箱 </template>
              {{ formValue.email }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 手机号 </template>
              {{ formValue.phone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 公司名称 </template>
              {{ formValue.company }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 反馈内容 </template>
              {{ formValue.content }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 回复内容 </template>
              {{ formValue.answer || '暂无回复' }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> IP </template>
              {{ formValue.ip }}
            </n-descriptions-item>
            <n-descriptions-item label="处理状态">
              <n-tag :type="getStatusType(formValue.status)" size="small" class="min-left-space">
                {{ getStatusLabel(formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 创建时间 </template>
              {{ formValue.createTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 修改时间 </template>
              {{ formValue.updateTime }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/web/feedback';
  import { State, newState, getStatusLabel, getStatusType } from './model';
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

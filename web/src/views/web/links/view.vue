<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="友情链接详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label> 用户ID </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 标题 </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 链接 </template>
              {{ formValue.link }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
              </template>
              {{ formValue.target }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> LOGO </template>
              {{ formValue.logo }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 权重 </template>
              {{ formValue.weigh }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 状态:disable=禁用,enable=启用,pending_trial=待审 </template>
              {{ formValue.status }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 备注 </template>
              {{ formValue.remark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 修改时间 </template>
              {{ formValue.updateTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 创建时间 </template>
              {{ formValue.createTime }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/web/links';
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

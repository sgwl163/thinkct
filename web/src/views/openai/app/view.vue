<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="AI应用详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                分类ID
              </template>
              {{ formValue.categoryId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应用名称
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应用图标
              </template>
              {{ formValue.icon }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应用描述
              </template>
              <span v-html="formValue.description"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                系统提示词
              </template>
              <span v-html="formValue.prompt"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                使用次数
              </template>
              {{ formValue.useCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                排序
              </template>
              {{ formValue.sort }}
            </n-descriptions-item>
            <n-descriptions-item label="状态(1:启用, 2:禁用)">
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
  import { View } from '@/api/openai/app';
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

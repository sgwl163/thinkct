<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="社区板块详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                上级板块ID
              </template>
              {{ formValue.pid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                板块名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                板块简介
              </template>
              <span v-html="formValue.description"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                板块图标
              </template>
              {{ formValue.image }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                排序
              </template>
              {{ formValue.sort }}
            </n-descriptions-item>
            <n-descriptions-item label="状态:1正常,2禁用">
              <n-tag
            :type="dict.getType('sys_normal_disable', formValue.status)"
            size="small"
            class="min-left-space"
          >
                {{ dict.getLabel('sys_normal_disable', formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题数量
              </template>
              {{ formValue.topicCount }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/forum/category';
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
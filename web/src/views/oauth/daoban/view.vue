<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="盗版管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                产品ID
              </template>
              {{ formValue.productId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                域名
              </template>
              {{ formValue.host }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                IP
              </template>
              {{ formValue.ip }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                版本号
              </template>
              {{ formValue.version }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                执行代码
              </template>
              {{ formValue.code }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据
              </template>
              <span v-html="formValue.data"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                浏览量
              </template>
              {{ formValue.views }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/oauth/daoban';
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
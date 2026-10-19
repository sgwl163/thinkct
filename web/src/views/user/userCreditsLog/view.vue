<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="资产明细详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                用户ID
              </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                资产类型:money=余额,score=积分
              </template>
              {{ formValue.creditType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变动前数量
              </template>
              {{ formValue.beforeNum }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变动数量
              </template>
              {{ formValue.num }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变动后数量
              </template>
              {{ formValue.afterNum }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                备注
              </template>
              {{ formValue.remark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                关联ID
              </template>
              {{ formValue.mapId }}
            </n-descriptions-item>
            <n-descriptions-item label="状态">
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
  import { View } from '@/api/user/userCreditsLog';
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
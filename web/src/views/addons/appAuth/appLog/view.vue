<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="升级记录详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item label="所属用户">
              <n-tag
                v-if="formValue.userId"
                :type="dict.getType('AppAuthUserOptions', formValue.userId)"
                size="small"
                :bordered="false"
              >
                {{ getUserLabel(formValue.userId) }}
              </n-tag>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="所属产品">
              <n-tag
                v-if="formValue.productId"
                :type="dict.getType('AppAuthProductOptions', formValue.productId)"
                size="small"
                :bordered="false"
              >
                {{ getProductLabel(formValue.productId) }}
              </n-tag>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                域名
              </template>
              {{ formValue.host }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                内容
              </template>
              <span v-html="formValue.content"></span>
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
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/appAuth/appLog';
  import { State, newState } from './model';
  import { getProductLabel, getUserLabel } from '../appAuth/model';
  import { useDictStore } from '@/store/modules/dict';
  import { adaModalWidth } from '@/utils/hotgo';

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
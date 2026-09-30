<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="授权管理详情" closable>
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
                产品ID
              </template>
              {{ formValue.productId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                名称
              </template>
              {{ formValue.name }}
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
                QQ
              </template>
              {{ formValue.qq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                电话
              </template>
              {{ formValue.phone }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                授权码
              </template>
              {{ formValue.syskey }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                版本
              </template>
              {{ formValue.versionId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                验证方式:0=单域名验证,1=IP双重验证
              </template>
              {{ formValue.authMode }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                验证域名:0=当前单域名方式,1=顶级泛域名方式
              </template>
              {{ formValue.authDomain }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                浏览量
              </template>
              {{ formValue.views }}
            </n-descriptions-item>
            <n-descriptions-item label="状态:1=启用,0=禁用">
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
                到期时间
              </template>
              {{ formValue.expiredAt }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/app/appAuth';
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
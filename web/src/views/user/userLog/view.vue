<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="操作记录详情" closable>
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
                用户名
              </template>
              {{ formValue.username }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                记录类型:login=登录,operation=操作
              </template>
              {{ formValue.logType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                操作标识
              </template>
              {{ formValue.action }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                操作名称
              </template>
              {{ formValue.actionName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                事件时间
              </template>
              {{ formValue.eventAt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                IP地址
              </template>
              {{ formValue.ip }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                UA信息
              </template>
              <span v-html="formValue.userAgent"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                结果消息
              </template>
              <span v-html="formValue.message"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                扩展数据
              </template>
              <span v-html="formValue.extra"></span>
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
  import { View } from '@/api/user/userLog';
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
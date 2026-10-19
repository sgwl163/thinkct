<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="产品管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                封面
              </template>
              {{ formValue.cover }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                关键词
              </template>
              {{ formValue.keywords }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                描述
              </template>
              <span v-html="formValue.description"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                别名
              </template>
              {{ formValue.alias }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                内容
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                价格
              </template>
              {{ formValue.price }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                演示站
              </template>
              {{ formValue.ysweb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                特点
              </template>
              {{ formValue.unique }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                权重
              </template>
              {{ formValue.weigh }}
            </n-descriptions-item>
            <n-descriptions-item label="状态:0=禁用,1=启用">
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
  import { View } from '@/api/addons/appAuth/appProducts';
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
<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="标签管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标签类型:default=default,success=success,info=info,warning=warning,danger=danger
              </template>
              {{ formValue.type }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                文档数
              </template>
              {{ formValue.documentCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                SEO标题
              </template>
              {{ formValue.seotitle }}
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
              {{ formValue.description }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                自动内链:0=关,1=开
              </template>
              {{ formValue.innerChain }}
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
            <n-descriptions-item>
              <template #label>
                修改时间
              </template>
              {{ formValue.updateTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                创建时间
              </template>
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
  import { View } from '@/api/web/tags';
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

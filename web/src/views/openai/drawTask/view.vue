<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="AI绘画任务详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1" bordered>
            <n-descriptions-item label="任务ID">{{ formValue.id }}</n-descriptions-item>
            <n-descriptions-item label="用户ID">{{ formValue.userId }}</n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag :type="getStatusType(formValue.status)" size="small">
                {{ getStatusLabel(formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="进度">
              <n-progress
                type="line"
                :percentage="formValue.progress"
                :status="formValue.status === 2 ? 'error' : formValue.status === 1 ? 'success' : 'default'"
              />
            </n-descriptions-item>
            <n-descriptions-item label="结果图" v-if="formValue.imageUrl">
              <n-image :src="formValue.imageUrl" width="200" object-fit="contain" />
            </n-descriptions-item>
            <n-descriptions-item label="参考图" v-if="formValue.referenceImage">
              <n-image :src="formValue.referenceImage" width="120" object-fit="contain" />
            </n-descriptions-item>
            <n-descriptions-item label="使用的模型">{{ formValue.model }}</n-descriptions-item>
            <n-descriptions-item label="模型版本">{{ formValue.version }}</n-descriptions-item>
            <n-descriptions-item label="图片比例">{{ formValue.ratio }}</n-descriptions-item>
            <n-descriptions-item label="RAW模式">{{ formValue.raw === 1 ? '开启' : '关闭' }}</n-descriptions-item>
            <n-descriptions-item label="绘画提示词">
              <span style="white-space: pre-wrap">{{ formValue.prompt }}</span>
            </n-descriptions-item>
            <n-descriptions-item label="负面提示词" v-if="formValue.negativePrompt">
              <span style="white-space: pre-wrap">{{ formValue.negativePrompt }}</span>
            </n-descriptions-item>
            <n-descriptions-item label="失败原因" v-if="formValue.failReason">{{ formValue.failReason }}</n-descriptions-item>
            <n-descriptions-item label="风格化">{{ formValue.stylize }}</n-descriptions-item>
            <n-descriptions-item label="多样性">{{ formValue.chaos }}</n-descriptions-item>
            <n-descriptions-item label="随机种子">{{ formValue.seed }}</n-descriptions-item>
            <n-descriptions-item label="MJ任务ID" v-if="formValue.mjTaskId">{{ formValue.mjTaskId }}</n-descriptions-item>
            <n-descriptions-item label="创建时间">{{ formValue.createdAt }}</n-descriptions-item>
            <n-descriptions-item label="更新时间">{{ formValue.updatedAt }}</n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { View } from '@/api/openai/drawTask';
  import { State, newState, getStatusLabel, getStatusType } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => adaModalWidth(620));

  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = newState(res);
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({ openModal });
</script>

<style lang="less" scoped></style>

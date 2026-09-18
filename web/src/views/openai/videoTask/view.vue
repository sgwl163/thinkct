<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="AI视频任务详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" :column="1" bordered class="py-2">
            <n-descriptions-item label="任务ID">{{ formValue.id }}</n-descriptions-item>
            <n-descriptions-item label="用户ID">{{ formValue.userId }}</n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag :type="getStatusType(formValue.status)" size="small">{{ getStatusLabel(formValue.status) }}</n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="使用的模型">{{ formValue.model }}</n-descriptions-item>
            <n-descriptions-item label="首帧参考图" v-if="formValue.referenceImage">
              <n-image :src="formValue.referenceImage" width="200" object-fit="contain" />
            </n-descriptions-item>
            <n-descriptions-item label="尾帧图片" v-if="formValue.lastFrame">
              <n-image :src="formValue.lastFrame" width="200" object-fit="contain" />
            </n-descriptions-item>
            <n-descriptions-item label="视频提示词">
              <span style="white-space: pre-wrap">{{ formValue.prompt }}</span>
            </n-descriptions-item>
            <n-descriptions-item label="配音音频" v-if="formValue.voiceDubUrl">
              <audio :src="formValue.voiceDubUrl" controls style="width:100%;max-width:360px" />
            </n-descriptions-item>
            <n-descriptions-item label="背景音乐" v-if="formValue.bgMusicUrl">
              <audio :src="formValue.bgMusicUrl" controls style="width:100%;max-width:360px" />
            </n-descriptions-item>
            <n-descriptions-item label="生成同步音频">
              {{ formValue.generateAudio === 1 ? '开启' : '关闭' }}
            </n-descriptions-item>
            <n-descriptions-item label="生成视频" v-if="formValue.videoUrl">
              <video :src="formValue.videoUrl" controls style="width:100%;max-width:480px;max-height:320px" />
            </n-descriptions-item>
            <n-descriptions-item label="失败原因" v-if="formValue.failReason">{{ formValue.failReason }}</n-descriptions-item>
            <n-descriptions-item label="远端任务ID" v-if="formValue.remoteTaskId">{{ formValue.remoteTaskId }}</n-descriptions-item>
            <n-descriptions-item label="创建时间">{{ formValue.createdAt }}</n-descriptions-item>
            <n-descriptions-item label="更新时间">{{ formValue.updatedAt }}</n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { View } from '@/api/openai/videoTask';
  import { State, newState, getStatusLabel, getStatusType } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => adaModalWidth(660));

  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => { formValue.value = newState(res); })
      .finally(() => { loading.value = false; });
  }

  defineExpose({ openModal });
</script>

<style lang="less" scoped></style>

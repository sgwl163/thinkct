<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="AI音乐任务详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" :column="1" bordered class="py-2">
            <n-descriptions-item label="任务ID">{{ formValue.id }}</n-descriptions-item>
            <n-descriptions-item label="用户ID">{{ formValue.userId }}</n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag :type="getStatusType(formValue.status)" size="small">{{ getStatusLabel(formValue.status) }}</n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="封面图" v-if="formValue.imageUrl">
              <n-image :src="formValue.imageUrl" width="120" object-fit="contain" />
            </n-descriptions-item>
            <n-descriptions-item label="歌名" v-if="formValue.title">{{ formValue.title }}</n-descriptions-item>
            <n-descriptions-item label="创作模式">{{ formValue.type === 1 ? '灵感模式' : '定制模式' }}</n-descriptions-item>
            <n-descriptions-item label="纯音乐">{{ formValue.makeInstrumental === 1 ? '是' : '否' }}</n-descriptions-item>
            <n-descriptions-item label="风格标签" v-if="formValue.tags">{{ formValue.tags }}</n-descriptions-item>
            <n-descriptions-item label="使用的模型">{{ formValue.model }}</n-descriptions-item>
            <n-descriptions-item label="提示词">
              <span style="white-space: pre-wrap">{{ formValue.prompt }}</span>
            </n-descriptions-item>
            <n-descriptions-item label="时长">{{ formValue.duration ? formValue.duration + ' 秒' : '-' }}</n-descriptions-item>
            <n-descriptions-item label="音频" v-if="formValue.audioUrl">
              <audio :src="formValue.audioUrl" controls style="width:100%;max-width:420px" />
            </n-descriptions-item>
            <n-descriptions-item label="视频" v-if="formValue.videoUrl">
              <video :src="formValue.videoUrl" controls style="width:100%;max-width:420px;max-height:280px" />
            </n-descriptions-item>
            <n-descriptions-item label="失败原因" v-if="formValue.failReason">{{ formValue.failReason }}</n-descriptions-item>
            <n-descriptions-item label="远端任务ID" v-if="formValue.remoteTaskId">{{ formValue.remoteTaskId }}</n-descriptions-item>
            <n-descriptions-item label="远端歌曲ID" v-if="formValue.remoteClipId">{{ formValue.remoteClipId }}</n-descriptions-item>
            <n-descriptions-item label="创建时间">{{ formValue.createdAt }}</n-descriptions-item>
            <n-descriptions-item label="更新时间">{{ formValue.updatedAt }}</n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { View } from '@/api/openai/musicTask';
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
      .then((res) => { formValue.value = newState(res); })
      .finally(() => { loading.value = false; });
  }

  defineExpose({ openModal });
</script>

<style lang="less" scoped></style>

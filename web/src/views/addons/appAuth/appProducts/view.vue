<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="产品详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item label="ID">{{ formValue.id }}</n-descriptions-item>
            <n-descriptions-item label="产品名称">
              {{ formValue.name || '-' }}
            </n-descriptions-item>
            <n-descriptions-item label="封面">
              <n-image
                v-if="formValue.cover"
                :src="formValue.cover"
                :fallback-src="fallbackSrc()"
                :width="100"
                :height="100"
                object-fit="cover"
              />
              <span v-else>-</span>
            </n-descriptions-item>
            <n-descriptions-item label="标题">{{ formValue.title || '-' }}</n-descriptions-item>
            <n-descriptions-item label="别名">{{ formValue.alias || '-' }}</n-descriptions-item>
            <n-descriptions-item label="价格">
              {{ formatPrice(formValue.price) }}
            </n-descriptions-item>
            <n-descriptions-item label="演示站">{{ formValue.ysweb || '-' }}</n-descriptions-item>
            <n-descriptions-item label="特点">
              <span class="app-products-text">{{ formValue.unique || '-' }}</span>
            </n-descriptions-item>
            <n-descriptions-item label="关键词">
              {{ formValue.keywords || '-' }}
            </n-descriptions-item>
            <n-descriptions-item label="描述">
              <div
                v-if="formValue.description"
                class="app-products-content"
                v-html="formValue.description"
              ></div>
              <span v-else>-</span>
            </n-descriptions-item>
            <n-descriptions-item label="权重">{{ formValue.weigh }}</n-descriptions-item>
            <n-descriptions-item label="状态">
              <n-tag
                :type="dict.getType('sys_normal_disable', formValue.status)"
                size="small"
                :bordered="false"
              >
                {{ dict.getLabel('sys_normal_disable', formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="创建时间">
              {{ formatDateTime(formValue.createdAt) }}
            </n-descriptions-item>
            <n-descriptions-item label="更新时间">
              {{ formatDateTime(formValue.updatedAt) }}
            </n-descriptions-item>
            <n-descriptions-item label="内容">
              <div
                v-if="formValue.content"
                class="app-products-content"
                v-html="formValue.content"
              ></div>
              <span v-else>-</span>
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { computed, ref, watch } from 'vue';
  import { View } from '@/api/addons/appAuth/appProducts';
  import { State, newState, formatPrice, formatDateTime } from './model';
  import { adaModalWidth, fallbackSrc } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => adaModalWidth(580));
  let requestId = 0;

  watch(
    showModal,
    (show) => {
      if (!show) requestId++;
    },
    { flush: 'sync' }
  );

  function openModal(state: State) {
    const currentRequest = ++requestId;
    formValue.value = newState(state);
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        if (currentRequest === requestId && showModal.value) {
          formValue.value = newState(res);
        }
      })
      .catch(() => {
        // 请求错误由项目HTTP封装提示。
      })
      .finally(() => {
        if (currentRequest === requestId) loading.value = false;
      });
  }

  defineExpose({ openModal });
</script>

<style lang="less" scoped>
  .app-products-text {
    white-space: pre-wrap;
    overflow-wrap: anywhere;
  }

  .app-products-content {
    overflow-wrap: anywhere;

    :deep(img),
    :deep(video) {
      max-width: 100%;
      height: auto;
    }
  }
</style>

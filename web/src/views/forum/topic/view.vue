<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="社区主题详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                板块ID
              </template>
              {{ formValue.categoryId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                发布用户ID
              </template>
              {{ formValue.userId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题关键词，英文逗号分隔
              </template>
              <span v-html="formValue.keywords"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题摘要
              </template>
              <span v-html="formValue.description"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题内容
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主题图片，英文逗号分隔
              </template>
              <span v-html="formValue.images"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                状态:normal,pending,rejected,hidden,closed
              </template>
              {{ formValue.status }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否置顶:0否,1是
              </template>
              {{ formValue.isTop }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否精华:0否,1是
              </template>
              {{ formValue.isEssence }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                评论数量
              </template>
              {{ formValue.commentCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点赞数量
              </template>
              {{ formValue.likeCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                收藏数量
              </template>
              {{ formValue.favoriteCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                浏览数量
              </template>
              {{ formValue.viewCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最后回复用户ID
              </template>
              {{ formValue.lastReplyUserId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最后回复时间
              </template>
              {{ formValue.lastReplyAt }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/forum/topic';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const message = useMessage();
  
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
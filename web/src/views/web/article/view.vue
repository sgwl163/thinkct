<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="文章管理详情" closable>
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
                管理员ID
              </template>
              {{ formValue.adminId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                频道
              </template>
              {{ formValue.channelId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                副频道
              </template>
              {{ formValue.channelIds }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                自定义名称
              </template>
              {{ formValue.diyname }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题样式
              </template>
              {{ formValue.titleStyle }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标志:top=置顶,hot=热门,recommend=推荐,new=最新
              </template>
              {{ formValue.flag }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                封面图片
              </template>
              <span v-html="formValue.images"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                内容
              </template>
              <span v-html="formValue.content"></span>
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
                标签
              </template>
              {{ formValue.tags }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                外部链接
              </template>
              <span v-html="formValue.url"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
              </template>
              {{ formValue.target }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                浏览量
              </template>
              {{ formValue.views }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                评论量
              </template>
              {{ formValue.comments }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点赞数
              </template>
              {{ formValue.likes }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点踩数
              </template>
              {{ formValue.dislikes }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                权重
              </template>
              {{ formValue.weigh }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线
              </template>
              {{ formValue.status }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                发布时间
              </template>
              {{ formValue.publishTime }}
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
  import { View } from '@/api/web/article';
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

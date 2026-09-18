<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="栏目管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                上级频道
              </template>
              {{ formValue.pid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                模型
              </template>
              {{ formValue.model }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                类型:cover=封面频道,list=列表频道,link=跳转链接
              </template>
              {{ formValue.type }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                频道名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                自定义名称
              </template>
              {{ formValue.diyname }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                模板地址
              </template>
              {{ formValue.template }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                跳转链接
              </template>
              {{ formValue.url }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
              </template>
              {{ formValue.target }}
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
                访问策略:all=全部,user=仅会员
              </template>
              {{ formValue.allowVisitGroups }}
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
  import { View } from '@/api/web/channel';
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

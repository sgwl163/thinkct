<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="权限规则详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                上级菜单
              </template>
              {{ formValue.pid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
              </template>
              {{ formValue.type }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                规则名称
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                路由路径
              </template>
              {{ formValue.path }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                图标
              </template>
              {{ formValue.icon }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                菜单类型:tab=选项卡,link=链接,iframe=Iframe
              </template>
              {{ formValue.menuType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                Url
              </template>
              {{ formValue.url }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                组件路径
              </template>
              {{ formValue.component }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                未登录有效:0=否,1=是
              </template>
              {{ formValue.noLoginValid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
              </template>
              {{ formValue.extend }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                备注
              </template>
              {{ formValue.remark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                权重
              </template>
              {{ formValue.weigh }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                状态:0=禁用,1=启用
              </template>
              {{ formValue.status }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/user/userRule';
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
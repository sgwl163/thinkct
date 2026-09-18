<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="会员管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label> 分组ID </template>
              {{ formValue.groupId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 用户名 </template>
              {{ formValue.username }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 昵称 </template>
              {{ formValue.nickname }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 邮箱 </template>
              {{ formValue.email }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 手机 </template>
              {{ formValue.mobile }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 头像 </template>
              <n-image
                style="margin-left: 10px; height: 100px; width: 100px"
                :src="formValue.avatar"
              />
            </n-descriptions-item>
            <n-descriptions-item label="性别:0=未知,1=男,2=女">
              <n-tag
                :type="dict.getType('sys_user_sex', formValue.gender)"
                size="small"
                class="min-left-space"
              >
                {{ dict.getLabel('sys_user_sex', formValue.gender) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 生日 </template>
              {{ formValue.birthday }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 余额 </template>
              {{ formValue.money }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 积分 </template>
              {{ formValue.score }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 加入时间 </template>
              {{ formValue.joinTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 签名 </template>
              <span v-html="formValue.motto"></span>
            </n-descriptions-item>
            <n-descriptions-item label="状态:enable=启用,disable=禁用">
              <n-switch
                v-model:value="formValue.status"
                :unchecked-value="2"
                :checked-value="1"
                :disabled="true"
              />
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label> 创建时间 </template> {{ formValue.createTime }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/user/user';
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

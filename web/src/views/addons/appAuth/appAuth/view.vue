<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="授权管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item label="ID">{{ formValue.id }}</n-descriptions-item>
            <n-descriptions-item label="授权名称">{{ formValue.name }}</n-descriptions-item>
            <n-descriptions-item label="所属产品">
              <n-tag
                v-if="formValue.productId"
                :type="dict.getType('AppAuthProductOptions', formValue.productId)"
                size="small"
                :bordered="false"
              >
                {{ getProductLabel(formValue.productId) }}
              </n-tag>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="所属用户">
              <n-tag
                v-if="formValue.userId"
                :type="dict.getType('AppAuthUserOptions', formValue.userId)"
                size="small"
                :bordered="false"
              >
                {{ getUserLabel(formValue.userId) }}
              </n-tag>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="授权域名">
              <n-text code v-if="formValue.host">{{ formValue.host }}</n-text>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="验证方式">
              <n-tag
                :type="dict.getType('AppAuthModeOptions', formValue.authMode)"
                size="small"
                :bordered="false"
              >
                {{ dict.getLabel('AppAuthModeOptions', formValue.authMode) || '-' }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="授权IP" v-if="requiresIp(formValue.authMode)">
              <n-text code v-if="formValue.ip">{{ formValue.ip }}</n-text>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="域名范围">
              <n-tag
                :type="dict.getType('AppAuthDomainOptions', formValue.authDomain)"
                size="small"
                :bordered="false"
              >
                {{ dict.getLabel('AppAuthDomainOptions', formValue.authDomain) || '-' }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="版本号">
              <n-tag v-if="formValue.versionId" size="small" :bordered="false">{{
                formValue.versionId
              }}</n-tag>
              <template v-else>-</template>
            </n-descriptions-item>
            <n-descriptions-item label="授权码">
              <n-input
                :value="formValue.syskey"
                type="password"
                show-password-on="click"
                placeholder="未设置授权码"
                readonly
              />
            </n-descriptions-item>
            <n-descriptions-item label="QQ">{{ formValue.qq || '-' }}</n-descriptions-item>
            <n-descriptions-item label="联系电话">{{ formValue.phone || '-' }}</n-descriptions-item>
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
            <n-descriptions-item label="到期时间">
              {{ formatDateTime(formValue.expiredAt) }}
            </n-descriptions-item>
            <n-descriptions-item label="浏览量">{{ formValue.views }}</n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { View } from '@/api/addons/appAuth/appAuth';
  import {
    State,
    newState,
    getProductLabel,
    getUserLabel,
    requiresIp,
    formatDateTime,
  } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  // 打开模态框
  function openModal(state: State) {
    formValue.value = newState(state);
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

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>

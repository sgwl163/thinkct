<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑授权管理 #' + formValue.id : '添加授权管理'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid
              cols="1 s:1 m:1 l:1 xl:1 2xl:1"
              responsive="screen"
            >
              <n-gi span="1">
                <n-form-item label="用户ID" path="userId">
                  <n-input-number
                    placeholder="请输入用户ID"
                    v-model:value="formValue.userId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="产品ID" path="productId">
                  <n-input-number
                    placeholder="请输入产品ID"
                    v-model:value="formValue.productId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="名称" path="name">
                  <n-input
                    placeholder="请输入名称"
                    v-model:value="formValue.name"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="域名" path="host">
                  <n-input
                    placeholder="请输入域名"
                    v-model:value="formValue.host"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="IP" path="ip">
                  <n-input
                    placeholder="请输入IP"
                    v-model:value="formValue.ip"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="QQ" path="qq">
                  <n-input
                    placeholder="请输入QQ"
                    v-model:value="formValue.qq"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="电话" path="phone">
                  <n-input
                    placeholder="请输入电话"
                    v-model:value="formValue.phone"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="授权码" path="syskey">
                  <n-input
                    placeholder="请输入授权码"
                    v-model:value="formValue.syskey"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="版本" path="versionId">
                  <n-input
                    placeholder="请输入版本"
                    v-model:value="formValue.versionId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="验证方式:0=单域名验证,1=IP双重验证" path="authMode">
                  <n-input
                    placeholder="请输入验证方式:0=单域名验证,1=IP双重验证"
                    v-model:value="formValue.authMode"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="验证域名:0=当前单域名方式,1=顶级泛域名方式" path="authDomain">
                  <n-input
                    placeholder="请输入验证域名:0=当前单域名方式,1=顶级泛域名方式"
                    v-model:value="formValue.authDomain"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="浏览量" path="views">
                  <n-input-number
                    placeholder="请输入浏览量"
                    v-model:value="formValue.views"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态:1=启用,0=禁用" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="创建时间" path="createTime">
                  <n-input-number
                    placeholder="请输入创建时间"
                    v-model:value="formValue.createTime"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="修改时间" path="updateTime">
                  <n-input-number
                    placeholder="请输入修改时间"
                    v-model:value="formValue.updateTime"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="到期时间" path="dueTime">
                  <n-input-number
                    placeholder="请输入到期时间"
                    v-model:value="formValue.dueTime"
                    />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" :disabled="!isFormValid" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/oauth/auth';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = ref(true);

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      
      return;
    }

    // 编辑
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

<style lang="less"></style>
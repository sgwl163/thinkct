<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑资产明细 #' + formValue.id : '添加资产明细'"
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
                <n-form-item label="资产类型:money=余额,score=积分" path="creditType">
                  <n-input
                    placeholder="请输入资产类型:money=余额,score=积分"
                    v-model:value="formValue.creditType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="变动前数量" path="beforeNum">
                  <n-input-number
                    placeholder="请输入变动前数量"
                    v-model:value="formValue.beforeNum"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="变动数量" path="num">
                  <n-input-number
                    placeholder="请输入变动数量"
                    v-model:value="formValue.num"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="变动后数量" path="afterNum">
                  <n-input-number
                    placeholder="请输入变动后数量"
                    v-model:value="formValue.afterNum"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <n-input
                    placeholder="请输入备注"
                    v-model:value="formValue.remark"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="关联ID" path="mapId">
                  <n-input-number
                    placeholder="请输入关联ID"
                    v-model:value="formValue.mapId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" />
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
  import { Edit, View } from '@/api/user/userCreditsLog';
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
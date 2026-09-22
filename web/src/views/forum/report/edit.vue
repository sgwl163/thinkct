<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑社区论坛举报表 #' + formValue.id : '添加社区论坛举报表'"
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
                <n-form-item label="举报对象类型:topic,comment" path="targetType">
                  <n-input
                    placeholder="请输入举报对象类型:topic,comment"
                    v-model:value="formValue.targetType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="举报对象ID" path="targetId">
                  <n-input-number
                    placeholder="请输入举报对象ID"
                    v-model:value="formValue.targetId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="举报用户ID" path="userId">
                  <n-input-number
                    placeholder="请输入举报用户ID"
                    v-model:value="formValue.userId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="举报原因" path="reason">
                  <n-input
                    placeholder="请输入举报原因"
                    v-model:value="formValue.reason"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="补充说明" path="description">
                  <Editor style="height: 450px" id="description" v-model:value="formValue.description" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="处理状态:pending,processed,rejected" path="status">
                  <n-input
                    placeholder="请输入处理状态:pending,processed,rejected"
                    v-model:value="formValue.status"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="处理管理员ID" path="handledBy">
                  <n-input-number
                    placeholder="请输入处理管理员ID"
                    v-model:value="formValue.handledBy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="处理时间" path="handledAt">
                  <DatePicker v-model:formValue="formValue.handledAt" type="datetime" />
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
  import { Edit, View } from '@/api/forum/report';
  import { State, newState, rules } from './model';
  import Editor from '@/components/Editor/editor.vue';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  
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
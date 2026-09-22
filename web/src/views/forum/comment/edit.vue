<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑评论管理 #' + formValue.id : '添加评论管理'"
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
                <n-form-item label="主题ID" path="topicId">
                  <n-input-number
                    placeholder="请输入主题ID"
                    v-model:value="formValue.topicId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="评论用户ID" path="userId">
                  <n-input-number
                    placeholder="请输入评论用户ID"
                    v-model:value="formValue.userId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="父评论ID" path="parentId">
                  <n-input-number
                    placeholder="请输入父评论ID"
                    v-model:value="formValue.parentId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被回复用户ID" path="replyUserId">
                  <n-input-number
                    placeholder="请输入被回复用户ID"
                    v-model:value="formValue.replyUserId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="评论内容" path="content">
                  <Editor style="height: 450px" id="content" v-model:value="formValue.content" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态:normal,pending,rejected,hidden" path="status">
                  <n-input
                    placeholder="请输入状态:normal,pending,rejected,hidden"
                    v-model:value="formValue.status"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="点赞数量" path="likeCount">
                  <n-input-number
                    placeholder="请输入点赞数量"
                    v-model:value="formValue.likeCount"
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
  import { Edit, View } from '@/api/forum/comment';
  import { State, newState, rules } from './model';
  import Editor from '@/components/Editor/editor.vue';
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
<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑版本更新 #' + formValue.id : '添加版本更新'"
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
                <n-form-item label="产品ID" path="productId">
                  <n-input-number
                    placeholder="请输入产品ID"
                    v-model:value="formValue.productId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="封面" path="cover">
                  <n-input
                    placeholder="请输入封面"
                    v-model:value="formValue.cover"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="标题" path="title">
                  <n-input
                    placeholder="请输入标题"
                    v-model:value="formValue.title"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="详情" path="content">
                  <Editor style="height: 450px" id="content" v-model:value="formValue.content" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="版本号" path="version">
                  <n-input
                    placeholder="请输入版本号"
                    v-model:value="formValue.version"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="文件" path="file">
                  <n-input
                    placeholder="请输入文件"
                    v-model:value="formValue.file"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态:1=启用,0=禁用" path="status">
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
  import { Edit, View } from '@/api/addons/appAuth/appVersion';
  import { State, newState, rules } from './model';
  import Editor from '@/components/Editor/editor.vue';
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
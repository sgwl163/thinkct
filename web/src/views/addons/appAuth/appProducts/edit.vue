<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑产品管理 #' + formValue.id : '添加产品管理'"
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
                <n-form-item label="关键词" path="keywords">
                  <n-input
                    placeholder="请输入关键词"
                    v-model:value="formValue.keywords"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="描述" path="description">
                  <n-input
                    type="textarea"
                    placeholder="描述"
                    v-model:value="formValue.description"
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
                <n-form-item label="别名" path="alias">
                  <n-input
                    placeholder="请输入别名"
                    v-model:value="formValue.alias"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="内容" path="content">
                  <Editor style="height: 450px" id="content" v-model:value="formValue.content" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="价格" path="price">
                  <n-input-number
                    placeholder="请输入价格"
                    v-model:value="formValue.price"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="演示站" path="ysweb">
                  <n-input
                    placeholder="请输入演示站"
                    v-model:value="formValue.ysweb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="特点" path="unique">
                  <n-input
                    placeholder="请输入特点"
                    v-model:value="formValue.unique"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="权重" path="weigh">
                  <n-input-number
                    placeholder="请输入权重"
                    v-model:value="formValue.weigh"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态:0=禁用,1=启用" path="status">
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
  import { Edit, View } from '@/api/addons/appAuth/appProducts';
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
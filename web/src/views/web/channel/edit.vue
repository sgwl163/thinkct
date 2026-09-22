<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑栏目 #' + formValue.id : '添加栏目'"
      :style="{ width: dialogWidth }"
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
            <n-grid cols="1 s:1 m:2 l:2 xl:2 2xl:2" responsive="screen" :x-gap="12">
              <n-gi span="2">
                <n-form-item label="上级频道" path="pid">
                  <n-tree-select
                    :options="treeOption"
                    v-model:value="formValue.pid"
                    key-field="id"
                    label-field="name"
                    clearable
                    filterable
                    default-expand-all
                    show-path
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="频道名称" path="name">
                  <n-input placeholder="请输入频道名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="自定义名称" path="diyname">
                  <n-input placeholder="请输入自定义名称" v-model:value="formValue.diyname" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="类型" path="type">
                  <n-select v-model:value="formValue.type" :options="channelTypeOptions" placeholder="请选择类型" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="模型" path="model" required>
                  <n-select v-model:value="formValue.model" :options="modelOptions" placeholder="请选择模型" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="模板地址" path="template" required>
                  <n-input placeholder="请输入模板地址" v-model:value="formValue.template" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="SEO标题" path="seotitle">
                  <n-input placeholder="请输入SEO标题" v-model:value="formValue.seotitle" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="关键词" path="keywords">
                  <n-input placeholder="请输入关键词，多个用逗号分隔" v-model:value="formValue.keywords" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="描述" path="description">
                  <n-input type="textarea" placeholder="请输入描述" v-model:value="formValue.description" :rows="3" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="跳转链接" path="url">
                  <n-input placeholder="请输入跳转链接（类型为跳转链接时填写）" v-model:value="formValue.url" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="跳转方式" path="target">
                  <n-select v-model:value="formValue.target" :options="targetOptions" placeholder="请选择跳转方式" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="访问策略" path="allowVisitGroups">
                  <n-select v-model:value="formValue.allowVisitGroups" :options="visitGroupOptions" placeholder="请选择访问策略" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="权重" path="weigh">
                  <n-input-number v-model:value="formValue.weigh" :min="0" style="width:100%" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" placeholder="请选择状态" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/web/channel';
  import { State, newState, treeOption, loadTreeOption, rules, channelTypeOptions, modelOptions, targetOptions, visitGroupOptions } from './model';
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
  const dialogWidth = computed(() => adaModalWidth(760));

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
          .finally(() => { formBtnLoading.value = false; });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  function openModal(state: State) {
    showModal.value = true;
    loadTreeOption();
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      return;
    }
    loading.value = true;
    View({ id: state.id })
      .then((res) => { formValue.value = newState(res); })
      .finally(() => { loading.value = false; });
  }

  defineExpose({ openModal });
</script>

<style lang="less"></style>

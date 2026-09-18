<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑绘画任务 #' + formValue.id : '添加绘画任务'"
      :style="{ width: dialogWidth }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="110"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:2 l:2 xl:2 2xl:2" responsive="screen" :x-gap="12">
              <n-gi span="2">
                <n-form-item label="绘画提示词" path="prompt">
                  <n-input
                    type="textarea"
                    placeholder="请输入绘画提示词（英文效果更佳）"
                    v-model:value="formValue.prompt"
                    :rows="4"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="负面提示词" path="negativePrompt">
                  <n-input
                    type="textarea"
                    placeholder="请输入负面提示词（不希望出现的内容）"
                    v-model:value="formValue.negativePrompt"
                    :rows="2"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="使用的模型" path="model">
                  <n-input placeholder="如: midjourney" v-model:value="formValue.model" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="模型版本" path="version">
                  <n-input placeholder="如: v6" v-model:value="formValue.version" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="图片比例" path="ratio">
                  <n-select
                    v-model:value="formValue.ratio"
                    :options="ratioOptions"
                    placeholder="请选择图片比例"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="RAW模式" path="raw">
                  <n-switch v-model:value="formValue.raw" :checked-value="1" :unchecked-value="0">
                    <template #checked>开启</template>
                    <template #unchecked>关闭</template>
                  </n-switch>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="风格化(0-1000)" path="stylize">
                  <n-slider v-model:value="formValue.stylize" :min="0" :max="1000" :step="50" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="多样性(0-100)" path="chaos">
                  <n-slider v-model:value="formValue.chaos" :min="0" :max="100" :step="5" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="随机种子" path="seed">
                  <n-input-number placeholder="留空随机" v-model:value="formValue.seed" :min="0" style="width:100%" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="参考图" path="referenceImage">
                  <UploadImage v-model:value="formValue.referenceImage" />
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
  import { Edit, View } from '@/api/openai/drawTask';
  import { State, newState, rules } from './model';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const ratioOptions = [
    { label: '1:1', value: '1:1' },
    { label: '4:3', value: '4:3' },
    { label: '3:4', value: '3:4' },
    { label: '16:9', value: '16:9' },
    { label: '9:16', value: '9:16' },
    { label: '2:3', value: '2:3' },
    { label: '3:2', value: '3:2' },
  ];

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
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
          .finally(() => {
            formBtnLoading.value = false;
          });
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
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      return;
    }
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = newState(res);
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({ openModal });
</script>

<style lang="less"></style>

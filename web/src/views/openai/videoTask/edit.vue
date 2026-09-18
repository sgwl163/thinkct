<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑AI视频任务 #' + formValue.id : '添加AI视频任务'"
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
              <n-gi span="1">
                <n-form-item label="用户ID" path="userId">
                  <n-input-number v-model:value="formValue.userId" :min="1" style="width:100%" placeholder="请输入用户ID" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="使用的模型" path="model">
                  <n-select
                    v-model:value="formValue.model"
                    :options="videoModelOptions"
                    placeholder="请选择模型"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="视频提示词" path="prompt">
                  <n-input
                    type="textarea"
                    v-model:value="formValue.prompt"
                    placeholder="请描述要生成的视频内容，支持中英文。可在末尾追加参数，如 --ratio 16:9 --dur 5"
                    :rows="4"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="首帧参考图" path="referenceImage">
                  <UploadImage v-model:value="formValue.referenceImage" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="尾帧图片" path="lastFrame">
                  <UploadImage v-model:value="formValue.lastFrame" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="配音音频" path="voiceDubUrl">
                  <UploadFile v-model:value="formValue.voiceDubUrl" :max-number="1" help-text="上传自定义配音（MP3/WAV）" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="背景音乐" path="bgMusicUrl">
                  <UploadFile v-model:value="formValue.bgMusicUrl" :max-number="1" help-text="上传背景音乐（MP3/WAV）" />
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="is15ProModel">
                <n-form-item label="生成同步音频" path="generateAudio">
                  <n-switch
                    :value="formValue.generateAudio === 1"
                    @update:value="(v) => (formValue.generateAudio = v ? 1 : 0)"
                  >
                    <template #checked>开启</template>
                    <template #unchecked>关闭</template>
                  </n-switch>
                  <span style="margin-left:8px;color:#999;font-size:12px">仅 Seedance 1.5 Pro 支持</span>
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
  import { Edit, View } from '@/api/openai/videoTask';
  import { State, newState, rules, videoModelOptions } from './model';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
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
  const dialogWidth = computed(() => adaModalWidth(720));

  // 是否为 1.5 Pro 模型（支持 generate_audio）
  const is15ProModel = computed(() => formValue.value.model?.includes('1-5'));

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
      .then((res) => { formValue.value = newState(res); })
      .finally(() => { loading.value = false; });
  }

  defineExpose({ openModal });
</script>

<style lang="less"></style>

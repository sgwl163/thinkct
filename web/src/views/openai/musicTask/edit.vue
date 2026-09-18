<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑AI音乐任务 #' + formValue.id : '添加AI音乐任务'"
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
                <n-form-item label="创作模式" path="type">
                  <n-select v-model:value="formValue.type" :options="musicTypeOptions" placeholder="请选择创作模式" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="提示词" path="prompt">
                  <n-input type="textarea" v-model:value="formValue.prompt" placeholder="灵感模式填写描述，定制模式填写歌词" :rows="4" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="歌名" path="title">
                  <n-input v-model:value="formValue.title" placeholder="请输入歌名" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="风格标签" path="tags">
                  <n-input v-model:value="formValue.tags" placeholder="如：pop, rock, electronic" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="使用的模型" path="model">
                  <n-input v-model:value="formValue.model" placeholder="如：chirp-v3-5" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="纯音乐" path="makeInstrumental">
                  <n-switch v-model:value="formValue.makeInstrumental" :checked-value="1" :unchecked-value="0">
                    <template #checked>是</template>
                    <template #unchecked>否</template>
                  </n-switch>
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
  import { Edit, View } from '@/api/openai/musicTask';
  import { State, newState, rules, musicTypeOptions } from './model';
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
  const dialogWidth = computed(() => adaModalWidth(700));

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

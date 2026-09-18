<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑AI模型 #' + formValue.id : '添加AI模型'"
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
              <n-gi span="1">
                <n-form-item label="模型别名" path="name">
                  <n-input placeholder="如：GPT-4" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="模型标识" path="model">
                  <n-input placeholder="如：gpt-4-turbo" v-model:value="formValue.model" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="所属分组ID" path="groupId">
                  <n-input-number v-model:value="formValue.groupId" :min="1" style="width:100%" placeholder="请输入接口分组ID" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="支持识图" path="isVision">
                  <n-switch v-model:value="formValue.isVision" :checked-value="1" :unchecked-value="0">
                    <template #checked>支持</template>
                    <template #unchecked>不支持</template>
                  </n-switch>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="支持文件" path="isFile">
                  <n-switch v-model:value="formValue.isFile" :checked-value="1" :unchecked-value="0">
                    <template #checked>支持</template>
                    <template #unchecked>不支持</template>
                  </n-switch>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="支持联网" path="isSearch">
                  <n-switch v-model:value="formValue.isSearch" :checked-value="1" :unchecked-value="0">
                    <template #checked>支持</template>
                    <template #unchecked>不支持</template>
                  </n-switch>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" placeholder="请选择状态" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="排序" path="sort">
                  <n-input-number v-model:value="formValue.sort" :min="0" style="width:100%" />
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
  import { Edit, View, MaxSort } from '@/api/openai/apiModel';
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
  const dialogWidth = computed(() => adaModalWidth(640));

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
      loading.value = true;
      MaxSort()
        .then((res) => { formValue.value.sort = res.sort; })
        .finally(() => { loading.value = false; });
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

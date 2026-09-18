<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑会员 #' + formValue.id : '添加会员'"
      :style="{ width: dialogWidth }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="90"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:2 l:2 xl:2 2xl:2" responsive="screen" :x-gap="12">
              <n-gi span="1">
                <n-form-item label="用户名" path="username">
                  <n-input placeholder="请输入用户名" v-model:value="formValue.username" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="昵称" path="nickname">
                  <n-input placeholder="请输入昵称" v-model:value="formValue.nickname" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手机号" path="mobile">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.mobile" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="email">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.email" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="密码" path="password">
                  <n-input
                    :placeholder="formValue.id > 0 ? '留空则不修改密码' : '请输入密码'"
                    v-model:value="formValue.password"
                    type="password"
                    show-password-on="click"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="性别" path="gender">
                  <n-radio-group v-model:value="formValue.gender" name="gender">
                    <n-radio-button
                      v-for="item in dict.getOptionUnRef('sys_user_sex')"
                      :key="item.value"
                      :value="item.value"
                      :label="item.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="生日" path="birthday">
                  <DatePicker v-model:formValue="formValue.birthday" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-select v-model:value="formValue.status" :options="userStatusOptions" placeholder="请选择状态" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="余额" path="money">
                  <n-input-number v-model:value="formValue.money" :min="0" style="width:100%" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="积分" path="score">
                  <n-input-number v-model:value="formValue.score" :min="0" style="width:100%" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="头像" path="avatar">
                  <UploadImage :maxNumber="1" v-model:value="formValue.avatar" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="签名" path="motto">
                  <n-input type="textarea" placeholder="请输入个人签名" v-model:value="formValue.motto" :rows="2" />
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
  import { Edit, View } from '@/api/user/user';
  import { State, newState, rules, userStatusOptions } from './model';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
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
  const dialogWidth = computed(() => adaModalWidth(720));

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
        formValue.value.password = ''; // 默认不显示密码
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({ openModal });
</script>

<style lang="less"></style>

<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑权限规则 #' + formValue.id : '添加权限规则'"
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
                <n-form-item label="上级菜单" path="pid">
                  <n-input-number
                    placeholder="请输入上级菜单"
                    v-model:value="formValue.pid"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮" path="type">
                  <n-input
                    placeholder="请输入类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"
                    v-model:value="formValue.type"
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
                <n-form-item label="规则名称" path="name">
                  <n-input
                    placeholder="请输入规则名称"
                    v-model:value="formValue.name"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="路由路径" path="path">
                  <n-input
                    placeholder="请输入路由路径"
                    v-model:value="formValue.path"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="图标" path="icon">
                  <n-input
                    placeholder="请输入图标"
                    v-model:value="formValue.icon"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="菜单类型:tab=选项卡,link=链接,iframe=Iframe" path="menuType">
                  <n-input
                    placeholder="请输入菜单类型:tab=选项卡,link=链接,iframe=Iframe"
                    v-model:value="formValue.menuType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="Url" path="url">
                  <n-input
                    placeholder="请输入Url"
                    v-model:value="formValue.url"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="组件路径" path="component">
                  <n-input
                    placeholder="请输入组件路径"
                    v-model:value="formValue.component"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="未登录有效:0=否,1=是" path="noLoginValid">
                  <n-input-number
                    placeholder="请输入未登录有效:0=否,1=是"
                    v-model:value="formValue.noLoginValid"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单" path="extend">
                  <n-input
                    placeholder="请输入扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"
                    v-model:value="formValue.extend"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <n-input
                    placeholder="请输入备注"
                    v-model:value="formValue.remark"
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
                  <n-input
                    placeholder="请输入状态:0=禁用,1=启用"
                    v-model:value="formValue.status"
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
  import { Edit, View } from '@/api/user/userRule';
  import { State, newState, rules } from './model';
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
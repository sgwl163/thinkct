<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑产品 #' + formValue.id : '添加产品'"
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
            <div class="edit-layout">
              <!-- 左栏：主要内容 60% -->
              <div class="edit-col-main">
                <!-- 基础信息 -->
                <div class="section-block">
                  <div class="section-title">基础信息</div>
                  <n-grid :cols="2" :x-gap="16">
                    <n-gi :span="2">
                      <n-form-item label="产品标题" path="title">
                        <n-input-group>
                          <n-input placeholder="请输入产品标题" v-model:value="formValue.title" :style="{
                            fontWeight: isBold ? 'bold' : 'normal',
                            color: titleColor || 'inherit'
                          }" />
                          <n-button :type="isBold ? 'primary' : 'default'" @click="toggleBold">
                            粗
                          </n-button>
                          <div style="width: 39px;height: 32px;">
                            <n-color-picker v-model:value="titleColor" style="width: 45px;" :modes="['hex']" :show-alpha="false" :render-label="() => ''" />
                          </div>
                        </n-input-group>
                      </n-form-item>
                    </n-gi>
                    <n-gi :span="1">
                      <n-form-item label="所属频道" path="channelId">
                        <n-tree-select
                          v-model:value="formValue.channelId"
                          :options="channelOptions"
                          key-field="id"
                          label-field="name"
                          children-field="children"
                          placeholder="请选择所属频道"
                          style="width: 100%"
                          clearable
                          default-expand-all
                        />
                      </n-form-item>
                    </n-gi>
                    <n-gi :span="1">
                      <n-form-item label="副频道" path="channelIds">
                        <n-tree-select
                          :value="channelIdsArray"
                          :options="channelOptions"
                          key-field="id"
                          label-field="name"
                          children-field="children"
                          placeholder="请选择副频道（可多选）"
                          style="width: 100%"
                          multiple
                          clearable
                          default-expand-all
                          @update:value="onChannelIdsChange"
                        />
                      </n-form-item>
                    </n-gi>
                    <n-gi :span="2">
                      <n-form-item label="封面" path="image">
                        <div style="width: 100%">
                          <div style="display: flex; gap: 8px; flex-wrap: wrap;">
                            <UploadImage
                              v-model:value="formValue.image"
                              :max-number="1"
                            />
                            <UploadImage
                              :value="imagesArray"
                              :max-number="20"
                              @update:value="onImagesChange"
                            />
                          </div>
                          <n-text depth="3" style="font-size: 12px; margin-top: 6px; display: block">
                            建议尺寸 800x800，第一张为封面，后面紧跟着为图集图片
                          </n-text>
                        </div>
                      </n-form-item>
                    </n-gi>
                  </n-grid>
                </div>

                <!-- 商品内容 -->
                <div class="section-block">
                  <div class="section-title">商品详情</div>
                  <n-form-item label=" " :show-label="false" path="content">
                    <Editor v-model:value="formValue.content" style="width: 100%; height: 500px" />
                  </n-form-item>
                </div>
              </div>

            
              <!-- 右栏：辅助设置 40% -->
              <div class="edit-col-side">
                <!-- 发布设置 -->
                <div class="section-block">
                  <div class="section-title">发布设置</div>
                  <n-form-item label="标志" path="flag">
                    <n-select
                      :value="flagArray"
                      :options="articleFlagOptions"
                      placeholder="请选择标志（可多选）"
                      multiple
                      clearable
                      @update:value="onFlagChange"
                    />
                  </n-form-item>
                  <n-form-item label="状态" path="status">
                    <n-select
                      v-model:value="formValue.status"
                      :options="articleStatusOptions"
                      placeholder="请选择状态"
                    />
                  </n-form-item>
                  <n-form-item label="发布时间" path="publishTime">
                    <n-date-picker
                      v-model:value="formValue.publishTime"
                      type="datetime"
                      clearable
                      placeholder="留空则发布时自动设为当前时间"
                      style="width: 100%"
                    />
                  </n-form-item>
                  <n-form-item label="权重" path="weigh">
                    <n-input-number v-model:value="formValue.weigh" :min="0" style="width: 100%" />
                  </n-form-item>
                  <n-form-item label="价格" path="price">
                    <n-input-number
                      v-model:value="formValue.price"
                      :min="0"
                      :precision="2"
                      style="width: 100%"
                      placeholder="0.00"
                    >
                      <template #prefix>￥</template>
                    </n-input-number>
                  </n-form-item>
                </div>

                <!-- SEO 优化 -->
                <div class="section-block">
                  <div class="section-title">SEO 优化</div>
                  <n-form-item label="SEO标题" path="seotitle">
                    <n-input placeholder="请输入SEO标题" v-model:value="formValue.seotitle" />
                  </n-form-item>
                  <n-form-item label="关键词" path="keywords">
                    <n-input placeholder="多个关键词用逗号分隔" v-model:value="formValue.keywords" />
                  </n-form-item>
                  <n-form-item label="描述" path="description">
                    <n-input
                      type="textarea"
                      placeholder="请输入描述"
                      v-model:value="formValue.description"
                      :rows="3"
                    />
                  </n-form-item>
                  <n-form-item label="标签" path="tags">
                    <n-select
                      :value="tagsArray"
                      :options="tagsOptions"
                      placeholder="请选择标签或输入新标签（按回车添加）"
                      multiple
                      clearable
                      filterable
                      tag
                      @update:value="onTagsChange"
                    />
                  </n-form-item>
                </div>

                <!-- 其他设置 -->
                <div class="section-block">
                  <div class="section-title">其他设置</div>
                  <n-form-item label="自定义名称" path="diyname">
                    <n-input placeholder="请输入自定义名称（用于URL链接）" v-model:value="formValue.diyname" />
                  </n-form-item>
                  <n-form-item label="外部链接" path="url">
                    <n-input placeholder="请输入外部链接URL" v-model:value="formValue.url" />
                  </n-form-item>
                  <n-form-item label="跳转方式" path="target">
                    <n-select
                      v-model:value="formValue.target"
                      :options="targetOptions"
                      placeholder="请选择跳转方式"
                    />
                  </n-form-item>
                </div>
              </div>
            </div>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm">保存</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, watch } from 'vue';
  import { cloneDeep } from 'lodash-es';
  import { Edit, View } from '@/api/web/products';
  import { Edit as TagsEdit } from '@/api/web/tags';
  import { TreeOption as ChannelTreeOption } from '@/api/web/channel';
  import { List as TagsList } from '@/api/web/tags';
  import { State, newState, rules, articleFlagOptions, articleStatusOptions, targetOptions } from './model';
  import Editor from '@/components/Editor/editor.vue';
  import UploadImage from '@/components/Upload/uploadImage.vue';
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
  const dialogWidth = computed(() => adaModalWidth(1400));

  // 频道树选项
  const channelOptions = ref<any[]>([]);
  function loadChannelOptions() {
    ChannelTreeOption({ model: 'products' }).then((res) => {
      channelOptions.value = res || [];
    });
  }
  loadChannelOptions();

  // 标题样式控制
  const isBold = ref(false);
  const titleColor = ref('#000000');

  // 监听 titleStyle 变化来解析加粗和颜色
  watch(() => formValue.value.titleStyle, (newVal) => {
    if (!newVal) {
      isBold.value = false;
      titleColor.value = '';
      return;
    }
    const parts = newVal.split(',');
    isBold.value = parts[0] === 'bold';
    titleColor.value = parts[1] || '';
  });

  // 更新 titleStyle 字段
  const updateTitleStyle = () => {
    formValue.value.titleStyle = `${isBold.value ? 'bold' : 'normal'},${titleColor.value || ''}`;
  };

  watch([isBold, titleColor], () => {
    updateTitleStyle();
  });

  const toggleBold = () => {
    isBold.value = !isBold.value;
  };

  // 副频道：逗号字符串 <-> 数字数组互转
  const channelIdsArray = computed<number[]>(() => {
    const v = formValue.value.channelIds;
    if (!v) return [];
    return v.split(',').map((s) => Number(s.trim())).filter(Boolean);
  });
  function onChannelIdsChange(list: number[]) {
    formValue.value.channelIds = list ? list.join(',') : '';
  }

  // 标签选项
  const tagsOptions = ref<{ label: string; value: string }[]>([]);
  function loadTagsOptions() {
    TagsList({ page: 1, perPage: 200 }).then((res) => {
      tagsOptions.value = (res?.list || []).map((item: any) => ({
        label: item.name,
        value: item.name,
      }));
    });
  }
  loadTagsOptions();

  // 图集：字符串 <-> 数组互转
  const imagesArray = computed<string[]>(() => {
    const v = formValue.value.images;
    if (!v) return [];
    return v.split(',').map((s) => s.trim()).filter(Boolean);
  });
  function onImagesChange(list: string | string[]) {
    formValue.value.images = Array.isArray(list) ? list.join(',') : list;
  }

  // 标志：逗号字符串 <-> 数组互转
  const flagArray = computed<string[]>(() => {
    const v = formValue.value.flag;
    if (!v) return [];
    return v.split(',').map((s) => s.trim()).filter(Boolean);
  });
  function onFlagChange(list: string[]) {
    formValue.value.flag = list ? list.join(',') : '';
  }

  // 标签：逗号字符串 <-> 数组互转
  const tagsArray = computed<string[]>(() => {
    const v = formValue.value.tags;
    if (!v) return [];
    return v.split(',').map((s) => s.trim()).filter(Boolean);
  });
  
  async function onTagsChange(list: string[]) {
    // 找出新增的标签 (不在当前 tagsOptions 里的)
    if (list && list.length > 0) {
      const existingValues = tagsOptions.value.map(t => t.value);
      const newTags = list.filter(tagName => !existingValues.includes(tagName));
      
      for (const newTag of newTags) {
        try {
          await TagsEdit({
            id: 0,
            name: newTag,
            status: 1,
            weigh: 10,
            type: 'default',
            innerChain: 0,
            documentCount: 0
          });
        } catch (e) {
          console.error("添加新标签失败:", e);
        }
      }
      if (newTags.length > 0) {
        loadTagsOptions(); // 重新加载标签列表
      }
    }
    formValue.value.tags = list ? list.join(',') : '';
  }

  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        const submitData = cloneDeep(formValue.value);
        if (submitData.publishTime) {
          submitData.publishTime = Math.floor(submitData.publishTime / 1000);
        }
        Edit(submitData)
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

<style lang="less" scoped>
  .edit-layout {
    display: grid;
    grid-template-columns: minmax(0, 3fr) minmax(0, 2fr);
    gap: 0 24px;
    align-items: start;

    @media (max-width: 900px) {
      grid-template-columns: 1fr;
    }
  }

  .edit-col-main,
  .edit-col-side {
    min-width: 0;
  }

  .section-block {
    margin-bottom: 20px;
    border: 1px solid #efeff5;
    border-radius: 6px;
    padding: 16px 20px 4px;
  }

  .section-title {
    font-size: 13px;
    font-weight: 600;
    color: #555;
    margin-bottom: 14px;
    padding-bottom: 10px;
    border-bottom: 1px solid #efeff5;
  }
</style>

<template>
  <div>
    <n-spin :show="show" description="正在生成配置信息...">
      <n-card class="proCard">
        <n-alert v-if="buildMessage" type="warning" class="build-message">
          {{ buildMessage }}
          <template v-if="loadFailed" #action>
            <n-button size="small" :loading="show" @click="initialize">重新加载</n-button>
          </template>
        </n-alert>
        <n-tabs
          v-if="loaded"
          type="card"
          class="card-tabs"
          :default-value="value"
          animated
          tab-style="min-width: 80px;"
          style="margin: 0 -4px"
          pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
          :on-update:value="updateTabs"
          ref="tabsRef"
          @close="handleClose"
          @add="handleAdd"
        >
          <n-tab-pane v-for="panel in panels" :key="panel" :name="panel">
            <template v-if="panel === '基本信息'">
              <BaseInfo v-model:value="genInfo" :selectList="selectList" />
            </template>

            <template v-if="panel === '主表字段'">
              <EditMasterCell v-model:value="genInfo" :selectList="selectList" />
            </template>
          </n-tab-pane>

          <n-tab-pane
            v-for="panel in slavePanels"
            :key="panel"
            :name="panel"
            v-show="slavePanels.length > 0"
          >
            <EditSlaveCell
              v-model:value="genInfo"
              :uuid="slaveMap[panel]"
              :selectList="selectList"
            />
          </n-tab-pane>

          <template #suffix>
            <n-space>
              <n-button type="default" :disabled="formBtnLoading" @click="handleBack"
                >返回列表</n-button
              >
              <n-button
                type="primary"
                :disabled="formBtnLoading"
                :loading="formBtnPreviewLoading"
                @click="preview"
                >预览代码</n-button
              >
              <n-button type="success" :loading="formBtnLoading" @click="submitBuild"
                >提交生成</n-button
              >
              <n-button type="info" dashed :disabled="formBtnLoading" @click="submitSave"
                >仅保存配置</n-button
              >
            </n-space>
          </template>
        </n-tabs>

        <n-modal
          v-model:show="showModal"
          :block-scroll="false"
          :mask-closable="false"
          :show-icon="false"
          preset="card"
          title="预览代码"
          style="width: 95%"
        >
          <PreviewTab :previewModel="previewModel" />
          <template #action>
            <n-space justify="end">
              <n-button @click="() => (showModal = false)">关闭</n-button>
              <n-button type="info" :loading="formBtnLoading" @click="submitBuild"
                >提交生成</n-button
              >
            </n-space>
          </template>
        </n-modal>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, onBeforeUnmount, ref, watch } from 'vue';
  import { useRouter } from 'vue-router';
  import { useDialog, useMessage, useNotification } from 'naive-ui';
  import BaseInfo from './components/BaseInfo.vue';
  import EditMasterCell from './components/EditMasterCell.vue';
  import EditSlaveCell from './components/EditSlaveCell.vue';
  import { Selects, View, Preview, Build, Edit } from '@/api/develop/code';
  import { selectListObj, newState, formatColumns } from '@/views/develop/code/components/model';
  import PreviewTab from '@/views/develop/code/components/PreviewTab.vue';
  import { isJsonString } from '@/utils/is';
  import { ResultEnum } from '@/enums/httpEnum';

  interface Props {
    genId?: number;
  }
  const props = withDefaults(defineProps<Props>(), { genId: 0 });
  const router = useRouter();
  const genId = Number(router.currentRoute.value.params.id ?? props.genId);
  const show = ref(false);
  const loaded = ref(false);
  const loadFailed = ref(false);
  const message = useMessage();
  const selectList = ref<any>(selectListObj);
  const genInfo = ref(newState(null));
  const tabsRef = ref();
  const value = ref('基本信息');
  const panels = ref(['基本信息', '主表字段']);
  const slaveMap = ref<any>([]);
  const slavePanels = ref<any>([]);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formBtnPreviewLoading = ref(false);
  const previewModel = ref<any>();
  const dialog = useDialog();
  const notification = useNotification();
  const buildMessage = ref('');
  let disposed = false;
  let checkingBuild = false;
  onBeforeUnmount(() => {
    disposed = true;
  });

  onMounted(initialize);

  async function initialize() {
    if (show.value || disposed) return;
    if (genId < 1 && props.genId < 1) {
      buildMessage.value = '生成ID不正确，请返回列表重新打开配置。';
      return;
    }
    show.value = true;
    loadFailed.value = false;
    try {
      for (let attempt = 0; attempt < 10 && !disposed; attempt++) {
        try {
          await loadSelect();
          if (disposed) return;
          await getGenInfo();
          if (disposed) return;
          loaded.value = true;
          return;
        } catch (error: any) {
          if (disposed) return;
          if (attempt === 9) {
            loadFailed.value = true;
            buildMessage.value = `生成配置加载失败：${error?.message || '服务暂时不可用'}。请稍后重新加载，无需再次提交生成。`;
          } else {
            buildMessage.value = '正在等待后端恢复并重新加载生成配置...';
            await new Promise((resolve) => window.setTimeout(resolve, 1500));
          }
        }
      }
    } finally {
      show.value = false;
    }
  }

  async function getGenInfo() {
    const result = await View({ id: genId }, { isShowErrorMessage: false });
    if (disposed) return;
    applyGenInfo(result);
    if (result.building) {
      formBtnLoading.value = true;
      void waitForBuildResult(true);
    }
  }

  function applyGenInfo(tmp) {
    // 导入主表数据
    tmp.masterColumns = formatColumns(tmp.masterColumns);

    // 导入生成选项
    if (isJsonString(tmp.options)) {
      tmp.options = JSON.parse(tmp.options);
    }
    tmp.options ||= {};

    // 预设流程
    if (!tmp.options.presetStep) {
      tmp.options.presetStep = {
        formGridCols: 1,
      };
    }

    // 树表
    if (!tmp.options.tree) {
      tmp.options.tree = {
        titleColumn: null,
        styleType: 1,
      };
    }

    genInfo.value = tmp;
    buildMessage.value = tmp.buildMessage || '';
  }

  watch(
    genInfo,
    (newVal, _oldVal) => {
      if (newVal.genType >= 10 && newVal.genType < 20) {
        handleAdd('主表字段');
      } else {
        handleClose('主表字段');
      }

      if (newVal && newVal.options && newVal.options.join !== undefined) {
        slavePanels.value = [];
        for (let i = 0; i < newVal.options.join.length; i++) {
          if (newVal.options.join[i]?.alias) {
            handleSlaveAdd(
              '关联表[ ' + newVal.options.join[i].alias + ' ]',
              newVal.options.join[i]
            );
          }
        }
      }
    },
    {
      deep: true, // 是否深度监听
    }
  );

  function updateTabs(value: string | number) {
    console.log('value:' + value);
  }

  function handleAdd(name: string) {
    const nameIndex = panels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) {
      panels.value.push(name);
    }
  }

  function handleSlaveAdd(name: string, join) {
    const nameIndex = slavePanels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) {
      slavePanels.value.push(name);
      slaveMap.value[name] = join.uuid;
    }
  }

  function handleClose(name: string) {
    const nameIndex = panels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) return;
    panels.value.splice(nameIndex, 1);
    if (name === value.value) {
      value.value = panels.value[Math.min(nameIndex, panels.value.length - 1)];
    }
  }

  const loadSelect = async () => {
    selectList.value = await Selects({}, { isShowErrorMessage: false });
  };

  function preview() {
    formBtnPreviewLoading.value = true;
    Preview(genInfo.value)
      .then((res) => {
        previewModel.value = res;
        showModal.value = true;
      })
      .finally(() => {
        formBtnPreviewLoading.value = false;
      });
  }

  function submitBuild() {
    if (formBtnLoading.value) return;
    dialog.info({
      title: '提示',
      content: '你确定要提交生成吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        // Naive UI 会等待返回的 Promise；同步返回，让确认框立即关闭。
        void runBuild();
      },
    });
  }

  async function runBuild() {
    if (formBtnLoading.value || disposed) return;
    formBtnLoading.value = true;
    showModal.value = false;
    buildMessage.value = '正在生成代码，请等待完成。使用热编译时服务可能重新加载，请勿重复提交。';
    try {
      const result = await Build(genInfo.value);
      if (!result || typeof result.code !== 'number') {
        throw new Error('生成响应无效，正在核对后台结果');
      }
      if (result.code !== ResultEnum.SUCCESS) {
        // 后台明确返回失败时保留本次错误，不能用上次成功的记录覆盖。
        const detail = result.message || '生成未完成，请查看后台日志';
        message.error(detail);
        try {
          applyGenInfo(await View({ id: genId }, { isShowErrorMessage: false }));
        } catch {
          // 服务重启时暂时无法刷新详情，仍保留后台已返回的本次生成错误。
        }
        buildMessage.value = detail;
        return;
      }
      buildSuccessNotify();
      // 已确认生成成功；详情查询只是同步文件和状态，重启期间失败应继续核对。
      await waitForBuildResult(false, '', true);
    } catch (error: any) {
      // 网络中断不能证明生成失败；服务恢复后查询后台真实状态。
      const detail = error?.message || '生成请求未确认完成';
      buildMessage.value = detail;
      await waitForBuildResult(false, detail);
    } finally {
      if (!checkingBuild) formBtnLoading.value = false;
    }
  }

  function submitSave() {
    dialog.info({
      title: '提示',
      content: '你确定要保存生成配置吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Edit(genInfo.value).then(async (_res) => {
          await getGenInfo();
          message.success('生成配置已保存');
        });
      },
    });
  }

  async function waitForBuildResult(
    observedRunning = false,
    requestError = '',
    confirmedSuccess = false
  ) {
    if (checkingBuild) return;
    checkingBuild = true;
    formBtnLoading.value = true;
    const deadline = Date.now() + 10 * 60 * 1000;
    if (confirmedSuccess) {
      buildMessage.value = '代码已生成，正在等待后端重新加载并同步文件状态，请勿重复提交。';
    }
    try {
      while (Date.now() < deadline && !disposed) {
        await new Promise((resolve) => window.setTimeout(resolve, 1500));
        if (disposed) return;
        try {
          const result = await View({ id: genId }, { isShowErrorMessage: false });
          if (disposed) return;
          if (result.building) {
            observedRunning = true;
            buildMessage.value = result.buildMessage;
            continue;
          }
          applyGenInfo(result);
          const statusOption = selectList.value.status.find((item) => item.value === result.status);
          if (result.buildSucceeded) {
            if (observedRunning && !confirmedSuccess) {
              message.success(statusOption?.label || '生成已完成，请返回列表查看结果');
            } else if (!confirmedSuccess) {
              // 没有观察到本次任务执行，后台的旧成功状态不能证明本次请求成功。
              buildMessage.value = `${requestError}；后台记录为${statusOption?.label || '已完成'}，本次请求结果未确认，请核对文件。`;
              message.warning(buildMessage.value);
            }
          } else {
            buildMessage.value = [requestError, result.buildMessage].filter(Boolean).join('；');
            message.warning(buildMessage.value || statusOption?.label || '请返回列表核对生成结果');
          }
          return;
        } catch {
          buildMessage.value = confirmedSuccess
            ? '代码已生成，后端暂时无法连接，正在等待服务恢复并同步文件状态...'
            : '后端暂时无法连接，正在等待服务恢复并核对生成结果...';
        }
      }
      if (!disposed) {
        buildMessage.value = confirmedSuccess
          ? '代码已生成，但暂时无法同步详情，请待后端恢复后刷新列表，无需重复提交生成。'
          : '暂时无法确认生成结果，请待后端恢复后刷新列表核对。已有文件不会因此被删除。';
      }
    } finally {
      checkingBuild = false;
      formBtnLoading.value = false;
    }
  }

  function buildSuccessNotify() {
    notification.success({
      title: '代码生成完成',
      content:
        '后台已完成生成。请待服务重新加载完成后刷新一次，以加载新菜单；未使用热编译时请先重启服务。',
      duration: 8000,
    });
  }

  function handleBack() {
    dialog.info({
      title: '提示',
      content: '你确定要返回生成列表？系统不会主动保存更改',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        router.push({ name: 'develop_code' });
      },
    });
  }
</script>
<style lang="less" scoped>
  .build-message {
    margin-bottom: 16px;
  }
  .card-tabs .n-tabs-nav--bar-type {
    padding-left: 4px;
  }
  ::v-deep(.alert-margin) {
    margin-bottom: 20px;
  }
  ::v-deep(.tag-margin) {
    margin-bottom: 10px;
  }

  ::v-deep(.code-scrollbar) {
    height: calc(100vh - 300px);
    background: #282b2e;
    color: #e0e2e4;
    padding: 10px;
  }
</style>

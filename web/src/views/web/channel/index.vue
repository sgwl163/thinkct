<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="栏目管理">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
        ref="searchFormRef"
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable
        ref="actionRef"
        openChecked
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
        :cascade="false"
        :expanded-row-keys="expandedKeys"
        @update:expanded-row-keys="updateExpandedKeys"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/channel/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            class="min-left-space"
            v-if="hasPermission(['/channel/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/channel/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
          <n-button type="primary" icon-placement="left" @click="handleAllExpanded" class="min-left-space">
            全部{{ expandedKeys.length ? '收起' : '展开' }}
            <template #icon>
              <div class="flex items-center">
                <n-icon size="14">
                  <AlignLeftOutlined />
                </n-icon>
              </div>
            </template>
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
    <View ref="viewRef" />
  </div>
</template>
<script lang="ts" setup>
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useDictStore } from '@/store/modules/dict';
  import { List, Export, Delete, Status } from '@/api/web/channel';
  import { PlusOutlined, ExportOutlined, DeleteOutlined, AlignLeftOutlined } from '@vicons/antd';
  import { columns, schemas, State, loadOptions, newState } from './model';
  import { adaTableScrollX, convertListToTree } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import View from './view.vue';

  const dict = useDictStore();
  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const viewRef = ref();
  const checkedIds = ref([]);
  const expandedKeys = ref([]);
  const allTreeKeys = ref([]);

  const actionColumn = reactive({
    width: 360,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: State) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/channel/edit'],
          },
          {
            label: '添加',
            onClick: handleAdd.bind(null, record),
            auth: ['/channel/edit'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/channel/status'],
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/channel/status'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/channel/delete'],
          },
        ],
        dropDownActions: [
          {
            label: '查看详情',
            key: 'view',
            auth: ['/channel/view'],
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          }
        },
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  // 加载普通数表数据
  const loadDataTable = async (res = {}) => {
    const params = { ...(searchFormRef.value?.formModel ?? {}), ...res, pagination: false };
    const dataSource = await List(params);
    allTreeKeys.value = expandedKeys.value = dataSource.list.map((item) => item.id);
    dataSource.list = convertListToTree(dataSource.list, 'id');
    return dataSource;
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 添加数据
  function addTable() {
    editRef.value.openModal(null);
  }

  // 添加树节点下级数据
  function handleAdd(record: Recordable) {
    const state = newState(null);
    state.pid = record.id;
    editRef.value.openModal(state);
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
  }

  // 单个删除
  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1){
      message.error('请至少选择一项要删除的数据');
      return;
    }

    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 导出
  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  // 修改状态
  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('设为' + dict.getLabel('sys_normal_disable', status) + '成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  // 收起/展开全部树节点
  function handleAllExpanded() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = allTreeKeys.value;
    }
  }

  // 更新展开的树节点
  function updateExpandedKeys(openKeys: never[]) {
    expandedKeys.value = openKeys;
  }

  onMounted(() => {
    loadOptions();
    
  });
</script>

<style lang="less" scoped></style>

<template>
  <n-drawer v-model:show="isDrawer" :width="width" placement="right">
    <n-drawer-content :title="title" :native-scrollbar="false">
      <div class="drawer">
        <n-divider title-placement="center">主题</n-divider>

        <div class="justify-center drawer-setting-item dark-switch">
          <n-tooltip placement="bottom">
            <template #trigger>
              <n-switch v-model:value="designStore.darkTheme" class="dark-theme-switch">
                <template #checked>
                  <n-icon size="14" color="#ffd93b">
                    <SunnySharp />
                  </n-icon>
                </template>
                <template #unchecked>
                  <n-icon size="14" color="#ffd93b">
                    <Moon />
                  </n-icon>
                </template>
              </n-switch>
            </template>
            <span>{{ designStore.darkTheme ? '深' : '浅' }}色主题</span>
          </n-tooltip>
        </div>

        <n-divider title-placement="center">系统主题</n-divider>

        <div class="drawer-setting-item align-items-top">
          <span
            class="theme-item"
            v-for="(item, index) in appThemeList"
            :key="index"
            :style="{ 'background-color': item }"
            @click="togTheme(item)"
          >
            <n-icon size="12" v-if="item === designStore.appTheme">
              <CheckOutlined />
            </n-icon>
          </span>
        </div>

        <n-divider title-placement="center">导航栏模式</n-divider>

        <div class="drawer-setting-item align-items-top">
          <div class="drawer-setting-item-style align-items-top">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/nav-theme-dark.svg"
                  @click="togNavMode('vertical')"
                  alt="左侧菜单模式"
                />
              </template>
              <span>左侧菜单模式</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-show="settingStore.navMode === 'vertical'" />
          </div>

          <div class="drawer-setting-item-style">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/nav-horizontal.svg"
                  alt="顶部菜单模式"
                  @click="togNavMode('horizontal')"
                />
              </template>
              <span>顶部菜单模式</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-show="settingStore.navMode === 'horizontal'" />
          </div>

          <div class="drawer-setting-item-style">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/nav-horizontal-mix.svg"
                  @click="togNavMode('horizontal-mix')"
                  alt="顶部菜单混合模式"
                />
              </template>
              <span>顶部菜单混合模式</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-show="settingStore.navMode === 'horizontal-mix'" />
          </div>
        </div>

        <n-divider title-placement="center">导航栏风格</n-divider>

        <div class="drawer-setting-item align-items-top">
          <div class="drawer-setting-item-style align-items-top">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/nav-theme-dark.svg"
                  alt="暗色侧边栏"
                  @click="togNavTheme('dark')"
                />
              </template>
              <span>暗色侧边栏</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-if="settingStore.navTheme === 'dark'" />
          </div>

          <div class="drawer-setting-item-style">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/nav-theme-light.svg"
                  alt="白色侧边栏"
                  @click="togNavTheme('light')"
                />
              </template>
              <span>白色侧边栏</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-if="settingStore.navTheme === 'light'" />
          </div>

          <div class="drawer-setting-item-style">
            <n-tooltip placement="top">
              <template #trigger>
                <img
                  src="~@/assets/images/header-theme-dark.svg"
                  @click="togNavTheme('header-dark')"
                  alt="暗色顶栏"
                />
              </template>
              <span>暗色顶栏</span>
            </n-tooltip>
            <n-badge dot color="#19be6b" v-if="settingStore.navTheme === 'header-dark'" />
          </div>
        </div>
        <n-divider title-placement="center">界面功能</n-divider>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 分割菜单 </div>
          <div class="drawer-setting-item-action">
            <n-tooltip :disabled="settingStore.navMode === 'horizontal-mix'" placement="left">
              <template #trigger>
                <span>
                  <n-switch
                    :disabled="settingStore.navMode !== 'horizontal-mix'"
                    v-model:value="settingStore.menuSetting.mixMenu"
                  />
                </span>
              </template>
              仅顶部菜单混合模式支持分割菜单
            </n-tooltip>
          </div>
        </div>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 固定顶栏 </div>
          <div class="drawer-setting-item-action">
            <n-switch v-model:value="settingStore.headerSetting.fixed" />
          </div>
        </div>

        <!--        <div class="drawer-setting-item">-->
        <!--          <div class="drawer-setting-item-title">-->
        <!--            固定侧边栏-->
        <!--          </div>-->
        <!--          <div class="drawer-setting-item-action">-->
        <!--            <n-switch v-model:value="settingStore.menuSetting.fixed" />-->
        <!--          </div>-->
        <!--        </div>-->

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 固定多页签 </div>
          <div class="drawer-setting-item-action">
            <n-switch
              :disabled="!settingStore.multiTabsSetting.show"
              v-model:value="settingStore.multiTabsSetting.fixed"
            />
          </div>
        </div>

        <n-divider title-placement="center">界面显示</n-divider>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 显示重载页面按钮 </div>
          <div class="drawer-setting-item-action">
            <n-switch v-model:value="settingStore.headerSetting.isReload" />
          </div>
        </div>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 显示面包屑导航 </div>
          <div class="drawer-setting-item-action">
            <n-switch v-model:value="settingStore.crumbsSetting.show" />
          </div>
        </div>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 显示面包屑图标 </div>
          <div class="drawer-setting-item-action">
            <n-switch
              :disabled="!settingStore.crumbsSetting.show"
              v-model:value="settingStore.crumbsSetting.showIcon"
            />
          </div>
        </div>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 显示多页签 </div>
          <div class="drawer-setting-item-action">
            <n-switch v-model:value="settingStore.multiTabsSetting.show" />
          </div>
        </div>
        <!--1.15废弃，没啥用，占用操作空间-->
        <!--        <div class="drawer-setting-item">-->
        <!--          <div class="drawer-setting-item-title"> 显示页脚 </div>-->
        <!--          <div class="drawer-setting-item-action">-->
        <!--            <n-switch v-model:value="settingStore.showFooter" />-->
        <!--          </div>-->
        <!--        </div>-->

        <n-divider title-placement="center">动画</n-divider>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 开启动画 </div>
          <div class="drawer-setting-item-action">
            <n-switch v-model:value="settingStore.isPageAnimate" />
          </div>
        </div>

        <div class="drawer-setting-item">
          <div class="drawer-setting-item-title"> 动画类型 </div>
          <div class="drawer-setting-item-select">
            <n-select
              v-model:value="settingStore.pageAnimateType"
              :options="animateOptions"
              :disabled="!settingStore.isPageAnimate"
            />
          </div>
        </div>

        <div class="drawer-setting-item">
          <n-alert type="info" :show-icon="false">
            配置修改后自动保存到当前浏览器，刷新页面后仍然生效。可通过“恢复默认”重置主题与布局。
          </n-alert>
        </div>
      </div>
      <template #footer>
        <n-button block @click="resetSettings">恢复默认</n-button>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts" setup>
  import { ref, watch } from 'vue';
  import { cloneDeep, omit } from 'lodash-es';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useDesignSettingStore } from '@/store/modules/designSetting';
  import { CheckOutlined } from '@vicons/antd';
  import { Moon, SunnySharp } from '@vicons/ionicons5';
  import { useDialog } from 'naive-ui';
  import { animates as animateOptions } from '@/settings/animateSetting';
  import projectSetting from '@/settings/projectSetting';
  import designSetting from '@/settings/designSetting';

  withDefaults(defineProps<{ title?: string; width?: number }>(), {
    title: '项目配置',
    width: 280,
  });

  const settingStore = useProjectSettingStore();
  const designStore = useDesignSettingStore();
  const dialog = useDialog();
  const isDrawer = ref(false);
  const appThemeList = designStore.appThemeList;

  watch(
    () => designStore.darkTheme,
    (to) => {
      settingStore.navTheme = to ? 'header-dark' : 'dark';
    },
    { flush: 'sync' }
  );

  function openDrawer() {
    isDrawer.value = true;
  }

  function closeDrawer() {
    isDrawer.value = false;
  }

  function togNavTheme(theme: string) {
    settingStore.navTheme =
      settingStore.navMode === 'horizontal' && theme === 'light' ? 'dark' : theme;
  }

  function togTheme(color: string) {
    designStore.appTheme = color;
  }

  function togNavMode(mode: string) {
    settingStore.$patch({ navMode: mode, menuSetting: { mixMenu: false } });
    togNavTheme(settingStore.navTheme);
  }

  function resetSettings() {
    dialog.warning({
      title: '恢复默认配置',
      content: '确定恢复默认主题和布局吗？此操作不会退出登录或清除其他数据。',
      positiveText: '恢复默认',
      negativeText: '取消',
      onPositiveClick: () => {
        // 仅重置界面偏好，保留移动端状态、权限模式和内置主题列表。
        designStore.$patch({
          darkTheme: designSetting.darkTheme,
          appTheme: designSetting.appTheme,
        });
        settingStore.$patch(cloneDeep(omit(projectSetting, ['isMobile', 'permissionMode'])));
      },
    });
  }

  defineExpose({ openDrawer, closeDrawer });
</script>

<style lang="less" scoped>
  .drawer {
    .n-divider:not(.n-divider--vertical) {
      margin: 10px 0;
    }

    &-setting-item {
      display: flex;
      align-items: center;
      padding: 12px 0;
      flex-wrap: wrap;

      &-style {
        display: inline-block;
        position: relative;
        margin-right: 16px;
        cursor: pointer;
        text-align: center;
      }

      &-title {
        flex: 1 1;
        font-size: 14px;
      }

      &-action {
        flex: 0 0 auto;
      }

      &-select {
        flex: 1;
      }

      .theme-item {
        width: 20px;
        min-width: 20px;
        height: 20px;
        cursor: pointer;
        border: 1px solid #eee;
        border-radius: 2px;
        margin: 0 5px 5px 0;
        text-align: center;
        line-height: 14px;

        .n-icon {
          color: #fff;
        }
      }
    }

    .align-items-top {
      align-items: flex-start;
      padding: 2px 0;
    }

    .justify-center {
      justify-content: center;
    }

    .dark-switch .n-switch {
      ::v-deep(.n-switch__rail) {
        background-color: #000e1c;
      }
    }
  }
</style>

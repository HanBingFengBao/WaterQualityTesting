<template>
  <div class="container" :style="`height: ${containerHeight}px`">
    <el-menu :default-active="activeIndex" router>
      <el-menu-item index="/backStage/user">
        <i class="el-icon-user"></i>
        <span slot="title">用户管理</span>
      </el-menu-item>
      <el-menu-item index="/backStage/map">
        <i class="el-icon-tickets"></i>
        <span slot="title">映射管理</span>
      </el-menu-item>
      <el-menu-item index="/backStage/file">
        <i class="el-icon-folder-opened"></i>
        <span slot="title">文件管理</span>
      </el-menu-item>
      <el-menu-item index="/backStage/data">
        <i class="el-icon-data-analysis"></i>
        <span slot="title">数据管理</span>
      </el-menu-item>
      <el-menu-item :index="logIndex">
        <i class="el-icon-notebook-1"></i>
        <span slot="title">日志记录</span>
      </el-menu-item>
    </el-menu>
    <div style="width: 100%">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      containerHeight: window.innerHeight - 56,
      activeIndex: "/backStage/user",
      logIndex: "/backStage/log/dataLog",
    };
  },
  created() {
    this.activeIndex = this.$route.path;
    if (this.activeIndex === "/backStage/log/dataLog") {
      this.logIndex = "/backStage/log/dataLog";
    } else if (this.activeIndex === "/backStage/log/fileLog") {
      this.logIndex = "/backStage/log/fileLog";
    } else if (this.activeIndex === "/backStage/log/mapLog") {
      this.logIndex = "/backStage/log/mapLog";
    }
  },
};
</script>

<style lang="less" scoped>
.container {
  display: flex;
  position: absolute;
  top: 56px;
  width: 100%;
}

.el-menu {
  height: 100%;
}
</style>

<template>
  <div class="container" :style="`height: ${containerHeight}px`">
    <div class="download">
      <Download sort="小时制" :list="hourList"></Download>
      <Download sort="月度制" :list="monthList"></Download>
    </div>
    <div class="upload">
      <UpLoad></UpLoad>
    </div>
  </div>
</template>

<script>
import Download from "@/components/load/DownLoad.vue";
import UpLoad from "@/components/load/UpLoad";
import { getFileList } from "@/api/file";
export default {
  data() {
    return {
      hourList: [],
      monthList: [],
      containerHeight: window.innerHeight - 57,
    };
  },
  components: {
    Download,
    UpLoad,
  },
  methods: {
    getFileList() {
      getFileList("path=/hour")
        .then((res) => {
          this.hourList = res.data.files;
        })
        .catch((err) => {
          this.$message.warning(err.message)
          console.log(err);
        });
      getFileList("path=/month")
        .then((res) => {
          this.monthList = res.data.files;
        })
        .catch((err) => {
          this.$message.warning(err.message)
          console.log(err);
        });
    },
  },
  created() {
    this.getFileList();
    window.addEventListener("resize", () => {
      this.containerHeight = window.innerHeight - 57;
    });
  },
};
</script>

<style lang="less" scoped>
.container {
  position: absolute;
  top: 57px;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  overflow: auto;
}
.download {
  display: flex;
  justify-content: center;
}
</style>

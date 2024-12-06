<script setup lang="ts">
import {ref, onMounted} from "vue";

const open = ref<boolean>(false);

const showModal = () => {
  open.value = true;
};

const handleOk = (e: MouseEvent) => {
  console.log(e);
  open.value = false;
};

// 图片地址
const imgSrc = '@/assets/img/a.jpg';

// 用来存储动态生成的图片
const images = ref<{ src: string, rotation: number }[]>([]);

// 初始化图片并设置随机的旋转角度
const generateImages = () => {
  const angles = [];
  let angle = 12; // 起始角度
  for (let i = 0; i < 13; i++) {
    angles.push(angle);
    angle = i % 2 === 0 ? angle + 10 : angle - 7;
  }

  // 创建图片数据
  images.value = angles.map(angle => ({
    src: imgSrc,
    rotation: angle,
  }));
};

onMounted(() => {
  generateImages();
});
</script>

<template>
  <!-- 图片详情 -->
  <a-modal v-model:open="open" title="图片详情" @ok="handleOk" :cancelButton="false">
    <img class="showImgInfo" src="@/assets/img/a.jpg" alt="154"/>
  </a-modal>
  <div id="cont">
    <!-- 动态生成的图片 -->
    <img v-for="(image, index) in images"
         :key="index"
         src="@/assets/img/a.jpg"
         class="ima"
         @click="showModal"
         :style="{ transform: 'rotate(' + image.rotation + 'deg)' }"
         alt="154"/>
  </div>
</template>

<style scoped>
.showImgInfo {
  width: 100%;
  height: 100%;
}

#cont {
  width: 100vw;
  height: 100vh; /* 设置一个固定的高度 */
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: flex-start; /* 设置为顶部对齐，以避免图片过多时有不必要的空白 */
  background: url("@/assets/login-bg.png") no-repeat center center fixed;
  overflow-y: auto; /* 开启垂直滚动条 */
  overflow-x: hidden; /* 开启垂直滚动条 */
}


.ima {
  width: 310px;
  padding: 18px;
  margin: 30px;
  background-color: #FFFFFF;
  transition: transform 0.8s;
}

.ima:hover {
  transform: scale(1.5) rotate(0deg) !important;
}
</style>

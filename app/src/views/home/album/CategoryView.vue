<script setup lang="ts">
import {ref, onMounted} from "vue";
import {useRouter} from "vue-router";
import {albumPhotoService} from "@/services/album/category.ts";
import ImageComponent from "@/components/ImageComponent.vue";

const router = useRouter();
const showModal = (category: number) => {
  // 跳转到 Photo 页面并传递 category 参数
  router.push({name: 'Photo', params: {category}});
};


// 用来存储动态生成的图片
const images = ref<{
  src: string,
  category: number,
  rotation: number,
  cover: number
}[]>([]);

// 初始化图片并设置随机的旋转角度
const generateImages = async () => {
  const res = await albumPhotoService()
  const angles = [];
  let angle = 12; // 起始角度
  res.forEach((item: any, i: number) => {
    angles.push(angle);
    angle = i % 2 === 0 ? angle + 10 : angle - 13;
    images.value.push({src: res, category: item.id, cover: item.cover, rotation: angle});
  })
};

onMounted(() => {
  generateImages();
});
</script>

<template>
  <div id="cont">
    <div class="img-box" v-for="(image) in images"
         :style="{ transform: 'rotate(' + image.rotation + 'deg)' }"
    >
      <image-component
          @click="showModal(image.category)"
          :params="{pid: image.cover,categoryId: image.category }"
      />
    </div>
  </div>
</template>

<style scoped>

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


.img-box {
  width: 310px;
  padding: 18px;
  margin: 30px;
  transition: transform 0.8s;
  background: #ddd;
}

.img-box:hover {
  transform: scale(1.2) rotate(0deg) !important;
}
</style>

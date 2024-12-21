<script lang="ts" setup>
import {ref, onMounted} from 'vue';
import {albumPhotoService} from "@/services/album/photo.ts";
import {message} from "ant-design-vue";  // 假设你的请求工具已经封装好了

// 使用 defineProps 定义 props
const props = defineProps<{
  params: {
    [key: string]: any;
  };
}>();

// Reactive state
const imageUrl = ref<string>('');

const fetchImage = async () => {
  const params = {...props.params};  // 获取所有传入的 props
  // 检查 params 中是否包含 pid 和 categoryId
  if (params.pid && params.categoryId) {
    // 如果包含 pid 和 categoryId，调用 albumPhotoService
    const res: any = await albumPhotoService(params);
    imageUrl.value = URL.createObjectURL(res); // 将返回的二进制数据转换为 URL
  } else {
    message.warn('参数错误');
  }
}

// 生命周期钩子
onMounted(fetchImage);
</script>

<template>
  <div>
    <!-- 如果 imageUrl 存在，就显示后端图片；否则显示本地默认图片 -->
    <div class="image">
      <img  v-if="imageUrl" :src="imageUrl" alt="后端图片"/>
      <img class="default-image" v-else src="@/assets/img/a.jpg" alt="本地默认图片"/>
    </div>
  </div>
</template>

<style scoped lang="scss">
.default-image {
  width: 50px;
}

/* 可以根据需要添加样式 */
.image {
  margin-bottom: 15px;

  img {
    width: 100%;
  }
}


</style>

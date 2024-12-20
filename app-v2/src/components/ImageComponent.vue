<script lang="ts" setup>
import {ref, onMounted} from 'vue';
import api from '@/services/api';  // 假设你的请求工具已经封装好了

// 使用 defineProps 定义 props
const props = defineProps<{
  pid: string;
  categoryId: string;
}>();

// Reactive state
const imageUrl = ref<string>('');
const loading = ref<boolean>(true);

const fetchImage = () => {

  // 获取二进制数据
  api.get('/album/photo',
      {pid: props.pid, categoryId: props.categoryId},
  ).then((res: any) => {
    // 如果 res.data 是 ArrayBuffer 类型，需要转换为 Blob
    const blob = new Blob([res.data], { type: 'image/jpeg' });  // 你可以根据实际图片类型来设置 mimeType

    // 创建一个 FileReader 来将 Blob 转换为 Base64
    const reader = new FileReader();

    // 读取成功后，将 Base64 数据赋给 imageUrl
    reader.onloadend = () => {
      imageUrl.value = reader.result as string;  // reader.result 是 Base64 格式
      console.log(imageUrl.value);  // 打印 Base64 数据，确保它是正确的
    };

    // 将 Blob 数据转为 Base64 格式
    reader.readAsDataURL(blob);  // 传入 Blob 对象
    console.log(imageUrl.value);  // 应该以 "data:image/jpeg;base64," 开头

  })

};

// 生命周期钩子
onMounted(fetchImage);
</script>

<template>
  <div >
    <img :src="imageUrl" alt="Fetched from backend"/>
  </div>
</template>

<style scoped>
/* 可以根据需要添加样式 */
img {
  max-width: 100%;
  height: auto;
}
</style>

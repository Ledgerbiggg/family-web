<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onMounted, onBeforeUnmount } from "vue";

const router = useRouter();

// 倒计时剩余时间
const countdown = ref(5);

// 创建一个函数来处理倒计时
const startCountdown = () => {
  const interval = setInterval(() => {
    countdown.value--; // 每秒减少1
    if (countdown.value <= 0) {
      clearInterval(interval); // 倒计时结束，清除定时器
      router.push({ name: "Home" }); // 跳转到登录页面
    }
  }, 1000);
};

onMounted(() => {
  startCountdown(); // 页面加载时开始倒计时
});

onBeforeUnmount(() => {
  // 在组件销毁时清除定时器
  clearInterval(startCountdown);
});
</script>

<template>
  <div class="container">
<!--    <img src="@/assets/404.png" class="bg" alt="404" />-->
    <div class="countdown">
      <p>页面将在 <strong>{{ countdown }}</strong> 秒后跳转到主页面</p>
    </div>
  </div>
</template>
<style scoped>
@import "@/styles/button.css";
* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

img {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.container {
  background-image: url("@/assets/404.png");
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  width: 100vw;
  height: 100vh;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: top center;
  margin: 0 auto;
}


.countdown {
  font-size: 18px;
  color: #fff;
  border-radius: 15px; /* 圆角 */
  display: inline-block;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* 阴影效果 */
  font-weight: bold;
}

.countdown strong {
  font-size: 24px; /* 放大倒计时数字 */
  color: #fff;
  font-family: 'Arial', sans-serif;
}

.countdown p {
  margin: 0;
  font-size: 18px;
  color: #fff;
}
</style>

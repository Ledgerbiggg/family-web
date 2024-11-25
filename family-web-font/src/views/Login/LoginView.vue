<script setup lang="ts">
import {ref} from "vue";
import router from "@/router";


// 定义 user 对象的类型
interface User {
  username: string;
  password: string;
  captcha: string;
}

// 获取 captchaImage 的引用
const captchaImage = ref<HTMLImageElement | null>(null);

// 创建一个响应式的 user 对象
const user = ref<User>({
  username: "",
  password: "",
  captcha: ""
});

// 登录方法
const login = () => {
  console.log(user.value);
}
// 刷新验证码
const refresh = () => {
  if (captchaImage.value && captchaImage.value.src) {
    captchaImage.value.src = "http://localhost:8001/captcha?_t=" + Date.now();
  }
}
// 注册方法
const register = () => {
  // 跳转到注册页面
  router.push("/register");
}

const forgotPassword = () => {
  // 跳转到忘记密码页面
  router.push("/forgot-password");
}
</script>

<template>
  <body>
  <div class="box">
    <h2>登录</h2>
    <div class="input-box">
      <label>手机号</label>
      <input v-model="user.username" type="text"/>
    </div>
    <div class="input-box">
      <label>密码</label>
      <input v-model="user.password" type="password"/>
    </div>
    <div class="input-box captcha-input-box">
      <label>验证码</label>
      <div class="captcha-box">
        <input v-model="user.captcha" type="text"/>
        <div class="img-box">
          <img ref="captchaImage" @click="refresh" src="http://localhost:8001/captcha" alt="captcha">
        </div>
      </div>
    </div>
    <div class="btn-box">
      <a @click="forgotPassword">忘记密码?</a>
      <div>
        <button @click="login">登录</button>
        <button @click="register">注册</button>
      </div>
    </div>
  </div>
  </body>
</template>

<style scoped>
@import '../../styles/login.css';
</style>
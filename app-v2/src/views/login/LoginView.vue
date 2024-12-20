<script setup lang="ts">
import {onMounted, ref} from "vue";
import {message} from "ant-design-vue";
import {useRouter} from "vue-router";
import {User, loginService} from "@/services/login/login.ts";

const router = useRouter()
// 创建一个响应式的 user 对象
const user = ref<User>({
  username: "",
  password: "",
  captcha: ""
});

// 在组件挂载后执行刷新验证码
onMounted(() => {
  refresh();
});

// 获取 captchaImage 的引用
const captchaImage = ref<HTMLImageElement | null>(null);
// 登录方法
const login = async () => {
  if (!user.value.username || !user.value.password || !user.value.captcha) {
    message.warn("请填写完整信息");
    return
  }
  if (await loginService(user.value)) {
    message.success("登录成功");
    await router.push("/home");
    return;
  }
  refresh()
}
// 刷新验证码
const refresh = () => {
  if (captchaImage.value && captchaImage.value.src) {
    captchaImage.value.src = import.meta.env.VITE_BASE_URL + "/captcha?_t=" + Date.now();
  }
}
// 注册方法
const register = () => {
  // 跳转到注册页面
  router.push({name: "Register"});
}

const forgotPassword = () => {
  // 跳转到忘记密码页面
  router.push({name: "ForgotPassword"});
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
        <input v-model="user.captcha" type="text" @keydown.enter="login"/>
        <div class="img-box">
          <img ref="captchaImage" @click="refresh" src="" alt="captcha">
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
@import '@/styles/login.css';
</style>
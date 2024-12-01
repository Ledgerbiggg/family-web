<script setup lang="ts">
import {onMounted, ref} from "vue";
import api from "@/services/api.ts";
import {message} from "ant-design-vue";
import {useRouter} from "vue-router";

const router = useRouter()
// 在组件挂载后执行刷新验证码
onMounted(() => {
  refresh();
});

// 获取 captchaImage 的引用
const captchaImage = ref<HTMLImageElement | null>(null);


// 定义 user 对象的类型
interface User {
  username: string;
  password: string;
  confirmPassword: string;
  captcha: string;
}

// 创建一个响应式的 user 对象
const user = ref<User>({
  username: "",
  password: "",
  confirmPassword: "",
  captcha: ""
});

// 刷新验证码
const refresh = () => {
  if (captchaImage.value && captchaImage.value.src) {
    captchaImage.value.src = import.meta.env.VITE_BASE_URL + "/captcha?_t=" + Date.now();
  }
}

// 登录方法
const backLogin = () => {
  router.push({name: "Login"});
}

// 注册方法
const register = () => {
  if (!user.value.username ||
      !user.value.captcha ||
      !user.value.password ||
      !user.value.confirmPassword ||
      user.value.password !==
      user.value.confirmPassword) {
    message.warn("请填写完整信息");
    return
  }
  api.post("/register", {...user.value}).then((res: any) => {
    // 跳转到登录
    message.success("注册成功,请登录");
    router.push("/login");
  }).catch((rea: any) => {
    refresh()
  })
}
</script>

<template>
  <body>
  <div class="box">
    <h2>注册</h2>
    <div class="input-box">
      <label>手机号</label>
      <input v-model="user.username" type="text"/>
    </div>
    <div class="input-box">
      <label>密码</label>
      <input v-model="user.password" type="password"/>
    </div>
    <div class="input-box">
      <label>确认密码</label>
      <input v-model="user.confirmPassword" type="password"/>
    </div>
    <div class="input-box captcha-input-box">
      <label>验证码</label>
      <div class="captcha-box">
        <input v-model="user.captcha" type="text"/>
        <div class="img-box">
          <img ref="captchaImage" @click="refresh" src="" alt="captcha">
        </div>
      </div>
    </div>
    <div class="btn-box">
      <div>
        <button @click="backLogin">返回登录</button>
        <button @click="register">注册</button>
      </div>
    </div>
  </div>
  </body>
</template>

<style scoped>
@import '@/styles/login.css';
</style>
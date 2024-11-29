<script setup lang="ts">
import {onMounted, ref} from "vue";
import router from "@/router";
import api from "@/services/api.ts";
import {message} from "ant-design-vue";

// 在组件挂载后执行刷新验证码
onMounted(() => {
  refresh();
});
// 获取 captchaImage 的引用
const captchaImage = ref<HTMLImageElement | null>(null);

// 定义 user 对象的类型
interface User {
  username: string;
  reaName: string;
  captcha: string;
}

// 创建一个响应式的 user 对象
const user = ref<User>({
  username: "",
  reaName: "",
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
  router.push("/login");
}

// 注册方法
const verify = () => {
  if (!user.value.username || !user.value.reaName || !user.value.captcha) {
    message.warn("请填写完整信息");
    return
  }
  api.post("/verify", {...user.value}).then((res: any) => {
    console.log("red")
  }).catch((rea: any) => {
    refresh()
  })
}
</script>

<template>
  <body>
  <div class="box">
    <h2>忘记密码</h2>
    <div class="input-box">
      <label>手机号</label>
      <input v-model="user.username" type="text"/>
    </div>
    <div class="input-box">
      <label>您的真实姓名</label>
      <input v-model="user.reaName" type="text"/>
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
        <button @click="verify">验证</button>
      </div>
    </div>
  </div>
  </body>
</template>

<style scoped>
@import '../../styles/login.css';
</style>
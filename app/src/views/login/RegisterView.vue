<script setup lang="ts">
import {onMounted, ref} from "vue";
import {message} from "ant-design-vue";
import {useRouter} from "vue-router";
import {registerService, RegisterUser} from "@/services/login/register.ts";

const router = useRouter()
// 在组件挂载后执行刷新验证码
onMounted(() => {
  refresh();
});

const captchaImage = ref<HTMLImageElement | null>(null);

const user = ref<RegisterUser>({
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
const register = async () => {
  if (!user.value.username ||
      !user.value.captcha ||
      !user.value.password ||
      !user.value.confirmPassword ||
      user.value.password !==
      user.value.confirmPassword) {
    message.warn("请填写完整信息");
    return
  }
  let res = await registerService(user.value)
  console.log(res, '0000')
  if (res) {
    // 跳转到登录
    message.success("注册成功,请登录");
    await router.push("/login");
  } else {
    refresh()
  }
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
        <input v-model="user.captcha" type="text" @keydown.enter="register"/>
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
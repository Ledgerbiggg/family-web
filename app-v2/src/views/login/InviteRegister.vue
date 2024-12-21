<script setup lang="ts">
import {onMounted, ref} from "vue";
import {message} from "ant-design-vue";
import {useRoute, useRouter} from "vue-router";
import {InviteInfo, inviteInfoService, inviteRegisterService, InviteUser} from "@/services/login/invite.ts";

const route = useRoute()
const router = useRouter()

// 在组件挂载后执行刷新验证码
onMounted(() => {
  refresh();
  // 获取uid
  getInviteInfo();
});
// 获取 captchaImage 的引用
const captchaImage = ref<HTMLImageElement | null>(null);

// 定义 inviteInfo 对象的类型
const InviteInfo = ref<InviteInfo>({
  inviterPhone: "",
  inviterRealName: "",
  invitedAdmin: false,
  description: ""
})

// 创建一个响应式的 user 对象
const user = ref<InviteUser>({
  inviteUid: "",
  realName: "",
  username: "",
  captcha: ""
});

// 刷新验证码
const refresh = () => {
  if (captchaImage.value && captchaImage.value.src) {
    captchaImage.value.src = import.meta.env.VITE_BASE_URL + "/captcha?_t=" + Date.now();
  }
}

// 注册方法
const verify = () => {
  if (!user.value.realName || !user.value.username || !user.value.captcha) {
    message.warn("请填写完整信息");
    return
  }
  user.value.inviteUid = String(route.query.uid || '');
  if (inviteRegisterService(user.value)) {
    message.success("注册成功,初始密码:123456");
    router.push({name: "Login"});
  } else {
    refresh()
  }
}
// 获取邀请信息
const getInviteInfo = async () => {
  let res = await inviteInfoService(route.query.uid)
  if (res) {
    InviteInfo.value = res
  } else {
    await router.push({name: "Login"});
  }
}
</script>

<template>
  <body>
  <div class="box">
    <h2>邀请注册</h2>
    <div class="invite-box">
      <div class="inviter">
        <div>邀请者: {{ InviteInfo.inviterRealName }}</div>
        <div>邀请者手机号: {{ InviteInfo.inviterPhone }}</div>
      </div>
      <div v-if="InviteInfo.invitedAdmin" class="admin-invite">邀请你成为管理员</div>
      <div class="description">{{ InviteInfo.description }}</div>
    </div>
    <div class="input-box">
      <label>您的<strong>真实姓名</strong></label>
      <input v-model="user.realName" type="text"/>
    </div>
    <div class="input-box">
      <label>您的<strong>手机号</strong></label>
      <input v-model="user.username" type="text"/>
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
        <button @click="verify">注册</button>
      </div>
    </div>
  </div>
  </body>
</template>

<style scoped lang="less">
@import '@/styles/login.css';

.box {
  width: 350px;
  height: 500px;
}

.invite-box {
  width: 70%;
  height: auto; /* 让内容自适应高度 */
  max-height: 80px; /* 限制最大高度 */
  overflow-y: auto; /* 允许纵向滚动 */
  border-radius: 12px; /* 圆角效果 */
  background: linear-gradient(145deg, #f0f0f0, #d9d9d9); /* 渐变背景 */
  box-shadow: -5px -5px 10px rgba(164, 164, 164, 0.4), 5px 5px 10px rgba(255, 255, 255, 0.8); /* 柔和阴影 */
  padding: 10px; /* 内边距 */
  margin: 0 0 15px 0; /* 外边距 */
  font-family: 'Arial', sans-serif; /* 字体 */
  color: #333; /* 字体颜色 */
  line-height: 1.6; /* 行高 */
  font-size: 14px; /* 基本字体大小 */
  transition: all 0.3s ease; /* 添加过渡效果 */
}

.invite-box:hover {
  box-shadow: -5px -5px 15px rgba(164, 164, 164, 0.6), 5px 5px 15px rgba(255, 255, 255, 0.9); /* 悬浮时更强的阴影效果 */
}

.invite-box div {
  margin-bottom: 8px; /* 每个子元素之间的间距 */
}

.invite-box .inviter {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold; /* 加粗邀请人文本 */
  color: #2d87f0; /* 邀请人用户名使用蓝色 */
}

.invite-box .admin-invite {
  color: #ff6347; /* 邀请成为管理员的文本使用橙红色 */
}

.invite-box .description {
  color: #666; /* 描述文字使用浅灰色 */
}

.invite-box::-webkit-scrollbar {
  width: 8px; /* 设置滚动条宽度 */
}

.invite-box::-webkit-scrollbar-thumb {
  background-color: #a0a0a0; /* 设置滚动条的颜色 */
  border-radius: 10px; /* 圆角效果 */
}

.invite-box::-webkit-scrollbar-track {
  background-color: #e0e0e0; /* 设置滚动条轨道的颜色 */
}


/* 隐藏滚动条 */
.invite-box::-webkit-scrollbar {
  display: none;
}

strong {
  color: red;
}
</style>
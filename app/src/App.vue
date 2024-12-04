<template>
  <div class="animation">
    <router-view v-slot="{ Component }">
      <transition name="horse-pull" :appear="true">
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<style scoped>
.animation {
  width: 100vw;
  height: 100vh;
  position: relative;
}

.horse-pull-enter-active,
.horse-pull-leave-active {
  transition: transform 0.5s ease-in-out;
}

.horse-pull-enter,
.horse-pull-leave-to /* .horse-pull-leave-active in <2.1.8 */ {
  transform: translateX(100%); /* 初始状态：页面在视口之外 */
}

.horse-pull-enter-to,
.horse-pull-leave {
  transform: translateX(0); /* 结束状态：页面正常显示 */
}

.horse-pull-enter-active {
  animation: horsePull 0.8s forwards;
}

@keyframes horsePull {
  0% {
    transform: translateX(100%); /* 页面从右侧进场 */
  }
  100% {
    transform: translateX(0); /* 页面平滑进入 */
  }
}
</style>

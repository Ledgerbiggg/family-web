<script setup lang="ts">
import {Swiper, SwiperSlide} from 'swiper/vue';
import 'swiper/swiper-bundle.css';


// 可以在这里配置 Swiper 的参数
const swiperOptions =
    {
      effect: 'coverflow',
      grabCursor: true,
      centeredSlides: true,
      slidesPerView: 3, // 同时显示 3 张卡片
      coverflowEffect: {
        rotate: 0,
        stretch: 10,
        depth: 160,
        modifier: 2,
        slideShadows: true
      },// 如果需要前进后退按钮
      pagination: {
        el: '.swiper-pagination',
        clickable: true
      },
      navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev'
      },
    };
</script>

<template>
  <div class="swiper-container">
    <Swiper v-bind="swiperOptions">
      <SwiperSlide>
        <div class="card">
          <div class="sliderText">
            <h3>01-Login</h3>
          </div>
          <div class="demo-img">
            <img src="@/assets/img/01-login.jpg" alt="">
          </div>
          <div class="content">
            <p>1. 请左右滑动选择不同卡片。如果你要获取源码，请记住图片上方的编号与名字。</p>
            <p>2. 点击“获取源码”获取该项目源码。</p>
            <p>3. 点击“效果演示”查看效果预览。</p>
            <p>4. 点击“查看视频”一起打卡学习~</p><br>
            <a href="https://github.com/local-host-8080/demo-html-css">获取源码</a>
            <a href="demo/01/index.html">效果演示</a>
            <a href="https://www.bilibili.com/video/av60343435">查看视频</a>
          </div>
        </div>
      </SwiperSlide>
      <SwiperSlide>
        <div class="card">
          <div class="sliderText">
            <h3>02-Login</h3>
          </div>
          <div class="demo-img">
            <img src="@/assets/img/01-login.jpg" :alt="''">
          </div>
          <div class="content">
            <p>1. 请左右滑动选择不同卡片。如果你要获取源码，请记住图片上方的编号与名字。</p>
            <p>2. 点击“获取源码”获取该项目源码。</p>
            <p>3. 点击“效果演示”查看效果预览。</p>
            <p>4. 点击“查看视频”一起打卡学习~</p><br>
            <a href="https://github.com/local-host-8080/demo-html-css">获取源码</a>
            <a href="demo/01/index.html">效果演示</a>
            <a href="https://www.bilibili.com/video/av60343435">查看视频</a>
          </div>
        </div>
      </SwiperSlide>
      <SwiperSlide>
        <div class="card">
          <div class="sliderText">
            <h3>03-Login</h3>
          </div>
          <div class="demo-img">
            <img src="@/assets/img/01-login.jpg" :alt="''">
          </div>
          <div class="content">
            <p>1. 请左右滑动选择不同卡片。如果你要获取源码，请记住图片上方的编号与名字。</p>
            <p>2. 点击“获取源码”获取该项目源码。</p>
            <p>3. 点击“效果演示”查看效果预览。</p>
            <p>4. 点击“查看视频”一起打卡学习~</p><br>
            <a href="https://github.com/local-host-8080/demo-html-css">获取源码</a>
            <a href="demo/01/index.html">效果演示</a>
            <a href="https://www.bilibili.com/video/av60343435">查看视频</a>
          </div>
        </div>
      </SwiperSlide>
      <SwiperSlide>
        <div class="card">
          <div class="sliderText">
            <h3>04-Login</h3>
          </div>
          <div class="demo-img">
            <img src="@/assets/img/01-login.jpg" :alt="''">
          </div>
          <div class="content">
            <p>1. 请左右滑动选择不同卡片。如果你要获取源码，请记住图片上方的编号与名字。</p>
            <p>2. 点击“获取源码”获取该项目源码。</p>
            <p>3. 点击“效果演示”查看效果预览。</p>
            <p>4. 点击“查看视频”一起打卡学习~</p><br>
            <a href="https://github.com/local-host-8080/demo-html-css">获取源码</a>
            <a href="demo/01/index.html">效果演示</a>
            <a href="https://www.bilibili.com/video/av60343435">查看视频</a>
          </div>
        </div>
      </SwiperSlide>
    </Swiper>
  </div>
</template>

<style scoped>
@import "@/styles/cards.css";

.swiper-container {
  width: 100%;
  height: 100vh;
  background-image: url("@/assets/home-bg.jpg");
}




</style>

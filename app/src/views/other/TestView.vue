<script setup lang="ts">
import {onMounted, ref} from 'vue';
import Button from "@/components/Button.vue";
import http from "@/services/api.ts";
import {useRouter} from "vue-router";

const router = useRouter()

// Define the items and their details (this should match your image data and text)
const items = ref([
      {
        title: "暂无",
        description: "这边什么都没有哦",
        image: new URL('@/assets/img/01.png', import.meta.url).href,
        path: '/home'
      },
    ]
);
let interval: any = null;
let currIndex = ref(0);
let width = 0;
let height = 0;
const intervalTime = 3000;

const shellSlider = ref(null);
const shellBox = ref(null);

const itemWidth = ref(0);
const itemHeight = ref(0);

// 根据窗口大小调整尺寸
const resize = () => {
  width = Math.max(window.innerWidth * 0.2, 275);
  height = Math.max(window.innerHeight * 0.5, 355);
  itemWidth.value = width;
  itemHeight.value = height;

  const totalWidth = itemWidth.value * items.value.length;
  if (shellSlider.value) {
    shellSlider.value.style.width = `${totalWidth}px`;
  }
};

// 根据当前索引移动滑动数据
const move = (index: number) => {
  if (index < 1) index = items.value.length;
  if (index > items.value.length) index = 1;
  currIndex.value = index;

  const allItems = shellSlider.value?.children || [];
  Array.from(allItems).forEach((item: HTMLElement, i: number) => {
    const box = item.querySelector('.frame') as HTMLElement;
    const front = box?.querySelector('.front') as HTMLElement;  // 获取 .front 元素
    const back = box?.querySelector('.back') as HTMLElement;  // 获取 .front 元素
    if (i === index - 1) {
      item.classList.add('item--active');
      box.style.transform = 'perspective(1200px)';
      if (front && shellBox.value && back) {
        shellBox.value.style.backgroundImage = window.getComputedStyle(front).backgroundImage;
        shellBox.value.style.backgroundImage = window.getComputedStyle(back).backgroundImage;
      }
    } else {
      item.classList.remove('item--active');
      box.style.transform = `perspective(1200px) rotateY(${i < index - 1 ? 40 : -40}deg)`;
    }
  });

  if (shellSlider.value) {
    shellSlider.value.style.transform = `translate3d(${(index * -width) + (width / 2) + window.innerWidth / 2}px, 0, 0)`;
  }
};

// 滑动到上一张
const prev = () => {
  move(--currIndex.value);
  timer();
};

// 滑动到下一张
const next = () => {
  move(++currIndex.value);
  timer();
};


// 启动定时器
const timer = () => {
  clearInterval(interval);
  interval = setInterval(() => {
    move(++currIndex.value);
  }, intervalTime);
};

// 清除滑动动画效果
const clearIntervalFn = () => {
  clearInterval(interval);
}

// 展示当前的slide
const showCurrentSlide = (c: number) => {
  move(c);
  currIndex.value = c
}

// 进入对应的路由
const toPath = (item: any) => {
  router.push(item.path);
}
// 获取后台的数据
const getCardData = () => {
  http.get("/home/cards").then((res: any) => {
    if (res) {
      //image: new URL('@/assets/img/01.png', import.meta.url).href,
      res.forEach((i: any) => {
        i.image = new URL(`/src/assets/img/${i.image}`, import.meta.url).href;
      });
      items.value = res
    }
  }).finally(() => {
    // 数据获取后执行这些操作
    startAnimation()
  })
}
// 开始动画效果
const startAnimation = () => {
  resize();
  move(1);
  window.addEventListener('resize', resize);
  timer();
}
const flipAndNavigate = (item: any) => {
  // 找到当前点击的item
  const currentItem = shellSlider.value?.children[currIndex.value - 1] as HTMLElement;
  const box = currentItem?.querySelector('.frame') as HTMLElement;
  const front = box?.querySelector('.front') as HTMLElement;

  // 确保找到 .frame 元素后进行动画
  if (box && front) {
    // 1. 设置 box 翻转效果
    box.style.transition = 'transform 1s'; // 翻转动画持续时间

    // 2. 设置 front 翻开效果
    front.style.transition = 'transform 1s'; // front 翻开动画持续时间
    front.style.transformOrigin = 'left'; // 设置旋转中心在左边
    front.style.transform = 'rotateY(-160deg)'; // 让 front 翻开（绕 Y 轴旋转 180 度）

    // 2. 等待翻转动画完成后执行缩放动画
    setTimeout(() => {
      // 变大动画
      box.style.transition = 'transform 1s, opacity 1s'; // 添加变大动画
      box.style.transform = 'scale(5)'; // 缩放到2倍大
      box.style.opacity = '0'; // 如果需要淡出，也可以设置透明度

      // 3. 等待动画完成后跳转页面
      setTimeout(() => {
        toPath(item); // 页面跳转
      }, 1000); // 变大动画时间为1s，确保在动画完成后跳转
    }, 1000); // 翻转动画时间为1s，确保在翻转后开始缩放
  }
};



onMounted(() => {
  // 数据获取后执行这些操作
  startAnimation()
  getCardData()
});
</script>

<template>
  <div class="shell-box" ref="shellBox">
    <div class="shell">
      <div class="shell_body">
        <div class="button">
          <div class="button-left">
            <Button @click="prev"></Button>
          </div>
          <div class="button-right">
            <Button @click="next"></Button>
          </div>
        </div>
        <div class="shell_slider" ref="shellSlider"
             @mouseleave="timer"
             @mouseenter="clearIntervalFn">
          <div class="item"
               v-for="(item, index) in items"
               :key="index"
               :style="{ width: itemWidth + 'px', height: itemHeight + 'px' }"
               @mouseover="showCurrentSlide(index+1)"
               @click="flipAndNavigate(item)">
            <div class="frame">
              <div class="box front" :style="{ backgroundImage: 'url(' + (item.image) + ')' }">
                <h1>{{ item.title }}</h1>
                <span>{{ item.description }}</span>
              </div>
              <div class="box back" :style="{ backgroundImage: 'url(' + (item.image) + ')' }">
                <h1>{{ item.title }}</h1>
                <span>{{ item.description }}</span>
              </div>
              <div class="box left" :style="{ backgroundImage: 'url(' + (item.image) + ')' }"></div>
              <div class="box right" :style="{ backgroundImage: 'url(' + (item.image) + ')' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
* {
  padding: 0;
  margin: 0;
  font-family: "Source Sans Pro", sans-serif;
}


/* 设置html和body元素为flex布局，水平和垂直居中对齐，高度为100vh，背景图大小为cover，溢出隐藏，背景图过渡动画时间为0.7秒 */
.shell-box {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-size: cover;
  overflow: hidden;
  transition: background-image .7s ease-in-out;
}

/* 设置.shell元素为相对定位，flex布局，水平和垂直居中对齐，宽度和高度为100%，盒模型为border-box，背景颜色为rgba(99, 99, 99, 0.8) */
.shell {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  background: rgba(99, 99, 99, 0.8);
}

/* 设置.button元素为flex布局，两端对齐，宽度为380px，绝对定位，左侧偏移量为50%，水平居中，底部偏移量为-80px */
.button {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 380px;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  bottom: -80px;
}

.button-left {
  transform: scaleX(-1) scale(2); /* 水平翻转并且放大 2 倍 */
  transform-origin: center; /* 确保翻转和缩放是从中心进行的 */
}

.button-right {
  transform: scale(2); /* 水平翻转并且放大 2 倍 */
  transform-origin: center; /* 确保翻转和缩放是从中心进行的 */
}


/* 设置.prev和.next元素中的i元素字体大小为90px，颜色为#fff，光标为指针，文字阴影为0 0 10px #ffffff */
.prev i, .next i {
  font-size: 90px;
  color: #fff;
  cursor: pointer;
  text-shadow: 0 0 10px #ffffff;
}

/* 设置.shell_body元素宽度为100%，缩放为0.8倍，上内边距为20px，下内边距为150px */
.shell_body {
  width: 100%;
  transform: scale(.8);
  padding: 20px 0 150px 0;
}

/* 设置.shell_slider元素为相对定位，过渡动画时间为1秒，背景为透明 */
.shell_slider {
  width: 100%;
  position: relative;
  transition: transform 1s ease-in-out;
  background: transparent;
}

/* 设置.item元素为相对定位，左浮动，左右外边距为20px */
.item {
  position: relative;
  float: left;
  cursor: pointer;
}

/* 设置.frame元素为相对定位，宽度和高度为100%，过渡动画时间为1秒，3D变换模式为保留3D效果 */
.frame {
  position: relative;
  width: 100%;
  height: 100%;
  transition: transform 1s ease-in-out;
  transform-style: preserve-3d;
}

/* 设置.frame元素的伪元素为绝对定位，底部偏移量为-16%，宽度为100%，高度为60px，背景颜色为#ffffff1c，盒阴影为0px 0px 15px 5px #ffffff1c，3D变换为绕X轴旋转90度并向上平移20px */
.frame:after {
  content: "";
  position: absolute;
  bottom: -16%;
  width: 100%;
  height: 60px;
  background: #ffffff1c;
  box-shadow: 0 0 15px 5px #ffffff1c;
  transform: rotateX(90deg) translate3d(0px, -20px, 0px);
}

/* 设置.box元素为flex布局，纵向排列，水平和垂直居中对齐，绝对定位，宽度和高度为100%，边框为4px实心白色，透视效果为1000px，3D变换模式为保留3D效果 */
.box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: absolute;
  width: 100%;
  height: 100%;
  border: 4px solid #fff;
  perspective: 1000px;
  transform-style: preserve-3d;
}

/* 设置.box元素中的h1和span元素颜色为#fff，Z轴平移距离为20px */
.box h1, .box span {
  color: #fff;
  transform: translateZ(20px);
}

/* 设置.box元素中的h1元素文字阴影为0 0 30px #1f05b4，字体大小为100px */
.box h1 {
  text-shadow: 0 0 30px #1f05b4;
  font-size: 100px;
}

/* 设置.box元素中的span元素为绝对定位，底部偏移量为20px，左右内边距为25px，文字阴影为0 0 10px #1f05b4 */
.box span {
  position: absolute;
  bottom: 20px;
  padding: 0 25px;
  text-shadow: 0 0 10px #1f05b4;
}

/* 设置.front、.left和.right元素的盒阴影为0 0 50px #ffffff，背景图大小为cover */
.front, .left, .right {
  box-shadow: 0 0 50px #ffffff;
  background-size: cover;
}

/* 设置.left和.right元素的顶部偏移量为0，宽度为60px，背面不可见 */
.right, .left {
  top: 0;
  width: 60px;
  backface-visibility: hidden;
}

/* 设置.left元素的左侧偏移量为0，左边框宽度为5px，3D变换为向右平移1px，Z轴平移-60px，绕Y轴逆时针旋转90度，变换原点为左侧 */
.left {
  left: 0;
  border-left-width: 5px;
  transform: translate3d(1px, 0, -60px) rotateY(-90deg);
  transform-origin: 0%;
}

/* 设置.right元素的右侧偏移量为0，右边框宽度为5px，3D变换为向左平移1px，Z轴平移-60px，绕Y轴顺时针旋转90度，变换原点为右侧 */
.right {
  right: 0;
  border-right-width: 5px;
  transform: translate3d(-1px, 0, -60px) rotateY(90deg);
  transform-origin: 100%;
}

</style>
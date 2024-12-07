<script setup lang="ts">
</script>

<template>
  <div class="shell">
    <div>
      <img src="@/assets/img/a.jpg" alt=""/>
      <span>1</span>
    </div>
    <div>
      <img src="@/assets/img/b.jpg" alt=""/>
      <span>2</span>
    </div>
    <div>
      <img src="@/assets/img/d.jpg" alt=""/>
      <span>3</span>
    </div>
    <div>
      <img src="@/assets/img/d.jpg" alt=""/>
      <span>3</span>
    </div>
    <div>
      <img src="@/assets/img/d.jpg" alt=""/>
      <span>3</span>
    </div>
    <div>
      <img src="@/assets/img/e.jpg" alt=""/>
      <span>4</span>
    </div>
    <div>
      <img src="@/assets/img/e.jpg" alt=""/>
      <span>4</span>
    </div>
    <div>
      <img src="@/assets/img/e.jpg" alt=""/>
      <span>4</span>
    </div>
    <div>
      <img src="@/assets/img/e.jpg" alt=""/>
      <span>4</span>
    </div>
    <div>
      <img src="@/assets/img/e.jpg" alt=""/>
      <span>4</span>
    </div>
  </div>
</template>
<style scoped>
/* 使用 grid 布局实现瀑布流效果 */
.shell {
  display: grid;
  grid-template-columns: repeat(3, 1fr); /* 三列布局 */
  gap: 20px; /* 设置列和行的间隔 */
  padding: 20px;
  overflow-y: auto; /* 启用竖向滚动条 */
  height: 100vh; /* 设置容器高度为视口高度 */
}

.shell > div {
  display: flex;
  flex-direction: column;
  break-inside: avoid; /* 防止图片被截断 */
  margin-bottom: 20px; /* 设置每个项目之间的间距 */
}

.shell img {
  width: 100%; /* 图片宽度自适应容器 */
  height: auto; /* 保持图片的原始高度和比例 */
  display: block;
}

.shell span {
  display: block;
  text-align: center;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  padding: 5px 0;
}
</style>
<script setup lang="ts">
</script>

<template>
  <div class="shell-box">
    <div class="shell">
      <div class="image"><img src="@/assets/img/a.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/a.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/a.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/a.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/b.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/d.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/c.jpg" alt=""></div>
      <div class="image"><img src="@/assets/img/a.jpg" alt=""></div>
    </div>
  </div>
</template>
<style scoped>
*{
  box-sizing: border-box;
}
.shell-box {
  background-color: rgba(130, 140, 250, 0.2);
  display: flex;
  justify-content: center;
}

.shell {
  max-width: 1300px;
  column-count: 5;
  column-gap: 15px;
}

.image {
  margin-bottom: 15px;
}

.image img {
  width: 100%;
}

@media (max-width: 1200px) {
  .shell {
    column-count: 4;
  }
}

@media (max-width: 850px) {
  .shell {
    column-count: 3;
  }
}

@media (max-width: 600px) {
  .shell {
    column-count: 2;
  }
}
</style>
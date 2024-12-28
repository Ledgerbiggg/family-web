<script setup lang="ts">
import ImageComponent from "@/components/ImageComponent.vue";
import {albumPhotoListService} from "@/services/album/photo.ts";
import {onMounted, ref} from "vue";
import {useRoute} from "vue-router";

const route = useRoute();
let currentCategory: number | null = null;  // 你可以定义为 number 或 null
const photos = ref<any[]>([]);

const getPhotoList = async () => {
  currentCategory = Number(route.params.category);  // 显式转换为 number
  photos.value = await albumPhotoListService(currentCategory);
};

onMounted(() => {
  getPhotoList()
})
</script>

<template>
  <div class="shell-box">
    <div class="shell">
      <div class="shell-item">
        <image-component
            :params="{pid: item.id, categoryId: item.categoryID}"
            v-for="item in photos"/>
      </div>
    </div>
  </div>
</template>
<style scoped lang="scss">
* {
  box-sizing: border-box;
}

.shell-box {
  height: 100vh;
  background-color: rgba(130, 140, 250, 0.2);
  display: flex;
  justify-content: center;

  .shell {
    max-width: 1300px;
    column-count: 5;
    column-gap: 15px;
  }
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
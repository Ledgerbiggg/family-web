<script lang="ts" setup>
import {
  Camera, User, FolderOpened, Picture, Refresh, UserFilled, Share
} from '@element-plus/icons-vue'
import {useRouter} from "vue-router";

const router = useRouter()

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const handleSelect = (key: string, _: string[]) => {
  let path = '/management/' + key
  router.push(path)
}
</script>
<template>
  <div class="menu">
    <el-row class="tac" type="flex" justify="start" align="top">
      <!-- 左侧导航栏 -->
      <el-col :span="2.3" class="sidebar">
        <el-menu
            default-active="2"
            class="el-menu-vertical-navigation"
            @open="handleOpen"
            @close="handleClose"
            @select="handleSelect"
        >
          <el-sub-menu index="album">
            <template #title>
              <el-icon>
                <Camera/>
              </el-icon>
              <span>相册</span>
            </template>
            <el-menu-item index="category-management">
              <el-icon>
                <FolderOpened/>
              </el-icon>
              <span>分类管理</span>
            </el-menu-item>
            <el-menu-item index="photo-management">
              <el-icon>
                <Picture/>
              </el-icon>
              <span>图片管理</span>
            </el-menu-item>
            <el-menu-item index="sync-data-management">
              <el-icon>
                <Refresh/>
              </el-icon>
              <span>同步数据</span>
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="user">
            <template #title>
              <el-icon>
                <User/>
              </el-icon>
              <span>人员</span>
            </template>
            <el-menu-item index="user-management">
              <el-icon>
                <UserFilled/>
              </el-icon>
              <span>人员管理</span>
            </el-menu-item>
            <el-menu-item index="invite-member-managements">
              <el-icon>
                <Share/>
              </el-icon>
              <span>邀请成员</span>
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-col>

      <!-- 右侧子路由视图 -->
      <el-col :span="21">
        <div class="content">
          <router-view/>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.menu {
  width: 100%;
  height: 100vh;
  overflow: hidden;

  .content {
    width: 100%;
    height: 88vh;
    overflow-y: auto;
    padding: 3vh;
    background: #d2f4ff;
  }
}

.el-menu-vertical-navigation:not(.el-menu--collapse) {
  width: 200px;
  height: 100vh;
}
</style>
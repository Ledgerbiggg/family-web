<script setup lang="ts">
import {ref} from 'vue'
import {Search} from '@element-plus/icons-vue'

// 搜索
const search = ref('') // 搜索关键字
const state = ref('')
const tags = ref([])
// 分页
const currentPage4 = ref(4) // 当前页
const pageSize4 = ref(100) // 每页条数
const totalPage = ref(400)
const handleSizeChange = (val: number) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
  console.log(`current page: ${val}`)
}

const tableData = [
  {
    name: '测试分类',
    coverPic: '1.png',
    description: 'test',
    viewCount: 0,
    status: 'archived',
    createdBy: '2',
    createdTime: '2024-12-07 14:21:33',
    updatedTime: '2024-12-21 20:48:16',
  },
]
const statesOptions = [
  {
    value: 'active',
    label: '激活',
  },
  {
    value: 'inactive',
    label: '非激活',
  },
  {
    value: 'archived',
    label: '已归档',
  },
]
const tagsOptions = [
  {
    value: 'active',
    label: '激活',
  },
  {
    value: 'inactive',
    label: '非激活',
  },
  {
    value: 'archived',
    label: '已归档',
  },
]
// 格式化status字段
const statusFormatter = (_: any, __: any, cellValue: string) => {
  const statusMap: { [key: string]: string } = {
    active: '激活',
    inactive: '非激活',
    archived: '已归档'
  };

  return statusMap[cellValue as keyof typeof statusMap] || cellValue;
}

</script>
<template>
  <div class="container">
    <div class="search">
      <!--   分类下拉筛选框   -->
      <el-select
          placeholder="分类状态"
          size="large"
          style="width: 240px;margin: 5px"
          v-model="state"
      >
        <el-option
            v-for="item in statesOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
      </el-select>
      <!--   标签下拉筛选框     -->
      <el-select
          placeholder="标签"
          size="large"
          style="width: 240px;margin: 5px"
          v-model="tags"
          multiple
          clearable
          collapse-tags
          :max-collapse-tags="1"
      >
        <el-option
            v-for="item in tagsOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
      </el-select>
      <el-input
          v-model="search"
          style="width: 240px;"
          size="large"
          placeholder="筛选过滤"
          :suffix-icon="Search"
      />
      <el-button size="large" type="primary" style="margin-left: 10px;">搜索</el-button>
    </div>

    <div class="table-container">
      <el-table border :data="tableData" style="width: 100%; height: 95%; max-width: 100%;">
        <el-table-column fixed prop="name" label="分类名字" width="150"/>
        <el-table-column prop="coverPic" label="分类封面" width="120"/>
        <el-table-column prop="description" label="分类描述"/>
        <el-table-column prop="viewCount" label="浏览次数" width="120"/>
        <el-table-column prop="status" label="分类状态" width="120" :formatter="statusFormatter"/>
        <el-table-column prop="createdTime" label="创建时间" width="120"/>
        <el-table-column prop="updatedTime" label="更新时间" width="120"/>
      </el-table>
    </div>

    <div class="pagination-container">
      <el-pagination
          v-model:current-page="currentPage4"
          v-model:page-size="pageSize4"
          :page-sizes="[10, 20, 30, 40]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="totalPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.search {
  padding: 10px;
  background-color: #f5f5f5;
}

.table-container {
  flex: 1;
  overflow-x: auto; /* 允许横向滚动 */
  overflow-y: hidden; /* 禁止纵向滚动 */
}

.pagination-container {
  display: flex;
  justify-content: flex-end; /* 使分页按钮对齐到右侧 */
  padding: 10px;
}

</style>

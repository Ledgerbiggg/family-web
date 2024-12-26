import { createApp } from 'vue'
import App from './App.vue'

import router from './router/index.ts'  // 引入路由
// 引入 Ant Design Vue 组件库
import Antd from 'ant-design-vue'
// 引入 Element Plus
import ElementPlus from 'element-plus';
// 引入 Element Plus 及其样式
import 'element-plus/dist/index.css';

const app = createApp(App)
app.use(Antd) // 在 Vue 应用中使用 Ant Design
app.use(ElementPlus);
app.use(router)  // 在 Vue 应用中使用 Vue Router

app.mount('#app')

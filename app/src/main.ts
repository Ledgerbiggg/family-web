import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index.ts'  // 引入路由

const app = createApp(App)

app.use(router)  // 在 Vue 应用中使用 Vue Router
app.mount('#app')

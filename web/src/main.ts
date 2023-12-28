import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import 'virtual:uno.css'

const app = createApp(App)

//引入router
import { router } from './routes'
app.use(router)

//引入pinia
import { createPinia } from 'pinia'
const pinia = createPinia()
app.use(pinia)

//引入i8n
import { i18n } from './i18n'
app.use(i18n)

//全局挂载lang
// import { useLanguageStore } from '@/pinia'
// app.config.globalProperties.$lang = useLanguageStore()

//挂载app实例
app.mount('#app')

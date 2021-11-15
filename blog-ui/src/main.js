import { createApp } from 'vue'
import store from './store'
import router from './router'
import App from './App.vue'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css';

let app = createApp(App)
app.use(Antd)
app.use(store)
app.use(router)
app.mount('#app')

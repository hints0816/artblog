import { createApp } from 'vue'
import App from './App.vue'
import axios from './plugin/http.js'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

const app = createApp(App).use(Antd)
app.config.globalProperties.$http = axios
app.mount('#app')
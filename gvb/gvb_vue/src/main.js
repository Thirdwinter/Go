import { createApp } from 'vue';
import App from './App.vue';
import { router } from './router/index';
import { createPinia } from "pinia";
import axios from 'axios';
import naive from 'naive-ui/es/preset';
import Antd from 'ant-design-vue';
import './assets/css/style.css';
import { UserOutlined,LockOutlined} from '@ant-design/icons-vue';



const pinia = createPinia();
axios.defaults.baseURL = 'http://localhost:8080/api/v1';
const app = createApp(App)
app.component('UserOutlined', UserOutlined);
app.component('LockOutlined', LockOutlined);

app.use(router)
app.use(pinia)
app.use(Antd)
app.use(naive)
app.provide('axios', axios)
app.mount('#app');



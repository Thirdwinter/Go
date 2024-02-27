import { createApp } from 'vue';
import App from './App.vue';
import { router } from './router/index';
import { createPinia } from "pinia";
import axios from 'axios';
import naive from 'naive-ui/es/preset';
import './assets/css/style.css';
import {createDiscreteApi} from "naive-ui";
//import { UserOutlined,LockOutlined} from '@ant-design/icons-vue';
//import Antd from 'ant-design-vue';

const {message,notification,dialog} = createDiscreteApi(["message","notification","dialog"]);
const pinia = createPinia();
axios.defaults.baseURL = 'http://localhost:8080/api/v1';
const app = createApp(App)
//app.component('UserOutlined', UserOutlined);
//app.component('LockOutlined', LockOutlined);

app.use(router)
app.use(pinia)
app.use(naive)
//app.use(Antd)

app.provide('axios', axios)
app.provide('message', message)
app.provide('notification', notification)
app.provide('dialog', dialog)
app.mount('#app');



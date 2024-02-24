import { createApp } from 'vue';
import App from './App';
import router from './router'
import axios from 'axios';
import Antd from 'ant-design-vue';
// import 'ant-design-vue/dist/reset.css';
import antui from '../src/plugins/antui'


const app = createApp(App)
axios.defaults.baseURL = 'http://localhost:8080/api/v1';
app.use(router).use(antui).use(Antd).provide('axios', axios).mount('#app');


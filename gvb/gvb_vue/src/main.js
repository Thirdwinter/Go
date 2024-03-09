import { createApp } from 'vue';
import App from './App.vue';
import { router } from './router/index';
import { createPinia } from "pinia";
import naive from 'naive-ui/es/preset';
import './assets/css/style.css';
import {createDiscreteApi} from "naive-ui"


const {message,notification,dialog} = createDiscreteApi(["message","notification","dialog"]);
const pinia = createPinia();
const app = createApp(App)

app.use(router)
app.use(pinia)
app.use(naive)
//app.use(Antd)

app.provide('message', message)
app.provide('notification', notification)
app.provide('dialog', dialog)
app.mount('#app');



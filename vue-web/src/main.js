import { createApp } from 'vue'
import App from './App.vue'
import store from "@/store";
// import {Form, Field, CellGroup, Button, } from "vant";
import Vant from 'vant';
import 'vant/lib/index.css';

createApp(App)
    .use(store)
    .use(Vant)
    // .use(Form)
    // .use(Field)
    // .use(CellGroup)
    // .use(Button)
    .mount('#app')

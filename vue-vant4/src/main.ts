import {createApp} from 'vue'
import App from './App.vue'
import router from './routers'
import vant from 'vant';
import 'vant/lib/index.css';
import store from "./store";


const app = createApp(App)
app.use(vant).use(router).use(store)
app.mount('#app')



import {createApp} from 'vue'
import App from './App.vue'
import router from './routers'
import './style.css'
import vant from 'vant';
import 'vant/lib/index.css';


const app = createApp(App)
app.use(vant).use(router).mount('#app');



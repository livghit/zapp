import { createApp } from 'vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import "./index.css"
import App from './App.vue'
import Home from './pages/Home.vue'

const routes = [
  { path: '/', component: Home }
]
const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

createApp(App).use(router).mount('#app')

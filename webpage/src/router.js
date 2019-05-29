import Vue from 'vue'
import Router from 'vue-router'
import welcome from './components/welcome'
import wxlogin from './pages/wxlogin'

Vue.use(Router)

const router = new Router({
    mode:"history",
    routes:[{
        path:"/login/:page?",
        component:wxlogin
    },{
        path:"/",
        component:welcome
    }]
})

export default router
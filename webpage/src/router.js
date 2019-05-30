import Vue from 'vue'
import Router from 'vue-router'
import welcome from './components/welcome'
import wxlogin from './pages/wxlogin'
import error from './pages/error'

Vue.use(Router)

const router = new Router({
    mode:"history",
    routes:[{
        path:"/login/:page?",
        component:wxlogin
    },{
        path:"/error",
        component:error
    },{
        path:"/",
        component:welcome
    }]
})

export default router
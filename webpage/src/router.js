import Vue from 'vue'
import Router from 'vue-router'
import welcome from './components/welcome'
import wxlogin from './pages/wxlogin'
import error from './pages/error'
import admin from './pages/admin/admin'

Vue.use(Router)

const router = new Router({
    mode:"history",
    routes:[{
        path:"/login/:page?",
        component:wxlogin
    },{
        path:"/admin",
        component:admin
    },{
        path:"/error",
        component:error
    },{
        path:"/",
        component:welcome
    }]
})

window.router = router
export default router
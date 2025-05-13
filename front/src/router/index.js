import Vue from 'vue'
import VueRouter from 'vue-router'
import store from "@/store";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: () => import("../views/login.vue")
  },
  {
    path: '/login',
    name: 'login',
    component: () => import("../views/login.vue")
  },
  {
    path: '/main',
    name: 'main',
    component: () => import("../views/main"),
    redirect: "/home",
    children: [
      {
        path: '/home',
        name: '主页',
        component: () => import("../views/Home.vue")
      },
      {
        path: '/user',
        name: '用户管理',
        component: () => import("../views/User.vue")
      },
      {
        path: '/device',
        name: '云桌面管理',
        component: () => import("../views/Device.vue")
      },
      {
        path: '/disk',
        name: '磁盘管理',
        component: () => import("../views/Disk.vue")
      },
      {
        path: '/gpu',
        name: 'GPU管理',
        component: () => import("../views/Gpu.vue")
      },
      {
        path: '/template',
        name: '模板配置',
        component: () => import("../views/Template.vue")
      },
      {
        path: '/license',
        name: '授权管理',
        component: () => import("../views/License.vue")
      },
    ]
  },

]

const router = new VueRouter({
  mode: 'history',
  routes
})

router.beforeEach((to, from, next) => {
  localStorage.setItem("currentPathName", to.name)
  store.commit("setPath")
  next()
})

export default router

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import User from '../pages/User.vue'
import UserInterface from '../pages/UserInterface.vue'
import Admin from '../pages/Admin.vue'
import AdminInterface from '../pages/AdminInterface.vue'
import AdminUser from '../pages/AdminUser.vue'

const routes = [
    { path: '/',component: Login },
    { path: '/register',component: Register},
    { path: '/user',component: UserInterface},
    { path: '/user/info',component: User},
    { path: '/admin',component: Admin},
    { path: '/admin/interface',component: AdminInterface},
    { path: '/admin/user',component: AdminUser},
];

export default routes

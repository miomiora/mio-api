<template>
    <h1>登录页面</h1>
    <div style="width: 30vw">
      <el-input v-model="user_account" placeholder="请输入用户名; 用户只能以字母开头，不能小于4位"  />
      <p></p>
      <el-input v-model="user_password" placeholder="请输入密码; 密码不能小于8位" show-password/>
      <p></p>
      <el-button type="primary" @click="login" >登录</el-button>
      <p></p>
      <el-button type="warning" @click="toRegister" >-> 注册页面</el-button>
    </div>


</template>

<script setup>
import {onMounted, ref} from 'vue'
import request from "../plugin/request";
import {ElMessage, ElMessageBox} from 'element-plus'

let user_account = ref('')
let user_password = ref('')


async function login() {
  request.post('/user/login', {
    user_account: user_account.value,
    user_password: user_password.value,
  }).then(res => {
    if( res.code === 0 ) {
      ElMessage.success('登录成功')
      // 设置Token
      if (res.data.user.role === 1) {
        window.location.href = 'admin'
      } else if (res.data.user.role === 0) {
        window.location.href = 'user'
      }
      localStorage.setItem('token', res.data.token);
    } else {
      ElMessage.error(res.description)
    }
  });
}

function toRegister() {
  window.location.href= 'register'
}

async function getCurrentUser() {
  request.get('/user/current').then(res => {
    if( res.code === 0 ) {
      ElMessage.success('用户已经登录')
      if(res.data.role === 0) {
        window.location.href ='user'
      } else if (res.data.role === 1) {
        window.location.href ='admin'
      }
    }
  });
}

onMounted( () => {
  getCurrentUser();
})

</script>

<style scoped>

</style>
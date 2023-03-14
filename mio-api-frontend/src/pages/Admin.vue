<template>
  <h1>管理员页面</h1>
  <el-button type="danger" @click="logout" style="position: absolute; top: 40px; right: 50px">退出登录</el-button>

  <el-button type="primary" @click="toAdminUser">用户管理</el-button>
  <el-button type="primary" @click="toAdminInterface">接口管理</el-button>
  <el-button type="primary" @click="toUser">个人中心</el-button>


</template>

<script setup>
import request from "../plugin/request.js";
import {ElMessage} from "element-plus";

async function logout() {
  request.post('/user/logout').then(res => {
    if (res.code === 0) {
      ElMessage.success('登出成功')
      window.location.href = '/'
    } else {
      ElMessage.error(res.message)
    }
  })
}

function toAdminUser() {
  window.location.href = '/admin/user'
}

function toAdminInterface() {
  window.location.href = '/admin/interface'
}

function toUser() {
  window.location.href = '/user'
}


</script>

<style scoped>

</style>
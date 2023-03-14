<template>
  <h1>在线接口开放平台 mio-api</h1>
  <el-button type="primary" @click="toInfo" style="position: absolute; top: 40px; right: 150px">修改个人资料</el-button>
  <el-button type="danger" @click="logout" style="position: absolute; top: 40px; right: 50px">退出登录</el-button>

  <p></p>
  <el-table :data="tableData" stripe style="width: 100%">
    <el-table-column prop="name" label="名称" width="200" />
    <el-table-column prop="method" label="请求方式" width="80" />
    <el-table-column prop="url" label="URL地址" width="350" />

    <el-table-column prop="description" label="描述" width="200" />

    <el-table-column prop="created_at" label="创建时间" width="150" />
    <el-table-column prop="updated_at" label="更新时间" width="150" />

    <el-table-column fixed="right" label="操作" width="80">
      <template v-slot="scope">
        <el-button type="primary" size="small" @click="detail(scope.row)">查看</el-button>

      </template>
    </el-table-column>

  </el-table>

  <el-pagination layout="prev, pager, next" :page-size="5" :total="total" @current-change="handleCurrentChange"/>

  <el-dialog
      v-model="centerDialogVisible"
      title="查看接口"
      width="50%"
      align-center
  >
    <el-descriptions
        :column="1"
        size="large"
        border
    >
      <el-descriptions-item label="名称">{{detailInterface.name}}</el-descriptions-item>

      <el-descriptions-item label="请求方式">
        <el-tag>
          {{detailInterface.method}}
        </el-tag>
      </el-descriptions-item>

      <el-descriptions-item label="URL">{{detailInterface.url}}</el-descriptions-item>
      <el-descriptions-item label="请求头">{{detailInterface.request_header}}</el-descriptions-item>
      <el-descriptions-item label="响应头">{{detailInterface.response_header}}</el-descriptions-item>
      <el-descriptions-item label="请求参数">{{detailInterface.request_params}}</el-descriptions-item>
      <el-descriptions-item label="描述">{{detailInterface.description}}</el-descriptions-item>
      <el-descriptions-item label="创建时间">{{detailInterface.created_at}}</el-descriptions-item>
      <el-descriptions-item label="更新时间">{{detailInterface.updated_at}}</el-descriptions-item>
    </el-descriptions>

    <p></p>
    <div v-if="detailInterface.request_params!='null'">
      <span style="float: left">请求参数</span>
      <el-input type="textarea" v-model="invokeInfo.body" ></el-input>
    </div>

    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="invoke(detailInterface.id)" style="float:left;">
          调用
        </el-button>
        <el-button @click="cancelDetail">关闭</el-button>
      </span>
    </template>

    <div v-if="resultInvisible">
      <p></p>
      <span style="float: left">返回结果：</span>
      <span>{{resultData}}</span>
    </div>
  </el-dialog>


</template>

<script setup>
import request from "../plugin/request";
import {onMounted, ref} from "vue";
import {ElMessage} from "element-plus";

const tableData = ref();
const currentPage = ref(1);
const total = ref(20);
const centerDialogVisible = ref(false)
const detailInterface = ref({
      id: 0,
      name: '',
      method: '',
      url: '',
      request_header: '',
      response_header: '',
      request_params: '',
      description: '',
      updated_at:'',
      created_at: '',
})
const resultInvisible = ref(false)
const resultData = ref()
const invokeInfo = ref({
  access_key: "",
  secret_key: "",
  body: "",

})

function handleCurrentChange(val) {
  currentPage.value = val
  getInterfaceList()
}

async function getInterfaceList() {
  request.get('/interface/user/5/'+currentPage.value).then(res => {
    if (res.code === 0) {
      tableData.value = res.data.table;
      total.value = res.data.total;
    } else if (res.code === 40100){

      ElMessage.error(res.message)
      window.location.href = '/'
    } else {
      ElMessage.error(res.description)
    }
  })
}

function detail(row) {
  centerDialogVisible.value = true
  detailInterface.value.id = row.id
  detailInterface.value.name = row.name
  detailInterface.value.method = row.method
  detailInterface.value.url = row.url
  detailInterface.value.request_header = row.request_header
  detailInterface.value.response_header = row.response_header
  detailInterface.value.description = row.description
  detailInterface.value.request_params = row.request_params
  detailInterface.value.created_at = row.created_at
  detailInterface.value.updated_at = row.updated_at
}


function toInfo() {
  window.location.href = "/user/info"
}

function cancelDetail() {
  centerDialogVisible.value = false
  resultInvisible.value = false
}

async function logout() {
  request.post('/user/logout').then(res => {
    if (res.code === 0) {
      ElMessage.success('登出成功')
      window.location.href ='/'
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function invoke(id) {
  request.post('/interface/invoke/'+id, invokeInfo.value).then(res=>{
    if (res.code === 0) {
      ElMessage.success('调用成功')
      resultInvisible.value = true
      resultData.value = res.data
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function getCurrentUser() {
  request.get('/user/current').then(res => {
    if( res.code === 0 ) {
      // ElMessage.success('用户已经登录')
      invokeInfo.value.access_key = res.data.access_key
      invokeInfo.value.secret_key = res.data.secret_key
    } else {
      ElMessage.error('无有效的accessKey')
    }
  });
}

onMounted( () => {
  getInterfaceList()
  getCurrentUser()
})

</script>

<style scoped>

</style>
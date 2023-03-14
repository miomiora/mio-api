<template>
  <h1>接口管理</h1>
  <el-button type="primary" @click="toAdmin" style="position: absolute; top: 40px; right: 150px">返回管理页</el-button>
  <el-button type="danger" @click="logout" style="position: absolute; top: 40px; right: 50px">退出登录</el-button>

  <el-button type="primary" @click="newInterface">新建接口</el-button>
  <el-table :data="tableData" stripe style="width: 100%">
    <el-table-column prop="id" label="id" width="50" />
    <el-table-column prop="user_id" label="创建者" width="150" />
    <el-table-column prop="name" label="名称" width="100" />
    <el-table-column prop="method" label="请求方式" width="80" />
    <el-table-column prop="url" label="URL地址" width="350" />
    <el-table-column
        prop="status"
        label="状态"
        width="80"
    >
      <template #default="scope">
        <el-tag v-if="scope.row.status === false" type="danger">关闭</el-tag>
        <el-tag v-if="scope.row.status === true" type="success">开启</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="request_header" label="请求头" width="150" />
    <el-table-column prop="response_header" label="响应头" width="150" />
    <el-table-column prop="description" label="描述" width="150" />
    <el-table-column prop="request_params" label="请求参数" width="150" />
    <el-table-column prop="created_at" label="创建时间" width="150" />
    <el-table-column prop="updated_at" label="更新时间" width="150" />

    <el-table-column fixed="right" label="操作" width="200">
      <template v-slot="scope">
        <el-button type="primary" size="small" @click="edit(scope.row)">修改</el-button>

        <el-button v-if="scope.row.status===false" type="success" size="small" @click="online(scope.row)">
          发布
        </el-button>
        <el-button v-if="scope.row.status===true" type="warning" size="small" @click="offline(scope.row)">
          下线
        </el-button>

        <el-popconfirm title="确定删除这个接口吗？" @confirm="confirmEvent(scope.row)">
          <template #reference>
            <el-button type="danger"  size="small" @click="visible = true">删除</el-button>
          </template>
        </el-popconfirm>

      </template>
    </el-table-column>

  </el-table>

  <el-pagination layout="prev, pager, next" :page-size="5" :total="total" @current-change="handleCurrentChange"/>

  <el-dialog
      v-model="dialogVisible"
      title="新建接口"
      width="40%"
  >

    <el-input v-model="addInterface.name" placeholder="名称"  />
    <p></p>
    <el-input v-model="addInterface.method" placeholder="请求方式" />
    <p></p>
    <el-input v-model="addInterface.url" placeholder="url" />
    <p></p>
    <el-input v-model="addInterface.request_header" placeholder="请求头" type="textarea"/>
    <p></p>
    <el-input v-model="addInterface.response_header" placeholder="响应头" type="textarea"/>
    <p></p>
    <el-input v-model="addInterface.description" placeholder="描述" type="textarea"/>
    <p></p>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitInterface">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog
      v-model="dialogVisible2"
      :title="'修改接口 id: '+updateInterface.id"
      width="40%"
  >

    <span style="float: left">接口名称</span>
    <el-input v-model="updateInterface.name" placeholder="名称"  />
    <p></p>


    <span style="float: left">请求方式</span>
    <el-input v-model="updateInterface.method" placeholder="请求方式" />
    <p></p>


    <span style="float: left">URL</span>
    <el-input v-model="updateInterface.url" placeholder="url" />
    <p></p>


    <span style="float: left">请求头</span>
    <el-input v-model="updateInterface.request_header" placeholder="请求头" type="textarea"/>
    <p></p>

    <span style="float: left">响应头</span>
    <el-input v-model="updateInterface.response_header" placeholder="响应头" type="textarea"/>
    <p></p>

    <span style="float: left">请求参数</span>
    <el-input v-model="updateInterface.request_params" placeholder="请求参数" type="textarea"/>
    <p></p>

    <span style="float: left">描述</span>
    <el-input v-model="updateInterface.description" placeholder="描述" type="textarea"/>
    <p></p>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible2 = false">Cancel</el-button>
        <el-button type="primary" @click="submitUpdate">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup>
import {onMounted, ref} from "vue";
import request from "../plugin/request";
import {ElMessage} from "element-plus";

const tableData = ref();
const currentPage = ref(1);
let dialogVisible = ref(false)
let dialogVisible2 = ref(false)
const visible = ref(false)
const total = ref(20)
const addInterface = ref({
  name: '',
  method: '',
  url: '',
  request_header: '',
  response_header: '',
  request_params: '',
  description: '',
})

const updateInterface = ref({
  id: 0,
  name: '',
  method: '',
  url: '',
  request_header: '',
  response_header: '',
  request_params: '',
  description: '',
  user_id: 0,
  status: false,
})

function handleCurrentChange(val) {
  currentPage.value = val
  getInterfaceList()
}


async function getInterfaceList() {
  request.get('/interface/5/'+currentPage.value).then(res => {
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

async function submitInterface() {
  request.post('/interface/create', addInterface.value).then(res => {
    if (res.code === 0) {
      ElMessage.success('添加成功！')
      window.location.href = ''
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function submitUpdate() {
  request.post('/interface/update/'+ updateInterface.value.id, updateInterface.value).then(res => {
    if (res.code === 0) {
      ElMessage.success('修改成功！')
      window.location.href = ''
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function online(row) {
  request.post('/interface/online/'+ row.id).then(res => {
    if (res.code === 0) {
      ElMessage.success('发布成功！')
      window.location.href = ''
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function offline(row) {
  request.post('/interface/offline/'+ row.id).then(res => {
    if (res.code === 0) {
      ElMessage.success('下线成功！')
      window.location.href = ''
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function confirmEvent(row) {
  request.post('/interface/delete/'+row.id).then(res => {
    if (res.code === 0) {
      ElMessage.success('删除成功')
      window.location.href = ''
    } else {
      ElMessage.error(res.message)
    }
  })
}

async function newInterface() {
  dialogVisible.value = true
}

function toAdmin () {
  window.location.href = '/admin'
}

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

function edit(row) {
  dialogVisible2.value = true
  updateInterface.value.id = row.id
  updateInterface.value.name = row.name
  updateInterface.value.method = row.method
  updateInterface.value.url = row.url
  updateInterface.value.request_header = row.request_header
  updateInterface.value.response_header = row.response_header
  updateInterface.value.request_params = row.request_params
  updateInterface.value.description = row.description
  updateInterface.value.user_id = row.user_id
  updateInterface.value.status = row.status
}

function del(row) {

}

onMounted( () => {
  getInterfaceList();
})
</script>

<style scoped>

</style>
<template>
  <el-card>
  <div style="margin: 10px 0">
    <el-input style="width: 200px" placeholder="请输入名称" suffix-icon="el-icon-search" v-model="name"></el-input>
    <el-button class="ml-5" type="primary" @click="load">搜索</el-button>
    <el-button type="warning" @click="reset">刷新</el-button>
  </div>

  <div style="margin: 10px 0">
    <el-button type="primary" @click="handleAdd" title="基于模板创建">快速新建 <i class="el-icon-circle-plus-outline"></i></el-button>
  </div>

  <el-table :data="tableData" :header-cell-class-name="headerBg" @selection-change="handleSelectionChange">
    <el-table-column prop="name" label="云桌面名称" ></el-table-column>
    <el-table-column prop="username" label="所属用户" ></el-table-column>
    <el-table-column prop="templateInfo" label="所属模板" ></el-table-column>
    <el-table-column prop="virtualIp" label="IP" ></el-table-column>
    <el-table-column prop="memoryInfo" label="内存" ></el-table-column>
    <el-table-column prop="cpuInfo" label="CPU" ></el-table-column>
    <el-table-column prop="gpuInfo" label="GPU" ></el-table-column>
    <el-table-column prop="status" label="状态" ></el-table-column>
    <el-table-column prop="createdTime" label="创建时间" ></el-table-column>
    <el-table-column label="操作">
      <template slot-scope="scope">
        <div class="button-container">
          <el-button round type="success" @click="openVM(scope.row.id)">开机</el-button>
          <el-popconfirm
              class="ml-5"
              confirm-button-text='确定'
              cancel-button-text='我再想想'
              icon="el-icon-info"
              icon-color="red"
              title="您确定关机吗？"
              @confirm="closeVM(scope.row.id)"
          >
            <el-button round type="danger" slot="reference">关机</el-button>
          </el-popconfirm>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="" >
      <template slot-scope="scope">
        <div class="button-container">
        <el-popconfirm
              class="ml-5"
              confirm-button-text='确定'
              cancel-button-text='我再想想'
              icon="el-icon-info"
              icon-color="red"
              title="您确定删除吗？删除后不可恢复，请谨慎操作！"
              @confirm="del(scope.row.id)"
          >
            <el-button round type="danger" slot="reference">删除</el-button>
          </el-popconfirm>
          <el-button round type="success" @click="handleEdit(scope.row)">编辑</el-button>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="" >
      <template slot-scope="scope">
        <div class="button-container">
          <el-popconfirm
              class="ml-5"
              confirm-button-text='确定'
              cancel-button-text='我再想想'
              icon="el-icon-info"
              icon-color="red"
              title="您确定重启吗？"
              @confirm="resetVM(scope.row.id)"
          >
            <el-button round type="danger" slot="reference">重启</el-button>
          </el-popconfirm>
          <el-button round type="danger" @click="unBindUser(scope.row.id)">用户解绑</el-button>
        </div>
      </template>
    </el-table-column>
    <el-table-column label=""  align="center">
      <template slot-scope="scope">
        <div class="button-container">
          <el-button round type="success" @click="bindGpu(scope.row)">GPU绑定</el-button>
          <el-button round type="danger" @click="unBindGPU(scope.row)">GPU解绑</el-button>
        </div>
      </template>
    </el-table-column>
  </el-table>
  <div style="padding: 10px 0">
    <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="pageNum"
        :page-sizes="[5, 10, 15, 20]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
    </el-pagination>
  </div>

  <el-dialog title="云桌面信息" :visible.sync="dialogFormVisible" width="50%" >
    <el-form label-width="80px" size="small">
      <el-form-item label="名称">
        <el-input v-model="addForm.name" autocomplete="off"></el-input>
      </el-form-item>

    <el-form-item label="模板" prop="srcVmPath">
      <el-select v-model="addForm.srcVmPath" placeholder="请选择" style="width: 100%">
          <el-option v-for="item in templates" :key="item" :value="item">
            {{ item }}
          </el-option>
        </el-select>
    </el-form-item>

    <el-form-item label="网络" prop="vmSwitch">
      <el-select v-model="addForm.vmSwitch" placeholder="请选择" style="width: 100%">
          <el-option v-for="item in switchs" :key="item" :value="item">
            {{ item }}
          </el-option>
        </el-select>
    </el-form-item>

    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="save">确 定</el-button>
    </div>
  </el-dialog>

  <el-dialog title="云桌面信息" :visible.sync="dialogUpdateFormVisible" width="30%" >
    <el-form label-width="80px" size="small">
      <el-form-item label="内存">
        <el-select v-model="editForm.memoryInfo" placeholder="请选择内存大小" style="width: 100%">
          <el-option v-for="item in memoryOptions" :key="item" :label="item" :value="item">
            {{ item }}
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="CPU">
        <el-select v-model="editForm.cpuInfo" placeholder="请选择CPU数量" style="width: 100%">
        <el-option v-for="item in cpuOptions" :key="item.value" :label="item.label" :value="item.value">
          {{ item.label }}
        </el-option>
      </el-select>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogUpdateFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="update(editForm.id)">修改</el-button>
    </div>
  </el-dialog>

  <el-dialog title="GPU信息" :visible.sync="dialogGpuVisible" width="30%" >
    <el-form label-width="80px" size="small">
      <el-form-item label="GPU">
        <el-select v-model="gpuForm.gpuId" placeholder="请选择GPU" style="width: 100%">
          <el-option v-for="item in gpus" :key="item.name" :label="item.name" :value="item.id">
            {{ item.name }}
          </el-option>
        </el-select>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogGpuVisible = false">取 消</el-button>
      <el-button type="primary" @click="gpuUpdate()">绑定</el-button>
    </div>
  </el-dialog>

  </el-card>
</template>

<script>

export default {
  name: "Device",
  data(){
    return{
      tableData:[],
      total:0,
      pageNum:1,
      pageSize:5,
      name:'',
      username: '',
      virtualIp: '',
      status: '',
      createdTime: '',
      gpuForm: {
        gpuId: '',
        deviceName: ''
      },
      addForm: {
        name:'',
        srcVmPath: '',
        vmSwitch: ''
      },
      editForm: {
        name:'',
        memoryInfo: '',
        cpuInfo: ''
      },
      configForm: {
        memoryInfo: '',
        cpuInfo: ''
      },
      dialogFormVisible: false,
      dialogUpdateFormVisible: false,
      dialogConfigFormVisible: false,
      dialogGpuVisible: false,
      menuDialogVis:false,
      menuData:[],
      props:{
        label: 'name',
      },
      expends:[],
      checks:[],
      roleId:1,
      multipleSelection: [],
      headerBg:"headerBg",
      userInfos: [],
      memoryOptions: ['2GB', '4GB', '8GB', '16GB'],
      cpuOptions: [
        {value: '1', label: '1'},
        {value: '2', label: '2'},
        {value: '4', label: '4'},
        {value: '8', label: '8'}
      ],
      timer: null,
      templates: [],
      gpus: [],
      switchs: []
    }
  },
  created() {
    this.load();
  },
  mounted() {
    this.timer = setInterval(this.load, 10000)
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  methods:{
    load() {
      this.request.get("/api/cloud/v1/devices?count="+this.pageSize+"&index="+this.pageNum+"&name="+this.name).then(res => {
        this.tableData = res.data.data.devices
        this.total = res.data.data.totalNum
      })
    },
    save() {
      console.log(this.addForm)
      this.request.post("/api/cloud/v1/vm", this.addForm).then(res => {
        if (res.status===200) {
          if (res.data.msg == "Device.Exist") {
            this.$message.error("云桌面已存在")
          } else if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else if (res.data.msg == "Device.TemplateNotExist") {
            this.$message.error("模板文件不存在，请检查后再重试!")
          } else {
            this.$message.success("云桌面创建中...")
          }
          this.dialogFormVisible = false
          this.load()
        } else {
          this.$message.error("保存失败")
        }
      })
    },

    update(id) {
      console.log("update vm")
      this.request.put("/api/cloud/v1/vm/" + id, this.editForm).then(res => {
        if (res.status===200) {
          if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else if (res.data.msg == "Device.IsRunning") {
            this.$message.error("云桌面处于运行状态，无法删除，请先关闭云桌面后再操作！")
          } else {
            this.$message.success("更新成功")
          }
          this.dialogUpdateFormVisible = false
          this.load()
        } else {
          this.$message.error("保存失败")
        }
      })
    },

    unBindUser(id) {
      let unBindForm = {
        deviceId: id
      }
      this.request.post("/api/cloud/v1/unbind_user", unBindForm).then(res => {
        if (res.status === 200) {
          if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else if (res.data.msg == "Device.UnBindUserFailed") {
            this.$message.error("云桌面用户解绑失败，请检查后再重试!")
          } else if (res.data.msg == "Device.MustBeRunning") {
            this.$message.error("云桌面必须是运行状态才能解绑!")
          } else if (res.data.msg == "Device.TemplateNotExist") {
            this.$message.error("模板文件不存在，请检查后再重试!")
          } else {
            this.$message.success("解绑成功")
          }
          this.load()
        } else {
          this.$message.error("解绑失败")
        }
      })
    },
    handleAdd() {
      this.request.get("/api/cloud/v1/device_templates").then(res => {
        this.templates = res.data.data.templates
      })
      this.request.get("/api/cloud/v1/device_switchs").then(res=> {
        this.switchs = res.data.data.switchs
      })
      this.dialogFormVisible = true
      this.addForm = {}
    },

    handleEdit(row) {
      this.editForm = row
      this.dialogUpdateFormVisible = true
    },
    bindGpu(row) {
      this.gpuForm.deviceName = row.name
      this.dialogGpuVisible = true
      this.request.get("/api/cloud/v1/gpus?count=1000&index=1").then(res => {
        this.gpus = res.data.data.gpus
      })
    },
    unBindGPU(row) {
      console.log(row)
      let unbind_form = {
        deviceId: row.id
      }
      this.request.post("/api/cloud/v1/unbind_gpu", unbind_form).then(res => {
        if (res.status === 200) {
          if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else {
            this.$message.success("解绑成功")
            this.gpuForm = {
              gpuId: '',
              deviceName: ''
            }
            this.dialogGpuVisible = false
          }
          this.load()
        } else {
          this.$message.error("解绑失败")
        }
      })
    },
    gpuUpdate() {
      console.log(this.gpuForm)
      this.request.post("/api/cloud/v1/bind_gpu", this.gpuForm).then(res => {
        if (res.status === 200) {
          if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else if (res.data.msg == "Gpu.NotExist") {
            this.$message.error("GPU不存在")
          } else if (res.data.msg == "Gpu.IsLimited") {
            this.$message.error("一个GPU最多只能被4个云桌面使用")
          } else if (res.data.msg == "Gpu.BindingExist") {
            this.$message.error("已绑定GPU,请先解绑再进行绑定")
          } else if (res.data.msg == "Gpu.BindFailed") {
            this.$message.error("绑定失败,请在设备管理器上检查GPU状态后再重试")
          } else {
            this.$message.success("绑定成功")
            this.gpuForm = {
              gpuId: '',
              deviceName: ''
            }
            this.dialogGpuVisible = false
          }
          this.load()
        } else {
          this.$message.error("绑定失败")
        }
      })
    },
    handleVMConfig(row) {
      this.configForm = row
      this.dialogConfigFormVisible = true
    },
    setVMConfig() {
      console.log(this.configForm)
    },
    openVM(id) {
      let obj = {
        "vm_id": id,
        "action": 1
      }
      this.request.post("/api/cloud/v1/vm/operate", obj).then(res => {
        if (res.status===200) {
          if (res.data.code == 2) {
            this.$message.error("无法操作，请检查云桌面的状态或联系管理员！")
          } else {
            this.$message.success("操作成功")
          }
          this.load()
        } else {
          this.$message.error("操作失败")
        }
      })
    },
    closeVM(id) {
      let obj = {
        "vm_id": id,
        "action": 2
      }
      this.request.post("/api/cloud/v1/vm/operate", obj).then(res => {
        if (res.status===200) {
          if (res.data.code == 2) {
            this.$message.error("无法操作，请检查云桌面的状态或联系管理员！")
          } else {
            this.$message.success("操作成功")
          }
          this.load()
        } else {
          this.$message.error("操作失败")
        }
      })
    },
    resetVM(id) {
      let obj = {
        "vm_id": id,
        "action": 3
      }
      this.request.post("/api/cloud/v1/vm/operate", obj).then(res => {
        if (res.status===200) {
          if (res.data.code == 2) {
            this.$message.error("无法操作，请检查云桌面的状态或联系管理员！")
          } else {
            this.$message.success("操作成功")
          }
          this.load()
        } else {
          this.$message.error("操作失败")
        }
      })
    },

    del(id) {
      this.request.delete("/api/cloud/v1/vm/" + id).then(res => {
        if (res.status===200) {
          if (res.data.msg == "Device.IsRunning") {
            this.$message.error("云桌面处于运行状态，无法删除，请先关闭云桌面后再操作！")
          } else {
            this.$message.success("删除成功")
          }
          this.load()
        } else {
          this.$message.error("删除失败")
        }
      })
    },

    handleSelectionChange(val) {
      console.log(val)
      this.multipleSelection = val
    },

    reset() {
      this.name = ""
      this.load()
    },

    handleSizeChange(pageSize) {
      console.log(pageSize)
      this.pageSize = pageSize
      this.load()
    },
    handleCurrentChange(pageNum) {
      console.log(pageNum)
      this.pageNum = pageNum
      this.load()
    }

  }
}
</script>

<style>
.headerBg {
  background: #eee !important;
}

.button-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.button-container > * {
  margin-bottom: 5px;
}

.button-container .el-button {
  width: 100px;
  height: 30px;
}

</style>
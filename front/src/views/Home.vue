<template>
  <div>

    <div>
      <el-card class="back">
        <div>
          <div style="font-size: 32px;margin: 20px">欢迎来到云桌面管理平台！</div>
        </div>
      </el-card>
    </div>

    <el-row :gutter="12" style="margin-bottom: 30px">
      <el-col :span="12">
        <el-card style="color: #E6A23C">
          <div><i class="el-icon-user" style="margin-right: 5px"></i>用户数</div>
          <div @click="gotoUser" style="text-align: center;font-size: 30px;font-weight: bold">{{ userNum }}</div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card style="color: #67C23A">
          <div><i class="el-icon-s-platform" style="margin-right: 5px"></i>云桌面数</div>
          <div @click="gotoDevice" style="text-align: center;font-size: 30px;font-weight: bold">{{ deviceNum }}</div>
        </el-card>
      </el-col>
    </el-row>
    
  </div>
</template>

<script>
import * as echarts from 'echarts'

export default {
  name: "Home",
  data(){
    return{
      userNum:0,
      deviceNum:0,
      count:10,
      article:[],
      dialogFormVisible: false,
      addForm: {
        machineCode: '',
        licenseCode: ''
      }
    }
  },
  created() {
    this.getcount()
    this.getLicense()
  },
  methods:{
    getcount(){
      console.log("get count")
      this.request.get("/api/cloud/v1/user_count").then(res => {
        this.userNum = res.data.data.num
      })
      this.request.get("/api/cloud/v1/device_count").then(res => {
        this.deviceNum = res.data.data.num
      })
    },
    gotoUser() {
      this.$router.push("/user")
    },
    gotoDevice() {
      this.$router.push("/device")
    },
    getLicense() {
      this.request.get("/api/cloud/v1/licenses").then(res => {
        if (res.data.msg == "User.LicenseExpired") {
          this.$message.error("授权已过期")
          this.$router.push("/login")
          return
        } else if (res.data.msg == "User.LicenseInvalid") {
          this.$message.error("因系统时间变更导致授权异常，请将系统时间调整正确后再重新授权")
          this.$router.push("/login")
          return
        }
        let data = res.data.data
        this.addForm.machineCode = data.machineCode
        if (!data.isChecked) {
          this.dialogFormVisible = true
        } else {
          this.dialogFormVisible = false
        }
      })
    },
    active() {
      console.log("激活授权码")
      this.request.post("/api/cloud/v1/licenses", this.addForm).then(res => {
        if (res.status===200) {
          if (res.data.msg == "User.LicenseExists") {
            this.$message.error("授权已使用过,授权码只能使用一次,请更换授权码再进行激活")
          } else if (res.data.msg == "User.LicenseActiveFailed") {
            this.$message.error("激活失败，请检查授权码是否正确!")
          } else if (res.data.msg == "Common.InvalidParam") {
            this.$message.error("参数异常，请检查参数后再重试！")
          } else {
            this.$message.success("激活成功")
            this.dialogFormVisible = false
          }
        } else {
          this.$message.error("激活失败")
        }
      })
      
    }
  }
}

</script>

<style scoped>
.back{
  background-image: url("../assets/dashboard.png");
  background-repeat: no-repeat;
  background-position:550px -200px;
  margin-bottom: 20px;
  height:200px;
}

</style>
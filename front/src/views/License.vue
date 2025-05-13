<template>
  <el-card>
  <!-- <div style="margin: 10px 0">
    <el-input style="width: 200px" placeholder="请输入名称" suffix-icon="el-icon-search" v-model="name"></el-input>
    <el-button class="ml-5" type="primary" @click="load">搜索</el-button>
    <el-button type="warning" @click="reset">重置</el-button>
  </div> -->

  <el-table :data="tableData"  :header-cell-class-name="headerBg" @selection-change="handleSelectionChange">
    <!-- <el-table-column prop="id" label="ID"></el-table-column> -->
    <el-table-column prop="machineCode" label="机器码" ></el-table-column>
    <el-table-column prop="licenseCode" label="授权码" ></el-table-column>
    <el-table-column prop="licenseType" label="授权类型" ></el-table-column>
    <el-table-column prop="expireTime" label="到期时间" ></el-table-column>
    <!-- <el-table-column label="操作"  align="center">
      <template slot-scope="scope">
        <el-button round type="success" @click="handleAllocation(scope.row)">分配</el-button>
      </template>
    </el-table-column> -->
  </el-table>

  </el-card>
</template>

<script>
export default {
  name: "Role",
  data(){
    return{
      tableData:[],
      total:0,
      pageNum:1,
      pageSize:5,
      name:'',
      headerBg:"headerBg"
    }
  },
  created() {
    this.load();
  },
  methods:{
    load() {
      this.request.get("/api/cloud/v1/licenses_all").then(res => {
        this.tableData = res.data.data.licenses
      })
    }
  }
}
</script>

<style>
.headerBg {
  background: #eee !important;
}
</style>
<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="机器码">
          <el-input v-model="searchInfo.code" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="el-icon-delete"
              size="mini"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户数" prop="userNum" width="120" />
        <el-table-column align="left" label="有效时间（试用、正式）" prop="validtime" width="120" />
        <el-table-column align="left" label="同时在线用户数" prop="concurrentUsers" width="120" />
        <el-table-column align="left" label="模式（1测试、2试用、3正式）" prop="model" width="120" />
        <el-table-column align="left" label="水印（开关）" prop="watermark" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.watermark) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户信息" prop="userInfo" width="120" />
        <el-table-column align="left" label="机器码" prop="code" width="120" show-overflow-tooltip="true" />
        <el-table-column align="left" label="软件类别  1.翻译机" prop="product" width="120" />
        <el-table-column
          align="left"
          label="最小启用版本（当小于最小启用版本则不生效）"
          prop="minVersion"
          width="240"
        />
        <el-table-column align="left" label="任务并发数" prop="concurrencyNum" width="120" />
        <el-table-column align="left" label="功能模块管理" prop="functionModule" width="120" />
        <el-table-column align="left" label="单次最大任务数" prop="onceTask" width="120" />
        <el-table-column
          align="left"
          label="每用户最大并发数（0：默认值，无限制；>0 ：限制数量）"
          prop="userConcurrencyPer"
          width="240"
        />
        <el-table-column
          align="left"
          label="并发模式（0：抢先式并发；1：平均并发）"
          prop="concurrencyModel"
          width="120"
        />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="el-icon-edit"
              size="small"
              class="table-button"
              @click="updateXadTemplate(scope.row)"
              >变更</el-button
            >
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="150px">
        <el-form-item label="用户数:">
          <el-input v-model.number="formData.userNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="有效时间（试用、正式）:">
          <el-input v-model.number="formData.validtime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="同时在线用户数:">
          <el-input v-model.number="formData.concurrentUsers" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="模式（1测试、2试用、3正式）:">
          <el-input v-model.number="formData.model" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="水印（开关）:">
          <el-switch
            v-model="formData.watermark"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          ></el-switch>
        </el-form-item>
        <el-form-item label="用户信息:">
          <el-input v-model="formData.userInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机器码:">
          <el-input v-model="formData.code" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="软件类别  1.翻译机:">
          <el-input v-model.number="formData.product" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最小启用版本（当小于最小启用版本则不生效）:">
          <el-input v-model="formData.minVersion" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="任务并发数:">
          <el-input v-model.number="formData.concurrencyNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="功能模块管理:">
          <el-input v-model="formData.functionModule" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单次最大任务数:">
          <el-input v-model.number="formData.onceTask" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="每用户最大并发数（0：默认值，无限制；>0 ：限制数量）:">
          <el-input v-model.number="formData.userConcurrencyPer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="并发模式（0：抢先式并发；1：平均并发）:">
          <el-input v-model.number="formData.concurrencyModel" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createXadTemplate,
  deleteXadTemplate,
  deleteXadTemplateByIds,
  updateXadTemplate,
  findXadTemplate,
  getXadTemplateList,
} from '@/api/xadTemplate' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'XadTemplate',
  mixins: [infoList],
  data() {
    return {
      listApi: getXadTemplateList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        userNum: 0,
        validtime: 0,
        concurrentUsers: 0,
        model: 0,
        watermark: false,
        userInfo: '',
        code: '',
        product: 0,
        minVersion: '',
        concurrencyNum: 0,
        functionModule: '',
        onceTask: 0,
        userConcurrencyPer: 0,
        concurrencyModel: 0,
      },
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.watermark === '') {
        this.searchInfo.watermark = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        this.deleteXadTemplate(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据',
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID)
        })
      const res = await deleteXadTemplateByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateXadTemplate(row) {
      const res = await findXadTemplate({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rexad_template
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        userNum: 0,
        validtime: 0,
        concurrentUsers: 0,
        model: 0,
        watermark: false,
        userInfo: '',
        code: '',
        product: 0,
        minVersion: '',
        concurrencyNum: 0,
        functionModule: '',
        onceTask: 0,
        userConcurrencyPer: 0,
        concurrencyModel: 0,
      }
    },
    async deleteXadTemplate(row) {
      const res = await deleteXadTemplate({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功',
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createXadTemplate(this.formData)
          break
        case 'update':
          res = await updateXadTemplate(this.formData)
          break
        default:
          res = await createXadTemplate(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
  },
}
</script>

<style></style>

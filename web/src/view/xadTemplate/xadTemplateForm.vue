<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="left" label-width="auto">
        <el-form-item label="用户数:">
          <el-input v-model.number="formData.userNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="有效时间（时间戳）:">
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
        <el-form-item label="最小启用版本">
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
        <el-form-item label="每用户最大并发数">
          <el-input v-model.number="formData.userConcurrencyPer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="并发模式（0：抢先式并发；1：平均并发）:">
          <el-input v-model.number="formData.concurrencyModel" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { createXadTemplate, updateXadTemplate, findXadTemplate } from '@/api/xadTemplate' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'XadTemplate',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findXadTemplate({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rexad_template
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    },
  },
}
</script>

<style></style>

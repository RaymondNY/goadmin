<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="客户编号:">
          <el-input v-model.number="formData.customerId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="产品编号:">
          <el-input v-model.number="formData.productId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model="formData.createUser" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="license地址:">
          <el-input v-model="formData.licenseUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公钥地址:">
          <el-input v-model="formData.publicKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="私钥地址:">
          <el-input v-model="formData.privateKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="机器码:">
          <el-input v-model="formData.machineCode" clearable placeholder="请输入" />
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
import {
  createXadLicense,
  updateXadLicense,
  findXadLicense
} from '@/api/xadLicense' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'XadLicense',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        customerId: 0,
        productId: 0,
        createUser: '',
        licenseUrl: '',
        publicKey: '',
        privateKey: '',
        machineCode: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findXadLicense({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rexadLicense
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
          res = await createXadLicense(this.formData)
          break
        case 'update':
          res = await updateXadLicense(this.formData)
          break
        default:
          res = await createXadLicense(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>

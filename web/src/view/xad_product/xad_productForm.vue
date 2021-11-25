<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="产品名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="硬件信息:">
          <el-input v-model="formData.hardInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:">
          <el-input v-model="formData.profile" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model="formData.createUser" clearable placeholder="请输入" />
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
  createXadProduct,
  updateXadProduct,
  findXadProduct
} from '@/api/xad_product' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'XadProduct',
  mixins: [infoList],
  data() {
    return {
      type: '',
      statusOptions: [],
      formData: {
        name: '',
        hardInfo: '',
        profile: '',
        status: undefined,
        createUser: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findXadProduct({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rexad_product
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('status')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createXadProduct(this.formData)
          break
        case 'update':
          res = await updateXadProduct(this.formData)
          break
        default:
          res = await createXadProduct(this.formData)
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

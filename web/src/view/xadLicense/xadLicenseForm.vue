<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="100px" :rules="rules" ref="authorityForm">
        <el-form-item label="客户选择:" prop="customerId">
          <el-select v-model="formData.customerId" placeholder="请选择">
            <el-option v-for="item in customerList" :key="item.value" :label="item.name" :value="item.ID"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="产品选择:" prop="productId">
          <el-select v-model="formData.productId" placeholder="请选择">
            <el-option v-for="item in productList" :key="item.value" :label="item.name" :value="item.ID"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="模板选择:" prop="templateId">
          <el-select v-model="formData.templateId" placeholder="请选择">
            <el-option v-for="item in templateList" :key="item.value" :label="item.userInfo" :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="创建人:" prop="createUser">
          <el-input v-model="formData.createUser" clearable placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="授权文件地址:">
          <el-input v-model="formData.licenseUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公钥地址:">
          <el-input v-model="formData.publicKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="私钥地址:">
          <el-input v-model="formData.privateKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="10">
            <el-form-item label="机器码:">
              <el-input v-model="formData.machineCode" clearable placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="输入请求">
              <el-input v-model="testurl" clearable placeholder="请输入" />
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-button type="primary" @click="test">Primary</el-button>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { createXadLicense, updateXadLicense, findXadLicense, getMachineCode } from '@/api/xadLicense' //  此处请自行替换地址
import { getXadProductList } from '@/api/xad_product'
import { getXadCustomerList } from '@/api/xadCustomer'
import { getXadTemplateList } from '@/api/xadTemplate'
import infoList from '@/mixins/infoList'
export default {
  name: 'XadLicense',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        customerId: '',
        productId: '',
        templateId: '',
        createUser: '',
        licenseUrl: '',
        publicKey: '',
        privateKey: '',
        machineCode: '',
      },
      testurl: '192.168.1.117',
      productList: [],
      customerList: [],
      templateList: [],
      rules: {
        templateId: [{ required: true, message: '请选择模板', trigger: 'blur' }],
        // createUser: [{ required: true, message: '请选择人', trigger: 'blur' }],
        productId: [{ required: true, message: '请选择产品', trigger: 'blur' }],
        customerId: [{ required: true, message: '请选择客户', trigger: 'blur' }],
      },
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
    let res2 = await getXadProductList()
    this.$data.productList = res2.data.list
    console.log(res2)

    let res3 = await getXadCustomerList()
    this.$data.customerList = res3.data.list
    console.log(res3)

    let res4 = await getXadTemplateList()
    this.$data.templateList = res4.data.list
    console.log(res4)
  },
  methods: {
    async save() {
      this.$refs.authorityForm.validate(async (valid) => {
        if (valid) {
          let res
          this.formData.createUser = this.$store.state.user.nickName
          switch (this.type) {
            case 'create':
              res = await createXadLicense(this.formData, this.testurl)
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
              message: '创建/更改成功',
            })
            this.$router.go(-1)
          }
        }
      })
    },
    back() {
      this.$router.go(-1)
    },
    async test() {
      console.log(this.$data.testurl)
      let res1 = await getMachineCode(this.$data.testurl)
      console.log(res1.data)
      if (res1.data == null || res1.data == '' || res1.code == 7) {
        console.log(11)
      } else {
        console.log(22)
        const d = JSON.parse(res1.msg)
        this.$data.formData.machineCode = d.value
        console.log(d.value)
        this.$message({
          type: 'success',
          message: '获取机器码',
        })
      }
    },
  },
}
</script>

<style></style>

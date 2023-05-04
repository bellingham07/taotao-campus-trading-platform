<template>
  <el-form :label-position="'left'" :scroll-to-error="true">
    <el-form-item style="height: 100%;alignment: center;" >
<!--      <el-image style="width: 100px; height: 100px" :src="image.url"/>-->
      <img src="@/assets/logo-dog.gif" style="margin: 0 auto" alt="logo">
    </el-form-item>
    <el-form-item label="账号">
      <el-input v-model="user.username" placeholder="请输入用户名"/>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="user.password" placeholder="请输入密码"/>
    </el-form-item>
    <el-form-item style="display: flex;align-items: flex-end">
      <el-button type="primary" @click="onSubmit" style="margin: auto">登入</el-button>
    </el-form-item>
  </el-form>
  <div class="register">
    <span class="form-footer">忘记密码</span>
    <span class="form-footer">没有账号</span>
  </div>
  <div class="footer">
    <el-row :gutter="10">
      <el-col :span="6"><div class="grid-content ep-bg-purple" />
        <span style="height: 40px;line-height: 40px">快速登陆</span>
      </el-col>
      <el-col :span="6"><div class="grid-content ep-bg-purple" />
        <el-avatar :size="40" src="/assets/icon/img.png" />
      </el-col>
      <el-col :span="6"><div class="grid-content ep-bg-purple" />
        <el-avatar :size="40" src="/assets/icon/img.png" />
      </el-col>
      <el-col :span="6"><div class="grid-content ep-bg-purple" />
        <el-avatar :size="40" src="/assets/icon/img.png" />
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "LoginForm",
  data(){
    return{
      user:{
        username:'',
        password:''
      },
      image:{
        url:'@/assets/logo-dog.gif'
      }
    }
  },
  methods:{
    onSubmit(){

    }
  }
}
</script>

<style scoped>
.register{
  flex-direction:row-reverse;
  display: flex;
}
.register .form-footer{
  /*相对于根元素设置字体大小，保证页面缩放后依然正常*/
  font-size: 0.5rem;
  color: #2c3e50;
  height: 0.6rem;
  line-height: 0.6rem;
  padding: 0;
  margin-left: 0.6rem;
}
.footer{
  position: absolute;
  bottom: 3rem;
  font-size: 0.7rem;
  margin-top: 2rem;
  width: 80%;
  margin-left: 10%;
}

</style>

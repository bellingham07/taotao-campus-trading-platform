<template>
  <div class="logo">
    <van-image
        width="10rem"
        height="10rem"
        fit="cover"
        position="center"
        :src="image.url"
    />
  </div>
  <van-form @submit="onSubmit">
    <van-cell-group inset class="van-hairline--surround">
      <van-field
          v-model="user.username"
          name="用户名"
          label="用户名"
          placeholder="用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
          v-model="user.password"
          type="password"
          name="密码"
          label="密码"
          placeholder="密码"
          :rules="[{ required: true, message: '请填写密码' }]"
      />
    </van-cell-group>
    <div style="margin: 16px;">
      <van-button round block type="primary" native-type="submit">
        提交
      </van-button>
    </div>
  </van-form>
  <div class="register">
    <span class="form-footer">忘记密码</span>
    <span class="form-footer">没有账号</span>
  </div>
<!--  <div class="footer">-->
<!--    <el-row :gutter="10">-->
<!--      <el-col :span="6"><div class="grid-content ep-bg-purple" />-->
<!--        <span style="height: 40px;line-height: 40px">快速登陆</span>-->
<!--      </el-col>-->
<!--      <el-col :span="6"><div class="grid-content ep-bg-purple" />-->
<!--        <el-avatar :size="40" src="/assets/icon/img.png" />-->
<!--      </el-col>-->
<!--      <el-col :span="6"><div class="grid-content ep-bg-purple" />-->
<!--        <el-avatar :size="40" src="/assets/icon/img.png" />-->
<!--      </el-col>-->
<!--      <el-col :span="6"><div class="grid-content ep-bg-purple" />-->
<!--        <el-avatar :size="40" src="/assets/icon/img.png" />-->
<!--      </el-col>-->
<!--    </el-row>-->
<!--  </div>-->

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
        url:require('@/assets/logo-dog.gif')
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
.logo{
  text-align: center;
  margin:3rem auto;
  width: 100%;
}
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
  margin-left: 0.6rem;
  padding: 0 1rem;
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

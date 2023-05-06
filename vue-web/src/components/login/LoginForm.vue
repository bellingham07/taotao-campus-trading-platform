<template>
  <div class="top">
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
  </div>

  <div class="footer">
    <van-space :size="30" style="margin: 0 auto">
    <van-image
        round
        width="3rem"
        height="3rem"
        src="https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg"
    />
    <van-image
        round
        width="3rem"
        height="3rem"
        src="https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg"
    />
    <van-image
        round
        width="3rem"
        height="3rem"
        src="https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg"
    />
    </van-space>
    <p>快速登陆</p>
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
@keyframes gradientBG {
  0% {
    background-position: 0 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0 50%;
  }
}
.top{
  /*background-color: lightslategray;*/
  background: linear-gradient(-45deg, #af9995, #655586, #23a6d5, #23d5ab);
  animation: gradientBG 5s ease infinite;
  background-size: 600% 600%;
  height: 100%;
  width: 100%;
}
.logo{
  text-align: center;
  margin:0 auto;
  padding: 2rem 0;
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
  width: 100%;
  text-align: center;
}
.footer p{
  padding: 1rem 0;
  font: 700 15px Arial;
  color: #2c3e50;
}

</style>

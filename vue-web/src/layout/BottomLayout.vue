<template>
  <van-tabbar v-model="active">
    <van-tabbar-item icon="home-o">首页</van-tabbar-item>
    <van-tabbar-item icon="bag-o" dot>标签</van-tabbar-item>
    <van-tabbar-item icon="chat-o" badge="5">标签</van-tabbar-item>
    <van-tabbar-item icon="setting-o">标签</van-tabbar-item>
  </van-tabbar>

</template>

<script>
export default {
  name: "BottomLayout",
  data(){
    return{
      active: 0
    }
  }
}
</script>

<style scoped>

</style>

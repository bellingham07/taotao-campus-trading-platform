<template>
  <van-tabbar v-model="active">
    <van-tabbar-item icon="home-o">淘淘</van-tabbar-item>
    <van-tabbar-item icon="search">圈子</van-tabbar-item>
    <van-tabbar-item class="add">➕</van-tabbar-item>
    <van-tabbar-item icon="setting-o">消息</van-tabbar-item>
    <van-tabbar-item icon="setting-o">我的</van-tabbar-item>
  </van-tabbar>
</template>

<script setup lang="ts">
import {ref} from "vue";

const active = ref(0)
</script>



<style scoped>
.add {
  font-size: 35px;
}
</style>
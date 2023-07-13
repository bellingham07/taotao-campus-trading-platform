<template>
  <van-tabbar v-model="active">
    <van-tabbar-item icon="home-o">
      <router-link to="/home/new">淘淘</router-link>
    </van-tabbar-item>
    <van-tabbar-item icon="search">
      <router-link to="/atcl">圈子</router-link>
    </van-tabbar-item>
    <van-tabbar-item class="add">➕
      <router-link to="/cmdty"></router-link>
    </van-tabbar-item>
    <van-tabbar-item icon="setting-o">
      <router-link to="/msg">消息</router-link>
    </van-tabbar-item>
    <van-tabbar-item icon="setting-o">
      <router-link to="/user">我的</router-link>
    </van-tabbar-item>
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
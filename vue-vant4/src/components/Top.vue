<template>
  <div class="top">
    <div>淘 淘</div>
    <van-search
        v-model="value"
        show-action
        label="地址"
        placeholder="想要什么呢 😃"
        @search="onSearch"
        background="#1989fa"
    >
      <template #action>
        <div @click="onClickButton">搜索</div>
      </template>
    </van-search>
    <div>头像</div>

  </div>
</template>

<script setup lang="ts">
import {LoggedUser} from "../models/user.ts";
import useStore from "../store";

useStore
const userStr = localStorage.getItem("user")
if (userStr != null) {
  const loggedUser: LoggedUser = JSON.parse(userStr)


}


</script>

<style scoped>
.top {
  padding: 5px 0 0 0;
  background-color: #1989fa;
  text-align: center;
  font-size: 1.1rem;
}
</style>
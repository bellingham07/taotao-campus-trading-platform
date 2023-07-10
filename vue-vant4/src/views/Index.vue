<template>
  <top-option></top-option>
  <van-search
      v-model="value"
      show-action
      label="地址"
      
      placeholder="想要什么呢😃"
      @search="onSearch"
      background="#1989fa"
  >
    <template #action>
      <div @click="onClickButton">搜索</div>
<!--      <button class="search-btn" type="primary" @click="onClickButton">搜索</button>-->
    </template>
  </van-search>


  <bottom-menu/>
</template>

<script setup lang="ts">
import BottomMenu from "../components/BottomMenu.vue";
import TopOption from "../components/TopOption.vue";


</script>

<style scoped>
.search-btn {
  margin:0;
}
</style>
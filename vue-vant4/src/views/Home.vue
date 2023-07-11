<template>
  <top></top>
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
    </template>
  </van-search>
  <div class="type-btn">
    <van-button round type="primary" size="mini">最新发布</van-button>
    <van-button round plain type="primary" size="mini">找主人</van-button>
    <van-button round plain type="primary" size="mini">求好物</van-button>
  </div>

  <BuyList></BuyList>

  <bottom/>
</template>

<script setup lang="ts">
import Bottom from "../components/Bottom.vue";
import Top from "../components/Top.vue";
import BuyList from "../components/home/BuyList.vue";


</script>

<style scoped>
.search-btn {
  margin: 0;
}

.type-btn {
  width: 100%;
  text-align: center;

  .van-button {
    margin: 10px 1vw 0 1vw;
    width: 30%;
  }
}


</style>
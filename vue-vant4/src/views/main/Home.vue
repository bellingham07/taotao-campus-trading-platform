<template>
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
  <div class="type-btn">
    <van-button round type="primary" size="mini">
      <router-link to="/home/new">最新发布</router-link>
    </van-button>
    <van-button round plain type="primary" size="mini">
      <router-link to="/home/sell">找主人</router-link>
    </van-button>
    <van-button round plain type="primary" size="mini">
      <router-link to="/home/want">求好物</router-link>
    </van-button>
  </div>

  <router-view></router-view>
</template>

<script setup lang="ts">


</script>

<style scoped>
.search-btn {
  margin: 0;
}

.type-btn {
  width: 100%;
  text-align: center;


}

.van-button {
    margin: 10px 1vw 0 1vw;
    width: 30%;
}
</style>
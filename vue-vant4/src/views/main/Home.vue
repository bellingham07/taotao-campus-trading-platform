<template>
  

  <div class="type-btn">
    <van-button plain type="primary" size="mini">
      <router-link to="/home/sell">找主人</router-link>
    </van-button>
    <van-button plain type="primary" size="mini">
      <router-link to="/home/want">求好物</router-link>
    </van-button>
  </div>

  <router-view></router-view>

  <bottom></bottom>
</template>

<script setup lang="ts">


import Bottom from "../../components/Bottom.vue";
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
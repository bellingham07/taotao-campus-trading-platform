<template>
  <van-list
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
  >
    <van-cell v-for="item in goodsList" :key="item" :title="item" />
  </van-list>

</template>

<script>
/***
 * List 组件通过 loading 和 finished 两个变量控制加载状态，
 * 当组件滚动到底部时，会触发 load 事件并将 loading 设置成 true。
 * 此时可以发起异步操作并更新数据，数据更新完毕后，将 loading 设置成 false 即可。
 * 若数据已全部加载完毕，则直接将 finished 设置成 true 即可。*/
export default {
  name: "IndexMain",
  data(){
    return{
      loading:false,
      finished:true,
      goodsList:[]
    }
  },
  methods:{
    onLoad(){
      //
    }
  }
}
</script>

<style scoped>

</style>

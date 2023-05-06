<template>
  <van-cell is-link>
    <!-- 使用 title 插槽来自定义标题 -->
    <template #title>
      <span class="custom-title">商品1</span>
    </template>
    <template #label>
      <van-row>
        <van-col span="4">
          <van-image
            width="3rem"
            height="3rem"
            src="https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg"
        />
        </van-col>
        <van-col span="16" offset="2">
          <van-text-ellipsis
              rows="2"
              content="这个商品咋滴咋滴咋滴咋滴，描述的话语，我这里是为了把文本撑多多多，然后查看几行比较合适"
          />
        </van-col>
      </van-row>

<!--      <span>这个商品咋滴咋滴咋滴咋滴，描述的话语，我这里是为了把文本撑多</span>-->
    </template>
  </van-cell>
</template>

<script>
export default {
  name: "GoodsListItem"
}
</script>

<style scoped>

</style>

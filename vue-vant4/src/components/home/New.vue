<template>
  <van-row justify="space-around">
    <van-col class="cmdty-item" span="10">
      <van-button class="type" plain disabled type="primary" size="mini">出售</van-button>
      <van-image
          width="10rem"
          height="10rem"
          fit="cover"
          src="https://taotao-trading.oss-cn-beijing.aliyuncs.com/images/2023-06/11/408301323265285/pic_router19%3A35%3A50.png?Expires=1689059686&OSSAccessKeyId=TMP.3Kji69psLB4efZh27uht3mVSdgyRNeUdqbiqn4TGKJztPbt8C5roY6V6R6XA9ZJ1bVutWrSffsVfkiLUUVZz8UJDzunme9&Signature=LC38wO1jqVyoTkrsnIVZQCnQJI4%3D"
      />
      <div class="font09">ROG龙骑士二代</div>
      <div>
        <span class="font07 gray">ASUS</span>
        <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
        <span class="font07 gray">几乎全新</span>
      </div>
      <div>
        <span class="price">￥1299</span>
        <span class="collect gray">38人想要</span>
      </div>
    </van-col>

    <van-col class="cmdty-item" span="10">
      <van-button class="type" plain disabled type="primary" size="mini">想要</van-button>
      <van-image
          width="10rem"
          height="10rem"
          fit="cover"
          src="https://my-halo-blog.oss-cn-nanjing.aliyuncs.com/pcis/QQ%E5%9B%BE%E7%89%8720230514210007.png"
      />
      <div class="font09">ROG龙骑士二代</div>
      <div>
        <span class="font07 gray">ASUS</span>
        <span class="gray">&nbsp;&nbsp;|&nbsp;&nbsp;</span>
        <span class="font07 gray">几乎全新</span>
      </div>
      <div>
        <span class="price">￥1299</span>
        <span class="collect gray">38人收藏</span>
      </div>
    </van-col>

  </van-row>

</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {cmdtyService} from '../../api'

onMounted(() => {
  async function asd() {
    await cmdtyService.get('/')
  }
})
</script>

<style scoped>
.cmdty-item {
  margin: 5px 0 0 10px;
  font-weight: 700;
}

.font09 {
  font-size: 0.9rem;
}

.font07 {
  font-size: 0.7rem;
}

.price {
  color: #ff1b1b;
  font-weight: 700;
}

.gray {
  color: #8d8d8d;
}

.collect {
  margin: 20px;
  font-size: 0.5rem;
}

.type {
  margin: 0 0 5px 0;
  width: 10rem;
}
</style>
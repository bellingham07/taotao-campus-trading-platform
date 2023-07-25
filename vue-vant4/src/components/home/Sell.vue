<template>
  <van-row justify="space-around">
    <van-col class="cmdty-item" span="10">
      <div>
        <div @click="toDetail(cs.id)">
          <van-image
              width="30vw"
              height="30vw"
              fit="cover"
              :src="'https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg'"
          />
          <div class="font09">ask到付哈四道口附近哈数据库·1</div>
          <div>
            <span class="font07 gray">{ dasfasdf }}</span>
            <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
            <span class="font07 gray">asdasd </span>
          </div>
          <div>
            <span class="price">￥{ cs.price }}</span>
            <span class="collect gray">{ cs.collect }}</span>
          </div>
        </div>
        <div>头像</div>
      </div>

      <div v-for="cs in cis.value" :key="cs.id">
        <van-image
            width="10rem"
            height="10rem"
            fit="cover"
            :src="cs.cover"
        />
        <div class="font09">{{ cs.briefIntro }}</div>
        <div>
          <span class="font07 gray">{{ cs.brand }}</span>
          <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
          <span class="font07 gray">{{ cs.old }}</span>
        </div>
        <div>
          <span class="price">￥{{ cs.price }}</span>
          <span class="collect gray">{{ cs.collect }}</span>
        </div>
      </div>
    </van-col>
  </van-row>

</template>

<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {listCisPageReq} from "../../api/cmdty";
import {CmdtyInfo} from "../../models/cmdty.ts";
import router from "../../routers";

// const data = reactive({
//       cis: []
//     })
let page = 1

const cis = ref<CmdtyInfo[]>([])

const getCis = async () => {
  const listCisPageResp = await listCisPageReq(1, page);
  if (listCisPageResp.code === 0) {
    cis.value = listCisPageResp.data.cis
  }
  return null
}
const toDetail = (id: number) => {
  router.push(`/cmdty/${id}`)
}

onMounted(() => {
  getCis()
})
</script>


<style scoped>
.cmdty-item {
  margin: 10px 0 0 10px;
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
</style>
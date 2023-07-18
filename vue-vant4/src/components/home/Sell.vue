<template>
  <van-row justify="space-around">
    <van-col class="cmdty-item" span="10">
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
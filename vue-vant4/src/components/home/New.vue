<template>
    <van-row justify="space-around">
        <van-col class="cmdty-item" span="10">
            <van-button class="type" plain disabled type="primary" size="mini">出售</van-button>
            <div v-for="cs in data.sell" :key="cs.id">
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

        <van-col class="cmdty-item" span="10">
            <van-button class="type" plain disabled type="primary" size="mini">想要</van-button>
            <div v-for="cw in data.want" :key="cw.id">
                <van-image
                        width="10rem"
                        height="10rem"
                        fit="cover"
                        :src="cw.cover"
                />
                <div class="font09">{{ cw.briefIntro }}</div>
                <div>
                    <span class="font07 gray">{{ cw.brand }}</span>
                    <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
                    <span class="font07 gray">{{ cw.old }}</span>
                </div>
                <div>
                    <span class="price">￥{{ cw.price }}</span>
                    <span class="collect gray">{{ cw.collect }}</span>
                </div>
            </div>
        </van-col>

    </van-row>

</template>

<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {newestResp} from "../../api/cmdty/type.ts";
import {cmdtyInfo} from "../../models/cmdty.ts";
import {cmdtyService} from "../../api";

let hasSell = true
let hasWant = true
const data = reactive({sell: [], want: []})

onMounted(async () => {

    await cmdtyService.get<any, newestResp>("/cache").then((resp) => {
        data.sell = resp.data.data.sell
        console.log(resp.data.data.sell)
        data.want = resp.data.want
        console.log(resp.data.data.want)
    })
    // if (resp.data.data.sell.length === 0) {
    //     hasSell = false
    // } else {
    // console.log(sell = resp.data.data.sell)
    // }
    // if (resp.data.data.want.length === 0) {
    //     hasWant = false
    // } else {
    //     want = resp.data.data.want
    // }

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
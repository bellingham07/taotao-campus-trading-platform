<template>
    <div>
        <van-image radius="100" fit="cover" width="100" height="100" :src=userInfo.avatar
        />
        <div>{{ userInfo.name }}</div>
        <div></div>
    </div>

    <div>{{ cmdtyInfo.price }}</div>
    <div>{{ cmdtyInfo.intro }}</div>

    <div class="pics">
        <div v-for="pic in cmdtyPics">
            <van-image fit="cover" :src=pic.url width="100" height="100"/>
        </div>

    </div>
    <van-action-bar>
        <van-action-bar-icon icon="chat-o" text="客服" color="#ee0a24"/>
        <van-action-bar-icon icon="cart-o" text="购物车"/>
        <van-action-bar-icon icon="star" text="已收藏" color="#ff5000"/>
        <van-button style="position: absolute; right: 5%" @click="toChat" round color="#3456fd" icon="chat-o">我想要
        </van-button>
        <!--        <van-action-bar-button @click="toChat" type="danger" icon="chat-o" color="#3456fd" text="我想要"/>-->
    </van-action-bar>


</template>

<script setup lang="ts">
import {useRoute} from "vue-router";
import {getInfoReq} from "../../api/cmdty";
import {onMounted, reactive, ref} from "vue";
import {createRoom} from "../../api/chat";
import {CreateRoomReq} from "../../api/chat/type.ts";
import {UserInfo} from "../../models/user.ts";
import {CmdtyInfo, CmdtyPics} from "../../models/cmdty.ts";
import userStorage from "../../store/user.ts";
import router from "../../routers";

const userStore = userStorage()

const cmdtyInfo: CmdtyInfo = ref<CmdtyInfo>({})
const userInfo: UserInfo = ref<UserInfo>({});
const cmdtyPics: CmdtyPics[] = ref<CmdtyPics[]>([])

const id: number = useRoute().params.id as number

const getInfo = async () => {
    const resp = await getInfoReq(id)
    cmdtyInfo.value = resp.data.cmdtyInfo
    userInfo.value = resp.data.userInfo
    cmdtyPics.value.push(...resp.data.cmdtyPics)
}

const toChat = () => {
    if (userStore.token == null || userStore.token === '') {
        router.push('/login')
        return
    }

    const req: CreateRoomReq = {
        cmdtyId: cmdtyInfo.id,
        sellerId: userInfo.id,
        seller: userInfo.name,
        buyerId: userStore.loggedUser.id,
        buyer: userStore.loggedUser.name,
        cover: cmdtyInfo.cover
    }

    const resp = createRoom(req)
}

onMounted(() => {
    getInfo()
})
</script>


<style scoped>
color {
    color: #3456fd
}
</style>
<script setup lang="ts">
import {useRoute} from "vue-router";
import {getInfoReq} from "../../api/cmdty";
import {onMounted, ref} from "vue";

const cmdtyInfo = ref({})
const id: number = useRoute().params.id as number
const getInfo = async () => {
    const resp = await getInfoReq(id)
    cmdtyInfo.value = resp.data.cmdtyInfo
    console.log(resp.data)
}

onMounted(() => {
    getInfo()
})
</script>

<template>
  <div>我是详情</div>
  <div></div>
</template>

<style scoped>

</style>
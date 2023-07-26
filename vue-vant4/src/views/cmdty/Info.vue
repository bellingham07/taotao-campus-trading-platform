<template>
  <div>
    <img :src=userInfo.avatar>
    <div>{{ userInfo.name}}</div>
    <div></div>
  </div>

  <div>{{ cmdtyInfo.price}}</div>
  <div>{{cmdtyInfo.intro}}</div>
</template>

<script setup lang="ts">
import {useRoute} from "vue-router";
import {getInfoReq} from "../../api/cmdty";
import {onMounted, ref} from "vue";

const cmdtyInfo = ref({})
const userInfo = ref({})
const id: number = useRoute().params.id as number

const getInfo = async () => {
  const resp = await getInfoReq(id)

  cmdtyInfo.value = resp.data.cmdtyInfo
  userInfo.value = resp.data.userInfo
}

onMounted(() => {
  getInfo()
})
</script>


<style scoped>

</style>
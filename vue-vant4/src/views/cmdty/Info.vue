<template>
  <div>头像</div>

  <div>{{ cmdtyInfo.price}}</div>
  <div>{{cmdtyInfo.intro}}</div>
</template>

<script setup lang="ts">
import {useRoute} from "vue-router";
import {getInfoReq} from "../../api/cmdty";
import {onMounted, ref} from "vue";

const cmdtyInfo = ref({})
const id: number = useRoute().params.id as number

const getInfo = async () => {
  const resp = await getInfoReq(id)

  cmdtyInfo.value = resp.data.cmdtyInfo
}

onMounted(() => {
  getInfo()
})
</script>


<style scoped>

</style>
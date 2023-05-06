<template>
  <header-layout/>
  <goods-list-item/>
  <bottom-layout/>
</template>

<script>
import HeaderLayout from "@/layout/HeaderLayout";
import BottomLayout from "@/layout/BottomLayout";
import GoodsListItem from "@/components/index/GoodsListItem";
export default {
  name: "MainLayout",
  components:{
    HeaderLayout,
    BottomLayout,
    GoodsListItem
  }
}
</script>

<style scoped>
</style>

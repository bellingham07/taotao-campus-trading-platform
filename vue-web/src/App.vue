<template>
  <keep-alive>
    <router-view></router-view>
  </keep-alive>
</template>

<script>
import {useStore} from 'vuex'

export default {
  name: 'App',
  components: {
  },
  mounted() {
    const store = useStore()
    console.log(store)
    // store.dispatch('hi','')
  }
}
</script>

<style>
*{
  padding: 0;
  margin: 0;
}
</style>

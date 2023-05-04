<template>
<!--  <LoginForm/>-->
  <MainLayout/>
</template>

<script>
// import LoginForm from './components/login/LoginForm'
import {useStore} from 'vuex'
import MainLayout from "@/layout/MainLayout";

export default {
  name: 'App',
  components: {
    // LoginForm,
    MainLayout

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

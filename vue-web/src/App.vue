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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

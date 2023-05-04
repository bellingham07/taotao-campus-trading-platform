<template>
  <div class="common-layout">
    <el-container>
      <el-header  class="header">Header</el-header>
      <el-container>
        <el-aside width="15%" class="aside">Aside</el-aside>
        <el-container>
          <el-main class="main">Main</el-main>
          <el-footer class="footer">Footer</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "MainLayout"
}
</script>

<style scoped>
.header{
  background-color: lightgreen
}
.aside{
  background-color: lightcoral
}
.main{
  background-color: lightslategray
}
.footer{
  background-color: lightseagreen
}
</style>

<template>
  <router-view></router-view>
</template>

<script>
import 'mint-ui/lib/style.css'
export default {
  name: 'app',
  components: {
  }
}
</script>

<style>
body{
  margin:0px;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.container{
  width: 100wh;
  margin: 10px;
}

.white-space{
  height: 20px;
}
</style>

<template>
  <router-view></router-view>
</template>

<script>
import "mint-ui/lib/style.css";
export default {
  name: "app",
  components: {}
};
</script>

<style>
body {
  margin: 0px;
}
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.container {
  width: 100wh;
  margin: 10px;
}

.white-space {
  height: 20px;
}

.title {
  text-align: center;
  font-size: 30px;
  font-weight: light;
  color: #333;
  margin: 30px 0 10px 0;
}
.line{
  border-top: solid rgb(217, 217, 217) 0.5px;
}
.subtitle {
  text-align: left;
  font-size: 16px;
  color: #888;
  text-indent: 2em;
}

</style>

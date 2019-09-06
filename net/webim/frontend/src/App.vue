<template>
  <div id="app">
    <config uri="ws://localhost:60012" name="wijie"/>
    <SendBox/>
  </div>
</template>

<script>

import Config from './components/Config.vue'
import SendBox from './components/Send.vue'
export default {
  name: 'app',
  components: {
    Config,
    SendBox
  }
}
</script>

<style lang="scss">
@import 'bulma/bulma.sass';
$fa-font-path: "~font-awesome/fonts";
@import 'font-awesome/scss/font-awesome.scss';
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

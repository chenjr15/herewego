<template>
  <div class="box">
    <div class="columns is-mobile">
      <div class="control column">
        <input class="input" id="name" type="text" :value="name" placeholder="Your Name" >
      </div>
      <div class="column control" >
        <div class="button is-info" onclick="">Set Name</div>
      </div>

    </div>
    <div class="columns is-mobile">
      <div class="control column">
        <input class="input" id="wsuri" type="url" :value="uri" placeholder="Web socket url">
      </div>
      <div class="column control">
        <a class="button is-info" onclick="connectSocket()">Connect</a>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: 'Config',
  props: {
    uri: String,
    name: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

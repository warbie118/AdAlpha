<template>
  <div id="app">
    <!--<img alt="Vue logo" src="./assets/logo.png">-->
    <portfolio></portfolio>
    <history v-show="this.showHistory"></history>
  </div>
</template>

<script>
import Portfolio from './components/Portfolio'
import History from './components/History'

export default {
  name: 'app',
  components: {
    Portfolio,
    History
  },
  data () {
    return {
      showHistory: true
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: whitesmoke;
  color: #2c3e50;
}
</style>

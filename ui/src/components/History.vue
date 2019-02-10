<template>
  <div id="history" class="ui vertical stripe segment">
    <div class="ui container">
      <div id="content" class="ui basic segment">
        <h3 class="ui header">History</h3>
        <vuetable ref="vuetable"
                  api-url="http://localhost:3030/history/investor/1"
                  :fields="fields"
                  data-path=""
        ></vuetable>
      </div>
    </div>
  </div>
</template>

<script>
import Vuetable from 'vuetable-2/src/components/Vuetable'

export default {
  components: {
    Vuetable
  },
  data () {
    return {
      fields: [
        'Instruction', 'Isin', 'Asset', { name: 'Asset_price', title: 'Asset Price(£)' },
        'Units', { name: 'CurrencyCode', title: 'Currency Code' }, 'Amount'
      ]
    }
  }
}
</script>

<style scoped>
  #history {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #045000;
  }
</style>

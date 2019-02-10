<template>
  <div id="portfolio" class="ui vertical stripe segment">
    <div class="ui container">
      <div id="content" class="ui basic segment">
        <h3 class="ui header">Portfolio</h3>
        <vuetable ref="vuetable"
                  api-url="http://localhost:3030/portfolio/investor/1"
                  :fields="fields"
                  data-path=""
        >
        <template slot="actions" scope="props">
          <div class="custom-actions">
            <button class="ui basic button"
                    style="padding: 4px"
                    @click="buy('buy-item', props.rowData, props.rowIndex)">Buy</button>
            <button class="ui basic button"
                    style="padding: 4px"
                    @click="invest('invest-item', props.rowData, props.rowIndex)">Invest</button>
            <button class="ui basic button"
                    style="padding: 4px"
                    @click="sell('sell-item', props.rowData, props.rowIndex)">Sell</button>
            <button class="ui basic button"
                    style="padding: 4px"
                    @click="raise('raise-item', props.rowData, props.rowIndex)">Raise</button>
          </div>
        </template>
        </vuetable>
        <buy-form :isin="selectedIsin"
          v-show="isBuyVisible"
          @close=this.closeBuy
        />
        <sell-form :isin="selectedIsin"
          v-show="isSellVisible"
          @close=this.closeSell
        />
        <invest-form :isin="selectedIsin"
          v-show="isInvestVisible"
          @close=this.closeInvest
        />
        <raise-form :isin="selectedIsin"
          v-show="isRaiseVisible"
          @close=this.closeRaise
        />
      </div>
    </div>
  </div>
</template>

<script>
import Vuetable from 'vuetable-2/src/components/Vuetable'
import BuyForm from './BuyForm'
import SellForm from './SellForm'
import InvestForm from './InvestForm'
import RaiseForm from './RaiseForm'

export default {
  components: {
    Vuetable,
    BuyForm,
    SellForm,
    InvestForm,
    RaiseForm
  },
  created () {
    this.interval = setInterval(() => this.$refs.vuetable.refresh(), 2000)
  },
  data () {
    return {
      isBuyVisible: false,
      isSellVisible: false,
      isRaiseVisible: false,
      isInvestVisible: false,
      selectedIsin: '',
      fields: [
        'Isin', 'Asset', { name: 'Current_price', title: 'Current Price(£)' }, 'Units', {
          name: '__slot:actions',
          title: 'Instructions',
          titleClass: 'center aligned',
          dataClass: 'center aligned'
        }
      ]
    }
  },
  methods: {
    buy (action, data, index) {
      console.log('slot action: ' + action, data.Isin, index)
      this.showBuy(data.Isin)
    },
    invest (action, data, index) {
      console.log('slot action: ' + action, data.units, index)
      this.showInvest(data.Isin)
    },
    sell (action, data, index) {
      console.log('slot action: ' + action, data.name, index)
      this.showSell(data.Isin)
    },
    raise (action, data, index) {
      console.log('slot action: ' + action, data.name, index)
      this.showRaise(data.Isin)
    },
    showBuy (isin) {
      this.isBuyVisible = true
      this.selectedIsin = isin
      this.$parent.showHistory = false
    },
    closeBuy () {
      this.isBuyVisible = false
      this.$parent.showHistory = true
    },
    showSell (isin) {
      this.isSellVisible = true
      this.selectedIsin = isin
      this.$parent.showHistory = false
    },
    closeSell () {
      this.isSellVisible = false
      this.$parent.showHistory = true
    },
    showInvest (isin) {
      this.isInvestVisible = true
      this.selectedIsin = isin
      this.$parent.showHistory = false
    },
    closeInvest () {
      this.isInvestVisible = false
      this.$parent.showHistory = true
    },
    showRaise (isin) {
      this.isRaiseVisible = true
      this.selectedIsin = isin
      this.$parent.showHistory = false
    },
    closeRaise () {
      this.isRaiseVisible = false
      this.$parent.showHistory = true
    }
  }
}
</script>

<style>
  #portfolio {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #045000;
  }
</style>

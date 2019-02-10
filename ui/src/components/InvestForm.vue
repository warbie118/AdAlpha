<template>
  <div class="modal-backdrop">
    <div class="modal">
      <header class="modal-header">
        <slot name="header">
          Invest Instruction
        </slot>
      </header>
      <section class="modal-body">
        <slot name="body">
          <vue-form
            :model="form"
            style="width: 500px;">
            <vue-form-item label="Isin">
              <vue-input
                :placeholder="isin"
                :disabled="true">
              </vue-input>
            </vue-form-item>
            <vue-form-item label="Currency Code">
              <vue-radio
                v-model="form.currencyCode"
                value="GBP">
                GBP
              </vue-radio>
              <vue-radio
                v-model="form.currencyCode"
                value="USD">
                USD
              </vue-radio>
              <vue-radio
                v-model="form.currencyCode"
                value="EUR">
                EUR
              </vue-radio>
              <vue-radio
                v-model="form.currencyCode"
                value="CNY">
                CNY
              </vue-radio>
            </vue-form-item>
            <vue-form-item label="Amount">
              <vue-input
                v-model="form.amount"
                style="width: 100%">
              </vue-input>
            </vue-form-item>
          </vue-form>
        </slot>
      </section>
      <footer class="modal-footer">
        <slot name="footer">
          <button
            type="button"
            class="btn-green"
            @click="invest"
          >Invest</button>
          <button
            type="button"
            class="btn-green"
            @click="close"
          >Cancel</button>
        </slot>
      </footer>
    </div>
  </div>
</template>

<script>
import 'vfc/dist/vfc.css'
import { Input, Form, Radio, FormItem } from 'vfc'
import VueForm from 'vfc/src/components/form/Form'
import VueInput from 'vfc/src/components/input/Input'
import axios from 'axios'

export default {
  name: 'modal',
  components: {
    VueInput,
    VueForm,
    [Input.name]: Input,
    [Radio.name]: Radio,
    [Form.name]: Form,
    [FormItem.name]: FormItem
  },
  props: {
    isin: {
      type: String,
      required: true
    }
  },
  methods: {
    close () {
      this.$emit('close')
    },
    invest () {
      axios.post(`http://localhost:3030/instruction/invest`, { 'investor_id': 1, 'isin': this.isin, 'currency_code': this.form.currencyCode, 'amount': Number(this.form.amount) },
        { headers: { 'Content-Type': 'application/json' } })
        .then(response => {})
        .then(this.close)
        .catch(e => {
          this.errors.push(e)
        })
    }
  },
  data () {
    return {
      form: {
        isin: '',
        currencyCode: '',
        amount: ''
      }
    }
  }
}
</script>

<style>
  .modal-backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: rgba(0, 0, 0, 0.3);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .modal {
    background: #FFFFFF;
    box-shadow: 2px 2px 20px 1px;
    overflow-x: auto;
    display: flex;
    flex-direction: column;
  }

  .modal-header,
  .modal-footer {
    padding: 15px;
    display: flex;
  }

  .modal-header {
    border-bottom: 1px solid #eeeeee;
    color: #4AAE9B;
    justify-content: space-between;
  }

  .modal-footer {
    border-top: 1px solid #eeeeee;
    justify-content: flex-end;
  }

  .modal-body {
    position: relative;
    padding: 20px 10px;
  }

  .btn-close {
    border: none;
    font-size: 20px;
    padding: 20px;
    cursor: pointer;
    font-weight: bold;
    color: #4AAE9B;
    background: transparent;
  }

  .btn-green {
    color: white;
    background: #4AAE9B;
    border: 1px solid #4AAE9B;
    border-radius: 2px;
  }
</style>

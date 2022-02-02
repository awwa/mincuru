<template>
  <div>
    <h1>クルマ</h1>
    <v-layout wrap>
      <v-spacer />
      <v-btn @click="add">追加</v-btn>
    </v-layout>
    <v-layout wrap>
      <v-select 
        label="メーカー"
        :items="$makers"
        item-text="value"
        item-value="value" />
      <v-select
        label="モデル" />
      <v-select
        label="グレード" />
      <v-text-field
        label="型式" />
    </v-layout>
    <v-layout wrap>
      <p>時期</p>
    </v-layout>

    <CarList :cars="cars" />
  </div>
</template>

<script>
import { DefaultApi, Configuration } from "../../../api-client"
export default {
  data() {
    return {
      query: {
        maker_name: "",
        model_name: "",
        grade_name: "",
        model_code: "",
        price_lower: 0,
        price_higher: 10000000,
        model_change_from: "1974-01-01",
        model_change_to: "2022-12-31",
        power_train: ["ICE"],
      }
    }
  },
  async asyncData({$axios}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.postCarsSearch(/*this.query*/)
    return {
      cars: resp.data
    }
  },
  methods: {
    add() {
      this.$router.push("/cars/add")
    }
  }
}
</script>
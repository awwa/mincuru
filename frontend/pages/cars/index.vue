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
        :items="makers"
        @change="selectMaker"
        v-model="query.maker_names"
        multiple />
      <v-select
        label="タイプ"
        :items="body_types"
        v-model="query.body_types"
        multiple />
      <v-select
        label="モデル"
        :items="models" />
      <!--
      <v-select
        label="グレード" />
      <v-text-field
        label="型式" />
      -->
    </v-layout>
    <v-layout wrap>
      <p>時期</p>
    </v-layout>
    <v-layout wrap>
      <v-spacer />
      <v-btn @click="search">検索</v-btn>
    </v-layout>

    <CarList :cars="cars" />
  </div>
</template>

<script>
import { DefaultApi, Configuration } from "../../../api-client"
const conf = new Configuration()
export default {
  data() {
    return {
      query: {
        body_types: [],
        maker_names: [],
        maker_name: "",
        model_name: "",
        grade_name: "",
        model_code: "",
        price_lower: 0,
        price_higher: 10000000,
        model_change_from: "1974-01-01",
        model_change_to: "2022-12-31",
        power_train: [],
      },
      makers: [],
      body_types: [],
    }
  },
  async asyncData({$axios}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    // メーカーリスト取得
    const respMakers = await api.getMakers()
    // ボディタイプリスト取得
    const respBodyTypes = await api.getCarsBodyTypes()
    return {
      makers: respMakers.data.map(item => {return {"value": item, "text": item}}),
      cars: [],
      models: [],
      body_types: respBodyTypes.data.map(item => {return {"value": item, "text": item}}),
    }
  },
  methods: {
    add() {
      this.$router.push("/cars/add")
    },
    async search() {
      console.log(this.query)
      const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
      const resp = await api.postCarsSearch(this.query)
      this.cars = resp.data
    },
    async selectMaker(e) {
      const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
      const resp = await api.getCarsMakersModels(e)
      this.models = resp.data.map(item => {return {"value": item, "text": item}})
    },
  }
}
</script>
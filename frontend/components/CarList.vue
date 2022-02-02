<template>
  <div>
    <!-- <v-text-field
      v-model="search"
      append-icon="mdi-magnify"
      label="検索"
      single-line
      hide-details
    /> -->
    <v-data-table
      :headers="headers"
      :items="cars"
      :items-per-page="5"
      @click:row="editItem"
    >
      <template v-slot:[`item.image_url`]="{ item }">
        <v-img :src="item.image_url"
          height="3vw" 
          min-height="80px"
          width="5.3vw" 
          min-width="160px"
          alt="イメージなし"
        />
      </template>
      <template v-slot:[`item.price`]="{ item }">
        {{$formatCurrencyJP(item.price)}}
      </template>
      <template v-slot:[`item.model_change_full`]="{ item }">
        {{$formatMonth(item.model_change_full)}}
      </template>
      <template v-slot:[`item.model_change_last`]="{ item }">
        {{$formatMonth(item.model_change_last)}}
      </template>
    </v-data-table>
  </div>
</template>

<script lang="js">
export default {
  props: {
    "cars": {
      type: Array,
      required: true,
    }
  },
  data() {
    return {
      search: "",
      headers: [
        {
          text: "イメージ",
          align: "start",
          value: "image_url",
        },
        {
          text: "メーカー",
          value: "maker_name",
        },
        {
          text: "モデル",
          value: "model_name",
        },
        {
          text: "グレード",
          value: "grade_name",
        },
        {
          text: "型式",
          value: "model_code",
        },
        {
          text: "小売価格",
          value: "price",
        },
        {
          text: "フルモデルチェンジ",
          value: "model_change_full",
        },
        {
          text: "最終モデルチェンジ",
          value: "model_change_last",
        },
      ],
    }
  },
  methods: {
    editItem(item) {
      this.$router.push(`/cars/${item.id}`)
    }
  }
}
</script>
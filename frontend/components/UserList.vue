<template>
  <div>
    <v-text-field
      v-model="search"
      append-icon="mdi-magnify"
      label="検索"
      single-line
      hide-details
    />
    <v-data-table
      :headers="headers"
      :items="users"
      :items-per-page="5"
      :search="search"
      @click:row="editItem"
    />
  </div>
</template>

<script lang="ts">
export default {
  props: {
    "users": {
      type: Array,
      required: true,
    }
  },
  data() {
    return {
      search: "",
      headers: [
        {
          text: "ID",
          align: "start",
          value: "id",
        },
        {
          text: "名前",
          value: "name",
        },
        {
          text: "メールアドレス",
          value: "email",
        },
        {
          text: "ロール",
          value: "role",
        },
      ],
    }
  },
  methods: {
    editItem(item) {
      this.$router.push(`/users/${item.id}`)
    }
  }
}
</script>
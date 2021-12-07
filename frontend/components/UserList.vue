<template>
  <v-data-table
    :headers="headers"
    :items="users"
    :items-per-page="5"
  >
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        {
          text: 'ID',
          align: 'start',
          value: 'id',
        },
        {
          text: '名前',
          value: 'name',
        },
        {
          text: 'メールアドレス',
          value: 'email',
        },
        {
          text: 'ロール',
          value: 'role',
        },
      ],
    }
  },
  props: ["users"],
}
</script>
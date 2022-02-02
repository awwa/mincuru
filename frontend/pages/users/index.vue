<template>
  <div>
    <h1>ユーザー</h1>
    <v-layout wrap>
      <v-spacer />
      <v-btn @click="add">追加</v-btn>
    </v-layout>
    <UserList :users="users" />
  </div>
</template>

<script>
import { DefaultApi, Configuration } from "../../../api-client"
export default {
  async asyncData({$axios}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.getUsers()
    return {
      users: resp.data
    }
  },
  methods: {
    add() {
      this.$router.push("/users/add")
    }
  }
}
</script>
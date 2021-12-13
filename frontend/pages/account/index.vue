<template>
  <div>
    <h1>アカウント設定</h1>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>ID</strong></label>
          </v-col>
          <v-col>
            <label>{{ this.user.id }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>名前</strong></label>
          </v-col>
          <v-col>
            <label>{{ this.user.name }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>メールアドレス</strong></label>
          </v-col>
          <v-col>
            <label>{{ this.user.email }}</label>
          </v-col>
        </v-row>
        <v-row v-if="$isAdmin()">
          <v-col cols="12" sm="2">
            <label><strong>ロール</strong></label>
          </v-col>
          <v-col>
            <label>{{ this.user.role }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-btn @click="edit">編集</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
import { DefaultApi, Configuration } from '../../../api-client'
export default {
  async asyncData({$axios}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.getUsersMe()
    return {
      user: resp.data
    }
  },
  // data() {
  //   return {
  //   }
  // },
  // computed() {
  //   return {
  //   }
  // },
  methods: {
    edit() {
      this.$router.push(`/account/edit`)
    },
  }
}
</script>
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
            <label>{{ this.$auth.user.id }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>名前</strong></label>
          </v-col>
          <v-col>
            <v-text-field
              v-model="user.name"
              required
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>メールアドレス</strong></label>
          </v-col>
          <v-col>
            <v-text-field
              v-model="user.email"
              required
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>パスワード</strong></label>
          </v-col>
          <v-col>
            <v-text-field
              v-model="user.password"
              required
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>パスワードの確認</strong></label>
          </v-col>
          <v-col>
            <v-text-field
              v-model="password_confirm"
              required
            />
          </v-col>
        </v-row>
        <v-row v-if="$isAdmin()">
          <v-col cols="12" sm="2">
            <label><strong>ロール</strong></label>
          </v-col>
          <v-col>
            <label>{{ user.role }}</label>
          </v-col>
        </v-row>
        <v-row>
          <!-- <v-btn @click="edit">編集</v-btn> -->
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
// import { DefaultApi, Configuration } from '../../../../api-client'
export default {
  // async asyncData({$axios, params}) {
  //   const conf = new Configuration()
  //   const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
  //   const resp = await api.getUserMe()
  //   return {
  //     user: resp.data
  //   }
  // },
  data() {
    return {
      user: this.$auth.user,
      password_confirm: ""
    }
  },
  computed: {
    // isAdmin() {
    //   return this.$isAdmin(this.$auth.user.role)
    // }
  },
  methods: {
    edit() {
      this.$router.push(`/account/edit`)
    },
  //   async del() {
  //     try {
  //       const conf = new Configuration()
  //       const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
  //       const resp = await api.deleteUser(this.$route.params.id)
  //       this.$router.go(-1)
  //     } catch (err) {
  //       this.error = "削除に失敗しました"
  //       this.hasError = true
  //     }
  //   }
  }
}
</script>
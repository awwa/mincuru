<template>
  <div>
    <h1>アカウント編集</h1>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>ID</strong></label>
          </v-col>
          <v-col>
            <label>{{ user.id }}</label>
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
          <v-btn @click="save">保存</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
import { DefaultApi, Configuration } from '../../../api-client'
export default {
  async asyncData({$axios, params}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.getUsersMe()
    return {
      user: resp.data
    }
  },
  data() {
    return {
  //     user: {
  //       id: this.$auth.user.id,
  //       name: this.$auth.user.name,
  //       email: this.$auth.user.email,
  //     },
      password_confirm: ""  // TODO 入力中に比較チェック
    }
  },
  computed: {
    // isAdmin() {
    //   return this.$isAdmin(this.$auth.user.role)
    // }
  },
  methods: {
    async save() {
      const conf = new Configuration()
      const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
      const params = {
        email: this.user.email,
        name: this.user.name,
        password: this.user.password,
      }
      console.log(params)
      const resp = await api.patchUsersMe(params)
      this.$router.push(`/account`)
    },
  }
}
</script>
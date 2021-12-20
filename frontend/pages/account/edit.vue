<template>
  <div>
    <h1>アカウント編集</h1>
    <v-form
      @submit.prevent="save"
      ref="form"
    >
      <v-container>
        <v-row>
          <v-text-field
            label="名前"
            v-model="user.name"
            required
          />
        </v-row>
        <v-row>
          <v-text-field
            label="メールアドレス"
            v-model="user.email"
            required
          />
        </v-row>
        <v-row>
          <v-text-field
            label="パスワード"
            v-model="user.password"
            :rules="passwordRules"
            required
            type="password"
          />
          <v-text-field
            label="パスワード確認"
            v-model="password_confirm"
            :rules="passwordConfirmRules"
            required
            type="password"
          />
        </v-row>
        <v-row v-if="$isAdmin()">
          <label>{{ $roleLabel(user.role) }}</label>
        </v-row>
        <v-row>
          <v-btn @click="save">保存</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script lang="js">
import { DefaultApi, Configuration } from '../../../api-client'
export default {
  async asyncData({$axios, params}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.getUserMe()
    return {
      user: resp.data
    }
  },
  data() {
    return {
      password_confirm: "",
      nameRules: [
        v => !!v || "名前を入力してください",
      ],
      emailRules: [
        v => !!v || "メールアドレスを入力してください",        
      ],
      passwordRules: [
        v => !!v || "パスワードを入力してください",
      ],
      passwordConfirmRules: [
        v => !!v || "パスワードを入力してください",
        v => v == this.user.password || "パスワード確認が一致していません",
      ],
    }
  },
  // computed: {
    // isAdmin() {
    //   return this.$isAdmin(this.$auth.user.role)
    // }
  // },
  methods: {
    async save() {
      try {
        const conf = new Configuration()
        const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
        const params = {
          email: this.user.email,
          name: this.user.name,
          password: this.user.password,
        }
        const resp = await api.patchUserMe(params)
        this.$router.push(`/account`)
      } catch (err) {
        this.error = "保存に失敗しました"
        this.hasError = true
      }
    },
  }
}
</script>
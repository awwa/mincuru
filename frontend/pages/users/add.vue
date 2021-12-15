<template>
  <div>
    <h1>ユーザー追加</h1>
    <v-form
      @submit.prevent="add"
      ref="form"
    >
      <v-container>
        <v-row>
          <v-text-field
            label="名前"
            v-model="user.name"
            :rules="nameRules"
            required
          />
        </v-row>
        <v-row>
          <v-text-field
            label="メールアドレス"
            v-model="user.email"
            :rules="emailRules"
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
        <v-row>
          <v-select
            label="ロール"
            :items="$roles"
            item-text="label"
            item-value="value"
            v-model="user.role"
            :rules="roleRules"
          />
        </v-row>
        <v-row>
          <v-btn type="submit">追加</v-btn>
          <v-btn @click="back">キャンセル</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
import { DefaultApi, Configuration } from '../../../api-client'
export default {
  data() {
    return {
      user: {
        name: "",
        email: "",
        role: "{value:'user', text:'ユーザー'}",
      },
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
      roleRules: [
        v => !!v || "ロールを入力してください",
      ],
    }
  },
  methods: {
    async add() {
      try {
        if (this.$refs.form.validate()) {
          const conf = new Configuration()
          const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
          const resp = await api.postUsers(this.user)
          this.$router.go(-1)
        }
      } catch (err) {
        this.error = "追加に失敗しました"
        this.hasError = true
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>
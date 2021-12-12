<template>
  <div>
    <h1>ユーザー追加</h1>
    <v-form>
      <v-container>
        <v-row>
          <v-col col="12" sm="2">
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
              type="password"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>ロール</strong></label>
          </v-col>
          <v-col>
            <v-select
              :items="roles"
              v-model="user.role"
            ></v-select>
          </v-col>
        </v-row>
        <v-row>
          <v-btn @click="add">追加</v-btn>
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
        role: "user",
      },
      roles: ["user", "admin"]
    }
  },
  methods: {
    async add() {
      try {
        const conf = new Configuration()
        const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
        const resp = await api.postUsers(this.user)
        this.$router.go(-1)
      } catch (err) {
        console.log(err)
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
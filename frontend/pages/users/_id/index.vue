<template>
  <div>
    <h1>ユーザー詳細</h1>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>名前</strong></label>
          </v-col>
          <v-col>
            <label>{{ user.name }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>メールアドレス</strong></label>
          </v-col>
          <v-col>
            <label>{{ user.email }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="2">
            <label><strong>ロール</strong></label>
          </v-col>
          <v-col>
            <label>{{ $roleLabel(user.role) }}</label>
          </v-col>
        </v-row>
        <v-row>
          <v-btn @click="edit">編集</v-btn>
          <v-spacer />
          <v-btn @click="del" :disabled="isMe">削除</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
import { DefaultApi, Configuration } from "../../../../api-client"
export default {
  async asyncData({$axios, params}) {
    const conf = new Configuration()
    const api = new DefaultApi(conf, $axios.defaults.baseURL, $axios)
    const resp = await api.getUser(params.id)
    return {
      user: resp.data
    }

  },
  data() {
    // return {
      // isMe: (this.$auth.user.id == this.params.id)
    // }
  },
  computed: {
    isMe() {
      return this.$auth.user.id == this.user.id
    },
  },
  methods: {
    edit() {
      this.$router.push(`/users/${this.user.id}/edit`)
    },
    async del() {
      try {
        const conf = new Configuration()
        const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
        const resp = await api.deleteUser(this.$route.params.id)
        this.$router.go(-1)
      } catch (err) {
        this.error = "削除に失敗しました"
        this.hasError = true
      }
    }
  }
}
</script>
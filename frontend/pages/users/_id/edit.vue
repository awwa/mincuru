<template>
  <div>
    <h1>ユーザー編集</h1>
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
          <v-btn @click="save">保存</v-btn>
          <v-btn @click="back">キャンセル</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script>
import { DefaultApi, Configuration } from '../../../../api-client'
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
    return {
      roles: ["user", "admin"]
    }
  },
  methods: {
    async save() {
      try {
        const conf = new Configuration()
        const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
        const resp = await api.patchUser(this.$route.params.id, this.user)
        this.$router.go(-1)
      } catch (err) {
        console.log(err)
        this.error = "更新に失敗しました"
        this.hasError = true
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>
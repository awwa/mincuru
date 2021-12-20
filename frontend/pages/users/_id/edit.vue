<template>
  <div>
    <h1>ユーザー編集</h1>
    <v-form
      @submit.prevent="save"
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
          <v-btn type="submit">保存</v-btn>
          <v-btn @click="back">キャンセル</v-btn>
        </v-row>
      </v-container>
    </v-form>
  </div>
</template>

<script lang="js">
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
      nameRules: [
        v => !!v || "名前を入力してください",
      ],
      emailRules: [
        v => !!v || "メールアドレスを入力してください",        
      ],
      roleRules: [
        v => !!v || "ロールを入力してください",
      ],
    }
  },
  methods: {
    async save() {
      try {
        if (this.$refs.form.validate()) {
          const conf = new Configuration()
          const api = new DefaultApi(conf, this.$axios.defaults.baseURL, this.$axios)
          const resp = await api.patchUser(this.$route.params.id, this.user)
          this.$router.go(-1)
        }
      } catch (err) {
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
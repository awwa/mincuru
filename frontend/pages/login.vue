<template>
<div>
  <div>
  <h1>ログインユーザ</h1>
  <form @submit.prevent="userLogin">
    <div class="form-group">
      <label for="email">メールアドレス</label>
      <input
        id="email"
        v-model="user.email"
        v-bind:class="{ error: hasError }">
    </div>
    <div class="form-group">
      <label for="password">パスワード</label>
      <input
        id="password"
        type="password"
        v-model="user.password"
        v-bind:class="{ error: hasError }">
    </div>
    <button id="submit" type="submit">ログイン</button>
  </form>
  <div>
    <label id="error" class="error">{{error}}</label>
  </div>
  </div>
</div>
</template>

<script lang="ts">
  import Vue from 'vue'
  export default Vue.extend({
    data() {
      return {
        user: {
          email:"",
          password:""
        },
        error: "",
        hasError: false
      }
    },
    methods: {
      async userLogin(): Promise<void> {
        try {
          await this.$auth.loginWith("local", { data: this.user })
          // TODO 要調査
          // nuxt/authの説明によると、
          // loginWith()で正常応答が返った後、
          // 自動的に"/"にリダイレクトするはずだが、
          // 自動リダイレクトしないので明示的にリダイレクトする
          window.location.href = "/"
        } catch (err) {
          // TODO 画面上にエラーメッセージ
          this.error = "ログインできませんでした。ユーザ名またはパスワードが間違っています。"
          this.hasError = true
        }
      },
    }
  })
</script>
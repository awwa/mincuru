<template>
  <div>
    <h1>ログイン</h1>
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
</template>

<script>
  // NuxtJS 2.8.xのTypeScriptサポートドキュメントによると
  // https://develop365.gitlab.io/nuxtjs-2.8.X-doc/ja/guide/typescript/
  // vue-property-decoratorや vue-class-componentを利用したサンプルの記載がある。
  // これは、JavaScript版のVue.JSやNuxt.JSの実装と大きく異なるため、
  // 今の段階では採用を見送る。
  // Vue.JSやNuxt.JSでドキュメント含めTypeScriptが第一級のサポート言語となった段階で
  // 本格的なTypeScript対応の検討を行う。
  // 現段階では、最低限のTypeScript化を行い、基本的な実装スタイルはJavaScript版に沿う形にする
  export default {
    layout: 'top',
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
      async userLogin() {
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
  }
</script>
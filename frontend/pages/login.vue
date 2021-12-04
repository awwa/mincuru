<template>
<div>
  <div>
  <h1>ログインユーザ</h1>
  <form @submit.prevent="userLogin">
    <div class="form-group">
      <label for="email">Email:</label>
      <input v-model="user.email">
    </div>
    <div class="form-group">
      <label for="password">Password:</label>
      <input type="password" v-model="user.password">
    </div>
    <button type="submit">ログイン</button>
  </form>
  </div>
</div>
</template>

<script lang="ts">
  import jwt_decode from "jwt-decode";
  // import { DefaultApi, LoginRequest, Configuration } from '../../api-client'
  // const conf = new Configuration({basePath: "http://localhost:8080"})
  export default {
    data() {
      return {
        user: {
          email:'',
          password:''
        }
      }
    },
    methods: {
      async userLogin() /*:Promise<void>*/ {
        try {
          // let response = await (this as any).$auth.loginWith('cookie', { data: (this as any).user })
          const response = await this.$auth.loginWith('local', { data: (this as any).user })
          // console.log(response)
          // // (this as any).$auth.setUser(response)
          // const decoded = jwt_decode(response.data.token)
          // const user = {
          //   id: (decoded as any).id,
          //   name: (decoded as any).name,
          //   email: (decoded as any).email,
          //   role: (decoded as any).role,
          // };
          // (this as any).$auth.setUser(decoded)
          // console.log(decoded)
          // console.log((this as any).$auth.user)
          // console.log((this as any).$auth)
          // console.log((this as any).$auth.$state.loggedIn)
          // console.log((this as any).$store.state.auth.loggedIn)
          // const cookies = (this as any).$auth.$storage.getCookies()
          // console.log(cookies)

          // this.$auth.loginWith('local',{
          //   data: this.user
          // })

        // api-clientを使ってログイン（ログイン状態は自前で管理する必要がある）
        //   const api = new DefaultApi(conf)
        //   const loginRequest: LoginRequest = {email: "hoge@example.com", password: "password"}
        //   const response = await api.postUsersLogin(loginRequest)
        //   console.log(response)
          // TODO 要調査。nuxt/authの説明によると、loginWith()で正常応答が返った後、自動的に"/"にリダイレクトするはずだが、今の実装だと自動リダイレクトしないので明示的にリダイレクトさせる
          window.location.href = "/";
        } catch (err) {
          console.log(err)
        }
      },
    }
  }
</script>
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
// import Vue from 'vue'
// import { AuthRequest } from '../../mincuru-api-client'
// // import { formatDate } from '@/utils'
// interface Data {
//   value: AuthRequest,
// }
// export default Vue.extend({
//   data: (): Data => {
//     return {
//       value: {
//         email: '',
//         password: ''
//       }
//     }
//   },
//   computed: {
//     // formatDate: () => formatDate,
//   },
//   // async asyncData({ app }): Promise<Data> {
//   //   const { authResponse } = await app.$mincuruApi.authAdmin()
//   //   return { authResponse }
//   // },
//   methods: {
//     async userLogin(): Promise<void> {
//       const { books } = await this.$bookApi.getBooks()
//       this.books = this.books.concat(books)
//     },
//   },
// })
  //import { DefaultApi, PostUsersLoginRequest, Configuration } from '../../mincuru-api-client'
  import { DefaultApi, LoginRequest, Configuration } from '../../api-client'
  const conf = new Configuration({basePath: "http://localhost:8080"})
  export default {
    data() {
      return {
        user: {
          email:'sabu@example.com',
          password:'password'
        }
      }
    },
    methods: {
      async userLogin() :Promise<void> {
        try {
          const api = new DefaultApi(conf)
          const loginRequest: LoginRequest = {email: "hoge@example.com", password: "password"}
          const response = await api.postUsersLogin(loginRequest)
          console.log("weeeeiiiii")
          console.log(response)
        } catch (err) {
          console.log(err)
        }
      },
    }
  }
</script>
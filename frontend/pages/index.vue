<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="logo py-4 d-flex justify-center">
        <NuxtLogo title="ううう" />
        <VuetifyLogo />
      </v-card>
      <v-card>
        <v-card-title class="headline">
          <h2>ログイン状態:{{ $auth.loggedIn }}</h2>
          <p>{{ $store.state.auth.user }}</p>
          <p>{{ $auth.user.email }}</p>
        </v-card-title>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="js">

</script>
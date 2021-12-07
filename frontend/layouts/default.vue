<template>
  <v-app>
    <!-- 左メインメニュー -->
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <!-- 上部バー -->
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <!-- 左ハンバーガーメニュー -->
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <!-- 左メインメニューの拡幅縮退切り替え（"<",">"） -->
      <v-btn
        icon
        @click.stop="miniVariant = !miniVariant"
      >
        <v-icon>mdi-{{ `chevron-${miniVariant ? 'right' : 'left'}` }}</v-icon>
      </v-btn>
      <!-- アプリタイトル -->
      <v-toolbar-title v-text="title" />
      <!-- スペーサ -->
      <v-spacer />
      <!-- アバターアイコンとユーザーメニュー -->        
      <v-menu offset-y>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            v-bind="attrs"
            v-on="on"
          >
            <v-icon large>mdi-account-circle</v-icon>
          </v-btn>
        </template>
        <v-list dense>
          <v-list-item link>
            <v-list-item-content>
              <v-list-item-title
                @click.stop="userLogout">
                ログアウト
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
    <!-- メインコンテンツ -->
    <v-main>
      <Breadcrumbs :path="this.$route.path" />
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <!-- フッタ -->
    <v-footer
      :absolute="!fixed"
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "LauoutDefault",
  data () {
    return {
      clipped: true,
      drawer: false,
      fixed: true,
      items: [
        {
          icon: 'mdi-apps',
          title: 'ホーム',
          to: '/'
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'ユーザー',
          to: '/users'
        },
      ],
      miniVariant: true,
      right: true,
      rightDrawer: false,
      title: this.$config.appTitle,
    }
  },
  methods: {
    async userLogout() {
      await this.$auth.logout("local")
      window.location.href = "/login"
    }
  }
}
</script>

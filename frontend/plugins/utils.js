// https://nuxtjs.org/ja/docs/directory-structure/plugins
// 管理者ユーザの場合true
const isAdmin = function() {
    return this.$auth.user.role == "admin"
}

export default ({}, inject) => {
    inject('isAdmin', isAdmin)
}
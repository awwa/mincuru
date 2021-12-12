// 管理者ユーザの場合true
const isAdmin = function(role) {
    return this.$auth.user.role == "admin"
}

export default ({app}, inject) => {
    inject('isAdmin', isAdmin)
}
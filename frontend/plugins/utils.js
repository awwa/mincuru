// https://nuxtjs.org/ja/docs/directory-structure/plugins
// 管理者ユーザの場合true
const isAdmin = function() {
    return this.$auth.user.role == "admin"
}

/*
 * ロールの値とラベルの一覧
 */
const roles = [
    {value: "user", label: "ユーザー"},
    {value: "admin", label: "管理者"},
]

//
// ロールの値に対応したラベルを返す
//
const roleLabel = function(role) {
    return roles.find(e => e.value == role).label;
}

export default ({}, inject) => {
    inject('isAdmin', isAdmin)
    inject("roles", roles)
    inject("roleLabel", roleLabel)
}
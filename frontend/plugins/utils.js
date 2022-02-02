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

//
// Dateオブジェクトを年月日フォーマット
//
const formatDate = function(text) {
    const d = new Date(text)
    return `${d.getFullYear()}年${d.getMonth()+1}月${d.getDate()}日`
}

// Dateオブジェクトを年月フォーマット
const formatMonth = function(text) {
    const d = new Date(text)
    return `${d.getFullYear()}年${d.getMonth()+1}月`
}

// 数値を通貨フォーマット
const formatCurrencyJP = function(value) {
    return new Intl.NumberFormat('ja-JP', { style: 'currency', currency: 'JPY' }).format(value)
}


// // テスト向けのexport
// export const plugin = {
//     install(Vue) {
//         Vue.prototype.$formatDate = formatDate
//     }
// };

export default ({app}, inject) => {
    inject('isAdmin', isAdmin)
    inject("roles", roles)
    inject("roleLabel", roleLabel)
    inject("formatDate", formatDate)
    inject("formatMonth", formatMonth)
    inject("formatCurrencyJP", formatCurrencyJP)
}
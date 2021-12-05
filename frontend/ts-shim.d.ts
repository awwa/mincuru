// https://github.com/nuxt/typescript/issues/153
declare module '*.vue' {
    import Vue from 'vue'
    export default Vue
}
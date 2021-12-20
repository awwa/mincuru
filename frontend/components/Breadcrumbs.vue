<template>
  <v-breadcrumbs
    :items="items"
    divider=">"
  ></v-breadcrumbs>
</template>

<script lang="ts">
interface Dictionary {
  [key: string]: string;
}
const DICTIONARY: Dictionary = {
  home: "ホーム",
  account: "アカウント設定",
  users: "ユーザー",
  grades: "グレード",
  edit: "編集",
  add: "追加",
}
type Path = {
  text: string,
  disabled: boolean,
  href: string,
}
export default {
  props: {
    path: {
      type: String,
      required: true,
    },
  },
  computed: {
    items: function(): Array<Path> {
      // 末尾のスラッシュ削除
      const effectivePath: string = (this as any).path.replace(/\/$/, "")
      const dirs:Array<string> = effectivePath.split("/")
      dirs[0] = "home"
      let href:string = ""
      let items:Array<Path> = []
      for(let i = 0; i < dirs.length; i++) {
        // 先頭はホーム
        let text: string = DICTIONARY[dirs[i]] ? DICTIONARY[dirs[i]] : dirs[i]
        if (i == 0) {
          href += "/"
        } else if (i == 1) {
          href += dirs[i]
        } else {
          href += "/" + dirs[i]
        }
        // 末尾のリンクは無効
        const disabled: boolean = (i == dirs.length - 1)
        items.push({
          text: text,
          disabled: disabled,
          href: href
        })
      }
      return items
    }
  }
}
</script>
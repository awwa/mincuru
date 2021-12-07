<template>
  <v-breadcrumbs
    :items="items"
    divider=">"
  ></v-breadcrumbs>
</template>

<script>
const dictionary = {
  home: 'ホーム',
  users: 'ユーザー',
  grades: 'グレード',
  edit: '編集',
  add: '追加',
}
export default {
  props: ['path'],
  computed: {
    items: function() {
      // pathの値が文字列であること
      if (typeof(this.path) != 'string') {
        throw new TypeError('no string')
      }
      // 先頭文字が'/'であること
      if (this.path[0] != '/') {
        throw new SyntaxError('First char is not slash')
      }
      // 末尾のスラッシュ削除
      const effectivePath = this.path.replace(/\/$/, '')
      const dirs = effectivePath.split('/')
      dirs[0] = 'home'
      let text = ''
      let href = ''
      let items = []
      for(let i = 0; i < dirs.length; i++) {
        // 先頭はホーム
        text = dictionary[dirs[i]] ? dictionary[dirs[i]] : dirs[i]
        if (i == 0) {
          href += '/'  
        } else if (i == 1) {
          href += dirs[i]
        } else {
          href += '/' + dirs[i]
        }
        // 末尾のリンクは無効
        const disabled = (i == dirs.length - 1)
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
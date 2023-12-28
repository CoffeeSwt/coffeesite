import { createI18n } from 'vue-i18n'

import { ja } from './lang/ja'
import { zh } from './lang/zh'
import { en } from './lang/en'

const messages = {
  ja,
  en,
  zh,
}

export const i18n = createI18n({
  locale: 'zh', // set locale
  fallbackLocale: 'en', // set fallback locale
  messages, // set locale messages
  legacy: false,
})

export const labelMap = (key: string) => {
  const map = new Map()
  map.set('zh', '中文-简体')
  map.set('ja', '日本語')
  map.set('en', 'English')
  return map.get(key)
}

import { defineStore } from 'pinia'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { labelMap } from '@/i18n'

export const useLanguageStore = defineStore('language', () => {
  const storageKey = 'company-lang-loacl'

  //i18n实例
  const i18n = useI18n()
  //可用语言列表
  const availableLocales = computed(() =>
    i18n.availableLocales.map((item) => {
      return {
        value: item,
        label: labelMap(item),
      }
    })
  )
  //当前语言
  const local = computed(() => {
    return labelMap(i18n.locale.value)
  })
  //设置当前语言
  const setLocal = (value: string) => {
    i18n.locale.value = value
    localStorage.setItem(storageKey, value)
  }
  //语言初始化
  const init = () => {
    const storage = localStorage.getItem(storageKey)
    if (storage) {
      i18n.locale.value = storage
    } else {
      localStorage.setItem(storageKey, i18n.locale.value)
    }
  }

  //t函数
  const t = (s: string) => i18n.t(s)

  return { i18n, availableLocales, local, setLocal, init, t }
})

<template>
    <div full-border class="group/item" rounded-border border-hidden hover:bg-gray-200>
        <div full-border flex p-1 justify-between flex-row-center relative cursor-pointer>
            <div class="i-fluent-mdl2:locale-language" w-4 h-4 shrink-0 />
            <div class="i-dashicons:arrow-down-alt2" shrink-0 w-3 h-3 ml-1 />
        </div>
        <ul p-2 class="group-hover/item:visible group-hover/item:opacity-100" bg-gray-100 invisible opacity-0 shrink-0
            list-none pl-0 absolute m-0 transition-opacity rounded-border border-gray-300>
            <template v-for="locale in lang.availableLocales">
                <li cursor-pointer mb-2 last:mb-0 hover-bg-gray-300 v-if="lang.i18n.locale != locale.value" pl-2
                    @click="lang.setLocal(locale.value)" w-full appearance-none>
                    <span>{{ locale.label }}</span>
                </li>
            </template>
        </ul>
    </div>
</template>

<script setup lang="ts">
import { useLanguageStore } from '@/pinia'
const lang = useLanguageStore()
console.log(lang.i18n.locale)

</script>

<style scoped></style>
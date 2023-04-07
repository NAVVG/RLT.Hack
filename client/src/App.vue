<template>
  <RouterView/>
  <TheHeader
  @openReg="openRegModal"
  @openLog="openLogModal"
  />
  <TheModal
  ref="modalRegRef"
  >
    Регистрация
  </TheModal>
  <TheModal
  ref="modalLogRef"
  >
    Авторизация
  </TheModal>
</template>


<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted, ref } from 'vue'
import TheModal from "@/components/TheModal.vue"
import TheHeader from "@/components/TheHeader.vue";


const modalRegRef = ref<InstanceType<typeof modalRegRef> | null>(null)
const modalLogRef = ref<InstanceType<typeof modalLogRef> | null>(null)

function openRegModal() {
  modalRegRef.value?.changeActiveModal()
}
function openLogModal() {
  modalLogRef.value?.changeActiveModal()
}
</script>









<style lang="scss">
@import '@/assets/reset.scss';
@import '@/assets/main.scss';
</style>

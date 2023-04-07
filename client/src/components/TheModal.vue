<template>
    <div class="popup"
    :class="{active : isActive}"
    @click="changeActiveModal"
    >
        <div class="popup__body"
        @click.stop
        >
            <div class="popup__content">
                <slot></slot>
                <div class="popup__close"
                @click="changeActiveModal"
                >
                    <img src="@/assets/icons/exit.svg">
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    
    import { ref } from 'vue'
    const isActive = ref(false)


    function changeActiveModal() {
        isActive.value = !isActive.value
    }
    defineExpose({
      changeActiveModal
    })
</script>

<style lang="scss" >
.popup {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  overflow-y: auto;
  overflow-x: hidden;
  display: none;
  z-index: 104;
  &.active {
    display: block;
  }
  &::before {
    content: "";
    background: #000000;
    opacity: 0.4;
    position: fixed;
    left: 0px;
    top: 0px;
    width: 100%;
    height: 100%;
  }
  &__body {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 30px 10px;
    min-height: 100%;
  }
  &__content {
    position: relative;
    background: #ebedea;
    border-radius: 15px;
    width: 95%;
    max-width: 300px;
    padding: 10px;
    z-index: 105;
  }
  &__close {
    position: absolute;
    cursor: pointer;
    right: 6px;
    top: 9px;
    &::after {
        content: "";
        position: absolute;
        left: 0;
        top: -5px;
        display: block;
        width: 16px;
        height: 14px;
    }
    & img {
      width: 15px;
      transition: .2s;
      &:hover {
        transform: scale(1.1);
      }
    }
  }
}
</style>
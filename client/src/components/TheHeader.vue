<template>
    <header class="header">
        <div class="header__content container">
            <a href="/">
                <IconLogo class="header__logo logo"/>
            </a>
            <div class="header__nav nav">
                <a href="ecp" class="header__nav-item">
                    Услуги и сервисы
                </a>
                <a href="procederes" class="header__nav-item">
                    Электронная подпись
                </a>
            </div>
            <div class="header__side">
                <div class="header__side-list">
                    <div class="header__side-item">
                        Помощь
                    </div>
                    <div class="header__side-item btn btn_blue"
                    @click="openReg"
                    >
                        Регистрация
                    </div>
                    <div class="header__side-item btn btn_blue"
                    @click="openLog"
                    >
                        Войти
                    </div>
                </div>
            </div>
        </div>
    </header>
</template>
<script setup lang="ts">
    import IconLogo from '@/components/icons/IconLogo.vue'

    import { ref } from 'vue'


    const emit = defineEmits(['openReg'])


    function openLog() {
        emit('openLog',false)
    }
    function openReg() {
        emit('openReg',false)
    }
</script>

<style lang="scss" scoped>
.header {
    padding: 20px 0px;
    border-bottom: 1px solid #969696;
    background: rgb(154 174 211 / 40%);
&__content {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 10px;
}
&__logo {
    max-width: 192px;
}

&__nav {
    display: flex;
    align-items: center;
    gap: 30px;
}

&__nav-item {
}

&__side {

}
&__side-list {
    display: flex;
    align-items: center;
    gap: 15px;
}
&__side-item {
}
}
.container {
}
.logo {
}
.nav {
}
.heeader {

&__side-list {
}
}
.btn {

&_blue {
}
}

</style>
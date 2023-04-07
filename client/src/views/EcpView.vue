<template>
    <div class="ecp-w">
        <div class="ecp">
            <div class="ecp__content container">
                <div class="ecp__text">
                    <div class="ecp__title title">
                        Электронная подпись
                        для любых задач
                    </div>
                    <p class="ecp__lead lead">
                        Росэлторг — доверенное лицо ФНС России.
                        Продление или выпуск новых электронных подписей генеральным директорам, ИП, их сотрудникам и физическим лицам.
                        Получить подпись
                    </p>
                    <div class="ecp__getEcp-btn btn btn_blue">
                        Получить подпись
                    </div>
                </div>
                <div class="ecp__pic">
                    <img src="@/assets/images/ecp.png">
                </div>
            </div>
            <div class="container">
                <EcpItem class="ecpItem_main"
                :ecp-item="mainEcpItem"
                />
            </div>
            <div class="ecp__list container">
                <EcpItem
                v-for="(item,index) of ecpItems" :key="index"
                :ecp-item="item"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import EcpItem from "@/components/EcpItem.vue"
    const mainEcpItem = {
        title:'Физическому лицу и самозанятому',
        subtitle:'Квалифицированная электронная подпись для участия в торгах, работы на госпорталах, Росреестра, банков и вузов',
        price: 1000
    }
    const ecpItems = [
        {title:'Директору и предпринимателю', subtitle:'Квалифицированная электронная подпись на 15 месяцев от ДУЦ ФНС России для работы на госпорталах и решения других бизнес-задач онлайн',price:1500},
        {title:'Сотруднику организации', subtitle:'Квалифицированная электронная подпись для участия в торгах, работы на госпорталах, ЭДО, отчетности и маркировки',price:4500},
        {title:'ЕГАИС', subtitle:'Квалифицированная электронная подпись для подключения и передачи данных об обороте алкогольной продукции в ЕГАИС по требованиям ФСРАР. Сертифицированный токен в комплекте!',price:4190},
        {title:'ГИИС ДМДК', subtitle:'Квалифицированная электронная подпись для регистрации в ГИИС ДМДК и учета драгоценных металлов, камней и изделий из них',price:1500},
        {title:'Для ЭДО', subtitle:'Неквалифицированная электронная подпись для внутреннего и внешнего ЭДО',price:2180},
    ]
</script>

<style lang="scss">
.ecp-w {
    background: #002160;
}
.ecp {
&__content {
    color: #fff;
    display: flex;
    gap: 20px;
    padding-top: 50px;
    padding-bottom: 50px
}

&__text {
}

&__title {
}
&__list {
    display: flex;
    gap: 20px;
    padding-top: 50px;
    flex-wrap: wrap;
}
&__lead {
    margin-top: 40px;
}

&__getEcp-btn {
    margin-top: 40px;
}

&__pic {
}
}
.container {
}
.btn {

&_blue {
}
}
.ecpItem_main {
    margin: 0 auto;
}
</style>

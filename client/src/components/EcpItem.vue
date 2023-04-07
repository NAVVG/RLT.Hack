<template>
    <div class="ecpItem">
        <div class="ecpItem__subtitle">
            {{ props.ecpItem.title }}
        </div>
        <div class="ecpItem__lead lead">
            {{ props.ecpItem.subtitle }}
        </div>
        <div class="ecpItem__footer">
            <div class="ecpItem__price">
            {{ props.ecpItem.price }} ₽
            </div>
            <div class="ecpItem__btn btn btn_blue">
                Оформить онлайн
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    interface ecpItemType {
        title: string,
        subtitle: string,
        price: number
    }
    const props = defineProps<{
        ecpItem: ecpItemType
    }>()
    
</script>

<style lang="scss">
.ecpItem {
    background: #fff;
    border-radius: 5px;
    padding: 40px;
    max-width: 500px;
    flex: 0 0 calc(33.3% - 20px);
    display: flex;
    justify-content: space-between;
    flex-direction: column;
&__subtitle {
    line-height: 1.19;
    color: #000000;
    font-size: 32px;
}
&__footer {
    margin-top: 20px;
}
&__lead {
    margin-top: 20px;
}

&__item {
}
&__price {
    font-size: 24px;
    line-height: 1.17;
}
&__btn {
    margin-top: 20px;
}
}
.lead {
}
.btn {

&_blue {
}
}

</style>
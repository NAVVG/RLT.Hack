<template>
    <div class="procedures__list-item">
        <div class="procedures__item-text">
            <div class="procedures__item-top">
                <div class="procedures__item-lot">
                <span class="procedures__item-status"></span>
                {{ props.procedurItem.lot }}
                </div>
                <div class="procedures__item-section">
                    {{ props.procedurItem.section }}
                </div>
            </div>
            <div class="procedures__item-name">
                {{ props.procedurItem.name }}
            </div>
            <div class="procedures__item-seller">
                {{ props.procedurItem.seller }}
            </div>
        </div>
        <div class="procedures__item-side">
            <div class="procedures__item-form">
                {{ props.procedurItem.form }}
            </div>
            <div class="procedures__item-price">
                {{ props.procedurItem.price }} ₽
            </div>
            <div class="procedures__item-time">
                {{ props.procedurItem.time }}
            </div>
            <div class="procedures__item-place">
                <img src="@/assets/icons/icon-map.svg">
                {{ props.procedurItem.place }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    interface proceduresItemType {
        lot: string,
        section: string,
        name: string,
        seller: string,
        price: number,
        time: string,
        form: string,
        place: string,
    }
    const props = defineProps<{
        procedurItem: proceduresItemType
    }>()
</script>

<style lang="scss" scoped>

.procedures {

&__list-item {
    box-shadow: 0 2px 10px 0 rgb(6 88 148 / 15%);
    padding: 24px 24px 24px 34px;
    position: relative;
    background: #fff;
    display: flex;
    gap: 20px;
    transition: .2s;
    &:hover {
        transform: translate(0px,-10px);
    }
    &:not(:first-child) {
        margin-top: 5px;
    }
}

&__item-text {
    flex: 0 0 70%;
}
&__item-status {
    width: 7px;
    height: 7px;
    background: #2ed618;
    border-radius: 50%;
    display: block;
    position: absolute;
    left: -13px;
    top: 4px;
}
&__item-top {
    display: flex;
    align-items: center;
    gap: 10px;
}

&__item-lot {
    color: #000;
    font-size:16px;
    font-weight: 600;
    position: relative;
}

&__item-section {
    color: #c7cdd2;
}

&__item-name {
    color: #0041CB;
    margin-top: 15px;
}

&__item-seller {
    margin-top: 15px;
}

&__item-side {
}

&__item-form {
    font-size:14px;
}

&__item-price {
    margin-top: 15px;
    font-size:24px;
}

&__item-time {
    font-size:14px;
    margin-top: 15px;
}

&__item-place {
    font-size:15px;
    margin-top: 15px;
    display: flex;
    gap: 4px;
    & img {
        width: 12px;
    }
}
}

</style>
<template>
    <div class="card">
        <div class="card__title">{{title}}</div>
        <div
            v-for="(key, index) in Object.keys(descNames)"
            :key="index"
            class="card__text"
        >
            {{`${descNames[key]}: ${descValues[key]}`}}
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: {
        cardData: {
            type: Object,
            default: () => ({})
        }
    },
    computed: {
        title() {
            const {title = ''} = this.cardData;
            return title;
        },
        descNames() {
            const {descNames = ''} = this.cardData;
            return descNames;
        },
        descValues() {
            const {descValues = ''} = this.cardData;
            return descValues;
        }
    }
}
</script>

<style lang="scss" scoped>
    .card {
        position: relative;
        margin: 0 10px 10px 0;
        padding: 30px;
        border: 1px solid #eaeaea;
        border-radius: 8px;
        width: 500px;
        background: #fff;
    }
    .card:hover {
        box-shadow: 0 4px 28px rgba(59,74,116,.15), 0 4px 10px rgba(59,74,116,.15);
    }
    .card__title {
        font-size: 14px;
    }
    .card__text {
        color: #a7b0bd;
        margin-top: 10px;
    }
</style>
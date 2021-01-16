<template>
    <div class="header-container">
        <h1 class="header-container__title">Мониторинг системы</h1>
        <header-menu />
    </div>
</template>

<script>
import HeaderMenu from './HeaderMenu';
export default {
    name: 'HeaderContainer',
    components: {
        HeaderMenu
    }
}
</script>

<style lang="scss" scoped>
    .header-container {
        display: flex;
        justify-content: space-between;
        padding: 0 25px;
        background: #fff;
        border-radius: 14px;
        align-items: center;
    }
    .header-container__title {
        font-size: 1.5385rem;
        font-weight: 400;
        margin: 16px 10px 16px 0;
    }
</style>
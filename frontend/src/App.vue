<template>
    <div id="app" class="app-container">
        <header-container />
        <router-view/>
    </div>
</template>

<script>
import HeaderContainer from './components/Header/HeaderContainer';
export default {
    components: {HeaderContainer}
}
</script>

<style lang="scss">
    body {
        margin: 0;
        font-family: Segoe UI SemiBold, Arial, sans-serif;
        color: #222;
        line-height: 1.4;
        background: #eaeaea;
    }
    .app-container {
        margin: 30px 30px 0 30px;
    }
    .btn-default {
        color: #000;
        background: #fe1060;
        cursor: pointer;
        pointer-events: auto;
        height: 32px;
        width: -webkit-min-content;
        width: -moz-min-content;
        width: min-content;
        padding: 4px 14px;
        border-radius: 4px;
        -webkit-transition: .3s;
        transition: .3s;
        border: 1px solid transparent;
        outline: none;
    }
</style>
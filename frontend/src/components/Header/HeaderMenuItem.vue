<template>
    <div class="header-menu-item">
        <router-link 
            :to="menuItemRout"
            class="header-menu-item__link"
        >
            {{menuItemName}}
        </router-link>
    </div>
</template>

<script>
export default {
    name: 'HeaderMenuItem',
    props: {
        menuItem: {
            type: Object,
            default: () => ({})
        }
    },
    computed: {
        menuItemName() {
            const {name = ''} = this.menuItem;
            return name;
        },
        menuItemRout() {
            const {rout = ''} = this.menuItem;
            return rout;
        }
    }
}
</script>

<style lang="scss" scoped>
    .router-link-active {
        color: #000;
        background: #fe1060;
        pointer-events: none;
    }
    .header-menu-item {
        margin: 0 10px 0 10px;
    }
    .header-menu-item__link {
        color: #000 !important;
        text-decoration: none;
        padding: 5px;
        border: 1px solid transparent;
        border-radius: 100px;
    }
    .header-menu-item__link:hover {
        color: #fe1060;
        border-color: #fe1060;
    }
</style>
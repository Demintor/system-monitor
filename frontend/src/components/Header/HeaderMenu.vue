<template>
    <div class="header-menu">
        <header-menu-item
            v-for="(menuItem, menuItemIndex) in menu"
            :key="menuItemIndex"
            :menuItem="menuItem"
        />
    </div>
</template>

<script>
import HeaderMenuItem from './HeaderMenuItem';
export default {
    name: 'HeaderMenu',
    components: {
        HeaderMenuItem
    },
    data: () => ({
        menu: [
            {
                name: 'Процессы',
                rout: '/process'
            },
            {
                name: 'Выполнить',
                rout: '/bash'
            },
            {
                name: 'Характеристики системы',
                rout: '/system-info'
            }
        ]
    })
}
</script>

<style lang="scss" scoped>
    .header-menu {
        display: flex;
    }
</style>
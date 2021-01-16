<template>
    <div class="terminal-container">
        <terminal 
            :isRun="true"
            @onClickRun="onClickRun"
        />
        <terminal :inputText="outData"/>
    </div>
</template>

<script>
import Terminal from './Terminal';
export default {
    name: 'TerminalContainer',
    components: {
        Terminal
    },
    props: {
        outData: {
            type: String,
            default: ''
        }
    },
    methods: {
        onClickRun(text = '') {
            this.$emit('onClickRun', text);
        }
    }
}
</script>

<style lang="scss" scoped>
    .terminal-container {
        display: flex;
        justify-content: space-between;
    }
</style>
<template>
    <div>
        <textarea
            v-model="text"
            class="terminal"
        >
        </textarea>
        <div 
            v-if="isRun" 
            class="teminal__actions"
        >
            <button
                class='btn-default'
                @click="onClickRun" 
            >
                Выполнить
            </button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Terminal',
    props: {
        isRun: {
            type: Boolean,
            default: false
        },
        inputText: {
            type: String,
            default: ''
        }
    },
    data: () => ({
        text: ''
    }),
    watch: {
        inputText: function () {
            this.text = this.inputText;
        }
    },
    methods: {
        onClickRun() {
            this.$emit('onClickRun', this.text);
        }
    }
}
</script>

<style lang="scss" scoped>
    .terminal {
        font-size: 12px;
        min-height: 500px;
        width: 850px;
        margin: 0;
        padding: 10px;
        resize: none;
        border-radius: 4px;
        background: #41444e;
        font-family: monospace;
        font-weight: 600;
        color: #fff;
    }
    .teminal__actions {
        display: flex;
        flex-direction: row-reverse;
        margin-top: 10px;
    }
</style>
<template>
    <div class="view-bash">
        <terminal-container
            :outData="outData"
            @onClickRun="onClickRun"
        />
    </div>
</template>

<script>
import ServiceApi from '../services/ServiceApi';
import TerminalContainer from '../components/Terminal/TerminalContainer';
export default {
    name: 'ViewBash',
    components: {TerminalContainer},
    data: () => ({
        outData: ''
    }),
    methods: {
        async postExecuteScript(requestBody = {}, callback = function () {}) {
            try {
                const {data = {}} = await ServiceApi.post(`execute-script`, requestBody);
                callback(data);
            }
            catch(error) {
                global.console.log(error);
                callback({Output: 'Error'});
            }
        },
        onClickRun(text = '') {
            if (text !== '') {
                let requestBody = {script: text};
                this.postExecuteScript(requestBody, (resData) => {
                    const {Output: output = ''} = resData;
                    this.outData = output;
                });
            }
        }
    }
}
</script>

<style lang="scss" scoped>
    .view-bash { 
        margin-top: 30px;
        padding: 20px 25px;
        background: #fff;
        border-radius: 14px;
    }
</style>
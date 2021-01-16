<template>
    <div class="view-process">
        <div class="view-process__actions">
            <button 
                class='btn-default'
                @click="getData"
            >
                Обновить
            </button>
        </div>
        <table-container 
            :columns="tableColumns"
            :rows="tableRows"
            @onClickRow="onClickRow"
        />
    </div>
</template>

<script>
import ServiceApi from '../services/ServiceApi';
import TableContainer from '../components/Table/TableContainer';
export default {
    name: 'ViewProcess',
    components: {TableContainer},
    data: () => ({
        tableRows: [],
        tableColumns: ['PID', 'PPID', 'UID', 'C', 'PSR', 'RSS', 'STIME', 'SZ', 'TIME', 'TTY', 'CMD']
    }),
    methods: {
        async getData() {
            try {
                const {data = {}} = await ServiceApi.get('get-process-table');
                this.tableRows = data;
            }
            catch(error) {
                global.console.log(error);
            }
        },
        async postKill(pid, callback = function () {}) {
            try {
                await ServiceApi.post(`kill-process/${pid}`);
                callback();
            }
            catch(error) {
                global.console.log(error);
            }
        },
        onClickRow(row = {}, rowIndex = -1) {
            let {PID: pid = -1} = row;
            this.postKill(pid, () => {
                this.tableRows.splice(rowIndex, 1);
            });
        }
    },
    mounted() {
        this.getData();
    }
}
</script>

<style lang="scss" scoped>
    .view-process {
        margin-top: 30px;
        padding: 20px 25px;
        background: #fff;
        border-radius: 14px;
    }
    .view-process__actions {
        display: flex;
        flex-direction: row-reverse;
        margin-bottom: 10px;
    }
</style>
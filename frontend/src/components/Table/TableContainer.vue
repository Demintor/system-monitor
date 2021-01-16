<template>
    <div class="table-container">
        <table>
            <thead>
                <tr>
                    <th>
                    </th>
                    <th
                        v-for="(column, columnIndex) in columns"
                        :key="columnIndex"
                    >
                        {{column}}
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr
                    v-for="(row, rowIndex) in rows"
                    :key="rowIndex"
                >
                    <td>
                        <div 
                            class="table-container__td_kill"
                            @click="onClickRow(row, rowIndex)"
                        >
                            x
                        </div>
                    </td>
                    <td
                        v-for="(column, columnIndex) in columns"
                        :key="columnIndex"
                    >
                        {{row[column]}}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    name: 'TableContainer',
    props: {
        columns: {
            type: Array,
            default: () => ([])
        },
        rows: {
            type: Array,
            default: () => ([])
        }
    },
    methods: {
        onClickRow(row = {}, rowIndex = -1) {
            this.$emit('onClickRow', row, rowIndex);
        }
    }
}
</script>

<style lang="scss" scoped>
    .table-container {
        position: relative;
        width: 100%;
        max-height: 550px;
        border-radius: 10px;
        overflow: auto;
        border: 3px solid #f5f7f9;
    }
    table {
        width: 100%;
        color: #818284;
        border-collapse: collapse;
        border-spacing: 0;
    }
    thead {
        border-bottom: 3px solid #f5f7f9;
    }
    th {
        padding: 16px 14px;
        color: #000;
        text-align: left;
    }
    td {
        padding: 16px 14px;
        border-bottom: 1px solid #f5f7f9;
        text-align: left;
    }
    table tbody tr:hover {
        background: #f5f7f9!important;
    }
    .table-container__td_kill {
        cursor: pointer;
        color: #000;
    }
</style>
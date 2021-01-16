<template>
    <div class="view-system-info">
        <card
            v-for="(cardData, cardIndex) in systemInfoCards"
            :key="cardIndex"
            :cardData="cardData"
        />
    </div>
</template>

<script>
import ServiceApi from '../services/ServiceApi';
import Card from '../components/Card/Card';
export default {
    name: 'ViewSystemInfo',
    components: {Card},
    data: () => ({
        systemInfoData: {}
    }),
    computed: {
        systemInfoCards() {
            return [this.systemInfoTotal, this.systemInfoCpu, this.systemInfoMemory];
        },
        systemInfoTotal() {
            const {CountUser = 0, RunningTime = '', CurrentTime = ''} = this.systemInfoData;
            const systemInfoTotal = {
                title: 'Общая информация',
                descNames: {
                    CountUser: 'Колличество пользователей',
                    RunningTime: 'Время работы',
                    CurrentTime: 'Текущее время'
                },
                descValues: {
                    CountUser,
                    RunningTime,
                    CurrentTime
                }
            };
            return systemInfoTotal;
        },
        systemInfoCpu() {
            const {CPUs = 0, CpuModelName = '', LoadLastOneMin = '', LoadLastFiveMin = '', LoadLastFifteenMin = ''} = this.systemInfoData;
            const systemInfoCpu = {
                title: 'Процессор',
                descNames: {
                    CPUs: 'Колличество ядер',
                    CpuModelName: 'Модель',
                    LoadLastOneMin: 'Cредняя загрузка процессора за последнюю минуту',
                    LoadLastFiveMin: 'Cредняя загрузка процессора за последние пять минут',
                    LoadLastFifteenMin: 'Cредняя загрузка процессора за последние пятнадцать минут'
                },
                descValues: {
                    CPUs,
                    CpuModelName,
                    LoadLastOneMin,
                    LoadLastFiveMin,
                    LoadLastFifteenMin
                }
            }
            return systemInfoCpu;
        },
        systemInfoMemory() {
            const {MemTotal = '', MemAvailable = ''} = this.systemInfoData;
            const systemInfoMemory = {
                title: 'Память',
                descNames: {
                    MemTotal: 'Всего kB',
                    MemAvailable: 'Свободно kB',
                },
                descValues: {
                    MemTotal,
                    MemAvailable,
                }
            };
            return systemInfoMemory;
        },
    },
    mounted() {
        this.getData();
    },
    methods: {
        async getData() {
            try {
                const {data = {}} = await ServiceApi.get('get-system-info');
                debugger;
                this.systemInfoData = data;
            }
            catch(error) {
                global.console.log(error);
            }
        },
    },
}
</script>

<style lang="scss" scoped>
    .view-system-info {
        margin-top: 30px;
        padding: 20px 25px;
        display: flex;
        justify-content: space-between;
    }
</style>
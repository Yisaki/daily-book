<template>
    <div>
        <SaveRecord
        :lastRecordDto="periodDto"
        :formDto="periodForm"
        :typeMap="periodTypeMap"
        :saveRecordFunc="saveRecordReq"
        :routeUrl="'/periodList'"
        
        ></SaveRecord>
    </div>
</template>

<script>
import SaveRecord from '@/components/SaveRecord'

import { getLastRecord, saveRecord, PeriodDto } from '@/network/period'

export default {
    name: 'Period',
    data() {
        return {
            //展示的对象
            periodDto: new PeriodDto(),
            //提交用的对象
            periodForm: new PeriodDto(),

            periodTypeMap:this.$store.state.periodTypeMap
        }
    },
    components:{
        SaveRecord
    },
    created() {
        this.getLastRecordReq()
    },
    methods: {
        saveRecordReq() {
            saveRecord(this.periodForm)
                .then((res) => {
                    const data = res.data
                    if (data.success) {
                        console.log(data)
                        this.getLastRecordReq()
                    }
                })
                .catch((err) => {})
        },
        getLastRecordReq() {
            getLastRecord()
                .then((res) => {
                    const data = res.data
                    if (data.success) {
                        this.periodDto = data.obj
                    }
                })
                .catch((err) => {})
        },

    },
    computed: {

    },
}
</script>

<style>
#last-rec {
    color: #909399;
    font-size: 8px;
}
.daily-row {
    height: 80px;
    padding-top: 20px;
}
</style>
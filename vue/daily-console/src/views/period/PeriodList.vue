<template>
    <div>
        <RecordList
            :typeMap="periodTypeMap"
            :recordTable="periodTable"
            :returnUrl="'/period'"
        ></RecordList>
    </div>
</template>

<script>
import RecordList from '@/components/RecordList'
import { listRecord, PeriodDto } from '@/network/period'
import { PageDto } from '@/common/utils'

export default {
    name: 'Period',
    data() {
        return {
            //展示的对象
            periodTable: [],
            pageDto: new PageDto(),
            //typeMap
            periodTypeMap: this.$store.state.periodTypeMap,
        }
    },
    components: {
        RecordList,
    },
    created() {
        this.listRecordReq()
    },
    methods: {
        listRecordReq() {
            listRecord(this.pageDto)
                .then((res) => {
                    const data = res.data
                    if (data.success) {
                        this.periodTable = data.obj
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
</style>
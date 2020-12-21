<template>
    <div>
        <RecordList
            :typeMap="petTypeMap"
            :recordTable="petTable"
            :returnUrl="'/pet'"
        ></RecordList>
    </div>
</template>

<script>
import { listRecord, PetDto } from '@/network/pet'
import { PageDto } from '@/common/utils'
import RecordList from '@/components/RecordList'

export default {
    name: 'PetList',
    data() {
        return {
            //展示的对象
            petTable: [],
            //分页dto
            pageDto: new PageDto(),
            //typeMap
            petTypeMap: this.$store.state.petTypeMap,
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
                        this.petTable = data.obj
                    }
                })
                .catch((err) => {})
        },
    },
    computed: {},
}
</script>

<style>
</style>
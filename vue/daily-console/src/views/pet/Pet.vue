<template>
    <div>
        <SaveRecord
        :lastRecordDto="petDto"
        :formDto="petForm"
        :typeMap="petTypeMap"
        :saveRecordFunc="saveRecordReq"
        :routeUrl="'/petList'"
        
        ></SaveRecord>
    </div>
</template>

<script>
import SaveRecord from '@/components/SaveRecord'
import { getLastRecord, saveRecord, PetDto } from '@/network/pet'

export default {
    name: 'Pet',
    components:{
        SaveRecord
    },
    data() {
        return {
            //展示的对象
            petDto: new PetDto(),
            //提交用的对象
            petForm: new PetDto(),
            //type与文案的map
            petTypeMap:this.$store.state.petTypeMap
        }
    },
    created() {
        this.getLastRecordReq()


    },
    methods: {
        saveRecordReq() {
            saveRecord(this.petForm)
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
                        this.petDto = data.obj
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
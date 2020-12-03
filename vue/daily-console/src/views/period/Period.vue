<template>
    <div>
        <el-row>
             <el-col :span="4" :offset="19">
            <el-button type="primary" plain @click="routeToList">===</el-button>
             </el-col>
        </el-row>
        <!-- 上次记录-->
        <el-row type="flex" justify="center" class="daily-shadow daily-row">
            <el-col :span="12">
                <div>
                    <div id="last-rec">上次记录时间</div>
                    <div>{{ periodDto.happenTime }}</div>
                    <div><span>类型:</span>{{ typeStr(periodDto.type) }}</div>
                </div>
            </el-col>
        </el-row>
        <!-- 记录-->
        <el-row type="flex" justify="center" class="daily-shadow daily-row">
            <el-col :span="12">
                <el-radio-group v-model="periodForm.type" size="medium">
                    <el-radio-button :label="0">开始</el-radio-button>
                    <el-radio-button :label="1">结束</el-radio-button>
                </el-radio-group>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center" class="daily-shadow daily-row" style="height:200px">
            
                <el-button type="primary" plain @click="saveRecordReq" style="width:60%;height:130px"
                    >记录</el-button
                >
           
        </el-row>
    </div>
</template>

<script>
import { getLastRecord, saveRecord, PeriodDto } from '@/network/period'

export default {
    name: 'Period',
    data() {
        return {
            //展示的对象
            periodDto: new PeriodDto(),
            //提交用的对象
            periodForm: new PeriodDto(),
        }
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
        routeToList() {
            this.$router.push('/periodList')
        },
    },
    computed: {
        typeStr() {
            return function (type) {
                return this.$store.state.periodTypeMap[type]
            }
        },
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
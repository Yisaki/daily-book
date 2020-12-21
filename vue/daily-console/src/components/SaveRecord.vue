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
                    <div>{{ lastRecordDto.happenTime }}</div>
                    <div>
                        <span>类型:</span>{{ typeFormat(lastRecordDto.type) }}
                    </div>
                </div>
            </el-col>
        </el-row>
        <!-- 记录-->
        <el-row type="flex" justify="center" class="daily-shadow daily-row">
            <el-col :span="20">
                <el-radio-group v-model="formDto.type" size="medium">
                    <el-radio-button
                        :key="k"
                        :label="parseInt(k)"
                        v-for="(v,k) in typeMap"
                        >{{v}}</el-radio-button
                    >
                </el-radio-group>
            </el-col>
        </el-row>
        <el-row
            type="flex"
            justify="center"
            class="daily-shadow daily-row"
            style="height: 200px"
        >
            <el-button type="primary" plain style="width: 60%; height: 130px"
                @click="saveRecordFunc">记录</el-button
            >
        </el-row>
    </div>
</template>

<script>
export default {
    name: 'SaveRecord',
    props: {
        //展示上次记录用dto
        lastRecordDto: Object,

        //提交本次记录用formDto
        formDto: Object,
        //展示type选项
        typeMap: Object,

        //保存记录
        saveRecordFunc: Function,
        //跳转url
        routeUrl:String
    },
    methods: {
        typeFormat(type){
            return this.typeMap[type]
        },
        routeToList() {
            this.$router.push(this.routeUrl)
        },
    },
    created() {
        /**for (var p in this.typeMap) {
            console.log("comp")
            console.log(p)
            console.log(this.typeMap[p])
        }**/
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#last-rec {
    color: #909399;
    font-size: 8px;
}
.daily-row {
    height: 80px;
    padding-top: 20px;
}
</style>

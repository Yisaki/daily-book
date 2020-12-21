<template>
    <div>
        <el-row>
            <el-button type="primary" plain @click="routeToSave"
                >返回</el-button
            >
        </el-row>
        <!-- 记录-->
        <el-row>
            <el-table :data="recordTable" style="width: 100%">
                <el-table-column prop="happenTime" label="日期" width="180">
                </el-table-column>
                <el-table-column
                    prop="type"
                    label="类型"
                    width="180"
                    :formatter="formatter"
                >
                </el-table-column>
                <el-table-column fixed="right" label="操作" width="60">
                    <template slot-scope="scope">
                        <el-button
                            @click.native.prevent="
                                deleteRow(scope.$index, recordTable)
                            "
                            type="text"
                            size="small"
                        >
                            移除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-row>
    </div>
</template>

<script>

export default {
    name: 'RecordList',
    props: {
        //type map
        typeMap:Object,

        //分页数据
        recordTable: Array,
        //跳转url
        returnUrl:String
    },
    created() {
       
    },
    methods: {

        routeToSave() {
            this.$router.push(this.returnUrl)
        },

        deleteRow(index, rows) {
            rows.splice(index, 1)
        },
        formatter(row, column) {
            return this.typeMap[row.type]
        },
    },
    computed: {

    },
}
</script>

<style>
</style>
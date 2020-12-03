<template>
    <div>
        <el-row>
            <el-button type="primary" plain @click="routeToSave"
                >返回</el-button
            >
        </el-row>
        <!-- 上次记录-->
        <el-row>
            <el-table :data="periodTable" style="width: 100%">
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
                                deleteRow(scope.$index, periodTable)
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
import { listRecord, PeriodDto } from '@/network/period'
import { PageDto } from '@/common/utils'

export default {
    name: 'Period',
    data() {
        return {
            //展示的对象
            periodTable: [],
            pageDto: new PageDto(),
        }
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
        routeToSave() {
            this.$router.push('/period')
        },

        deleteRow(index, rows) {
            rows.splice(index, 1)
        },
        formatter(row, column) {
            return this.$store.state.periodTypeMap[row.type]
        },
    },
    computed: {

    },
}
</script>

<style>
</style>
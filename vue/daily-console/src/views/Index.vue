<template>
    <div>
        <el-row type="flex" justify="center" style="margin-top:40%;height:50px">
            <el-col :span="12">
                <el-input v-model="password" placeholder="口令"></el-input>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center" style="margin-top:8px;height:50px">
            <el-col :span="8">
                <el-button type="primary" round @click="checkPassword">确定</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<script>


export default {
    name: 'Index',
    components:{

    },
    data() {
        return {
            password:''
        }
    },
    created() {
        let isLogin=localStorage.getItem('isLogin')
        if(isLogin){
           this.$router.push('/period') 
        }
    },
    methods: {

        checkPassword(){
            if(this.password==undefined&&this.password==''){
                return
            }
            if(this.password=='sb250'){
                localStorage.setItem('isLogin',true)
                this.$router.push('/period')      
            }
        }

    },
    computed: {

    },
}
</script>

<style>

</style>
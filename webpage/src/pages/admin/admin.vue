<template lang="pug">
    div 权限分配
</template>

<style lang="less" scoped>

</style>

<script>
import {Indicator, Toast} from 'mint-ui'
export default {
    data(){
        return {
            myLevel:0,
            list:[]
        }
    },
    async created(){
        Indicator.open("获取管理员列表")
        let res = await this.$api.get('/admin/admin')
        Indicator.close()
        if(res.data.success){
            this.myLevel = res.data.result.myLevel
            this.list = res.data.result.list
        } else {
            Toast("数据获取失败")
        }
    }
}
</script>


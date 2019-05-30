<template lang="pug">
    div#admin.container 
        .title 授权管理
        .subtitle 设置/撤销用户的管理员权限，管理员可以调整部门设置、人员从属关系。
        .white-space
        Field(label="一卡通号" placeholder="请输入新管理员一卡通号" type="number" v-model="cardnum")
        .line
        .white-space
        MtButton(type="primary" @click="addAdmin") 添加管理员
        .white-space
        div(v-for="item in list" :key="item.cardnum")
            .line
            .item
                .cardnum {{item.cardnum}} - {{item.name}}
                MtButton(type="danger" size="small" :disabled="item.level >= myLevel" @click="fireAdmin(item.cardnum)") 撤销
        .line
            

</template>

<style lang="less" scoped>
    #admin{
        display: flex;
        flex-direction: column;
        align-items: stretch;
        .field{
            border: none;
        }
        .title{
            text-align: center;
            font-size: 30px;
            font-weight: light;
            color:#333;
            margin:30px 0 10px 0;
        }
        .subtitle{
            text-align: left;
            font-size: 16px;
            color:#888;
            text-indent: 2em;
        }
        .item{
            display: flex;
            align-items: center;
            padding: 10px 10px;
            .cardnum{
                flex-grow: 1;
            }
        }
    }
</style>

<script>
import {Indicator, Toast, Field, Button} from 'mint-ui'
export default {
    components:{
        Field, MtButton:Button
    },
    data(){
        return {
            myLevel:0,
            list:[],
            cardnum:''
        }
    },
    methods:{
        async addAdmin(){
            if(this.cardnum.length === 9){
                Indicator.open("设置管理员")
                let res = await this.$api.post('/admin/admin',{cardnum:this.cardnum})
                Indicator.close()
                if(res.data.success){
                    this.updateList()
                } else {
                    Toast(res.data.reason)
                }
            }  
        },
        async updateList(){
            Indicator.open("获取管理员列表")
            let res = await this.$api.get('/admin/admin')
            Indicator.close()
            if(res.data.success){
                this.myLevel = res.data.result.myLevel
                this.list = res.data.result.list
            } else {
                Toast("数据获取失败")
            }
        },
        async fireAdmin(cardnum){
            let res = await this.$api.delete(`/admin/admin?cardnum=${cardnum}`)
            this.updateList()
        }
    },
    async created(){
        this.updateList()
    }
}
</script>


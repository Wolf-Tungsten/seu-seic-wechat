<template>
    <div v-if="needUserInfo" class="container">
    <p class="hint">为了给您提供及时有效的服务，请补全您的个人信息：</p>
    <field label="姓名" placeholder="请填写真实姓名" v-model="name"></field>
    <field label="一卡通号" placeholder="提交后无法修改" type="number" v-model="cardnum"></field>
    <field label="联系电话" placeholder="故障处理人员能够及时联系您" type="number" v-model="phoneNumber"></field>
    <div class="white-space"></div>
    <mt-button type="primary" @click="updateUserInfo" :disabled="loading">提交并继续</mt-button>
    </div>
</template>

<style lang="less" scoped>
    .hint{
        padding:10px;
    }
    .container{
        display: flex;
        flex-direction: column;
    }
</style>

<script>
import {Indicator, Toast} from 'mint-ui';
import {Field, Button} from 'mint-ui'

export default {
    name:'wxlogin',
    data(){
        return {
        needUserInfo:false,
        name:'',
        cardnum:'',
        phoneNumber:'',
        loading:false
    }},
    components:{
        field:Field,
        'mt-button':Button
    },
    methods:{
        async updateUserInfo(){
            if(!this.name || !this.cardnum ){
                Toast({
                    message: '请填写完整',
                    position: 'bottom',
                    duration: 5000
                })
                this.loading = false
                return
            }
            let res = await this.$api.post('/login', {name:this.name, cardnum:this.cardnum, phoneNumber:this.phoneNumber})
            if(res.data.success){
                this.$router.replace('/'+this.$route.params.page)
            }
        }
    },
    async created() {
        Indicator.open("微信登录")
        let code = this.$route.query.code
        let res = await this.$api.get(`/login?code=${code}`)
        Indicator.close()
        if(res.data.success){
            this.$setSessionToken(res.data.result.sessionToken)
            this.needUserInfo = res.data.result.needUserInfo
            if(!this.needUserInfo){
                this.$router.replace('/'+this.$route.params.page)
            }
        } else {
            this.$router.replace('/error')
        }
    }
}
</script>


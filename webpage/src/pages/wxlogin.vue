<template>
    <div>微信登录中</div>
</template>

<style lang="less" scoped>

</style>

<script>
import {Indicator} from 'mint-ui';

export default {
    name:'wxlogin',
    created() {
        Indicator.open('核对身份')
        console.log(this.$route)
        //this.$router.replace(`/${this.$route.params.page}`)
    }
}
</script>


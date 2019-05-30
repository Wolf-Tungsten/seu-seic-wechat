<template lang="pug">
    .container 
        img(src='../assets/error.png')
        p 出现错误
</template>

<style lang="less" scoped>
    .container{
        display: flex;
        flex-direction: column;
        align-items: center;
        padding-top: 100px;
    }

    img{
        width:64px;
    }

    p{
        color:#888;
    }
</style>

<script>

export default {
}
</script>


<template>
    <div class="main" style="height: 100%">
        <iframe width="100%" height="100%" :src="url+'view/form/#/'" frameborder="0"></iframe>
    </div>
</template>

<script>

    export default {
        name: "form",
        data() {
            return {
                url: ''
            }
        },
        created() {
            switch (process.env.NODE_ENV) {
                case 'development'://本地地址
                    this.url = this.$Config.localUrl //这里是本地的请求url
                    break
                case 'alpha': //测试地址
                    this.url = this.$Config.testUrl
                    break
                case 'production'://生产地址
                    this.url = this.$Config.url //生产环境url
                    break
            }
        }

    }
</script>

<style scoped>

</style>
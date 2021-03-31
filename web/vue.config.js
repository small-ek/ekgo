const Timestamp = new Date().getTime()
module.exports = {
    /* 部署生产环境和开发环境下的URL：可对当前环境进行区分，baseUrl 从 Vue CLI 3.3 起已弃用，要使用publicPath */
    /* baseUrl: process.env.NODE_ENV === 'production' ? './' : '/' */
    publicPath: '/',
    // webpack配置11
    // see https://github.com/vuejs/vue-cli/blob/dev/docs/webpack.md
    chainWebpack: () => {
        // config.optimization.minimize(true);
        // config.optimization.splitChunks({
        // 	chunks: 'all'
        // })
    },
    /* 输出文件目录：在npm run build时，生成文件的目录名称 */
    outputDir: 'dist',
    /* 放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录 */
    assetsDir: "assets",
    /* 是否在构建生产包时生成 sourceMap 文件，false将提高构建速度 */
    productionSourceMap: false,
    /* 默认情况下，生成的静态资源在它们的文件名中包含了 hash 以便更好的控制缓存，你可以通过将这个选项设为 false 来关闭文件名哈希。(false的时候就是让原来的文件名不改变) */
    filenameHashing: false,
    /* 代码保存时进行eslint检测 */
    lintOnSave: true,
    /* webpack-dev-server 相关配置 */
    css: {
        // 是否使用css分离插件 ExtractTextPlugin
        extract: true,
        // 开启 CSS source maps?
        requireModuleExtension: true,
        sourceMap: true,
        // css预设器配置项
        loaderOptions: {
            less: {
                lessOptions: {
                    javascriptEnabled: true,
                    modifyVars: {
                        'vab-color-blue': '#1890ff',
                        'vab-margin': '20px',
                        'vab-padding': '20px',
                        'vab-header-height': '65px',
                    },
                },
            },
        },
        // 启用 CSS modules for all css / pre-processor files.
        modules: false
    },
    devServer: {
        /* 自动打开浏览器 */
        open: true,
        /* 设置为0.0.0.0则所有的地址均能访问 */
        host: '0.0.0.0',
        port: 8080,
        https: false,
        hotOnly: false,
        /* 使用代理 */
        // proxy: {
        // 	'/api': {
        // 		/* 目标代理服务器地址 */
        // 		target: '',
        // 		/* 允许跨域 */
        // 		changeOrigin: true,
        // 	},
        // },
    },
    chainWebpack: config => {
        if (process.env.NODE_ENV === 'production') {
            // 给js和css配置版本号
            config.output.filename('js/[name].' + Timestamp + '.js').end();
            config.output.chunkFilename('js/[name].' + Timestamp + '.js').end();
            config.plugin('extract-css').tap(() => [{
                filename: `css/[name].${Timestamp}.css`,
                chunkFilename: `css/[name].${Timestamp}.css`
            }])
        }
        //最小化代码测试时候会卡顿
        //config.optimization.minimize(true);
        //分割代码
        // config.optimization.splitChunks({
        //     chunks: 'all'
        // });
    },
    // 修改webpack的配置
    //安装 npm install compression-webpack-plugin --save-dev
    configureWebpack: config => {

    }
}


const CompressionWebpackPlugin = require('compression-webpack-plugin');
const productionGzipExtensions = ['js', 'css'];

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
		sourceMap: false,
		// css预设器配置项
		loaderOptions: {},
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
		//最小化代码
		config.optimization.minimize(true);
		//分割代码
		config.optimization.splitChunks({
			chunks: 'all'
		});
	},
	// 修改webpack的配置
	//安装 npm install compression-webpack-plugin --save-dev
	configureWebpack: config => {
		config.plugins.push(new CompressionWebpackPlugin({
			algorithm: 'gzip',
			test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
			threshold: 0, // 只处理比这个值大的资源，按字节算
			minRatio: 0.8, //只有压缩率比这个值小的文件才会被处理，压缩率=压缩大小/原始大小，如果压缩后和原始文件大小没有太大区别，就不用压缩
			deleteOriginalAssets: false //是否删除原文件，最好不删除，服务器会自动优先返回同名的.gzip资源，如果找不到还可以拿原始文件
		}))
		config.externals = {
			'vue': 'Vue',
			"vue-router": "VueRouter",
			'vuex': 'Vuex',
			"element-ui": "ELEMENT",
			'axios': 'axios',
			'nprogress': 'NProgress',
			'vue-i18n': 'VueI18n',
			'crypto-js': 'CryptoJS'
		}
	}
}

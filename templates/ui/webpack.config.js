const path = require("path")
const HtmlWebpackPlugin = require("html-webpack-plugin")

module.exports = {
    devtool: "inline-source-map",// debug 调试定位到源代码
    entry: {
        // 多页面，多入口
        login: './src/pages/login/login.ts',
        home: './src/pages/home/home.ts',
		append: './src/pages/append/append.ts',
    },
    output: {
        // 输出文件名配置
        filename: '[name].[contenthash].js',
        path: path.resolve(__dirname, 'dist'),
        // 每次构建清除旧的（上次构建）文件
		clean: true,
    },
    resolve: {
        extensions: ['.ts', '.js'],
    },
    plugins: [
        /*
            生成 html 插件配置：
            template： 指定模版生成 html
            filename： 输出路径及文件名
            chunks： 生成的 html 使用那个 js 文件，对应上面 entry 配置的key
        */ 
        new HtmlWebpackPlugin({
			template: './src/pages/login/login.html',
			filename: 'pages/login.html',
			chunks: ['login'],
		}),
		new HtmlWebpackPlugin({
			template: './src/pages/home/home.html',
			filename: 'pages/home.html',
			chunks: ['home'],
        }),
        new HtmlWebpackPlugin({
			template: './src/pages/append/append.html',
			filename: 'pages/append.html',
			chunks: ['append'],
		}),
    ],
    module: {
        rules: [
			{
				test: /\.ts$/,
				use: "ts-loader",
				include: [path.resolve(__dirname, 'src')],
			},
			{
				test: /\.(css|less)$/,
				use: [
					'style-loader',
					'css-loader',
					{
                        // Css 兼容配置, 另外需要在 package.json 添加额外配置
                        /**
                            "browserslist":[
                                "> 1%", --> 市场份额大于 1%
                                "last 2 version"  --> 支持每个浏览器最新的两个版本
                            ]
                         */

						loader: 'postcss-loader',
						options: {
							postcssOptions: {
                                // 添加浏览器前缀配置：-webkit-、-moz-、-ms- -o-
								plugins: [require('postcss-preset-env'), require('autoprefixer')],
							},
						},
                    },
                    "less-loader"
				],
			},
			{
                test: /\.(png|jpg|jpeg|gif|svg)$/i,
                // webpack5 新出配置，将图片资源打包，不需要在配置 url-loader、file-loader...
                type: "asset/resource",
                // 打包后资源存放路径
                generator:{
                    filename: "images/[contenthash].[ext]"
                }
            }
		]
    },
    devServer:{
        // devServer 根目录，意思就是在那个目录下启动server，类似 http-server / live-server
        static: path.resolve(__dirname, "./dist"),
        port: 3000,
        open: true,
        hot: true,  // 热替换，新增、删除模块局部更新，非常棒的功能，比如 form 表单填了好多数据，突然脑回路改了一下js或css，再回来看，填的数据没了，哭～，热替换就解决此类问题
        liveReload: true, // 热更新，自动给刷新，但是不会影响hot
        compress: true, // 采用 gzip 压缩，提高传输速度
    }
}
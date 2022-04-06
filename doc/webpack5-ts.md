# 使用 webpack5 和 typescript 构建 Millionare 记账项目前端部分

<span style="font-weight: bold;color:#80c8f5;">技术栈：</span>
**使用 webpack5 和 ts 构建一个多页面的前端项目，直接采用 ts 开发项目，不采用任何 JS 库或者框架，也不是任何 CSS 库**

### 构建步骤
1. 在 millionare 目下 mkdir -p templates/ui && cd templates/ui
1. tsc --init
``` json
// 修改 module 值
{
  "compilerOptions": {
    //   ...
    "module": "es2015",
    // ....
  }
}

```
1. npm init -y 
1. npm i webpack webpack-cli ts-loader 
1. npm i typescript -D 
1. touch webpack.config.js
1. mkdir -p src/assets src/pages src/styles 
1. cd src/pages && mkdir home
1. cd home && touch home.html home.ts home.less 
1. cd .. && mkdir login && cd login && touch login.html login.ts login.less
- <span style="font-weight: bold;color:#80c8f5;">说明：</span>

**一次性安装配置所有插件库** <br>
1. npm i style-loader css-loader less less-loader -D
1. npm i -D postcss postcss-loader postcss-preset-env autoprefixer
1. npm i html-webpack-plugin -D
1. npm i webpack-dev-server -D
- <span style="font-weight: bold;color:#80c8f5;">说明：</span>

**在这里不使用babel和相关库，因为 tsconfig.json 配置了`{"target": "es5"}`，最终 ts 代码编译成 ECMAScript es5 版本。**

**紧接着是编辑 webpack.config.js**

``` js
const path = require("path")
const HtmlWebpackPlugin = require("html-webpack-plugin")

module.exports = {
    devtool: "source-map", // debug 调试定位到源代码
    entry: {
        // 多页面，多入口
        login: './src/pages/login/login.ts',
		home: './src/pages/home/home.ts',
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
```

**然后 package.json 修改 script，添加browserslist**
``` json
 "scripts": {
    "dev": "webpack-dev-server --mode=development --open",
    "build": "webpack --mode=production"
  },
"browserslist":[
    "> 1%",
    "last 2 version"
  ]
```
**最后就是编写代码测试配置是否成功，此处就不在做记录** <br>

<span style="font-weight: bold;color:#80c8f5;">
配置的完整代码可查看代码分支：01-webpack5-config
</span><br>

**那么如何验证配置是否成功呢？主要关注以下几点：**
- 执行 npm run dev 查看页面内容是否完整，图片，样式是否生效
    - 完整浏览器地址：http://localhost:3000/pages/login.html
- 查看 css 的 display: flex 是否加了前缀
- 检查 network 是否加了压缩，查找 "Content-Encoding: gzip" 字眼
---
- 在执行 npm run build 
- 检查打包是否成功，又没dist目录，dist 下面是否有images，pages，login.xxx.js 文件夹和文件等
- 检查 login.xxx.js 的 async await 是否被编译了

### 初始化 git 仓库
- 在 millionare 目录下 git init .
- touch .gitignore
``` .gitignore
node_modules
.DS_Store
```
- 防止 git 监听 node_modules 目录先： git add .gitignore
``` 

```
- git add .
- git commit -m"add: webpack5 ts config"
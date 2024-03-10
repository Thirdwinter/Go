const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true,
//   devServer: {
//     	//代理axios
//         proxy: 'http://localhost:9000',
//         //vue自己启动的端口
//         port:9001
//     }
// })
module.exports = {
    devServer: {
      // 自动打开浏览器
      //open: true,
      port: 8081,
      proxy: {
        '/Api': {
          target: `http://47.93.40.133:8080/api/v1`,
          changeOrigin: true,
          pathRewrite: {
            '^/Api': ''
          }
        }
      }
    }
  }
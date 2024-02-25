const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    	//代理axios
        proxy: 'http://localhost:9000',
        //vue自己启动的端口
        port:9001
    }
})

const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  publicPath: './',
  parallel: false,
  transpileDependencies: true,
  lintOnSave: false,
	chainWebpack: config => {
    config.module
      .rule('md')
      .test(/\.md$/)
      .use('vue-loader')
      .loader('vue-loader')
      .end()
      .use('vue-markdown-loader')
      .loader('vue-markdown-loader/lib/markdown-compiler')
      .options({
        raw: true
      })
  }
})

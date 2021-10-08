module.exports = {
  assetsDir: 'static',
  publicPath: '/',

  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到GinBlog'
      return args
    })
  },

  devServer: {
    port: 8081,   
  },

  productionSourceMap: false,

  pluginOptions: {
    quasar: {
      importStrategy: 'kebab',
      rtlSupport: false
    }
  },

  transpileDependencies: [
    'quasar'
  ]
}
  

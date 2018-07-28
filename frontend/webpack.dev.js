const webpack = require('webpack');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  devtool: 'inline-source-map',
  devServer: {
    compress: false,
    host: '0.0.0.0',
    index: 'index.html',
    port: 5001,
    publicPath: '/',
  },
  mode: 'development',
});

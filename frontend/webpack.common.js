const path = require('path');
const webpack = require('webpack');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');

const distPath = path.resolve(__dirname, 'dist');
console.log(distPath);

module.exports = {
  entry: {
    app: './js/entry.js',
    style: './css/entry.js',
  },
  plugins: [
    new CleanWebpackPlugin([distPath]),
    new VueLoaderPlugin(),
  ],
  output: {
    filename: '[name].bundle.js',
    path: distPath,
  },
  module: {
    rules: [
      {
        enforce: 'pre',
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loader: 'eslint-loader',
      },
      {
        test: /\.vue$/,
        exclude: /node_modules/,
        loader: 'vue-loader',
      },
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        options: {
          'presets': [
            ['env', {
              'targets': {
                'browsers': [
                  'last 2 versions',
                ],
              },
            }],
          ],
          'plugins': [],
        },
      },
      {
        test: /\.(sass|scss)$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader',
        ],
      },
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader',
        ],
      },
    ],
  },
  resolve: {},
};

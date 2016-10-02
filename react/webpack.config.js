var debug = process.env.NODE_ENV !== "production";
var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry:  "./app.jsx",
  output: {
    path: __dirname,
    filename: "bundle.js"
  },
    resolveLoader: {
        root: path.join(__dirname, 'node_modules')
    },
  module: {
    loaders: [
      {
        test: /.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ['es2015', 'react', 'stage-1']
        }
      },
      { 
        test: /\.css$/,
        exclude: /node_modules/, 
        loader: "style-loader!css-loader"
      },
      {
        test: /\.less$/,
        exclude: /node_modules/,       
        loader: ExtractTextPlugin.extract("style-loader", "css-loader!less-loader")
      }
    ]
  },
plugins: [
        new ExtractTextPlugin("styles.css"),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
        })
    ]
};

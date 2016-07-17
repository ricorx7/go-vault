var path = require('path');
var webpack = require('webpack');

module.exports = {
  entry:  {
    hello: "./hello.jsx",
    listFixed: "./watertestFixedTable.jsx",
    listComp: ["./watertestListComp.jsx", "./header.jsx"],
    listCompTest: ["./watertestListCompTest.jsx", "./header.jsx"],
    watertest: ["./watertestList.jsx", "./header.jsx"]
 },
  output: {
    path: __dirname,
    filename: "bundle.[name].js"
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
          presets: ['es2015', 'react']
        }
      }
    ]
  },
plugins: [
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
        })
    ]
};

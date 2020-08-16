module.exports = {
    productionSourceMap: false,
    devServer: {
        proxy: "http://0.0.0.0:9999"
    },
    outputDir: "../build/static",
    pages: {
        index: "src/main.js"
    }
};
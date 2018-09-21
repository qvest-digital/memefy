const pkg = require('./package')
const webpack = require('webpack')

module.exports = {
    mode: 'spa',

    /*
    ** Headers of the page
    */
    head: {
        title: pkg.name,
        meta: [
            {charset: 'utf-8'},
            {name: 'viewport', content: 'width=device-width, initial-scale=1'},
            {hid: 'description', name: 'description', content: pkg.description}
        ],
        link: [
            {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
        ]
    },

    /*
    ** Customize the progress-bar color
    */
    loading: {color: '#3B8070'},

    /*
    ** Global CSS
    */
    css: [
        'bootstrap/dist/css/bootstrap.css',
        'admin-lte/dist/css/AdminLTE.min.css',
        'admin-lte/dist/css/skins/skin-black.min.css',
        'font-awesome/css/font-awesome.min.css',
    ],

    /*
    ** Plugins to load before mounting the App
    */
    plugins: [
        '~/plugins/bootstrap',
        '~/plugins/global_components',
    ],

    /*
    ** Nuxt.js modules
    */
    modules: [
        // Doc: https://github.com/nuxt-community/axios-module#usage
        '@nuxtjs/axios'
    ],

    /*
    ** Axios module configuration
    */
    axios: {
        // See https://github.com/nuxt-community/axios-module#options
        proxy: true,
    },

    proxy: {
        '/meme': {
            target: process.env.MEME_SERVER_URL || 'http://localhost:8080',
            pathRewrite: {'^/meme/': ''},
            logLevel: 'debug'
        }
    },

    /*
    ** Build configuration
    */
    build: {
        vendor: ['bootstrap', 'admin-lte'],
        plugins: [
            // set shortcuts as global for bootstrap
            new webpack.ProvidePlugin({
                jQuery: 'jquery',
                'window.jQuery': 'jquery',
                $: 'jquery',
            })
        ],
        /*
        ** You can extend webpack config here
        */
        extend(config, ctx) {

        }
    }
}

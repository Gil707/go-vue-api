module.exports = {
    head: {
        title: 'go-nuxt',

        meta: [
            {charset: 'utf-8'},
        ],

        link: [
            { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
            { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Lato:300' }
        ],

        base: {
            target: '_blank',
            href: '/'
        },

        noscript: [
            {innerHTML: 'This website requires JavaScript.'}
        ],
    },
    css: [
        '~/assets/css/style.css',
    ],
    modules: [
        'bootstrap-vue/nuxt',
    ]
};
var app = new Vue({
    el: '#app',
    data: {
        title: "Users/Products",
        time_users: null,
        time_products: null,
        time_user: null,
        selectedUser: null,
        users:[],
        products:[],
        user: null,
        avUsers: [
            1,2,3,4,5,6,7,8
        ]
    },
    methods: {
        getUsers: function () {
            let start = new Date();
            this.$http.get('http://localhost:9090/api/v1/users').then(res => {
                console.log('res:', res);
                this.time_users = (new Date() - start) + ' ms';
                this.users = res.data.users
            }, res => {
                console.log('error:', res)
            });
        },
        getUser: function () {
            let start = new Date();
            this.$http.get('http://localhost:9090/api/v1/users/' + this.selectedUser).then(res => {
                console.log('res:', res);
                this.time_user = (new Date() - start) + ' ms';
                this.user = res.data.user
            }, res => {
                console.log('error:', res)
            });
        },
        getProducts: function () {
            let start = new Date();
            this.$http.get('http://localhost:9090/api/v1/products').then(res => {
                console.log('res:', res);
                this.time_products = (new Date() - start) + ' ms';
                this.products = res.data.products
            }, res => {
                console.log('error:', res)
            });
        },
    }
});
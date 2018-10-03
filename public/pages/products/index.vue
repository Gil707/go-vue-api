<template>
    <div>
        <navbar></navbar>
        <div class="container-fluid">
            <div class="row">
                <div class="col-12">
                    <table class="table table-condensed table-striped">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Name</th>
                            <th scope="col">Basket</th>
                            <th scope="col">Count</th>
                            <th scope="col">Price</th>
                            <th scope="col">Delivery</th>
                            <th scope="col">Comment</th>
                            <th scope="col">Status</th>
                            <th scope="col">Pushed</th>
                            <th scope="col">Updated_at</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr v-for='(product, index) in products' :key='product.p_id'>
                            <th scope="row">{{product.p_id}}</th>
                            <td><a :href="product.link">{{product.name}}</a></td>
                            <td>{{product.basket}}</td>
                            <td>{{product.count}}</td>
                            <td>{{product.price}} ¥</td>
                            <td>{{product.delivery_price}}</td>
                            <td>{{product.editor_comm}}</td>
                            <td>{{product.status}}</td>
                            <td>{{product.pushed}}</td>
                            <td>{{product.updated_at}}</td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import Navbar from '../../components/Navbar'
    import axios from 'axios'

    export default {
        components: {
            Navbar
        },
        async asyncData () {
            const { data } = await axios.get('http://localhost:9090/api/v1/products');
            return { products: data.products }
        }
    }
</script>
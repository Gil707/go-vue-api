<template>
    <div>
        <navbar></navbar>
        <div class="container-fluid">
            <div class="row">
                <div class="col-8">
                    <table class="table table-condensed table-striped">
                        <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Login</th>
                            <th scope="col">Name</th>
                            <th scope="col">Surname</th>
                            <th scope="col">Age</th>
                            <th scope="col">Type</th>
                            <th scope="col">Car</th>
                            <th scope="col">Status</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr v-for='(user, index) in users' :key='user.id'>
                            <th scope="row">{{user.id}}</th>
                            <td>{{user.login}}</td>
                            <td>{{user.name}}</td>
                            <td>{{user.surname}}</td>
                            <td>{{user.age}}</td>
                            <td>{{user.type}}</td>
                            <td>{{user.car}}</td>
                            <td>{{user.status}}</td>
                        </tr>
                        </tbody>
                    </table>
                </div>

                <div class="col-4">
                    <h5>Select specific user_id</h5>
                    <select class="form-control form-control-sm" v-model='selectedUser' @change='getUser'>
                        <option v-for='u in avUsers'>{{u}}</option>
                    </select>
                    <div class="card" v-if="user">
                        <div class="card-body">
                            <h5 class="card-title">Login: {{user.login}}</h5>
                            <p class="card-text">Id: <span>{{user.id}}</span></p>
                            <p class="card-text">Name: <span>{{user.name}}</span></p>
                            <p class="card-text">Surname: <span>{{user.surname}}</span></p>
                            <p class="card-text">Age: <span>{{user.age}}</span></p>
                            <p class="card-text">Type: <span>{{user.type}}</span></p>
                            <p class="card-text">Car: <span>{{user.car}}</span></p>
                            <p class="card-text">Status: <span>{{user.status}}</span></p>
                            <p class="card-text">Registered: <span>{{user.created_at}}</span></p>
                        </div>
                    </div>
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
            const { data } = await axios.get('http://localhost:9090/api/v1/users');
            return { users: data.users }
        },
        data () {
            return {
                selectedUser: null,
                user: null,
                avUsers: [
                    1,2,3,4,5,6,7,8
                ]
            }
        },
        methods: {
            async getUser () {
                const { data } = await axios.get('http://localhost:9090/api/v1/users/' + this.selectedUser);
                this.user = data.user;
            },
        }
    }
</script>
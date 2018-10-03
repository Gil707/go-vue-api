<template>
    <div>
        <navbar></navbar>
        <nuxt/>
        <div class="container-fluid">
            <div class="row">
                <div class="col-lg-5">
                    <h4 class="title">Simple Nuxt.js app with GoLang back</h4>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import Navbar from '../components/Navbar'

    export default {
        components: {
            Navbar
        }
    }
</script>
<template>
        <nav class="navbar navbar-dark bg-dark" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
                <span>GORESTAPI</span>
            </div>
            <div class="navbar-menu">
                <div class="navbar-end">
                    <nuxt-link class="navbar-item" to="/">Home</nuxt-link>
                    <span class="divider">|</span>
                    <nuxt-link class="navbar-item" to="/users">Users</nuxt-link>
                    <span class="divider">|</span>
                    <nuxt-link class="navbar-item" to="/products">Products</nuxt-link>
                </div>
            </div>
        </nav>
</template>
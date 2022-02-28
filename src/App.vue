<template>
  <div id="app">
    <b-navbar toggleable="lg" type="light">
      <b-navbar-brand>
        <a href="/home">
          <img
              alt="logo"
              height="50"
              src="https://raw.githubusercontent.com/soekarnohatta/PilEk-Vue/main/logo.png"
              href="/home"
          />
        </a>
      </b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <button v-if="!currentUser" class="btn btn-outline-success my-2 my-sm-0" type="submit" href="/login">Masuk</button>
            <button v-if="currentUser" class="btn btn-outline-success my-2 my-sm-0" href type="submit" @click.prevent="logOut">Keluar</button>
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <br/>
    <div class="container">
      <router-view/>
    </div>
    <br>
  </div>
</template>

<script>
import Vue from "vue";
import {NavbarPlugin} from 'bootstrap-vue'

export default {
  created() {
    Vue.use(NavbarPlugin);
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    },
  }
};
</script>
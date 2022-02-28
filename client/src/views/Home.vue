<template>
  <div class="container-fluid">
    <header class="jumbotron">
      <h5 class="text-muted mb-0 mt-10">Selamat Datang,</h5>
      <h3 v-if="!currentUser"><b>Laman Aplikasi Pusdatin OSIS SMAIA 19</b></h3>
      <h3 v-if="currentUser" class="mb-0"><b>{{ currentUser.nama }}</b></h3>
    </header>

    <div class="content">
      <h5 class="text-muted">Pilih Aplikasi</h5>
    </div>

    <div class="row gutters-tiny ">
      <div v-if="!currentUser" class="col-12 col-lg-3 mb-10 mt-3">
        <div v-if="!currentUser" class="card h-100">
          <div class="card-body">
            <div class="row">
              <div class="col-2 text-center">
                <i class="fas fa-check-circle fa-2x text-primary-dark"></i>
              </div>
              <div class="col-10">
                <h6 class="card-title mb-0">
                  <router-link to="/login">
                    <a class="stretched-link text-black text-decoration-none">Login</a>
                  </router-link>
                </h6>
                <p class="card-text text-muted text-truncate">
                  Silahkan Login Disini</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="currentUser" class="col-12 col-lg-3 mb-10 mt-3">
        <div v-if="currentUser" class="card h-100">
          <div class="card-body">
            <div class="row">
              <div class="col-2 text-center">
                <i class="fas fa-check-circle fa-2x text-primary-dark"></i>
              </div>
              <div class="col-10">
                <h6 class="card-title mb-0">
                  <router-link to="/vote">
                    <a class="stretched-link text-black text-decoration-none">
                      Voting</a>
                  </router-link>
                </h6>
                <p class="card-text text-muted text-truncate">
                  Silahkan Voting Disini</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="currentUser" class="col-12 col-lg-3 mb-10 mt-3">
        <div v-if="currentUser" class="card h-100">
          <div class="card-body">
            <div class="row">
              <div class="col-2 text-center">
                <i class="fas fa-user fa-2x text-primary-dark"></i>
              </div>
              <div class="col-10">
                <h6 class="card-title mb-0">
                  <router-link to="/profile">
                    <a class="stretched-link text-black text-decoration-none">
                      Profil</a>
                  </router-link>
                </h6>
                <p class="card-text text-muted text-truncate">
                  Silahkan Profil Disini</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="currentUser" class="col-12 col-lg-3 mb-10 mt-3">
        <div v-if="currentUser" class="card h-100">
          <div class="card-body">
            <div class="row">
              <div class="col-2 text-center">
                <i class="fas fa-info fa-2x text-primary-dark"></i>
              </div>
              <div class="col-10">
                <h6 class="card-title mb-0">
                  <router-link to="/settings">
                    <a class="stretched-link text-black text-decoration-none">
                      Pengaturan</a>
                  </router-link>
                </h6>
                <p class="card-text text-muted text-truncate">
                  Silahkan Atur Disini</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {dom, library} from '@fortawesome/fontawesome-svg-core';
import {fas} from '@fortawesome/free-solid-svg-icons';
import NProgress from 'nprogress/nprogress'

export default {
  name: 'Home',
  created() {
    NProgress.done()
    dom.watch()
    library.add(fas);
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    }
  },
};
</script>

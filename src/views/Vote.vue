<template>
  <div class="container-fluid">
    <header class="jumbotron">
      <h5 class="text-muted mb-0 mt-10">Laman Aplikasi</h5>
      <h3>
        <strong>Pemilihan Ketua dan Wakil Ketua OSIS</strong>
      </h3>
    </header>

    <!-- API Error -->
    <div v-if="error">
      <b-alert show variant="danger">
        <h4 class="alert-heading"><strong>Terjadi Masalah Pada API Pusdatin OSIS SMAIA 19 !</strong></h4>
        <p>
          Terjadi permasalahan saat mengambil data dari API Pusdatin OSIS SMAIA 19. Silahkan hubungi admin
          terkait untuk menyelesaikan permasalahan ini. Kami mohon maaf yang sebesar besarnya atas galat
          yang terjadi.
        </p>
        <hr>
        <p class="mb-0">{{ error }}</p>
      </b-alert>
    </div>

    <!-- Vote Start -->
    <div v-if="!error" class="content">
      <h5 class="text-muted">Silahkan Pilih Kandidat</h5>
    </div>
    <div v-if="!error" class="container mt-3">
      <div class="mt-3">
        <b-card-group deck>
          <div v-for="(kandidat) in content" :key="kandidat.idkandidat" class="col-md-3">
            <b-card :img-src="kandidat.image" class="mt-3" img-alt="Image" img-height="245px" img-top style="object-fit:cover;">
              <b-card-title class="text-center">
                <h4><strong>{{ kandidat.nama }}</strong></h4>
              </b-card-title>
              <b-card-text class="mb-0 text-center align-items-center">
                <b>Visi :
                  <b-badge v-b-modal:[`visi-modal-${kandidat.idkandidat}`] block pill variant="light">Lihat Disini</b-badge>
                </b>
              </b-card-text>
              <b-card-text class="mb-0 text-center align-items-center">
                <b>Misi :
                  <b-badge v-b-modal:[`misi-modal-${kandidat.idkandidat}`] block pill variant="light">Lihat Disini</b-badge>
                </b>
              </b-card-text>
              <template #footer>
                <button :disabled="active[kandidat.idkandidat]" class="btn btn-primary btn-block" @click="handleVote(kandidat)">
                  <span v-if="active[kandidat.idkandidat]" class="spinner-border spinner-border-sm"></span>
                  <span>Pilih</span>
                </button>
              </template>

              <!-- Modals -->
              <b-modal :id="`visi-modal-${kandidat.idkandidat}`" centered header-bg-variant="dark" header-text-variant="light"
                       scrollable
                       title="Visi">
                <p class="my-4">
                  {{ kandidat.visi }}
                </p>
              </b-modal>
              <b-modal :id="`misi-modal-${kandidat.idkandidat}`" centered header-bg-variant="dark" header-text-variant="light" scrollable title="Misi">
                <p class="my-4">{{ kandidat.misi }}</p>
              </b-modal>
            </b-card>
          </div>
        </b-card-group>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import KandidatService from "@/services/kandidat.service";
import UserService from "@/services/user.service";
import {CardPlugin, BadgePlugin, ModalPlugin, AlertPlugin} from 'bootstrap-vue/src'
import NProgress from 'nprogress/nprogress'

export default {
  name: "Vote",
  data() {
    return {
      content: [],
      error: "",
      active: []
    };
  },
  created() {
    Vue.use(CardPlugin);
    Vue.use(BadgePlugin);
    Vue.use(ModalPlugin);
    Vue.use(AlertPlugin);
  },
  beforeCreate() {
    UserService.getUser().then(response => {
      if (response.data.userlogin.status === "Sudah Memilih") {
        NProgress.done()
        this.$router.push('/sudah');
      }
    });
    KandidatService.getAll().then(
        response => {
          this.content = response.data.kandidat;
          NProgress.done()
        },
        error => {
          this.error = (error.response && error.response.data && error.response.data.errors.body) || error.message;
        }
    );
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    }
  },
  methods: {
    handleVote(kandidat) {
      if (this.active[kandidat.idkandidat]) {
        this.active[kandidat.idkandidat] = false;
      } else if (!this.active[kandidat.idkandidat]) {
        Vue.set(this.active, kandidat.idkandidat, true);
        this.$store.dispatch('vote/vote', kandidat.nomorurut).then(
            () => {
              this.$router.push('/sudah')
            },
            error => {
              if (error.response.status === 400) {
                this.$router.push('/sudah');
              } else {
                this.error = (error.response && error.response.data && error.response.data.errors.body) || error.message;
                this.active[kandidat.idkandidat] = false;
              }
            }
        );
      }
    }
  }
}
</script>
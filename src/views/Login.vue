<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          alt="Login"
          class="profile-img-card"
          src="https://yt3.ggpht.com/ytc/AKedOLQaf7Xc6BXReUa9pePkrz7hxs7_083D0dqQVNNIMw=s900-c-k-c0x00ffffff-no-rj"
      />
      <form name="form" @submit.prevent="handleLogin">
        <div class="form-group">
          <label>Username</label>
          <input
              v-model="user.username"
              v-validate="'required'"
              class="form-control"
              name="username"
              type="text"
          />
          <div
              v-if="vErrors.has('username')"
              class="alert alert-danger"
              role="alert"
          >Username is required!
          </div>
        </div>
        <div class="form-group">
          <label>Password</label>
          <input
              v-model="user.password"
              v-validate="'required'"
              class="form-control"
              name="password"
              type="password"
          />
          <div
              v-if="vErrors.has('password')"
              class="alert alert-danger"
              role="alert"
          >Password is required!
          </div>
        </div>
        <div class="form-group">
          <button :disabled="loading" class="btn btn-primary btn-block">
            <span v-show="loading" class="spinner-border spinner-border-sm"></span>
            <span>Login</span>
          </button>
        </div>
        <div class="form-group">
          <div v-if="message" class="alert alert-danger" role="alert">{{ message }}</div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import User from '../models/user';
import VeeValidate from 'vee-validate';
import Vue from "vue";
import NProgress from 'nprogress/nprogress'

Vue.use(VeeValidate, {errorBagName: 'vErrors'});

export default {
  name: 'Login',
  data() {
    return {
      user: new User('', ''),
      loading: false,
      message: '',
    };
  },
  created() {
    NProgress.done()
    if (this.loggedIn) {
      this.$router.push('/home')
    }
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.statusUser.loggedIn;
    }
  },
  methods: {
    async handleLogin() {
      this.loading = true;
      this.$validator.validateAll().then(isValid => {
        if (!isValid) {
          this.loading = false;
          return;
        }

        if (this.user.username && this.user.password) {
          this.$store.dispatch('auth/login', this.user).then(
              async data  => {
                if (data.aktif ){
                  await this.$router.push('/home');
                }

                this.loading = false;
                this.message = 'Akun Ada Dinonaktifkan. Silahkan Hubungi Admin Untuk Mengaktifkan Kembali.'
              },
              async error => {
                this.loading = false;
                this.message =  (error.response && error.response.data && error.response.data.errors.body) || error.message;
              }
          );
        }
      });
    }
  }
};
</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}

.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 50px auto 25px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}
</style>
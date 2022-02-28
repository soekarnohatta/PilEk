<template>
  <div class="container-fluid">
    <header class="jumbotron">
      <h5 class="text-muted mb-0 mt-10">Profil</h5>
      <h3>
        <strong>{{ user.nama }}</strong>
      </h3>
    </header>
    <div class="card bd-0 bd-md-x bd-md-y mg-t-10">
      <div class="card-header">
        Data Pribadi
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-xl-9 order-last order-xl-first">
            <div v-for="(val,key,idx) in user" :key=idx>
              <div class="form-group row">
                <label class="col-sm-3 col-form-label text-md-right text-capitalize"><b>{{ key }}</b></label>
                <label class="col-sm-8 col-form-label">{{ val }}</label>

              </div>
            </div>
          </div>
          <div class="col-xl-3 order-first order-xl-last text-center text-xl-left mb-5 mb-xl-0">
            <div class="photo-frame">
              <img src="//ssl.gstatic.com/accounts/ui/avatar_2x.png" width="200" height="300" alt="Foto">
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserService from "@/services/user.service";
import User from "@/models/user";
import NProgress from 'nprogress/nprogress'

export default {
  name: 'Profile',
  data() {
    return {user: new (User)}
  },
  beforeMount() {
    UserService.getUser().then(response => {
      this.user = response.data.userlogin
      NProgress.done()
    });
    if (!this.user) {
      this.$router.push('/login');
    }
  }
};
</script>
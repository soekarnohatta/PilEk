<template>
  <div class="content container-fluid">
    <div class="container ht-100p">
      <div class="ht-100p d-flex flex-column justify-content-center">
        <div class="row align-items-center">
          <div class="col-md-6 order-md-last mt-6">
            <img alt="Error 404" class="img-fluid" src="https://portal.ltmpt.ac.id/assets/media/img/illustration-404.png">
          </div>
          <div class="col-md-6 order-md-first mg-t-50 mg-md-t-0 mt-6">
            <h1 id="title" class="text-primary font-weight-bold mg-b-0">Oops!</h1>
            <h5 id="subtitle" class="text-primary text-uppercase font-weight-bold">Terjadi Kesalahan</h5>
            <p class="mg-t-60 mg-b-20 tx-18 text-primary">Halaman yang kamu cari tidak ditemukan! Kembali ke laman utama.</p>
            <a href="/home">
              <button class="btn-primary btn-block rounded-pill" href="/home"><b>Kembali Ke Laman Utama</b></button>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NProgress from 'nprogress/nprogress'

export default {
  name: 'error',
  created() {
    NProgress.done()
  }
}
</script>
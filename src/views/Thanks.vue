<template>
  <div class="container text-center" style="margin-top: 20vh">
    <h1 class="display-3">Terima Kasih!</h1>
    <p class="lead"><strong>Anda sudah menggunakan hak pilih anda</strong><br>Kembali ke laman utama dalam <span
        id="time"></span></p>
    <hr>
    <p class="lead">
      <a class="btn btn-primary btn-sm" href="/home" role="button">Laman Utama</a>
    </p>
  </div>
</template>

<script>
import NProgress from 'nprogress/nprogress'

export default {
  name: "Thanks",
  mounted() {
    this.onWindowLoad();
  },

  methods: {
    onWindowLoad() {
      NProgress.done()
      const fiveSeconds = 5,
            display = document.querySelector('#time');
      startTimer(fiveSeconds, display);
    },
  },
}

function startTimer(duration, display) {
  let start = Date.now(),
      diff,
      minutes,
      seconds;

  function timer() {
    // get the number of seconds that have elapsed since
    // startTimer() was called
    diff = duration - (((Date.now() - start) / 1000) | 0);

    // does the same job as parseInt truncates the float
    minutes = (diff / 60) | 0;
    seconds = (diff % 60) | 0;

    minutes = minutes < 10 ? "0" + minutes : minutes;
    seconds = seconds < 10 ? "0" + seconds : seconds;

    display.textContent = seconds;

    if (diff <= 0) {
      // add one second so that the count down starts at the full duration
      // example 05:00 not 04:59
      start = Date.now() + 1000;
      window.location.href = "./home";
    }
  }

  // we don't want to wait a full second before the timer starts
  timer();
  setInterval(timer, 1000);
}
</script>
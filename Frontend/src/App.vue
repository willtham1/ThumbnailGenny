<template>
  <div id="app" class='container'>
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail for your project!</h1>
        <form @submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
              <input v-model="websiteUrl" type="text" id="website-input" class='form-control' placeholder="Enter your website url">
          </div>
          <div class="form-group">
              <button class ="btn btn-primary">Genny Generate!</button>
          </div>
        </form>
        <img :src="thumbnailUrl" class="picture"/>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',

  data() { return {
    websiteUrl: '',
    thumbnailUrl: ''
  } },

  methods: {
    makeWebsiteThumbnail() {
      axios.post(`https://shot.screenshotapi.net/screenshot`, {
        token: '7CE50YV-ARY4S9R-QH2QGFJ-6619ANX',
        url: this.websiteUrl,
        width: 1920,
        height: 1080,
        output: 'json',
        thumbnail_width: 700
      })
      .then(response => {
        this.thumbnailUrl = response.data.screenshot
      })
      .catch(error => {
        window.alert(`The API returned an error:  ${error}`)
      })
    }
  }
}
</script>

<style>
.container {
  outline: 1px solid #ccc;
}
</style>

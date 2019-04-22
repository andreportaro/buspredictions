<template>
  <div id="app" class="mw6 mt5 pa4 center bg-near-white bg-white">
    <h2>Predictions</h2>
    <form action="#">

<label class="f3 mb2 dib w-100">Route
    <input type="text" v-model="route"
      class="w-100 mt2 pl2 dib input-reset b--silver ba h2 f4 gray"
    >
</label>
<label class="f3 mb2 dib w-100">Stop

    <input type="text" v-model="stop_id"
      class="w-100 mt2 pl2 dib input-reset b--silver ba h2 f4 gray"
    >
</label>

    <button @click.prevent="search" 
    :disabled="busy"
    class="fr mt2 b--transparent h2 bg-blue white ph3 copy pointer">Search</button>

    </form>
    <div class="w-100 dib">

    <ul class="list pl0">
      <li v-for="prediction in predictions" :key="prediction._id"
          class="f4 gray h3 ba b--silver  pa3"
      >
        Vehicle {{ prediction.Vehicle }}  will arrive in {{ prediction.Minutes }} minutes
      </li>
    </ul>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'app',

  data() {
    return {
      route: 80,
      stop_id: 11777,
      predictions: [],
      busy: false,
    }
  },

  mounted() {
    this.fetchSearches();
  },

  methods: {
    search() {
      this.getPredictions();
    },

    fetchSearches() {
      axios.get(`/history`).then((response) => this.history = response.data.History)
    },

    getPredictions() {
      this.busy = true;
      
          axios.get(`/search?r=${this.route}&s=${this.stop_id}`)
    .then((response) => {
      this.predictions = response.data.Directions[0].Predictions;

    }).catch(() => 
      this.predictions = []
    ).finally(this.busy = false);
    }
  },
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>

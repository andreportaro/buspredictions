<template>
  <div id="app" class="mw6 mt5 pa4 center bg-near-white bg-white">
    <h2>Predictions</h2>
    <form action="#">
      <label class="f3 mb2 dib w-100">
        Route
        <input
          type="text"
          v-model="route"
          class="w-100 mt2 pl2 dib input-reset b--silver ba h2 f4 gray"
        >
      </label>
      <label class="f3 mb2 dib w-100">
        Stop
        <input
          type="text"
          v-model="stop_id"
          class="w-100 mt2 pl2 dib input-reset b--silver ba h2 f4 gray"
        >
      </label>

      <button
        @click.prevent="getPredictions"
        :disabled="busy"
        class="fr mt2 b--transparent h2 bg-blue white ph3 copy pointer"
      >Search</button>
    </form>
    <div class="w-100 dib">
      <div v-if="history.length > 0">
        <h3 class="w-100 db">From your history</h3>
        <div
          v-for="hist in history"
          :key="hist.uuid"
          @click="search(hist)"
          class="mt4 w-25 ba b--blue pa3 h-input blue pointer"
        >{{ hist.route_id }} / {{ hist.stop_id }}</div>
      </div>
      <ul class="list pl0">
        <li
          v-for="prediction in predictions"
          :key="prediction._id"
          class="f4 gray h3 ba b--silver pa3"
        >Vehicle {{ prediction.Vehicle }} will arrive in {{ prediction.Minutes }} minutes</li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "app",

  data() {
    return {
      route: 80,
      stop_id: 11777,
      predictions: [],
      busy: false,
      searching: false,
      history: []
    };
  },

  mounted() {
    this.history = window.searches || [];
  },

  methods: {
    search() {
      this.getPredictions();
    },

    fetchHistory() {
      axios.get(`/history`).then((response) => {
        this.history = response.data || []; 
      });
    },

    getPredictions() {
      this.busy = true;

      this.busy = false;

      axios
        .get(`/search?r=${this.route}&s=${this.stop_id}`)
        .then(response => {
          this.predictions = response.data.Directions[0].Predictions;

          let indexAt = this.history.filter(
            i => i.route_id == this.route && i.stop_id === this.stop_id
          );

          if (!indexAt.length) {
            this.history.push({
              route_id: this.route,
              stop_id: this.stop_id
            });
          }
        })
        .catch(() => (this.predictions = []))
        .finally((this.busy = false));
    },

    search(hist) {
      this.route = hist.route_id;
      this.stop_id = hist.stop_id;
      this.getPredictions();
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>

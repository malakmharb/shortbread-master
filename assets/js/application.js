require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {

});

import Vue from "vue";
import VueRouter from "router";
Vue.use(VueRouter);

import Parking from "./components/parking.vue";
import Residents from "./components/residents.vue"
import ParkingSecond from "./components/parking-second.vue"

const routes = [
  { path: "/parking-first-floor", component: Parking },

  { path: "/resident-table", component: Residents },
    
  { path: "/parking-second-floor", component: ParkingSecond },

  // catch-all
  { path: "*", redirect: "/parking-first-floor" }
];

import Banner from "./components/banner.vue";
Vue.component("banner", Banner);


const router = new VueRouter({
  mode: "history",
  routes
});

const app = new Vue({
  router
}).$mount("#app");

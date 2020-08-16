import Vue from "vue";
import Vuex from "vuex";
import table from "./modules/table"
import ticker from "./modules/ticker"
Vue.use(Vuex);
let store = new Vuex.Store({
    modules: {
        table,
        ticker
    }
})

export default store
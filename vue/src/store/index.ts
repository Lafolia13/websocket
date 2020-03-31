import Vue from "vue";
import Vuex from "vuex";
import { ChatI } from "./modules/chat";

Vue.use(Vuex);

export interface RootI {
  chat: ChatI;
}

const store = new Vuex.Store<RootI>({});

export default store;

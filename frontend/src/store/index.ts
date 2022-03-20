import { createStore, Store } from 'vuex';
import { InjectionKey } from 'vue';

import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import main, { MainStore } from './main';
import word, { WordStore } from './word';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  word: WordStore;
};

// define injection key
export const key: InjectionKey<Store<MainTree>> = Symbol();

export default createStore({
  modules: { main, modal, word },
});

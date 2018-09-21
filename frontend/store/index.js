import Vuex from "vuex";
import meme from './meme'

const createStore = () => {
    let store = new Vuex.Store({
        modules: {
            meme
        }
    })

    store.dispatch('meme/init')

    return store
}

export default createStore;

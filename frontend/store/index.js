import Vuex from "vuex";
import meme from './meme'

const createStore = () => {
    return new Vuex.Store({
        modules: {
            meme
        }
    })
}

export default createStore;

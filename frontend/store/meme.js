
export const state = () => ({
    memes: []
})

export const mutations = {
    addMeme(state, meme) {
        state.memes.push(meme)
    }
}

const actions = {
    init(vuexContext) {
        return new Promise((resolve, reject) => {
            //TODO: ask for memes

            console.log("HAllo ?")

            vuexContext.commit('addMeme', {
                id: 'id1',
                name: 'Testname',
                pic: 'https://media.giphy.com/media/gSIz6gGLhguOY/giphy.gif',
                sound: 'http://...'
            })

            setInterval(() => {
                vuexContext.commit('addMeme', {
                    id: 'id2',
                    name: 'Testname2',
                    pic: 'https://media.giphy.com/media/xUPGcA1SkYqVLDtxiU/giphy.gif',
                    sound: 'http://...'
                })
            }, 1000)

            resolve()
        });
    },
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}

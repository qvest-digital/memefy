
export const state = () => ({
    memes: []
})

export const mutations = {
    addMeme(state, meme) {
        state.memes.push(meme)
    },
    updateMeme(state, meme) {
        //update a existing meme or add a new one
        for(let curMeme of state.memes) {
            if(curMeme.name === meme.name) {
                curMeme.name = meme.name
                curMeme.pic = meme.pic
                curMeme.sound = meme.sound
                return;
            }
        }

        //new one..
        state.meme.push(meme)
    }
}

const actions = {
    init(ctx) {
        return new Promise((resolve, reject) => {
            //TODO: ask for memes
            ctx.commit('addMeme', {
                id: 'id1',
                name: 'Testname',
                pic: 'https://media.giphy.com/media/gSIz6gGLhguOY/giphy.gif',
                sound: 'http://...'
            })

            setTimeout(() => {
                ctx.commit('addMeme', {
                    id: 'id2',
                    name: 'Testname2',
                    pic: 'https://media.giphy.com/media/xUPGcA1SkYqVLDtxiU/giphy.gif',
                    sound: 'http://...'
                })
            }, 1000)

            resolve()
        });
    },

    saveMeme(ctx, meme) {
        ctx.commit("updateMeme", meme)
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}

export const state = () => ({
    memes: []
})

export const mutations = {
    addMeme(state, meme) {
        state.memes.push(meme)
    },
    updateMeme(state, meme) {
        //update a existing meme or add a new one
        for (let curMeme of state.memes) {
            if (curMeme.name === meme.name) {
                curMeme.name = meme.name
                curMeme.pic = meme.pic
                curMeme.sound = meme.sound
                curMeme.meta = meme.meta
                return
            }
        }

        //new one..
        state.memes.push(meme)
    },
    setMemeMeta(state, {name, meta}) {
        //update a existing meme or add a new one
        for (let curMeme of state.memes) {
            if (curMeme.name === name) {
                curMeme.meta = meta
                return
            }
        }
    }
}

const actions = {
    init(ctx) {
        return this.$axios.get('/meme/')
            .then((result) => result.data)
            .catch((err) => console.log(err))
            .then((memes) => {
                for (let curMeme of memes) {
                    ctx.commit('addMeme', {
                        name: curMeme.name,
                        pic: `/meme/${curMeme.pic}`,
                        sound: `/meme/${curMeme.sound}`,
                    })

                    //ask for meta
                    this.$axios.get(`/meme/${curMeme.meta}`)
                        .then((result) => result.data)
                        .catch((err) => console.log(err))
                        .then(metaContent => {
                            ctx.commit('setMemeMeta', {
                                name: curMeme.name,
                                meta: metaContent
                            })
                        })
                }
            })
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

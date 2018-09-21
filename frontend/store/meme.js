import cfg from '../static/config.js'

const buildMemeDataKey = (memeName, key) => {
    return `${memeName}.${key}`
}

const getDataFromLS = (memeName, key, defaultValue = null) => {
    if (!localStorage[buildMemeDataKey(memeName, key)]) {
        return defaultValue
    }

    return localStorage[buildMemeDataKey(memeName, key)]
}

const saveDataIntoLS = (memeName, key, data) => {
    localStorage.setItem(buildMemeDataKey(memeName, key), data)
}

const state = () => ({
    memes: []
})

const mutations = {
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
        state.memes.push({
            ...meme,
            hasCoolDown: false,
            cooldown: -1,
        })
    },
    setMemeMeta(state, {name, meta}) {
        //update a existing meme or add a new one
        for (let curMeme of state.memes) {
            if (curMeme.name === name) {
                curMeme.meta = meta
                return
            }
        }
    },
    setMemeCoolDown(state, name) {
        for (let curMeme of state.memes) {
            if (curMeme.name === name) {
                curMeme.lastTrigger = new Date()
                saveDataIntoLS(name, 'trigger', curMeme.lastTrigger.toISOString())

                return
            }
        }
        return null
    },
    setMemeCoolDownTime(state, {name, time}) {
        for (let curMeme of state.memes) {
            if (curMeme.name === name) {
                curMeme.cooldown = time
                return
            }
        }
    }
}

const getters = {
    getMemeByName: (state) => (name) => {
        for (let curMeme of state.memes) {
            if (curMeme.name === name) {
                return curMeme
            }
        }
        return null
    },

    getMemesInCooldown: (state, getters) => () => {
        return state.memes.filter(m => getters.hasMemeCooldown(m.name))
    },

    getMemeCoolDown: (state, getters) => (name) => {
        return new Date() - getters.getMemeByName(name).lastTrigger;
    },

    hasMemeCooldown: (state, getters) => (name) => {
        return getters.getMemeCoolDown(name) <= cfg.meme.cooldown;
    },

    cooldownLeft: (state, getters) => (name) => {
        return cfg.meme.cooldown - getters.getMemeByName(name).cooldown;
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
                        lastTrigger: new Date(getDataFromLS(curMeme.name, 'trigger', "2000-01-01T00:00:00.000Z")),

                        /*
                        //god dammed: it is REALY fucking important(!) to define ALL needed properties
                        //otherwise vue(x) has no chance to build the property-proxies and inform other components
                        //for changes ... -.-
                        */
                        cooldown: -1,
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
    },

    coolDownMeme(ctx, memeName) {
        ctx.commit("setMemeCoolDown", memeName)
    },

    updateMemeCoolDownTime(ctx, memeName) {
        ctx.commit("setMemeCoolDownTime", {
            name: memeName,
            time: ctx.getters.getMemeCoolDown(memeName)
        })
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions
}

import cfg from '../static/config'

const checkCooldown = (store) => {
    for(let meme of store.getters['meme/getMemesInCooldown']()) {
        store.dispatch('meme/updateMemeCoolDownTime', meme.name)
    }
}

export default function ({ store }) {
    setInterval(() => checkCooldown(store), cfg.meme.cooldownWatch)
}

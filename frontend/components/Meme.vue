<template>
    <div @click="triggerMeme()" class="meme thumbnail" :class="{ 'cooldown': hasCooldown }">
        <img :src="meme.pic" :alt="meme.name" style="width:100%">
        <div class="caption">
            <span>{{meme.name}}</span>
            <span class="pull-right" v-if="hasCooldown">{{cooldownInSec}}</span>
        </div>
    </div>
</template>

<script>
    import { mapActions, mapGetters } from 'vuex'
    import cfg from '../static/config'
    import moment from 'moment'

    export default {
        name: "meme",
        props: {
            meme: {
                type: Object,
            },
        },
        methods: {
            ...mapActions({
                coolDownMeme: 'meme/coolDownMeme'
            }),
            triggerMeme() {
                if(this.hasCooldown) return;

                this.$axios.get(`/meme/play?name=${this.meme.name}`).then((result) => {
                    this.coolDownMeme(this.meme.name)
                }, (err) => {
                    console.log(err)
                })
            }
        },
        computed: {
            ...mapGetters({
                hasMemeCooldown: 'meme/hasMemeCooldown',
                cooldownLeftForMeme: 'meme/cooldownLeft',
                getMemeByName: 'meme/getMemeByName',
            }),
            hasCooldown() {
                return this.hasMemeCooldown(this.meme.name)
            },
            cooldownInSec() {
                let d = moment.duration(cfg.meme.cooldown - this.meme.cooldown, 'milliseconds')
                return moment.utc(d.as('milliseconds')).format('HH:mm:ss')
            },
        },
        watch: {
        }
    }
</script>

<style scoped>
    div.meme {
       cursor: pointer;
    }
    div.cooldown {
        cursor: not-allowed;
        filter: grayscale(100%);
    }
    p.cooldown {
        text-align: right;
    }
    img {
        width: 100%;
    }
</style>
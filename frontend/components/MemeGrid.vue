<template>
    <div>
        <div class="box">
            <div class="box-body row">
                <div class="col-md-2">
                    <div class="btn-group" @click="filter.hideCooldown = !filter.hideCooldown">
                        <label class="btn" :class="{'btn-success': !filter.hideCooldown, 'btn-danger': filter.hideCooldown}">
                            <span class="glyphicon glyphicon-check" v-if="!filter.hideCooldown"></span>
                            <span class="glyphicon glyphicon-minus" v-else></span>
                        </label>
                        <label class="btn btn-default" :class="{'actice': filter.hideCooldown}">
                            <span v-if="filter.hideCooldown">Keine Cooldowns</span>
                            <span v-else>Auch Cooldowns</span>
                        </label>
                    </div>
                </div>
                <div class="col-md-2">
                    <div class="input-group">
                        <div class="input-group-btn">
                            <label class="btn btn-danger">Exklusive</label>
                        </div>
                        <input class="form-control" type="text" v-model="filter.exclusive">
                    </div>
                </div>
                <div class="col-md-2">
                    <div class="input-group">
                        <div class="input-group-btn">
                            <label class="btn btn-success">Inklusive</label>
                        </div>
                        <input class="form-control" type="text" v-model="filter.inclusive">
                    </div>
                </div>
                <div class="col-md-2 pull-right">
                    <button class="btn btn-danger pull-right" type="button" @click="clearFilter()">Filter Zurücksetzen</button>
                </div>
            </div>
        </div>
        <div class="row">
            <div v-for="meme in filteredMemes" class="col-lg-3 col-md-6">
                <meme :meme="meme"></meme>
            </div>
        </div>
    </div>
</template>

<script>
    import { mapState } from 'vuex';
    import Meme from './Meme';
    import cfg from '../static/config';

    export default {
        name: "meme-grid",
        components: {
            Meme
        },
        data(){
            return {
                filter: {
                    hideCooldown: true,
                    inclusive: '',
                    exclusive: '',
                }
            }
        },
        methods: {
            clearFilter() {
                this.filter.hideCooldown= true
                this.filter.inclusive = ''
                this.filter.exclusive = ''
            }
        },
        computed: {
            ...mapState({
                memes: state => state.meme.memes,
            }),
            filteredMemes() {
                let result = this.memes;

                if(this.filter.hideCooldown) {
                    result = result.filter(m => {
                        return m.cooldown <= 0 || m.cooldown >= cfg.meme.cooldown
                    })
                }

                if(this.filter.exclusive !== '') {
                    result = result.filter(m => {
                        return !m.name.match(new RegExp(this.filter.exclusive))
                    })
                }

                if(this.filter.inclusive !== '') {
                    result = result.filter(m => {
                        return m.name.match(new RegExp(this.filter.inclusive))
                    })
                }

                return result
            }
        },
    }
</script>

<template>
    <div @click="triggerMeme()" class="meme thumbnail">
        <img :src="meme.pic" :alt="meme.name" style="width:100%">
        <div class="caption">
            <p>{{meme.name}}</p>
        </div>
    </div>
</template>

<script>
    export default {
        name: "meme",
        props: {
            meme: {
                type: Object,
            },
        },
        methods: {
            triggerMeme() {
                this.$axios.get(`/meme/play?name=${this.meme.name}`).then((result) => {
                    console.log("success")
                }, (err) => {
                    console.log("errr", err)
                })
            }
        }
    }
</script>

<style scoped>
    div.meme {
       cursor: pointer;
    }
    img {
        width: 100%;
    }
</style>
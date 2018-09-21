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
                //TODO: trigger meme :D
                alert(`Trigger ${JSON.stringify(this.meme)}`)
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
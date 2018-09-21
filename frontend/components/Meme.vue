<template>
    <div @click="triggerMeme()" class="meme">
        <img :src="meme.pic" />
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
</style>
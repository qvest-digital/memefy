<template>
    <div class="row">
        <div v-for="meme in memes" class="col-lg-3 col-md-6">
            <meme :meme="meme"></meme>
        </div>
    </div>
</template>

<script>
    import { mapState } from 'vuex';
    import Meme from './Meme';

    export default {
        name: "meme-grid",
        components: {
            Meme
        },
        computed: {
            ...mapState({
                memes: state => state.meme.memes,
            }),
        },
    }
</script>

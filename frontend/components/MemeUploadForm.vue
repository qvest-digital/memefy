<template>
    <div class="box">
        <div class="box-header with-border">
            <h3 class="box-title">MEME Upload</h3>
        </div>

        <form role="form" @submit.prevent="upload">
            <div class="box-body">
                <message-success v-if="uploadStatus === 'success'">MEME erfolgreich gespeichert</message-success>
                <message-success v-if="uploadStatus === 'fail'">MEME konnte nicht gespeichert werden!</message-success>

                <div class="form-group">
                    <label for="meme-name">Name</label>
                    <input class="form-control" id="meme-name" placeholder="MEME Name" v-model="meme.name">
                </div>
                <div class="form-group">
                    <label for="meme-pic">MEME</label>
                    <input id="meme-pic" type="file" @change="setMemePicture">
                </div>
                <div class="form-group">
                    <label for="meme-sound">Sound</label>
                    <input id="meme-sound" type="file" @change="setMemeSound">
                </div>
            </div>
            <!-- /.box-body -->

            <div class="box-footer">
                <button type="submit" class="btn btn-primary">Speichern</button>
            </div>
        </form>
    </div>
</template>

<script>
    import { mapActions } from 'vuex'

    export default {
        name: "meme-upload-form",
        data() {
            return {
                meme: {
                    pic: null,
                    sound: null,
                    name: ''
                },
                uploadStatus: null
            }
        },
        methods: {
            ...mapActions({
                saveMeme: 'meme/saveMeme'
            }),
            setMemePicture(event) {
                this.meme.pic = event.target.files[0]
            },
            setMemeSound(event) {
                this.meme.sound = event.target.files[0]
            },
            upload() {
                const formData = new FormData();
                formData.append("name", this.meme.name);
                formData.append("pic", this.meme.pic);
                formData.append("sound", this.meme.sound);

                this.$axios.post("/meme/", formData)
                    .then((result) => {
                        this.uploadStatus = 'success'
                        // this.saveMeme(result.data)

                        //TODO: remove testcode
                        this.saveMeme({
                            name: "Testname",
                            pic: "https://media.giphy.com/media/wWue0rCDOphOE/giphy.gif",
                        });
                    }, function (error) {
                        this.uploadStatus = 'fail'

                        console.log(error);
                    });
            }
        }
    }
</script>

<style scoped>
</style>
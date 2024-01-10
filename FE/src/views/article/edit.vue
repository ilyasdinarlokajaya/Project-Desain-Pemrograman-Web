<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-article">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Update Article</h1>
                    <p class="mt5 mb20 tc-green">Make your article looks good</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Article Image</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 10%; height: auto;" />
                    <img v-else :src="article.image" alt="Uploaded Image" style="max-width: 10%; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Article Name & <br>Content</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Product Name" maxLength="20" v-model="article.judul_artikel"/>
                        <div class="p-kanan p-kanan-name" ><img src="../../assets/icons/x.png" width="30%"  @click="removeArticleName"></div>
                    <div class="character-count1 ">
                        | {{ judul_artikel_length }} / 20
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="w20">
                </div>
                <div class="w80">
                    <textarea class="form-text" cols="30" rows="10" placeholder="Content (Min. 1000 Char)"
                    v-model="article.konten">
                    </textarea>
                <div class="p-kanan2 p-kanan-desk2" ><img src="../../assets/icons/x.png" width="30%"  @click="removeContent"></div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="article.jadwal_broadcast">
                        <option disabled selected value="">Pilih Hari</option>
                        <option value="1">Senin</option>
                        <option value="2">Selasa</option>
                        <option value="3">Rabu</option>
                        <option value="4">Kamis</option>
                        <option value="5">Jumat</option>
                        <option value="6">Sabtu</option>
                        <option value="7">Minggu</option>
                    </select>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Valid From</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="date"
                        v-model="article.valid_from"
                    >
                </div>
            </div>
             <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Valid Until</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="date"
                        v-model="article.valid_until"
                    >
                </div>
            </div>
            
            <div class="flex mt30 mb40">
                <router-link to="/list-article">
                    <button class="bt-submit-green" @click="onEditArticle">Update Article</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-article',
    data(){
        return {
            imageUrl: null,
            imageData: null,
            judul_artikel_length: 5
        }
    },
	computed: {
        ...mapState({ // get data from state
			article: state => state.article.detail
		})
	},
    mounted() {
        this.detailArticle(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailArticle',
            'editArticle'
		]),
        onEditArticle() {
            this.editArticle({
                id: this.article.id,
                judul_artikel: this.article.judul_artikel,
                konten: this.article.konten,
                jadwal_broadcast: this.article.jadwal_broadcast,
                valid_from: this.article.valid_from,
                valid_until: this.article.valid_until,
                image: this.imageData != null? this.imageData : this.article.image
            })
        },
        handleInput() {
            this.judul_artikel_length = this.article.judul_artikel.length;
        },
        openFileInput() {
            // Trigger click on the file input when the button is clicked
            this.$refs.fileInput.click();
        },
        handleFileChange(event) {
            const file = event.target.files[0];

            if (file) {
                // Handle the selected file, for example, display it as an image
                this.imageUrl = URL.createObjectURL(file);
                const reader = new FileReader();
                reader.onloadend = () => {
                    this.imageData = reader.result; // Save the Base64-encoded image data
                };
                reader.readAsDataURL(file);
            }
        },
        removeArticleName() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.article.judul_artikel = ""
        },
        removeContent() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.article.konten = ""
        },
    }
}
</script>
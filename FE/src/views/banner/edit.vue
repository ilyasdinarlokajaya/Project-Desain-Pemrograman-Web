<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-banner">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Update Banner</h1>
                    <p class="mt5 mb20 tc-green">Make your customers tempted to shop via banners</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Banner Image</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <img v-else :src="banner.image" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Banner Name</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input Banner Name"
                        v-model="banner.bannername"
                       
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removeBannername"></div>

                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input Schedule"
                        v-model="banner.broadcastday"
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeBroadcastDay"></div>
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
                        v-model="banner.validfrom"
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
                        v-model="banner.validuntil"
                    >
                </div>
            </div>
            
            <div class="flex mt30 mb40">
                <router-link to="/list-banner">
                    <button class="bt-submit-green" @click="onEditBanner">Update Banner</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-banner',
    data(){
        return {
            imageUrl: null,
            imageData: null
        }
    },
	computed: {
        ...mapState({ // get data from state
			banner: state => state.banner.detail
		})
	},
    mounted() {
        this.detailBanner(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailBanner',
            'editBanner'
		]),
        onEditBanner() {
            this.editBanner({
                id: this.banner.id,
                bannerName: this.banner.bannerName,
                broadcastDay: this.banner.broadcastDay,
                validFrom: this.banner.validFrom,
                validUntil: this.banner.validUntil,
                image: this.imageData != null? this.imageData : this.banner.image
            })
        },
        removeBannername() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.banner.bannername = ""
        },
        removeBroadcastDay() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.banner.broadcastday = ""
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
    }
}
</script>
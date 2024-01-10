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
                    <h1 class="m0">Add Banner</h1>
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
                    <button v-if="imageUrl === null" class="bt-image" @click="openFileInput"><img src="../../assets/icons/upload_image.svg"></button>
                    <div v-if="imageUrl">
                        <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                        <button class="bt-change-image" @click="openFileInput">Change Image</button>
                    </div>
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
                        placeholder="Input your description"
                        v-model="bannerName"
                    >
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
                        placeholder="Input your day"
                        v-model="broadcastDay"
                    >
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
                        v-model="validFrom"
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
                        v-model="validUntil"
                    >
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-banner">
                    <button class="bt-submit-green" @click="onAddBanner">Add Banner</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'add-banner',
    data(){
        return {
            id: "",
            bannerName: "",
			broadcastDay: "",
			validFrom: "",
			validUntil: "",
            imageUrl: null,
            imageData: null
        }
    },
    methods: {
        ...mapActions([
			'addBanner',
		]),
        onAddBanner(){
            this.addBanner({
                bannerName: this.bannerName,
                broadcastDay: this.broadcastDay,
                validFrom: this.validFrom,
                validUntil: this.validUntil,
                image: this.imageData
            })
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
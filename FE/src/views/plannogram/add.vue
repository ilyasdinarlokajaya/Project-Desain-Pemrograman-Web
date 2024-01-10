<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-plannogram">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Add Plannogram</h1>
                    <p class="mt5 mb20 tc-green">Increase the visibility of promotional products in our store highlights</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Background Image</label>
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
                    <label class="form-label">Plannogram Name</label>
                </div>
                <div class="w80 relative">
                    <input class="form-text" @input="handleInput" type="text" minlength="5" maxlength="20"
                        placeholder="Masukkan Plannogram Name" v-model="plannogramName">
                    <div class="character-count ">
                        | {{ plannogramName_length }} / 20
                    </div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="broadcastDay">
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
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">List Product</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="listproduct">
                        <option disabled selected value="">--- Select Product ---</option>
                        <option value="1">Electronic</option>
                        <option value="2">Fashion</option>
                        <option value="3"></option>
                    </select>
                </div>
            </div>

            <div class="flex mt30 mb40">
                <router-link to="/list-plannogram">
                    <button class="bt-submit-green" @click="onAddPlannogram">Add Plannogram</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'add-plannogram',
    data(){
        return {
            id: "",
            plannogramName: "",
			broadcastDay: "",
			validFrom: "",
			validUntil: "",
            listproduct: "",
            imageUrl: null,
            imageData: null,
            plannogramName_length: 5,
        }
    },
    methods: {
        ...mapActions([
			'addPlannogram'
		]),
        onAddPlannogram(){
            this.addPlannogram({
                plannogramName: this.plannogramName,
                broadcastDay: this.broadcastDay,
                validFrom: this.validFrom,
                validUntil: this.validUntil,
                listproduct: this.listproduct,
                image: this.imageData,
            })
        },
        handleInput() {
            this.plannogramName_length = this.plannogramName.length;
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

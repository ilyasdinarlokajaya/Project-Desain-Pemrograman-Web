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
                    <h1 class="m0">Update Plannogram</h1>
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
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <img v-else :src="plannogram.image" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Plannogram Name</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input Plannogram"
                        v-model="plannogram.plannogramname"
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removePlannogramName"></div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="plannogram.broadcastday">
                        <option disabled selected value="">Pilih Hari</option>
                        <option value="1">Senin</option>
                        <option value="2">Selasa</option>
                        <option value="3">Rabu</option>
                        <option value="4">Kamis</option>
                        <option value="5">Jumat</option>
                        <option value="6">Sabtu</option>
                        <option value="7">Minggu</option>
                    </select>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removeBroadcastDay"></div>
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
                        v-model="plannogram.validfrom"
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
                        v-model="plannogram.validuntil"
                    >
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">List Product</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="plannogram.listproduct">
                        <option disabled selected value="">--- Select Product ---</option>
                        <option value="1">Electronic</option>
                        <option value="2">Fashion</option>
                        <option value="3"></option>
                    </select>
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-plannogram">
                    <button class="bt-submit-green" @click="onEditPlannogram">Update Plannogram</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-plannogram',
    data(){
        return {
            imageUrl: null,
            imageData: null
        }
    },
	computed: {
        ...mapState({ // get data from state
			plannogram: state => state.plannogram.detail
		})
	},
    mounted() {
        this.detailPlannogram(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailPlannogram',
            'editPlannogram'
		]),
        onEditPlannogram() {
            this.editPlannogram({
                id: this.plannogram.id,
                plannogramName: this.plannogram.plannogramname,
                broadcastDay: this.plannogram.broadcastday,
                validFrom: this.plannogram.validfrom,
                validUntil: this.plannogram.validuntil,
                listproduct: this.plannogram.listproduct,
                image: this.imageData != null? this.imageData : this.plannogram.image
            })
        },
        removePlannogramName() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.plannogram.plannogramname = ""
        },
        removeBroadcastDay() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.plannogram.broadcastday = ""
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
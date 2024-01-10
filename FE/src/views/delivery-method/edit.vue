<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-delivery-method">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Update Delivery Method</h1>
                    <p class="mt5 mb20 tc-green">Make your customers have many choices of delivery methods</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Delivery Type Image</label>
                </div>
                <div class="w20">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 80%; height: 80%;" />
                    <img v-else :src="delivery_method.image" alt="Uploaded Image" style="max-width: 80%; height: 80%;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Delivery Name</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Input Delivery Name" maxLength="15" v-model="delivery_method.delivery_name"/>
                        <div class="p-kanan p-kanan-name" ><img src="../../assets/icons/x.png" width="30%"  @click="removeName"></div>
                    <div class="character-count1 ">
                        | {{ delivery_name_length }} / 15
                    </div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Delivery Fee</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Input Delivery Fee" v-model="delivery_method.delivery_fee"/>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removeFee"></div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Estimation Delivery</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Input Estimation Delivery" v-model="delivery_method.estimation"/>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removeEstimation"></div>
                </div>
            </div>
            
            <div class="flex mt30 mb40">
                <router-link to="/list-delivery-method">
                    <button class="bt-submit-green" @click="onEditDeliveryMethod">Update Delivery Method</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-delivery-method',
    data(){
        return {
            imageUrl: null,
            imageData: null,
            delivery_name_length: 0
        }
    },
	computed: {
        ...mapState({ // get data from state
			delivery_method: state => state.delivery_method.detail
		})
	},
    mounted() {
        this.detailDeliveryMethod(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailDeliveryMethod',
            'editDeliveryMethod'
		]),
        onEditDeliveryMethod() {
            this.editDeliveryMethod({
                id: this.delivery_method.id,
                delivery_name: this.delivery_method.delivery_name,
                delivery_fee: this.delivery_method.delivery_fee,
                estimation: this.delivery_method.estimation,
                image: this.imageData != null? this.imageData : this.delivery_method.image
            })
        },
        handleInput() {
            this.delivery_name_length = this.delivery_method.delivery_name.length;
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
        removeName() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.delivery_method.delivery_name = ""
        },
        removeFee() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.delivery_method.delivery_fee = ""
        },
        removeEstimation() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.delivery_method.estimation = ""
        }
    }
}
</script>
<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-delivery-method">
                        <img src="../../assets/icons/back.svg" />
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Add Delivery Method</h1>
                    <p class="mt5 mb20 tc-green">
                        Make your customers have many choices of delivery methods
                    </p>
                </div>
            </div>
            <hr class="uline-grey" />

            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Delivery Type Image</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none;" />
                    <button v-if="imageUrl === null" class="bt-image" @click="openFileInput">
                        <img src="../../assets/icons/upload_image.svg" />
                    </button>
                    <div v-if="imageUrl">
                        <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                        <button class="bt-change-image" @click="openFileInput">
                            Change Image
                        </button>
                    </div>
                </div>
            </div>

            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Delivery Name</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Input Delivery Name" maxLength="15" v-model="delivery_name"/>
                    <div class="character-count">
                        | {{ delivery_name_length }} / 15
                    </div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Delivery Fee</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Input Delivery Fee" v-model="delivery_fee"/>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Estimation Delivery</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Input Estimation Delivery" v-model="estimation"/>
                </div>
            </div>
            <div class="flex mt30">
                <router-link to="/list-delivery-method">
                    <button class="bt-submit-green" @click="onAddDeliveryMethod">Add Delivery Method</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'add-delivery-method',
    data(){
        return {
            delivery_name: "",
            delivery_fee: "",
            estimation: "",
            imageUrl: null, 
            imageData: null,
            delivery_name_length: 15,
        }
    },
    methods: {
        ...mapActions([
			'addDeliveryMethod',
		]),
        onAddDeliveryMethod(){
            this.addDeliveryMethod({
                delivery_name: this.delivery_name,
                delivery_fee: this.delivery_fee,
                estimation: this.estimation,
                image: this.imageData
            })
        },
        handleInput() {
            this.delivery_name_length = this.delivery_name.length;
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
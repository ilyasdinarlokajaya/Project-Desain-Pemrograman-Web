<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-product">
                        <img src="../../assets/icons/back.svg" />
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Add Product</h1>
                    <p class="mt5 mb20 tc-green">
                        Make your product looking good for customers
                    </p>
                </div>
            </div>
            <hr class="uline-grey" />

            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Product Image</label>
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
                    <label class="form-label">Product Code</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Product Code" maxLength="8" v-model="product_code"/>
                    <div class="character-count">
                        | {{ product_code_length }} / 8
                    </div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Product Name</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Product Name" maxLength="20" v-model="product_name"/>
                    <div class="character-count ">
                        | {{ product_name_length }} / 20
                    </div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Purchase</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Price" v-model="purchase"/>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Product Description</label>
                </div>
                <div class="w80">
                    <textarea style="width: 1067px; height: 86px;" placeholder="Masukkan Product Description" v-model="productdescription"></textarea>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Unit</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Unit" v-model="unit"/>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Min Purchase</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Masukkan Min Purchase" v-model="minpurchase"/>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Max Purchase</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Masukkan Max Purchase" v-model="maxpurchase"/>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Category</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Pilih Category" v-model="category"/>
                </div>
            </div>
            <div class="flex mt30">
                <router-link to="/list-product">
                    <button class="bt-submit-green" @click="onAddProduct">Add Product</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'add-product',
    data(){
        return {
            product_name: "",
            product_code: "",
            purchase: "",
            productdescription: "",
            unit: "",
            minpurchase: "",
            maxpurchase: "",
            category: "",
            imageUrl: null, 
            imageData: null,
            product_code_length: 0,
            product_name_length: 5,
        }
    },
    methods: {
        ...mapActions([
			'addProduct',
		]),
        onAddProduct(){
            this.addProduct({
                product_code: this.product_code,
                product_name: this.product_name,
                purchase: parseInt(this.purchase),
                productdescription: this.productdescription,
                unit: parseInt(this.unit),
                minpurchase: parseInt(this.minpurchase),
                maxpurchase: parseInt(this.maxpurchase),
                category: this.category,
                image: this.imageData
            })
        },
        handleInput() {
            this.product_code_length = this.product_code.length;
            this.product_name_length = this.product_name.length;
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
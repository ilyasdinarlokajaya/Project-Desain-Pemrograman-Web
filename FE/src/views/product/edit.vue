<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-product">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Update Product</h1>
                    <p class="mt5 mb20 tc-green">Make your product looking good for customer</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Product Image</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <img v-else :src="product.image" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label ">Product Code</label>
                </div>
                <div class="w80">
                    <input
                        disabled
                        class="form-text"
                        type="text"
                        value="Data Tersimpan"
                        v-model="product.product_code"
                    >
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Price</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Product Price"
                        v-model="product.purchase"
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removePrice"></div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Product Description</label>
                </div>
                <div class="w80">
                    <textarea style="width: 1067px; height: 86px;" placeholder="Masukkan Product Description" v-model="product.productdescription"></textarea>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Unit</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Unit" v-model="product.unit"/>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeUnit"></div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Min Purchase</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Masukkan Min Purchase" v-model="product.minpurchase"/>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeMinPurchase"></div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Max Purchase</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Masukkan Max Purchase" v-model="product.maxpurchase"/>
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeMaxPurchase"></div>
                </div>
            </div>
            <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Category</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder="Pilih Category" v-model="product.category"/>
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-product">
                    <button class="bt-submit-green" @click="onEditProduct">Update Product</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-notification',
    data(){
        return {
            imageUrl: null,
            imageData: null
        }
    },
	computed: {
        ...mapState({ // get data from state
			product: state => state.product.detail
		})
	},
    mounted() {
        this.detailProduct(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailProduct',
            'editProduct'
		]),
        onEditProduct() {
            this.editProduct({
                id: this.product.id,
                product_code: this.product.product_code,
                product_name: this.product.product_name,
                purchase: this.product.purchase,
                productdescription: this.product.productdescription,
                unit: this.product.unit,
                minpurchase: this.product.minpurchase,
                maxpurchase: this.product.maxpurchase,
                category: this.product.category,
                image: this.imageData != null? this.imageData : this.product.image
            })
        },
        removePrice() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.product.purchase = ""
        },
        removeUnit() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.product.unit = ""
        },
        removeMinPurchase() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.product.minpurchase = ""
        },
        removeMaxPurchase() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.product.maxpurchase = ""
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
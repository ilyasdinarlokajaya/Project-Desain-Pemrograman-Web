<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-voucher">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Update Voucher</h1>
                    <p class="mt5 mb20 tc-green">Make your customer shop using voucher</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Voucher Image</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 200px; height: 100px;" />
                    <img v-else :src="voucher.image" alt="Uploaded Image" style="max-width: 200px; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Voucher Code</label>
                </div>
                <div class="w80">
                    <input 
                        disabled
                        id="txt" 
                        :maxlength="6"
                        class="form-text txt" 
                        type="text" 
                        placeholder="Input your voucher code"
                        v-model="voucher.vouchercode"
                    >        
                    <div class="p-kanan" >| {{ voucher.vouchercode.length }}/6</div>  
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Voucher Name</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input your voucher name"
                        v-model="voucher.vouchername"
                    >
                    
                    <div class="p-kanan" >| {{ voucher.vouchername.length }}/20</div> 
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Description</label>
                </div>
                <div class="w80">
                    <textarea 
                        :maxlength="500"
                        class=" form-text description-textarea" 
                        placeholder="Input your voucher description"
                        v-model="voucher.description"
                    >{{voucher.description}}</textarea>
                    <div class="p-kanan p-kanan-desk" >{{voucher.description.length}} | 500</div>                        
                </div>
            </div>
            <div>
                <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Type</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="voucher.type">
                        <option disabled selected>--- Select Type ---</option>
                        <option value="1">ONGKOS KIRIM</option>
                        <option value="2">FLASH SALE</option>
                        <option value="3">TAHUN BARU</option>
                    </select>
                    <div class="p-kanan" ><img class="mt4" src="../../assets/icons/pen-icon.png" alt="" height="17px"></div>
                </div>
            </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Min. Transaction</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input minimum transaction"
                        v-model="voucher.mintransaction"
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeTransaction"></div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Amount</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder=" Input amount"
                        v-model="voucher.amount"
                    >
                    <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeAmount"></div>
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
                        placeholder="Input date"
                        v-model="voucher.validfrom"
                    >
                    <!-- <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeValidFrom"></div> -->
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
                        placeholder="Input date"
                        v-model="voucher.validuntil"
                    >
                    <!-- <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%" @click="removeValidUntil"></div> -->
                </div>
            </div>
            <div class="flex mt30">
                <router-link to="/list-voucher">
                    <button class="bt-submit-green" @click="onEditVoucher">Update Voucher</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
    name: 'edit-voucher',
    data(){
        return{
            imageUrl: null,
            imageData: null,
            letterCount: 0,
            letterCount2: 0,
            letterCount3: 0,
        }
    },
    computed: {
        ...mapState({ // get data from state
			voucher: state => state.voucher.detail
		})
	},
    mounted() {
        this.detailVoucher(this.$route.params.id)

    },
    methods: {

        ...mapActions([
			'detailVoucher',
            'editVoucher'
		]),
        onEditVoucher() {
            this.editVoucher({
                id: this.voucher.id,
                vouchercode: this.voucher.vouchercode,
				vouchername: this.voucher.vouchername,
				description: this.voucher.description,
				type: this.voucher.type,
				mintransaction: parseInt(this.voucher.mintransaction),
				amount: parseInt(this.voucher.amount),
				validfrom: this.voucher.validfrom,
				validuntil: this.voucher.validuntil,
				image: this.imageData != null? this.imageData : this.voucher.image
            })
        },
        
        removeTransaction() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.voucher.mintransaction = ""
        },
        removeAmount() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            
            this.voucher.amount =""
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

<style>
.xPojok{
  text-align: right;
} 

.p-kanan{
    font-family: sans-serif;
    z-index: 1;
    margin-top: -25px;
    float: right;
    margin-right: 10px;
}


.p-kanan-desk{
    margin-top: -110px;
}

input[type="file"] {
   content: "";
   opacity: 0;
   width: 80px;
   height: 80px;
}

.voucher-image{ 
    display: block;
    width: 10%;
    padding: 10px;
    height: 80px;
    border-radius: 5px;
    border: 1px solid #ddd;
}

.gg{
    background-image: url('../../assets/icons/addimg.jpeg');
    background-repeat: no-repeat;
    background-size: cover;
}

.description-textarea {
    width: 64,5em; /* Lebar textarea */
    height: 100px; /* Tinggi textarea */
    padding: 10px; /* Padding untuk jarak teks dari tepi textarea */
    user-select: none;
    resize: none;
    font-family: sans-serif;
}
</style>
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
                    <h1 class="m0">Add Voucher</h1>
                    <p class="mt5 mb20 tc-green">make your customer shop using voucher</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Voucher Image</label>
                </div>
                <div class="w80">           
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <button v-if="imageUrl === null" class="bt-image" @click="openFileInput"><img :src="imageUrl || defaultImageUrl"></button>
                    <div v-if="imageUrl">
                        <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="width: 200px; height: auto;" />
                        <button class="bt-change-image" @click="openFileInput">Change Image</button>
                    </div>
                </div>
                
            </div>
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Voucher code</label>
                </div>
                <div class="w80">
                        <input 
                            id="txt" 
                            v-model="inputText" :maxlength="6"
                            class="form-text txt" 
                            type="text" 
                            placeholder="Input your voucher code" >
                            <div class="p-kanan" >| {{ letterCount }}/6</div>     
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Voucher Name</label>
                </div>
                <div class="w80">
                    <input 
                            id="txt" 
                            v-model="inputText2" :maxlength="20"
                            class="form-text txt" 
                            type="text" 
                            placeholder="Input your voucher name" >
                            <div class="p-kanan" >| {{ letterCount2 }}/20</div>   
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Description</label>
                </div>
                <div class="w80">
                    <textarea 
                    v-model="inputText3" :maxlength="500"
                    class=" form-text description-textarea" 
                    placeholder="Input your voucher description"></textarea>
                    <div class="p-kanan p-kanan-desk" >{{ letterCount3 }} | 500</div>   
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Type</label>
                </div>
                <div class="w80">        
                    <select class="form-select offappea" v-model="type">
                        <option disabled selected>--- Select Type ---</option>
                        <option value="1">ONGKOS KIRIM</option>
                        <option value="2">FLASH SALE</option>
                        <option value="3">TAHUN BARU</option>
                    </select>
                    <div class="p-kanan" ><img class="" src="../../assets/icons/pen-icon.png" alt="" height="17px"></div>   
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
                        v-model="minTransaction"
                    >
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
                        placeholder="Input amount"
                        v-model="amount"
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
                        placeholder="Input date"
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
                        placeholder="Input date"
                        v-model="validUntil"
                    >
                </div>
            </div>
            <div class="flex mt30 bawah">
                <router-link to="/list-voucher">
                    <button class="bt-submit-green" @click="onClickAddVoucher">Add Voucher</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

    export default {
        name: 'add-voucher',

        data(){
            return{
                inputText: "",
                hasil: 0,
                inputText2: "",
                inputText3: "",
                letterCount: 0,
                letterCount2: 0,
                letterCount3: 0,
                file:'',
                imageUrl: null,
                imageData: null,
                defaultImageUrl: require('../../assets/icons/addimg.jpeg'),
                type: null,
                minTransaction:"",
                amount:"",
                validFrom:"",
                validUntil:""
                
            };
        }, 
        watch: {
            inputText() {
                this.countLetters();          
            },
            inputText2() {
                this.countLetters2();
            },
            inputText3() {
                this.countLetters3();
            }
        },
        methods: {
            ...mapActions([
			'addVoucher',
            'createVoucher'
		    ]),
            onClickAddVoucher(){
                this.createVoucher({
                    vouchercode: this.inputText,
                    vouchername: this.inputText2,
                    description: this.inputText3,
                    type: this.type,
                    minTransaction: parseInt(this.minTransaction),
                    amount: parseInt(this.amount),
                    validfrom: this.validFrom,
                    validuntil: this.validUntil,
                    image: this.imageData
                })
            },
            changeFile() {
                let file = e.target.files[0];
                this.file = file;
                console.log(this.file);
            },
            countLetters() {
                this.letterCount = this.inputText.length;   
            },
            countLetters2() {
                this.letterCount2 = this.inputText2.length; 
            },
            countLetters3() {
                this.letterCount3 = this.inputText3.length; 
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
                        console.log(this.imageData)
                    };
                    reader.readAsDataURL(file);
                }
            },
        }
    };

</script>


<style>
.border{    
  background-size: cover;
  margin: 0 auto;
}

.p-kanan{
    font-family: sans-serif;
    z-index: 1;
    margin-top: -30px;
    float: right;
    margin-right: 20px;
}


.p-kanan-desk{
    margin-top: -110px;
}

/*  My codee   */
.label-merahdarah {
    background-color: #e81313; 
    padding: 10px 20px; 
    border-radius: 5px; 
    font-size: 14px; 
    color: #014935; 
    font-weight: 500;
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
    /* background-image: url('../../assets/icons/addimg.jpeg'); */
    /* background-repeat: no-repeat; */
    background-size: cover;
    
}

.offappea{
    appearance: none;
}

.bawah{
    margin-bottom: 20px;
}

.hgh{
    text-align: left;   
    padding-bottom: 200px; 
}

input .hgh::placeholder{
    text-align: start;
}

.text-box {
    width: 200px; /* Lebar kotak teks */
    height: 100px; /* Tinggi kotak teks */
    border: 1px solid #000; /* Tampilan garis kotak */
    text-align: left; /* Mengatur teks ke kiri */
    padding: 10px; /* Padding untuk jarak teks dari tepi kotak */
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
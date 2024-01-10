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
                    <h1 class="m0">Detail Voucher</h1>
                    <p class="mt5 mb20 tc-green">Your voucher details</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 right">
                <img class="imgdetail" :src="voucher.image" >
                <div class="mp ">
                    <p class="mp abu">{{ voucher.vouchercode }}</p>

                    <h2 class="mp">{{ voucher.vouchername }}</h2>

                    <p class="detail-title-bold mp">Voucher Description</p>
                    <P class="mp abu desk">{{ voucher.description }}</P>
                    
                    <div class="container">
                        <div class="in">
                            <p class="detail-title-bold space">Voucher Category : </p>
                            <p class="detail-title-thin" v-if="voucher.type === '1'" >ONGKOS KIRIM</p>
                            <p class="detail-title-thin" v-if="voucher.type === '2'" >TAHUN BARU</p>
                            <p class="detail-title-thin" v-if="voucher.type === '3'" >FLASH SALE</p>
                        </div>
                        <div class="in">
                            <p class="detail-title-bold space">Min Transaction : </p>
                            <p class="abu">Rp. {{ voucher.mintransaction }}</p>
                        </div>
                        <div class="in">
                            <p class="detail-title-bold space">Amount : </p>
                            <p class="abu">Rp. {{ voucher.amount }}</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="exp">
                <div class="container2">
                    <p class="detail-title-bold">Voucher Status</p>
                    <hr class="tebal">
                    <label class="label-green c" v-if=" date() >= new Date(voucher.validfrom) && date() <= new Date(voucher.validuntil)">Active</label>
                    <label class="label-merahdarah" v-else>Inactive</label>
                </div>
                <div class="container2">
                    <p class="detail-title-bold">Valid From</p>
                    <hr class="tebal">
                    <label class="label-green c">{{ voucher.validfrom }}</label>
                </div>
                <div class="container2">
                    <p class="detail-title-bold">Valid Until</p>
                    <hr class="tebal">
                    <label class="label-green c">{{ voucher.validuntil }}</label>
                </div>
            </div>
        </div>
        <div class="admin container">
            <p>👤 Create by Admin 1<br>👤 Update by Admin 2</p>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-voucher',
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
			'detailVoucher'
		]),
        
        date(){
            return new Date();
        }
    }
}

</script>

<style>


.exp{
    display: flex;   
}

.c{
    width:100px;
    display: flex;
}

body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
}

.tebal{
    border: 1px solid black;
    width: 100%;
    margin: 0;
    margin-bottom: 16px;
    margin-top: -10px;
}

.container2 {
    margin-right: 60px;
   
}

.container {
    margin-left: 30px;
   
}

.space{
    margin-right: 10px;
}

.desk{
    line-height: 2;
}

.in {
    display: flex;
    margin-bottom: -20px;
}

.right{
    float: right;
    text-align: left;
    display: flex;
}

.mp{
    margin-left: 30px;
    margin-bottom: 20px;
}

.abu{
    color: darkgray;
}

.admin{
    margin-top: 40px;
    
}

.label-merahdarah {
    background-color: #e81313; 
    padding: 10px 20px; 
    border-radius: 5px; 
    font-size: 14px; 
    color: #014935; 
    font-weight: 500;
}
</style>
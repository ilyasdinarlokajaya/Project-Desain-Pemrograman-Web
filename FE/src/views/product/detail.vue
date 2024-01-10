
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
                    <h1 class="m0">Detail Product</h1>
                    <p class="mt5 mb20 tc-green">Your product details</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="detailgrid">
                <div class="maindvg">
                    <div class=" ">
                        <img class="imgdetail" height="250" :src="product.image" alt="Uploaded Image" />
                    </div>
                    <div class="belowmaingrid">
                        <div class=" belowkiri">
                            <p class="detail-bold">Status:</p> <hr class="tebal">
                            <p class="label-green">Publish</p>
                        </div>
                        <div class="belowtengah">
                            <p class="detail-bold">Product Status:</p><hr class="tebal">
                            <div v-if="product.unit > 0">
                                <p class="label-green">{{ "Available" }}</p>
                            </div>
                            <div v-if="product.unit < 0">
                                <p class="label-red">{{ "Unavailable" }}</p>
                            </div>
                        </div>
                    </div>
                    <p>👤 Create by Admin 1<br>👤 Update by Admin 2</p>
                </div>
                <div class="asidedvg">
                        <div class="detail-title-thin">
                        <p>{{ product.product_code }}</p>
                        </div>
                        <div class="mt40">
                            <p class="detail-title-bold">{{ product.product_name }}</p>
                        </div>
                        <div class="mt40">
                            <p class="detail-title-bold">Rp. {{ product.purchase }}</p>
                        </div>
                        <div class="mt40">
                            <p >
                            <span class="center detail-title-bold mr10">Product Description: </span>
                            <span class="center detail-title-thin ">{{ product.productdescription }}</span>
                            </p>
                        </div>
                        <div class="mt40">
                            <p>
                            <span class="detail-title-bold mr10">Product Category: </span> <span class="detail-title-thin">{{ product.category }}</span>
                            </p>
                        </div>
                        <div class="mt40">
                            <p>
                            <span class="detail-title-bold mr10">Unit: </span> <span class="detail-title-thin">{{ product.unit }} Unit</span>
                            </p>
                        </div>
                        <div class="mt40">
                            <p>
                            <span class="detail-title-bold mr10">Min Purchase: </span> <span class="detail-title-thin">{{ product.minpurchase }} Unit</span>
                            </p>
                        </div>
                        <div class="mt40">
                            <p>
                            <span class="detail-title-bold mr10">Max Purchase: </span> <span class="detail-title-thin">{{ product.maxpurchase }} Unit</span>
                            </p>
                        </div>
                    </div>  
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-product',
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
			'detailProduct'
		])
    }
}
</script>
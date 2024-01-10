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
                    <h1 class="m0">Detail Delivery Method</h1>
                    <p class="mt5 mb20 tc-green">Your delivery method details</p>
                </div>
            </div>
           <hr class="uline-grey">
           <main>
            <div class="center box1">
                <img class="rectangle" height="250" :src="delivery_method.image" alt="Uploaded Image" />
            </div>
            <div class="content1">
                <p class="detail-title-thin">Delivery Name:</p>
                <h2>{{ delivery_method.delivery_name }}</h2>
            </div>
            <div class="content2">
                <p class="detail-title-thin ">Delivery Fee:</p>
                <h2>{{ delivery_method.delivery_fee }}</h2>
            </div>
            <div class="content3">
                <p class="detail-title-thin ">Estimation Delivery:</p>
                <h2>{{ delivery_method.estimation }}</h2>
            </div>
        
            </main>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-delivery-method',
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
			'detailDeliveryMethod'
		])
    }
}
</script>
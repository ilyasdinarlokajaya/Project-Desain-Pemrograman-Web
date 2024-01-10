<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-banner">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Detail Banner</h1>
                    <p class="mt5 mb20 tc-green">Your banner details</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="detail-banner-grid">
                <div class="main-banner">
                    <div class="center mt20">
                        <img class="imgdetail" height="250" :src="banner.image" alt="Uploaded Image" />
                    </div>
                    <div class="center">
                        <h2>{{ banner.bannername }}</h2>
                    </div>
                    <div class="center">
                        <p class="detail-title-bold">Broadcast Day:
                            <label class="detail-title-thin">{{ banner.broadcastday }}</label>
                        </p>
                    </div>
                    <p class="center">👤 Create by Admin 1<br>👤 Update by Admin 2</p>
                    <div class="container1">
                        <div class="test">
                            <div class="center">
                                <p class="detail-title-bold">Valid From:</p><hr class="bold-banner">
                                <label class="label-green">{{ banner.validfrom }}</label>
    
                            </div>
                        </div>
                        <div class="test">
                            <div class="center">
                                <p class="detail-title-bold">Valid Until:</p><hr class="bold-banner">
                                <label class="label-green">{{ banner.validuntil }}</label>
    
                            </div>
                        </div>
                        <div class="test">
                            <div class="center">
                                <p class="detail-title-bold">Banner status</p><hr class="bold-banner">
                                <div v-if="banner.validuntil >= banner.validfrom">
                                    <label class="label-green">{{ banner.bannerstatus  = "Active"}}</label>
                                </div>
                                <div v-else="banner.validuntil <= banner.validfrom">
                                    <label class="label-red">{{ banner.bannerstatus = "Inactive" }}</label>
                                </div>
                            </div>
                    </div>
                </div>
            </div>

            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-banner',
	computed: {
        ...mapState({ // get data from state
			banner: state => state.banner.detail
		})
	},
    mounted() {
        this.detailBanner(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailBanner'
		])
    }
}
</script>

<style>
 .test{
    width: 220px;
    height: -10px;
} 
.container1 {
    display: flex;
    flex-direction: row;
    /* flex-wrap: wrap; */
    justify-content: center; /* Mengatur container agar berada di tengah secara horizontal */
    align-items: center; /* Mengatur container agar berada di tengah secara vertikal */
}

</style>
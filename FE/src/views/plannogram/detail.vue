<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-plannogram">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Detail Plannogram</h1>
                    <p class="mt5 mb20 tc-green">Your plannogram details</p>
                </div>
            </div>
           <hr class="uline-grey">
           <!-- <main>
            <div class="center box1">
                <img class="rectangle" height="250" :src="plannogram.image" alt="Uploaded Image" />
            </div>
            <div class="content1">
                <p class="detail-title-thin ">Plannogram Name:</p>
                <h2>{{ plannogram.plannogramname }}</h2>
            </div>
            <div class="content2">
                <p class="detail-title-bold">Broadcast Day:
                <label v-if="parseInt(plannogram.broadcastday) === 1" class="detail-title-thin">Senin</label>
                <label v-if="parseInt(plannogram.broadcastday) === 2" class="detail-title-thin">Selasa</label>
                <label v-if="parseInt(plannogram.broadcastday) === 3" class="detail-title-thin">Rabu</label>
                <label v-if="parseInt(plannogram.broadcastday) === 4" class="detail-title-thin">Kamis</label>
                <label v-if="parseInt(plannogram.broadcastday) === 5" class="detail-title-thin">Jumat</label>
                <label v-if="parseInt(plannogram.broadcastday) === 6" class="detail-title-thin">Sabtu</label>
                <label v-if="parseInt(plannogram.broadcastday) === 7" class="detail-title-thin">Minggu</label>
                </p>
            </div>
            <div class="content4 ">
                <p class="detail-title-bold uline-black">Valid From:</p>
                <label Class="label-green">{{ plannogram.validfrom }}</label>
            </div>
            <div class="content5">
                <p class="detail-title-bold uline-black">Valid Until:</p>
                <label Class="label-green">{{ plannogram.validuntil }}</label>
            </div>
             <div class="content3">
                <p class="detail-title-bold uline-black">Plannogram Status:</p>
                <div v-if="plannogram.validuntil >= plannogram.validfrom">
                    <label class="label-green center">{{ plannogram.plannogram_status = "Active" }}</label>
                </div>
                <div v-else="plannogram.validuntil <= plannogram.validfrom">
                    <label class="label-red">{{ plannogram.plannogram_status = "Inactive" }}</label>
                </div>
            </div>
            </main> -->
            <div class="detail-plannogram-grid">
                <div class="maindvg">
                    <div class=" ">
                        <img class="imgdetail-plannogram" :src="plannogram.image" alt="Uploaded Image" />
                    </div>
                </div>
                <div class="asidedvg">
                    <div class="mt10">
                        <p class="detail-title-thin">Plannogram Name:</p>
                        <h2>{{ plannogram.plannogramname }}</h2>
                    </div>
                    <div class="mt10">
                        <p class="detail-title-bold">Broadcast Day:
                            <label v-if="parseInt(plannogram.broadcastday) === 1" class="detail-title-thin">Senin</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 2" class="detail-title-thin">Selasa</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 3" class="detail-title-thin">Rabu</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 4" class="detail-title-thin">Kamis</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 5" class="detail-title-thin">Jumat</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 6" class="detail-title-thin">Sabtu</label>
                            <label v-if="parseInt(plannogram.broadcastday) === 7" class="detail-title-thin">Minggu</label>
                        </p>
                        <div class="belowmaingrid">
                            <div class=" belowkiri">
                                <p class="detail-title-bold">Valid Until:</p><hr class="bold-plannogram">
                                <label Class="label-green">{{ plannogram.validuntil }}</label>
                            </div>
                            <div class="belowtengah">
                                <p class="detail-title-bold">Valid From:</p><hr class="bold-plannogram">
                                <label class="label-green">{{ plannogram.validfrom }}</label>
                            </div>
                            <div class="belowkanan">
                                <p class="detail-title-bold">Plannogram Status:</p><hr class="bold-plannogram">
                                <div v-if="plannogram.validuntil >= plannogram.validfrom">
                                    <label class="label-green">{{ plannogram.plannogram_status = "Active" }}</label>
                                </div>
                                <div v-else="plannogram.validuntil <= plannogram.validfrom">
                                    <label class="label-red">{{ plannogram.plannogram_status = "Inactive" }}</label>
                                </div>
                            </div>
                        </div>
                        <p>👤 Create by Admin 1<br>👤 Update by Admin 2</p>
                        <p class="mt30 detail-title-bold">List Product</p>
                        <ul>
                            <li class="detail-title-thin" v-if="plannogram.listproduct === '1'">Elektronik</li>
                            <li class="detail-title-thin" v-if="plannogram.listproduct === '2'">Fashion</li>
                        </ul>
                    </div>
                    
                </div>  
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-plannogram',
	computed: {
        ...mapState({ // get data from state
			plannogram: state => state.plannogram.detail
		})
	},
    mounted() {
        this.detailPlannogram(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailPlannogram'
		])
    }
}
</script>
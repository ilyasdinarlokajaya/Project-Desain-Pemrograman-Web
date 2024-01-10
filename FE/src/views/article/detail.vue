<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-article">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Detail Article</h1>
                    <p class="mt5 mb20 tc-green">Your article details</p>
                </div>
            </div>
           <hr class="uline-grey">
        <div class="detail-plannogram-grid">
                <div class="maindvg">
                    <div>
                        <img class="imgdetail-plannogram" height="250" :src="article.image" alt="Uploaded Image" />
                    </div>
                    <div class="belowmaingrid">
                        <div class=" belowkiri">
                            <p class="detail-title-bold">Article Status:</p><hr class="bold-plannogram">
                            <div v-if="article.valid_until >= article.valid_from">
                                <label class="label-green center">{{ article.article_status = "Active" }}</label>
                            </div>
                            <div v-else="article.valid_until < article.valid_from">
                                <label class="label-red">{{ article.article_status = "Inactive" }}</label>
                            </div>
                        </div>
                        <div class="belowtengah">
                            <p class="detail-title-bold">Valid From:</p><hr class="bold-plannogram">
                            <label class="label-green">{{ article.valid_from }}</label>
                        </div>
                        <div class="belowkanan"> 
                            <p class="detail-title-bold">Valid Until:</p><hr class="bold-plannogram">
                            <label class="label-green">{{ article.valid_until }}</label>
                        </div>
                    </div>
                    <p>👤 Create by Admin 1<br>👤 Update by Admin 2</p>
                </div>
                <div class="asidedvg">
                        <div class="mt40">
                            <p class="detail-title-bold">Broadcast Day:
                            <label v-if="parseInt(article.jadwal_broadcast) === 1" class="detail-title-thin">Senin</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 2" class="detail-title-thin">Selasa</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 3" class="detail-title-thin">Rabu</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 4" class="detail-title-thin">Kamis</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 5" class="detail-title-thin">Jumat</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 6" class="detail-title-thin">Sabtu</label>
                            <label v-if="parseInt(article.jadwal_broadcast) === 7" class="detail-title-thin">Minggu</label>
                            </p>
                        </div>
                        <div class="mt40">
                            <p class="detail-title-thin">Article Name:</p>
                            <h2>{{ article.judul_artikel }}</h2>
                        </div>
                        <div class="mt40">
                            <p class="detail-title-bold ">Content:</p>
                            <p class="detail-title-thin">{{ article.konten }}</p>
                        </div>
                        
                    </div>  
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-article',
	computed: {
        ...mapState({ // get data from state
			article: state => state.article.detail
		})
	},
    mounted() {
        this.detailArticle(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailArticle'
		])
    }
}
</script>

<style>
main{
    display: grid;
    height: 80vh;
    grid-template-columns: 0.4fr 0.4fr 0.4fr 0.4fr;
    grid-template-rows: 40px 40px 40px 40px; 
    grid-template-areas:
        "foto header header header"
        "foto content1 content1 content1"
        "foto content2 content2 content2"
        "foto content3";
    grid-gap: 50px;
}
 .test{
    width: 547px;
    height: 494px;
} 
.container1{
    display: flex;
    flex-direction: row;
    /* justify-content: space-evenly; */
    align-items: flex-start;
    gap: 1em;
}
</style>
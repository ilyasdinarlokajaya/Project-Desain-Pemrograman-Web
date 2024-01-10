<template>
    <div>
        <Menu/>
        <div class="container">
            <h1 class="m0">Article</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your Article</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-article">
                        <button class="bt-green">+ Add Article</button>
                    </router-link>
                </div>
            </div>
            <div class="row bg-blue mt10 p10">
                <div class="col">
                    <input
                        class="search-field"
                        type="text"
                        id="deskripsi"
                        placeholder="Search by invoice number, name, amount..."
                        v-model="search"
                    />
                </div>
                <div class="col flex">
                    <button @click="onSearchArticle" class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>ARTICLE NAME</th>
							<th>BROADCAST SCHEDULE</th>
							<th>VALID FROM</th>
							<th>VALID UNTIL</th>
							<th>ARTICLE STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_article" :key="item.id">
                            <td class="center">{{ item.judul_artikel }}</td>
                            <td class="center">
                                <div v-if="parseInt(item.jadwal_broadcast) === 1">
                                    <label class="label-blue">Senin</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 2">
                                    <label class="label-green">Selasa</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 3">
                                    <label class="label-red">Rabu</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 4">
                                    <label class="label-blue">Kamis</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 5">
                                    <label class="label-yellow">Jumat</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 6">
                                    <label class="label-red">Sabtu</label>
                                </div>
                                <div v-if="parseInt(item.jadwal_broadcast) === 7">
                                    <label class="label-green">Minggu</label>
                                </div>
                            </td>
                            <td class="center">
                                <label class="label-green">{{ item.valid_from }}</label>
                            </td>
                            <td class="center">
                                <label class="label-green">{{ item.valid_until }}</label>
                            </td>
                            <td class="center">
                                <div v-if=" date() >= new Date(item.valid_from) && date() <= new Date(item.valid_until)">
                                    <label class="label-green">{{ item.article_status = "Active" }}</label>
                                </div>
                                <div v-else>
                                    <label class="label-red">{{ item.article_status = "Inactive" }}</label>
                                </div>
                            </td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link :to="{ name: 'edit-article', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-article', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteArticle(item.id)" src="../../assets/icons/delete.svg">
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script>
import Menu from "../../components/menu.vue"
import { mapActions, mapState } from 'vuex'

export default {
	name: 'list-article',
    data(){
        return {
            search: ""
        }
    },
	components: {
		Menu
	},
    computed: {
        ...mapState({ // get data from state
			list_article: state => state.article.list
		})
	},
	mounted() {
		this.listArticle() // render items for the first
	},
	methods: {
		...mapActions([
			'listArticle',
            'deleteArticle',
		]),
        onSearchArticle(){
            console.log("asdasd")
            this.$store.commit('setSearchArticle', this.search)
            this.listArticle()
        }, 
		date(){
            return new Date();
        }
	},
}
</script>
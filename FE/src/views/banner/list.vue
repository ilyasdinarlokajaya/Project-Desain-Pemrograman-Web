<template>
    <div>
        <Menu/>
        <div class="container">
            <h1 class="m0">Banner</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your Banner to atrract visitor's attention</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-banner">
                        <button class="bt-green">+ Add Banner</button>
                    </router-link>
                </div>
            </div>
            <div class="row bg-blue mt10 p10">
                <pre>{{ search }}</pre>
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
                    <button @click="onSearchBanner" class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>BANNER NAME</th>
                            <th>BROADCAST SCHEDULE</th>
                            <th>VALID FROM</th>
                            <th>VALID UNTIL</th>
                            <th>BANNER STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_banner" :key="item.id">
                            <td class="center">{{ item.bannername }}</td>
                            <td class="center">{{ item.broadcastday }}</td>
                            <td class="center">
                                <label class="label-green">{{ item.validfrom }}</label>
                            </td>
                            <td class="center">
                                <label class="label-green">{{ item.validuntil }}</label>
                            </td>
                            <td class="center">
                                <div v-if=" date() >= new Date(item.validfrom) && date() <= new Date(item.validuntil)">
                                    <label class="label-green">{{ item.bannerstatus = "Active" }}</label>
                                </div>
                                <div v-else>
                                    <label class="label-red">{{ item.bannerstatus = "Inactive" }}</label>
                                </div>
                            </td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link :to="{ name: 'edit-banner', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-banner', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteBanner(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-banner',
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
			list_banner: state => state.banner.list
		})
	},
	mounted() {
		this.listBanner() // render items for the first
	},
	methods: {
		...mapActions([
			'listBanner',
            'deleteBanner',
		]),
        onSearchBanner(){
            console.log("asdasd")
            this.$store.commit('setSearchBanner', this.search)
            this.listBanner()
        },
        date(){
            return new Date();
        }
	},
}
</script>
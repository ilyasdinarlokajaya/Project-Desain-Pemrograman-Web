<template>
    <div>
        <Menu/>
        <div class="container">
            <h1 class="m0">Plannogram</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your Plannogram</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-plannogram">
                        <button class="bt-green">+ Add Plannogram</button>
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
                    />
                </div>
                <div class="col flex">
                    <button class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>PLANNOGRAM NAME</th>
                            <th>BACKGROUND IMAGE</th>
                            <th>BROADCAST SCHEDULE</th>
                            <th>VALID FROM</th>
                            <th>VALID UNTIL</th>
                            <th>PLANNOGRAM STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_plannogram" :key="item.id">
                            <td class="center">{{ item.plannogramname }}</td>
                            <td class="center box">
								<img class="rectangle" :src="item.image" alt="Product Image" />
							</td>
                            <td class="center">
                                <div v-if="parseInt(item.broadcastday) === 1">
                                    <label class="label-blue">Senin</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 2">
                                    <label class="label-green">Selasa</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 3">
                                    <label class="label-red">Rabu</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 4">
                                    <label class="label-blue">Kamis</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 5">
                                    <label class="label-yellow">Jumat</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 6">
                                    <label class="label-red">Sabtu</label>
                                </div>
                                <div v-if="parseInt(item.broadcastday) === 7">
                                    <label class="label-green">Minggu</label>
                                </div>
                            </td>
                            <td class="center"><label Class="label-green">{{ item.validfrom }}</label></td>
                            <td class="center"><label Class="label-green">{{ item.validuntil }}</label></td>
                            <td class="center">
                                <div v-if="item.validuntil >= item.validfrom">
                                    <label class="label-green">{{ item.plannogram_status = "Active" }}</label>
                                </div>
                                <div v-else="item.validuntil <= item.validfrom">
                                    <label class="label-red">{{ item.plannogram_status = "Inactive" }}</label>
                                </div>
                            </td>
                            <td class="center ">
                                <div class="mt5" >
                                    <router-link :to="{ name: 'edit-plannogram', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-plannogram', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deletePlannogram(item.id)" src="../../assets/icons/delete.svg">
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
import Menu from "../../components/menu.vue";
import { mapActions, mapState } from "vuex";
export default {
  name: "list-plannogram",
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
			list_plannogram: state => state.plannogram.list
		})
	},
	mounted() {
		this.listPlannogram() // render items for the first
	},
	methods: {
		...mapActions([
			'listPlannogram',
            'deletePlannogram',
		]),
        onSearchPlannogram(){
            console.log("asdasd")
            this.$store.commit('setSearchPlannogram', this.search)
            this.listPlannogram()
        }
	},
};
</script>


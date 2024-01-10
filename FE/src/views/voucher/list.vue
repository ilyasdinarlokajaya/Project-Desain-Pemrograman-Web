<template>
    <div>
        <Menu/>
		<div class="container">
            <h1 class="m0">Voucher</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your voucher promotion</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-voucher">
                        <button class="bt-green">+ Add</button>
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
                            <th>VOUCHER CODE</th>
                            <th>VOUCHER NAME</th>
                            <th>VALID FROM</th>
                            <th>VALID UNTIL</th>
                            <th>VOUCHER STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
						<tr v-for="item of list_voucher" :key="item.id">
                            <td class="center">{{ item.vouchercode }}</td>
                            <td class="center" >{{ item.vouchername }}</td>
                            <td class="center" ><label class="label-green">{{ item.validfrom }}</label></td>
                            <td class="center" ><label class="label-green">{{ item.validuntil }}</label></td>
                            <!-- <img :src="item.image" alt="" srcset="" width="70px"  class=""> -->
                            <td class="center">
                                <div v-if=" new Date(item.validuntil) >= new Date(item.validfrom)">
                                    <label class="label-green">Active</label>
                                </div>
                                <div v-else>
                                    <label class="label-merahdarah">Inactive</label>
                                </div>
                            </td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link :to="{ name: 'edit-voucher', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-voucher', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteVoucher(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-voucher',
	components: {
		Menu
	},
    data(){
        return {
            search: ""
        }
    },
    computed: {
        ...mapState({ // get data from state
			list_voucher: state => state.voucher.list
		})
	},
	mounted() {
		this.listVoucher() // render items for the first
	},
	methods: {
		...mapActions([
			'listVoucher',
            'deleteVoucher',
		]),
        onSearchUser(){
            console.log("asdasd")
            this.$store.commit('setSearchVoucher', this.search)
            this.listVoucher()
        },
        date(){
            return new Date();
        }
	}
}
</script>

<style>

.label-merahdarah {
    background-color: #e81313; 
    padding: 10px 20px; 
    border-radius: 5px; 
    font-size: 14px; 
    color: #014935; 
    font-weight: 500;
}

</style>
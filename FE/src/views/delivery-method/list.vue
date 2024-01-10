<template>
	<div>
		<Menu />
		<div class="container">
			<h1 class="m0">Delivery Method</h1>
			<div class="row">
				<div class="col">
					<p class="mt5 mb20 tc-green">Manage your Delivery Methods for the convenience of your customers</p>
				</div>
				<div class="col flex">
					<router-link to="/add-delivery-method">
						<button class="bt-green">+ Add</button>
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
                        placeholder="Search by Name"
                        v-model="search"
                    />
                </div>
                <div class="col flex">
                    <button @click="onSearchDeliveryMethod" class="bt-blue">Filter</button>
                </div>
            </div>
			<div class="mt50">
				<table>
					<thead>
						<tr>
							<th style="font-size: 14px;">DELIVERY METHOD</th>
							<th style="font-size: 14px;">DELIVERY FEE</th>
							<th style="font-size: 14px;">ESTIMATION DELIVERY</th>
							<th style="font-size: 14px;">DELIVERY TYPE IMAGE</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="item of list_delivery_method" :key="item.id">
							<td class="center">{{ item.delivery_name }}</td>
							<td class="center"><p class="detail-title-bold">Rp. {{ item.delivery_fee }}</p></td>
							<td class="center"><label class="label-green text-wrapper">{{ item.estimation }}</label></td>
							<td class="center box">
							  <img class="rectangle" :src="item.image" alt="Delivery Image" />
							</td>
							<!-- <td class="center">
							</td> -->
							 <td class="center ">
                                <div class="mt5" >
                                    <router-link :to="{ name: 'edit-delivery-method', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-delivery-method', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteDeliveryMethod(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-delivery-method',
    data(){
        return {
            search: "",
        }
    },
	components: {
		Menu
	},
    computed: {
        ...mapState({ // get data from state
			list_delivery_method: state => state.delivery_method.list
		})
	},
	mounted() {
		this.listDeliveryMethod() // render items for the first

	},
	methods: {
		...mapActions([
			'listDeliveryMethod',
            'deleteDeliveryMethod',
		]),
        onSearchDeliveryMethod(){
            console.log("asdasd")
            this.$store.commit('setSearchDeliveryMethod', this.search)
            this.listDeliveryMethod()
        },
	},
}
</script>
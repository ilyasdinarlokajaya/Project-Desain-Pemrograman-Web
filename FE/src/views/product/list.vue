<template>
	<div>
		<Menu />
		<div class="container">
			<h1 class="m0">Product</h1>
			<div class="row">
				<div class="col">
					<p class="mt5 mb20 tc-green">Manage your product details</p>
				</div>
				<div class="col flex">
					<router-link to="/add-product">
						<button class="bt-green">+ Add Product</button>
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
                    <button @click="onSearchProduct" class="bt-blue">Filter</button>
                </div>
            </div>
			<div class="mt50">
				<table>
					<thead>
						<tr>
							<th style="font-size: 14px;">PRODUCT CODE</th>
							<th style="font-size: 14px;">PRODUCT NAME</th>
							<th style="font-size: 14px;">PRICE </th>
							<th style="font-size: 14px;">STATUS</th>
							<th style="font-size: 14px;">PRODUCT <br> CATEGORY</th>
							<th style="font-size: 14px;">UNIT</th>
							<th style="font-size: 14px;">PRODUCT <br> STATUS</th>
							<th style="font-size: 14px;">PRODUCT IMAGE</th>

						</tr>
					</thead>
					<tbody>
						<tr v-for="item of list_product" :key="item.id">
							<td class="center">{{ item.product_code }}</td>
							<td class="center">{{ item.product_name }}</td>
							<td class="center"><p class="detail-title-bold">Rp. {{ item.purchase }}</p></td>
							<td class="center"><label class="label-green text-wrapper">Publish{{ item.status }}</label></td>
							<td class="center">{{ item.category }}</td>
							<td class="center"><p class="detail-title-bold">{{ item.unit }}</p></td>
							<td class="center">
								<div >
									<label class="label-green">{{ item.productstatus ="Availability"}}</label>
								</div>
								<td class="center box">
								  <img class="rectangle" :src="item.image" alt="Product Image" />
								</td>
							</td>
							 <td class="center ">
                                <div class="mt5" >
                                    <router-link :to="{ name: 'edit-product', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-product', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteProduct(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-product',
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
			list_product: state => state.product.list
		})
	},
	mounted() {
		this.listProduct() // render items for the first

	},
	methods: {
		...mapActions([
			'listProduct',
            'deleteProduct',
		]),
        onSearchProduct(){
            console.log("asdasd")
            this.$store.commit('setSearchProduct', this.search)
            this.listProduct()
        },
	},
}
</script>
<template>
    <div>
        <Menu/>
        <div class="container">
            <h1 class="m0">User</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage user data here</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-user">
                        <button class="bt-green">+ Add User</button>
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
                    <button @click="onSearchUser" class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>NAME</th>
                            <th>USERNAME</th>
                            <th>EMAIL</th>
                            <th>ROLES</th>
                            <th>ACTION</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_user" :key="item.id">
                            <td class="center">{{ item.name }}</td>
                            <td class="center">{{ item.username }}</td>
                            <td class="center">{{ item.email }}</td>
                            <td class="center">
                                <div v-if="item.role === '1'">
                                    <label class="label-green">Super Admin</label>
                                </div>
                                <div v-if="item.role === '2'">
                                    <label class="label-blue">Sales</label>
                                </div>
                                <div v-if="item.role === '3'">
                                    <label class="label-yellow">Marketing</label>
                                </div>
                            </td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link :to="{ name: 'edit-user', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-user', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteUser(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-user',
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
			list_user: state => state.user.list
		})
	},
	mounted() {
		this.listUser() // render items for the first
	},
	methods: {
		...mapActions([
			'listUser',
            'deleteUser',
		]),
        onSearchUser(){
            console.log("asdasd")
            this.$store.commit('setSearchUser', this.search)
            this.listUser()
        }
	},
}
</script>
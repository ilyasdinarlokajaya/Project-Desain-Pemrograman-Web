<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-user">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Detail User</h1>
                    <p class="mt5 mb20 tc-green">Your user details</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50">
                <p>Image Profile:</p>
                <img height="250" :src="user.image" alt="Uploaded Image" />
            </div>
            <div class="m10">
                <p>Name:</p>
                <h2>{{ user.name }}</h2>
            </div>
            <div class="mt10">
                <p class="detail-title-bold">Username:</p>
                <p>{{ user.username }}</p>
            </div>
            <div class="mt10">
                <p class="detail-title-bold">Email:</p>
                <p>{{ user.email }}</p>
            </div>
            <div>
                <p class="detail-title-bold">Roles</p>
                <label v-if="user.role === 1" class="label-green">Super Admin</label>
                <label v-if="user.role === 2" class="label-blue">Sales</label>
                <label v-if="user.role === 3" class="label-yellow">Marketing</label>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-user',
	computed: {
        ...mapState({ // get data from state
			user: state => state.user.detail
		})
	},
    mounted() {
        this.detailUser(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailUser'
		])
    }
}
</script>
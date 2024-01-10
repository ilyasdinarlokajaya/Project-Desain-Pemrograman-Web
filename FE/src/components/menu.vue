<template>
    <div class="container">
        <div class="row">
            <div class="col center">
                <router-link to="/list-user">
                    <div @click="selectedMenu('user')">
                        User
                        <hr v-if="selected_menu === 'user'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-product">
                    <div @click="selectedMenu('product')">
                        Product
                        <hr v-if="selected_menu === 'product'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-voucher">
                    <div @click="selectedMenu('voucher')">
                        Voucher
                        <hr v-if="selected_menu === 'voucher'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-banner">
                    <div @click="selectedMenu('banner')">
                        Banner
                        <hr v-if="selected_menu === 'banner'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-plannogram">
                    <div @click="selectedMenu('plannogram')">
                        Plannogram
                        <hr v-if="selected_menu === 'plannogram'" class="uline-green" >
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-article">
                    <div @click="selectedMenu('article')">
                        Article
                        <hr v-if="selected_menu === 'article'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-notification">
                    <div @click="selectedMenu('notification')">
                        Notification
                        <hr v-if="selected_menu === 'notification'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
            <div class="col center">
                <router-link to="/list-delivery-method">
                    <div @click="selectedMenu('delivery-method')">
                        Delivery Method
                        <hr v-if="selected_menu === 'delivery-method'" class="uline-green">
                        <hr v-else class="uline-grey">
                    </div>
                </router-link>
            </div>
        </div>
    </div>
</template>
<script>
import { mapState } from 'vuex'

export default {
    name: 'Menu',
    computed: {
        ...mapState({ // get data from state
			selected_menu: state => state.selected_menu
		})
	},
    methods: {
        selectedMenu(menu){
            this.$store.commit('setSelectedMenu', menu)
        }
    }
}
</script>
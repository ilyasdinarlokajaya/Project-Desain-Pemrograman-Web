<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-notification">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Detail Notification</h1>
                    <p class="mt5 mb20 tc-green">Your notification details</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50">
                <p>Title:</p>
                <h1>{{ notification.title }}</h1>
            </div>
            <div class="m30">
                <p>Content:</p>
                <h4 class="detail-title-thin">{{ notification.konten }}</h4>
            </div>
            <td class="m50">
                <p class="mb20 uline-black">Broadcast Schedule:</p>
                    <label class="label-green ">{{ notification.broadcast_day }}</label>
                <!-- <div v-if="item.notification_status === 1"> -->
                    <!-- </div> -->
                    <!-- <div v-if="notification.notificationstatus === 2">
                        <label class="label-red">{{ notification.broadcastDay }}</label>
                    </div> -->
                    
            </td>

        
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'detail-notification',
	computed: {
        ...mapState({ // get data from state
			notification: state => state.notification.detail
		})
	},
    mounted() {
        this.detailNotification(this.$route.params.id)

    },
    methods: {
        ...mapActions([
			'detailNotification'
		])
    }
}
</script>
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
                    <h1 class="m0">Update Notification</h1>
                    <p class="mt5 mb20 tc-green">Make notifications to your customers so they know about discount products and other things in your shop</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                <label class="form-label w80">Title & Content</label>
            </div>
                <div class="w80 relative">
                <input class="form-text" type="text" placeholder="Title Article" v-model="notification.title" />
                <div class="p-kanan xPojok" ><img src="../../assets/icons/x.png" width="30%"  @click="removeName"></div>
                <textarea class="form-text" cols="30" rows="10" placeholder="Content (Min. 1000 Char)"
                    v-model="notification.konten"></textarea>
                <div class="p-kanan2 p-kanan-desk2" ><img src="../../assets/icons/x.png" width="30%"  @click="removeContent"></div>
            </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="datetime-local"
                        v-model="notification.broadcast_day"
                    >
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-notification">
                    <button class="bt-submit-green" @click="onEditNotification">Update Notification</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-notification',
    data(){
        return {
        }
    },
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
			'detailNotification',
            'editNotification'
		]),
        onEditNotification() {
            this.editNotification({
                id: this.notification.id,
                title: this.notification.title,
                konten: this.notification.konten,
                broadcast_day: this.notification.broadcast_day
            })
        },
        removeName() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.notification.title = ""
        },
        removeContent() {
            // Fungsi ini akan dipanggil saat gambar diklik
            // Mengatur nilai displayText menjadi string kosong
            this.notification.konten = ""
        },
    }
}
</script>
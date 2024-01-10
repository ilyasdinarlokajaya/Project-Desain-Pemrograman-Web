<template>
    <div>
        <Menu/>
        <div class="container">
            <h1 class="m0">Notification</h1>
            <div class="row">
                <div class="col">
                    <p class="mt5 mb20 tc-green">Manage your Notification</p>
                </div>
                <div class="col flex">
                    <router-link to="/add-notification">
                        <button class="bt-green">+ Add Notification</button>
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
                    <button @click="onSearchNotification" class="bt-blue">Filter</button>
                </div>
            </div>
            <div class="mt50">
                <table>
                    <thead>
                        <tr>
                            <th>TITLE</th>
							<th>BROADCAST SCHEDULE</th>
							<th>ARTICLE STATUS</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item of list_notification" :key="item.id">
							<td class="center">{{ item.title }}</td>
							<td class="center">
								<div v-if="isValidTime(item.broadcast_day)">
								    <label class="label-green">{{ formattedBroadcastDay(item.broadcast_day) }}</label>
								</div>
                                <div v-else>
								    <label class="label-red">{{ formattedBroadcastDay(item.broadcast_day) }}</label>
								</div>
							</td>
							<td class="center">
								<div v-if="isValidTime(item.broadcast_day)"> 
								    <label class="label-green">Active</label>
								</div>
								<div v-else>
								    <label class="label-red">Inactive</label>
								</div>
							</td>
                            <td class="center">
                                <div class="mt5">
                                    <router-link :to="{ name: 'edit-notification', params: { id: item.id } }">
                                        <img  src="../../assets/icons/edit.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <router-link :to="{ name: 'detail-notification', params: { id: item.id } }">
                                        <img src="../../assets/icons/detail.svg">
                                    </router-link>
                                </div>
                                <div class="mt5">
                                    <img @click="deleteNotification(item.id)" src="../../assets/icons/delete.svg">
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
	name: 'list-notification',
    data(){
        return {
            search: "",
            isValidDateTime: true
        }
    },
	components: {
		Menu
	},
    computed: {
        ...mapState({ // get data from state
			list_notification: state => state.notification.list
		})
	},
	mounted() {
		this.listNotification() // render items for the first
	},
	methods: {
		...mapActions([
			'listNotification',
            'deleteNotification',
		]),
        onSearchNotification(){
            console.log("asdasd")
            this.$store.commit('setSearchNotification', this.search)
            this.listNotification()
        },
        // isValidTime(broadcastDay) {
        //     // Check if the selected date is in the future
        //     const currentTime = new Date();
        //     const broadcastTime = new Date(broadcastDay);
            
        //     // Format the date as 'DD-MM-YYYY'
        //         const formattedDate = `${currentTime.getDate() < 10 ? '0' : ''}${currentTime.getDate()}-${currentTime.getMonth() + 1 < 10 ? '0' : ''}${currentTime.getMonth() + 1}-${currentTime.getFullYear()}`;
        //         // Format the time as 'hh:mm A'
        //         const hours = currentTime.getHours();
        //         const hours2 = broadcastTime.getHours();
        //         const minutes = currentTime.getMinutes();
        //         const minutes2 = broadcastTime.getMinutes();
        //         const ampm = hours >= 12 ? 'PM' : 'AM';
        //         const ampm2 = hours2 >= 12 ? 'PM' : 'AM';
        //         const formattedHours = hours % 12 || 12; // Use 12 if hours is 0
        //         const formattedHours2 = hours2 % 12 || 12; // Use 12 if hours is 0
        //         const formattedTime = `${formattedHours < 10 ? '0' : ''}${formattedHours}:${minutes < 10 ? '0' : ''}${minutes} ${ampm}`;
        //         // Combine the formatted date and time
        //         const isAfterCurrentTime = hours2 > hours || (hours2 == hours && minutes2 >= minutes);
        //         this.formattedDateTime = `${formattedDate} ${formattedTime}`;
        //         const formattedBroadcastDay = `${broadcastTime.getMonth() + 1 < 10 ? '0' : ''}${broadcastTime.getMonth() + 1}-${broadcastTime.getDate() < 10 ? '0' : ''}${broadcastTime.getDate()}-${broadcastTime.getFullYear()} ${formattedHours2 < 10 ? '0' : ''}${formattedHours2}:${minutes2 < 10 ? '0' : ''}${minutes2} ${ampm2}`;
        //         console.log(broadcastTime.getMonth()+1);
        //         console.log(currentTime.getDate());
        //         // console.log(this.formattedDateTime);
        //     return this.formattedBroadcastDay >= this.formattedDateTime || isAfterCurrentTime || (hours2 == 0 && broadcastTime.getMonth() + 1 > currentTime.getDate());
        // },
        // updateTime() {
        //     // This method can be used to update isValidDateTime for all items
        //     this.list_notification.forEach(item => {
        //         item.isValidDateTime = this.isValidTime(item.broadcast_day);
        //     });
        // }
        formattedBroadcastDay(broadcastDay) { 
            const date = new Date(broadcastDay);
            const formattedMonth = date.getMonth() + 1 < 10 ? `0${date.getMonth()+1}` : `${date.getMonth()+1}`;
            const formattetDate = date.getDate() < 10 ? `0${date.getDate()}` : `${date.getDate()}`;
            const formattedHour = date.getHours() < 10 ? `0${date.getHours()}` : `${date.getHours()}`;
            const formattedMin = date.getMinutes() < 10 ? `0${date.getMinutes()}` : `${date.getMinutes()}`;
            const period = date.getHours() < 12 ? 'AM' : 'PM';
            return `${formattedMonth}-${formattetDate}-${date.getFullYear()} ${formattedHour}:${formattedMin} ${period}`;
        },
        isValidTime(broadcastDay){
            return new Date(broadcastDay) >= new Date()
        },
	},
}
</script>
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
                    <h1 class="m0">Add Notification</h1>
                    <p class="mt5 mb20 tc-green">Make notifications to your customers so they know about discount
            products and other things in your shop</p>
                </div>
            </div>
             <hr class="uline-grey" />
            <div class="mt50 row">
                <div class="w20">
                <label class="form-label w80">Title & Content</label>
            </div>
                <div class="w80 relative">
                <input class="form-text" type="text" placeholder="Title Article" v-model="title" />
                <textarea class="form-text" cols="30" rows="10" placeholder="Content (Min. 1000 Char)"
                    v-model="konten"></textarea>
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
                        placeholder="Input your day"
                        v-model="broadcast_day"
                        @input="updateTime"
                    >
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-notification">
                    <button class="bt-submit-green" @click="onAddNotification">Add Notification</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'add-notification',
    data(){
        return {
            id: "",
            title: "",
            konten: "",
			broadcast_day: "",
            currentFormattedDateTime: "",
        }
    },
    mounted(){
        this.updateFormattedDateTime();
    },
    methods: {
        updateTime() {
      this.updateFormattedDateTime();
    },
    // updateFormattedDateTime() {
    //    if (this.broadcast_day) {
    //         const inputDate = new Date(this.broadcast_day);
    //         // Format the date as 'DD-MM-YYYY'
    //         const formattedDate = `${inputDate.getDate()}-${inputDate.getMonth() + 1}-${inputDate.getFullYear()}`;

    //         // Format the time as 'hh:mm A'
    //         const hours = inputDate.getHours();
    //         const minutes = inputDate.getMinutes();
    //         const ampm = hours >= 12 ? 'PM' : 'AM';
    //         const formattedHours = hours % 12 || 12;
    //         const formattedTime = `${formattedHours < 10 ? '0' : ''}${formattedHours}:${minutes < 10 ? '0' : ''}${minutes} ${ampm}`;

    //         // Combine the formatted date and time
    //         this.currentFormattedDateTime = `${formattedDate} ${formattedTime}`;
    //         // this.currentFormattedDateTime = formattedTime;
    //     }

    // },
    updateFormattedDateTime() {
      if (this.broadcast_day) {
        // Parse the formatted date back to the original format4
        console.log(this.broadcast_day);
        const parts = this.broadcast_day.split('T');
        const datePart = parts[0];
        const timePart = parts[1];

        // Rearrange the date to the desired format
        const [year, month, day] = datePart.split('-');
        const [hour, minute] = timePart.split(':');

        // Create the final formatted date
        const period = hour < 12 ? 'AM' : 'PM';
        const hour12 = (hour % 12) || 12;

        const formattedHour = hour12 < 10 ? `0${hour12}` : `${hour12}`;

        // Create the final formatted date
        const formatted = `${month}-${day}-${year} ${formattedHour}:${minute} ${period}`;
        this.currentFormattedDateTime = formatted;
        console.log(this.currentFormattedDateTime);
        } 
    },
        ...mapActions([
			'addNotification',
            'createNotification'
		]),
        onAddNotification(){
            this.addNotification({
                title: this.title,
                konten: this.konten,
                broadcast_day: this.currentFormattedDateTime,
            })
        },
        onClickAddNotification(){
            this.createNotification({
                title: this.title,
                konten: this.konten,
                broadcast_day: this.currentFormattedDateTime,
                
            })
        },
    }
}
</script>
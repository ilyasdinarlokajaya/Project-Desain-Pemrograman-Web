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
                    <h1 class="m0">Update User</h1>
                    <p class="mt5 mb20 tc-green">Update user data</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-label">Image Profile</label>
                </div>
                <div class="w80">
                    <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
                    <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <img v-else :src="user.image" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
                    <button class="bt-change-image" @click="openFileInput">Change Image</button>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Name</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input your full name"
                        v-model="user.name"
                    >
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Username</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input your username"
                        v-model="user.username"
                    >
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Email</label>
                </div>
                <div class="w80">
                    <input
                        class="form-text"
                        type="text"
                        placeholder="Input your email"
                        v-model="user.email"
                    >
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-label">Roles</label>
                </div>
                <div class="w80">
                    <select class="form-select" v-model="user.role">
                        <option disabled selected>--- Select Role ---</option>
                        <option value="1">Super Admin</option>
                        <option value="2">Sales</option>
                        <option value="3">Marketing</option>
                    </select>
                </div>
            </div>
            <div class="flex mt30 mb40">
                <router-link to="/list-user">
                    <button class="bt-submit-green" @click="onEditUsser">Update User</button>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
	name: 'edit-user',
    data(){
        return {
            imageUrl: null,
            imageData: null
        }
    },
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
			'detailUser',
            'editUser'
		]),
        onEditUsser() {
            this.editUser({
                id: this.user.id,
                name: this.user.name,
                username: this.user.username,
                email: this.user.email,
                role: parseInt(this.user.role),
                image: this.imageData != null? this.imageData : this.user.image
            })
        },
        openFileInput() {
            // Trigger click on the file input when the button is clicked
            this.$refs.fileInput.click();
        },
        handleFileChange(event) {
            const file = event.target.files[0];

            if (file) {
                // Handle the selected file, for example, display it as an image
                this.imageUrl = URL.createObjectURL(file);
                const reader = new FileReader();
                reader.onloadend = () => {
                    this.imageData = reader.result; // Save the Base64-encoded image data
                };
                reader.readAsDataURL(file);
            }
        },
    }
}
</script>
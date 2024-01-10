<template>
  <div>
    <div class="container">
      <div class="row">
        <div class="mr40 mt20">
          <router-link to="/list-article">
            <img src="../../assets/icons/back.svg" />
          </router-link>
        </div>
        <div>
          <h1 class="m0">Add Article</h1>
          <p class="mt5 mb20 tc-green">
            Make your article looks good
          </p>
        </div>
      </div>
      <hr class="uline-grey" />
      <div class="mt50 row">
        <div class="w20">
          <label class="form-label">Article Image</label>
        </div>
        <div class="w80">
          <input type="file" ref="fileInput" @change="handleFileChange" style="display: none" />
          <button v-if="imageUrl === null" class="bt-image" @click="openFileInput"><img
              src="../../assets/icons/upload_image.svg"></button>
          <div v-if="imageUrl">
            <img v-if="imageUrl" :src="imageUrl" alt="Uploaded Image" style="max-width: 100%; height: auto;" />
            <button class="bt-change-image" @click="openFileInput">Change Image</button>
          </div>
        </div>
      </div>
      <div class="mt20 row">
                <div class="w20">
                    <label class="form-label">Article Name & <br>Content</label>
                </div>
                <div class="w80">
                    <input @input="handleInput" class="form-text" type="text"
                        placeholder="Product Name" maxLength="20" v-model="judul_artikel"/>
                    <div class="character-count1 ">
                        | {{ judul_artikel_length }} / 20
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="w20">
                </div>
                <div class="w80">
                    <textarea class="form-text" cols="30" rows="10" placeholder="Content (Min. 1000 Char)"
                    v-model="konten"></textarea>
                </div>
            </div>
      <div class="mt30 row">
        <div class="w20">
            <label class="form-label">Broadcast Schedule</label>
        </div>
        <div class="w80 custom-select-container">
            <select class="form-select" v-model="jadwal_broadcast" >
                <option disabled selected value="">Pilih Hari</option>
                <option value="1">Senin</option>
                <option value="2">Selasa</option>
                <option value="3">Rabu</option>
                <option value="4">Kamis</option>
                <option value="5">Jumat</option>
                <option value="6">Sabtu</option>
                <option value="7">Minggu</option>
            </select>
            <img width="15" class="addon-right-icon" src="../../assets/images/edit-2.png" alt="" />
        </div>
    </div>
      <div class="mt30 row">
        <div class="w20">
          <label class="form-label">Valid From</label>
        </div>
        <div class="w80">
          <input class="form-text" type="date" placeholder="Masukkan Tanggal" v-model="valid_from" />
        </div>
      </div>
      <div class="mt30 row">
        <div class="w20">
          <label class="form-label">Valid Until</label>
        </div>
        <div class="w80">
          <input class="form-text" type="date" placeholder="Masukkan Tanggal" v-model="valid_until" />
        </div>
      </div>
      <div class="flex mt30">
        <router-link to="/list-article">
          <button class="bt-submit-green" @click="onAddArticle">Add Article</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
export default {
  name: "add-article",
  data() {
    return {
      id: "",
      judul_artikel: "",
      konten: "",
      jadwal_broadcast: "",
      valid_from: "",
      valid_until: "",
      judul_artikel_length: 5,
      imageUrl: null,
      imageData: null
    }
  },
  methods: {
    ...mapActions([
			'addArticle',
		    ]),
    onAddArticle(){
        this.addArticle({
            judul_artikel: this.judul_artikel,
            konten: this.konten,
            jadwal_broadcast: this.jadwal_broadcast,
            valid_from: this.valid_from,
            valid_until: this.valid_until,
            image: this.imageData
      })
    },
    handleInput() {
        this.judul_artikel_length = this.judul_artikel.length;
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
  },
};
</script>


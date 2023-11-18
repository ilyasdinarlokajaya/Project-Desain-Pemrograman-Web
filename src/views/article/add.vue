<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="mr40 mt20">
                    <router-link to="/list-Article">
                        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
                        <img src="../../assets/icons/back.svg">
                    </router-link>
                </div>
                <div>
                    <h1 class="m0">Add Article</h1>
                    <p class="mt5 mb20 tc-green">Make your article looks good</p>
                </div>
            </div>
            <hr class="uline-grey">
            <div class="mt50 row">
                <div class="w20">
                    <label class="form-labell">Article Image</label>
                </div>
                <div class="upload-container">
                    <input type="file" id="file-input" accept="image/*">
                    <label class="tulisan" for="file-input"><i class="far fa-image fa-lg"></i>Tambahkan Foto</label>
                    <div id="uploaded-image"></div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-labell">
                        <tr>Article Name &</tr>
                        <tr>Content</tr>
                    </label>
                </div>
                <div class="w80">
                    <input class="form-text" type="text" placeholder=" Judul Article">
                    <input class="form-text-content" type="text" placeholder="  Content (Min. 1000 Char)">
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-labell">Broadcast Schedule</label>
                </div>
                <div class="w80">
                    <div class="select-container">
                        <select>
                            <option>Pilih Hari</option>
                            <option>Senin</option>
                            <option>Selasa</option>
                            <option>Rabu</option>
                            <option>Kamis</option>
                            <option>Jumat</option>
                            <option>Sabtu</option>
                            <option>Minggu</option>
                        </select>
                        <div class="select-icon">
                            <i class="fa-light fa-pen"></i>
                        </div>
                    </div>
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-labell">Valid From</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="date" placeholder="Masukkan Tanggal">
                    <!-- <input type="date" id="input-date" name="input-date"> -->
                </div>
            </div>
            <div class="mt30 row">
                <div class="w20">
                    <label class="form-labell">Valid Until</label>
                </div>
                <div class="w80">
                    <input class="form-text" type="date" placeholder="Masukkan Tanggal">
                    <!-- <select class="form-select">
                            <option disabled selected>--- Select Role ---</option>
                            <option value="1">Super Admin</option>
                            <option value="2">Marketing</option>
                            <option value="3">Sales</option>
                        </select> -->
                </div>

            </div>
            <div class="flex mt30">
                <router-link to="/list-Article">
                    <button class="bt-submit-green">Add Article</button>
                </router-link>
            </div>
        </div>
    </div>
</template>
<script>
// Menangani perubahan pada input file
document.getElementById('file-input').addEventListener('change', function (event) {
    var fileInput = event.target;
    var uploadedImage = document.getElementById('uploaded-image');

    // Pastikan ada file yang dipilih
    if (fileInput.files.length > 0) {
        var reader = new FileReader();

        reader.onload = function (e) {
            // Tampilkan gambar yang diunggah
            uploadedImage.innerHTML = '<img src="' + e.target.result + '" alt="Uploaded Image">';
        };

        // Baca file sebagai URL data
        reader.readAsDataURL(fileInput.files[0]);
    } else {
        // Kosongkan tampilan jika tidak ada file yang dipilih
        uploadedImage.innerHTML = '';
    }
});
</script>

    <script>
    export default {
        name: 'add-Article'
    }
    </script>
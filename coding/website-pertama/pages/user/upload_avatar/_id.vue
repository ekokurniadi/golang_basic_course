<template>
  <div>
    <h2>Upload Avatar</h2>
    <div class="card mb-4">
      <div class="card-body">
        <nuxt-link to="/user" class="btn btn-danger">Back</nuxt-link>

        <div class="form-group">
          <label for="nama_user" class="label-control col-md-2"
            >Nama User</label
          >
          <div class="col-md-12">
            <input
              type="text"
              class="form-control"
              name="nama_user"
              id="nama_user"
              v-model="nama_user"
              readonly
            />
          </div>
        </div>

        <div class="form-group">
          <label for="avatar" class="label-control col-md-2"
            >Profile Picture</label
          >
          <div class="col-md-12">
            <input
              type="file"
              class="form-control"
              name="profile_picture"
              id="profile_picture"
              @change="handleFileUpload"
            />
          </div>
        </div>

        <div class="form-group">
          <div class="col-md-12">
            <button class="btn btn-primary" @click="submit()">Save</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      nama_user: '',
      profile_picture: '',
    }
  },
  mounted() {
    this.getUser()
  },
  methods: {
    getUser() {
      this.$axios
        .get(process.env.API_BASE_URL + '/users/' + this.$route.params.id)
        // http://localhost:8080/api/v1/users/1
        .then((response) => {
          this.nama_user = response.data.data.nama_lengkap
        })
    },
    handleFileUpload(event) {
      this.profile_picture = event.target.files[0]
    },
     async submit() {
      let formData = new FormData()
      formData.append('id',this.$route.params.id)
      formData.append('profile_picture', this.profile_picture)

      await this.$axios
        .post(
          process.env.API_BASE_URL + '/users/avatar',formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
          },

        )
        .then((response) => {
          alert('Avatar berhasil diupload')
          this.$router.push('/user')
        })
        .catch((error) => {
         alert(error);
        })
    },
  },
}
</script>

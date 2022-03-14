<template>
  <div>
    <h2>Edit User</h2>
    <div class="card mb-4">
      <div class="card-body">
        <nuxt-link to="/user" class="btn btn-danger">Back</nuxt-link>

        <div class="form-group">
          <label for="nama_lengkap" class="label-control col-md-2"
            >Nama Lengkap</label
          >
          <div class="col-md-12">
            <input
              type="text"
              class="form-control"
              name="nama_lengkap"
              id="nama_lengkap"
              v-model="nama_lengkap"
              required
            />
          </div>
        </div>

        <div class="form-group">
          <label for="jenis_kelamin" class="label-control col-md-2"
            >Jenis Kelamin</label
          >
          <div class="col-md-12">
            <select
              name="jenis_kelamin"
              id="jenis_kelamin"
              class="form-control"
              v-model="jenis_kelamin"
            >
              <option value="">Choose an option</option>
              <option
                value="Laki-Laki"
                selected="'jenis_kelamin==Laki-Laki'"
              >
                Laki-Laki
              </option>
              <option
                value="Perempuan"
                selected="'jenis_kelamin==Perempuan'"
              >
                Perempuan
              </option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="username" class="label-control col-md-2">Username</label>
          <div class="col-md-12">
            <input
              type="text"
              class="form-control"
              name="username"
              id="username"
              v-model="user_name"
              required
              readonly
            />
          </div>
        </div>

        <div class="form-group">
          <label for="password" class="label-control col-md-2">Password</label>
          <div class="col-md-12">
            <input
              type="text"
              class="form-control"
              name="password"
              id="password"
              v-model="password"
              required
            />
          </div>
        </div>

        <div class="form-group">
          <label for="role" class="label-control col-md-2">Role</label>
          <div class="col-md-12">
            <select name="role" id="role" class="form-control" v-model="role">
              <option v-bind:value="role" selected>
                <div v-if="role == 'administrator'">Administrator</div>
                <div v-if="role == 'operator1'">Operator1</div>
                <div v-if="role == 'operator2'">Operator2</div>
              </option>
              <option value="administrator" >Administrator</option>
              <option value="operator1" >Operator1</option>
              <option value="operator2" >Operator2</option>
            </select>
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
      nama_lengkap: '',
      jenis_kelamin: '',
      user_name: '',
      password: '',
      role:'',
      forms: {
        nama_lengkap: '',
        jenis_kelamin: '',
        user_name: '',
        password: '',
        role:''
      },
    }
  },
  methods: {
    async getDataUser() {
      await this.$axios
        .get(`${process.env.API_BASE_URL}/users/${this.$route.params.id}`)
        .then((response) => {
          this.nama_lengkap = response.data.data.nama_lengkap
          this.jenis_kelamin = response.data.data.jenis_kelamin
          this.user_name = response.data.data.user_name
          this.password = response.data.data.password
          this.role = response.data.data.role

        })
        .catch((err) => {
          alert(err)
        })
    },
    async submit() {
      this.forms.nama_lengkap = this.nama_lengkap
      this.forms.jenis_kelamin = this.jenis_kelamin
      this.forms.user_name = this.user_name
      this.forms.password = this.password
      this.forms.role= this.role
      await this.$axios
        .put(
          `${process.env.API_BASE_URL}/users/${this.$route.params.id}`,
          this.forms
        )
        .then((response) => {
          alert(response.data.meta.message)
          this.$router.push('/user')
        })
        .catch((err) => {
          alert(err)
        })
    },
  },
  mounted(){
    this.getDataUser()
  }
}
</script>

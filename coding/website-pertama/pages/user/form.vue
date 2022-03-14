<template>
  <div>
    <h2>Form User</h2>
    <div class="card mb-4">
      <div class="card-body">
        <nuxt-link to="/user" class="btn btn-danger">Back</nuxt-link>

        <div class="form-group">
          <label for="nama_lengkap" class="label-control col-md-2">Nama Lengkap</label>
          <div class="col-md-12">
            <input type="text" name="nama_lengkap" id="nama_lengkap" class="form-control" placeholder="Nama Lengkap" required v-model="form.nama_lengkap">
          </div>
        </div>

        <div class="form-group">
          <label for="jenis_kelamin" class="label-control col-md-2">Jenis Kelamin</label>
          <div class="col-md-12">
            <select name="jenis_kelamin" id="jenis_kelamin" class="form-control" required v-model="form.jenis_kelamin">
              <option value="">Choose an option</option>
              <option value="Laki-Laki">Laki-Laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="user_name" class="label-control col-md-2">Username</label>
          <div class="col-md-12">
            <input type="text" name="user_name" id="user_name" class="form-control" placeholder="Username" required v-model="form.user_name">
          </div>
        </div>

        <div class="form-group">
          <label for="password" class="label-control col-md-2">Password</label>
          <div class="col-md-12">
            <input type="password" name="password" id="password" class="form-control" required placeholder="Password" v-model="form.password">
          </div>
        </div>

        <div class="form-group">
          <label for="role" class="label-control col-md-2">Role</label>
          <div class="col-md-12">
            <select name="role" id="role" class="form-control" required v-model="form.role">
              <option value="">Choose an option</option>
              <option value="administrator">Administrator</option>
              <option value="operator1">Operator1</option>
              <option value="operator2">Operator2</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <div class="col-md-12">
            <button class="btn btn-success" @click="save()"><i class="fa fa-disk"></i> Save</button>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      form:{
        nama_lengkap:'',
        jenis_kelamin:'',
        user_name:'',
        password:'',
        role:''
      }
    }
  },
  methods:{
    async save(){
      await this.$axios.post(`${process.env.API_BASE_URL}/users`,this.form,{
        headers:{
          'Content-Type':'application/json'
        }
      }).then((response)=>{
        alert(response.data.meta.message)
        this.$router.push('/user')
      }).catch((error)=>{
        alert(error)
      })
    }
  }
}
</script>

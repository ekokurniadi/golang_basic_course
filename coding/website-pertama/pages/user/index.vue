<template>
  <div>
    <h2>List Of Users</h2>
      <div class="card mb-4">
        <div class="card-body">

          <nuxt-link to="/user/form" class="btn btn-primary"><i class="fa fa-plus"></i> Add New</nuxt-link>

          <!-- komponen search / pencarian data -->
          <div class="row">
            <div class="col-md-3 offset-8">
              <input type="text" placeholder="Pencarian data" class="form-control" v-model="search" @keyup="retrieveData()">
            </div>
            <div class="col-md-1">
              <button class="btn btn-primary btn-sm" @click="clearSearch()">Reset</button>
            </div>
          </div>
          <!-- komponen jumlah data per halaman -->
          <div class="row">
            <div class="col-md-3">
              <v-select v-model="pageSize" :items="pageSizeItem" label="Data per page" @change="handlePageSize">
              </v-select>
            </div>
          <!-- komponen paging / tombol untuk pindah ke halaman table next atau previous -->
            <div class="col-md-9">
              <v-pagination
                v-model="page"
                :length="totalPages"
                total-visible="7"
                next-icon="mdi-menu-right"
                prev-icon="mdi-menu-left"
                @input="handlePageChange"
              >
              </v-pagination>
            </div>
          </div>

          <!-- komponen table -->
          <div class="row">

           <div class="col-md-12">
             <v-data-table
             :headers="headers"
             :items="users"
             :loading="loading"
             :hide-default-footer="true"
             class="elevation-1 spacing-playground pa-6">
            <template v-slot:[`item.profile_picture`]="{item}">
              <v-img :src="item.profile_picture" height="100" width="100" class="img-thumbnail"></v-img>
            </template>
              <template v-slot:[`item.actions`]="{item}">
                <v-icon
									small
									class="mr-2 btn btn-success btn-sm"
									style="color: black"
									@click="uploadPicture(item.id)">
									 mdi-wrench
								</v-icon>
                <v-icon
									small
									class="mr-2 btn btn-warning btn-sm"
									style="color: black"
									@click="editData(item.id)">
									mdi-pencil
								</v-icon>
                <v-icon
									small
									class="mr-2 btn btn-danger btn-sm"
									style="color: white"
									@click="deleteData(item.id)">
									mdi-delete
								</v-icon>
              </template>
             </v-data-table>
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
      users:[],
      headers:[
        {text:'Avatar',value:'profile_picture',align:'center'},
        {text:'Nama' ,sortable:false,value:'nama_lengkap'},
        {text:'Jenis Kelamin',sortable:false,value:'jenis_kelamin'},
        {text:'Username',sortable:false,value:'user_name'},
        {text:'Password',sortable:false,value:'password'},
        {text:'Role',sortable:false,value:'role'},
        {text:'Action',sortable:false,value:'actions'}
      ],
      search:'',
      pageSize:10,
      totalPages:0,
      pageSizeItem:[10],
      page:1,
      loading:true,
    }
  },
  watch:{
    options:{
      handler(){
        this.retrieveData()
      },
      deep:true
    },
  },
  methods:{
    uploadPicture(id){
      this.$router.push('/user/upload_avatar/'+id)
      // localhost:3000/user/upload_avatar/1
    },
    async deleteData(id){
      if(confirm('Apakah anda yakin?')){
        await this.$axios.delete(`${process.env.API_BASE_URL}/users/${id}`)
        .then((response)=>{
          alert('Data berhasil dihapus')
          this.retrieveData()
        })
        .catch((err)=>{
          alert(err)
        })
      }
    },
    editData(id){
      this.$router.push('/user/'+id)
    },
    clearSearch(){
      this.search =''
      this.retrieveData()
    },
    getRequestParams(search, page, pageSize) {
			let params = {}
			if (search) {
				params['search'] = search
			}
			if (page) {
				params['page'] = page - 1
			}
			if (pageSize) {
				params['size'] = pageSize
			}
			return params
		},
    async getAll(params){
     return await this.$axios.get(`${process.env.API_BASE_URL}/users`,{params})
    //  localhost:8080/api/v1/users
    },
    retrieveData(){
      const params = this.getRequestParams(
        this.search,
        this.page,
        this.pageSize
      )
      this.getAll(params)
      .then((response)=>{
        this.loading = true;
        const {data,totalPages}= response.data;
        this.users = data.map(this.showData);
        this.totalPages =totalPages;
        this.loading = false;
      })
      .catch((err)=>{
        alert(err);
      })
    },
    handlePageChange(value) {
			this.page = value
			this.retrieveData()
		},
		handlePageSize(size) {
			this.pageSize = size
			this.page = 1
			this.retrieveData()
		},
    showData(data){
      return {
        id:data.id,
        nama_lengkap: data.nama_lengkap,
        jenis_kelamin : data.jenis_kelamin,
        user_name : data.user_name,
        password : data.password,
        role : data.role,
        profile_picture: process.env.IMAGE_URL +data.profile_picture,
      }
    },
    refreshData(){
      this.retrieveData()
    }
  },
  mounted(){
    this.refreshData()
  }
}
</script>
<style>
.theme--light.v-pagination .v-pagination__item--active {
	color: red;
}
</style>

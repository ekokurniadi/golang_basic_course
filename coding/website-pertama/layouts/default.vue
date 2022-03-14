<template>
  <div>
    <app-header></app-header>

    <div class="d-flex">

       <app-sidebar></app-sidebar>

        <div class="content p-4">
          <Nuxt/>
        </div>

    </div>

    <app-footer></app-footer>
  </div>
</template>

<script>
import ComponentHeader from '@/components/Header.vue'
import ComponentSideBar from '@/components/SideBar.vue'
import ComponentFooter from '@/components/Footer.vue'
export default {
  components:{
    'app-header':ComponentHeader,
    'app-sidebar':ComponentSideBar,
    'app-footer':ComponentFooter
  }
}
</script>

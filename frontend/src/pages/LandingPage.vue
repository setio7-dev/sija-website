<script setup lang="ts">
import Navbar from '../components/landing/Navbar.vue'
import Home from '../components/landing/Home.vue'
import ITC from '../components/landing/ITC.vue'
import Modul from '../components/landing/Modul.vue'
import Achievement from '../components/landing/Achievement.vue'
import Partner from '../components/landing/Partner.vue'
import Footer from '../components/Footer.vue'
</script>

<template>
  <div class="min-h-screen overflow-x-hidden">
    
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div class="absolute top-20 -left-40 w-96 h-96 bg-primary/20 rounded-full blur-3xl animate-pulse"></div>
      <div class="absolute bottom-20 -right-40 w-96 h-96 bg-secondary/20 rounded-full blur-3xl animate-pulse" style="animation-delay: 1s"></div>
      <div class="absolute top-1/2 left-1/2 w-96 h-96 bg-tertiary/10 rounded-full blur-3xl animate-pulse" style="animation-delay: 2s"></div>
    </div>

    <Navbar/>
    <Home/>
    <ITC/>
    <Modul/>
    <Achievement/>
    <Partner/>
    <Footer/>    
  </div>
</template>
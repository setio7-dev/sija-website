<script setup lang="ts">
import { ref, onMounted } from 'vue'

const scrolled = ref(false)
const mobileMenuOpen = ref(false)

const handleScroll = () => {
    scrolled.value = window.scrollY > 50
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

</script>

<template>
    <nav class="fixed top-0 left-0 right-0 z-50 transition-all duration-500 backdrop-blur-xl py-6 shadow-2xl shadow-primary/10" :class="scrolled ? 'bg-slate-900/80' : 'bg-transparent'">
        <div class="container mx-auto px-6">
            <div class="flex items-center justify-between">

                <div class="flex items-center gap-4 group cursor-pointer">
                    <div class="relative">
                        <div
                            class="absolute inset-0 bg-linear-to-r from-primary to-secondary rounded-2xl blur-lg group-hover:blur-xl transition-all duration-300 opacity-75">
                        </div>
                        <div
                            class="relative w-14 h-14 rounded-2xl bg-linear-to-br from-primary via-tertiary to-secondary flex items-center justify-center text-white text-2xl font-poppins-semibold shadow-2xl transform group-hover:scale-110 group-hover:rotate-6 transition-all duration-300">
                            S
                        </div>
                    </div>
                    <div>
                        <h1
                            class="text-2xl font-poppins-semibold bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">
                            SIJA</h1>
                        <p class="text-xs text-slate-400 font-poppins-regular">SMKN 26 Jakarta</p>
                    </div>
                </div>

                <div class="hidden lg:flex items-center gap-8">
                    <a href="#welcome"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>Beranda</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                    <a href="#tentang"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>Tentang</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                    <a href="#itclub"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>IT Club</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                    <a href="#modul"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>Modul</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                    <a href="#prestasi"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>Prestasi</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                    <a href="#partner"
                        class="relative cursor-pointer text-slate-300 hover:text-white transition-all duration-300 font-poppins-medium group">
                        <span>Partner</span>
                        <span
                            class="absolute bottom-0 left-0 w-0 h-0.5 bg-linear-to-r from-primary to-secondary group-hover:w-full transition-all duration-300"></span>
                    </a>
                </div>

                <div class="hidden lg:block">
                    <button class="relative group">
                        <div
                            class="absolute inset-0 bg-linear-to-r from-primary to-secondary rounded-xl blur opacity-75 group-hover:opacity-100 transition-all duration-300">
                        </div>
                        <div
                            class="relative px-8 py-3 bg-linear-to-r from-primary to-secondary rounded-xl font-poppins-semibold text-white transform group-hover:scale-105 transition-all duration-300 shadow-lg shadow-primary/50">
                            Masuk
                        </div>
                    </button>
                </div>

                <button @click="mobileMenuOpen = !mobileMenuOpen" class="lg:hidden text-white p-2">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            :d="mobileMenuOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'" />
                    </svg>
                </button>
            </div>

            <div v-if="mobileMenuOpen"
                class="lg:hidden mt-6 pb-6 border-t border-slate-800 pt-6 space-y-4 animate-fadeIn">
                <a href="#welcome"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">Beranda</a>
                <a href="#tentang"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">Tentang</a>
                <a href="#itclub"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">IT
                    Club</a>
                <a href="#modul"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">Modul</a>
                <a href="#prestasi"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">Prestasi</a>
                <a href="#partner"
                    class="block cursor-pointer w-full text-left text-slate-300 hover:text-white transition-colors font-poppins-medium py-2">Partner</a>
                <button
                    class="w-full px-8 py-3 bg-linear-to-r from-primary to-secondary rounded-xl font-poppins-semibold text-white mt-4">
                    Masuk
                </button>
            </div>
        </div>
    </nav>
</template>
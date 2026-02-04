<script setup lang="ts">
import { Partners } from '../../data/LandingPageData';
</script>

<template>
    <section id="partner" class="relative py-32 px-6 bg-linear-to-b from-slate-950 to-slate-900">
        <div class="container mx-auto relative z-10">
            <div class="text-center mb-20 animate-fadeInUp">
                <div
                    class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Trusted
                        Partners</span>
                </div>
                <h2 class="text-5xl lg:text-6xl font-poppins-semibold mb-6">
                    <span class="text-white">Rekan</span>
                    <br />
                    <span
                        class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Perusahaan</span>
                </h2>
                <p class="text-xl text-slate-400 font-poppins-regular max-w-3xl mx-auto">
                    Bekerja sama dengan perusahaan teknologi global untuk memberikan pengalaman industri terbaik
                </p>
            </div>

            <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
                <div v-for="(partner, index) in Partners" :key="index"
                    class="group relative bg-white rounded-3xl p-8 hover:shadow-2xl hover:shadow-primary/20 transition-all duration-500 transform hover:scale-110 animate-fadeInUp"
                    :style="{ animationDelay: `${index * 0.1}s` }">
                    <div
                        class="absolute inset-0 bg-linear-to-br from-primary/0 to-secondary/0 group-hover:from-primary/5 group-hover:to-secondary/5 rounded-3xl transition-all duration-500">
                    </div>

                    <div class="relative aspect-square flex items-center justify-center">
                        <img :src="partner.image" :alt="partner.name"
                            class="w-full h-full object-contain filter grayscale group-hover:grayscale-0 transition-all duration-500 transform group-hover:scale-110" />
                    </div>
                    <p
                        class="text-center mt-4 font-poppins-semibold text-slate-800 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                        {{ partner.name }}</p>
                </div>
            </div>

            <div class="mt-16 text-center">
                <p class="text-slate-400 font-poppins-regular mb-4">Dan masih banyak perusahaan lainnya...</p>
                <button class="group relative inline-block">
                    <div
                        class="absolute inset-0 bg-linear-to-r from-primary to-secondary rounded-xl blur opacity-75 group-hover:opacity-100 transition-all duration-300">
                    </div>
                    <div
                        class="relative px-8 py-4 bg-linear-to-r from-primary to-secondary rounded-xl font-poppins-semibold text-white transform group-hover:scale-105 transition-all duration-300">
                        Lihat Semua Partner
                    </div>
                </button>
            </div>
        </div>
    </section>
</template>
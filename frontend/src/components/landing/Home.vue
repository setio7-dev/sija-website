<script setup lang="ts">
import { Stats } from '../../data/LandingPageData';
import sija from "../../../public/images/sija-building.jpeg"
</script>

<template>
    <section id="welcome" class="relative min-h-screen flex items-center justify-center pt-20 px-6" data-aos="fade-up" data-aos-duration="1000">
        <div class="container mx-auto relative z-10">
            <div class="grid lg:grid-cols-2 gap-16 items-center">
                <div class="space-y-8">
                    <div
                        class="inline-flex items-center gap-2 px-6 py-3 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium animate-pulse">
                        <span class="w-2 h-2 bg-primary rounded-full animate-ping"></span>
                        <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Sistem
                            Informasi Jaringan & Aplikasi</span>
                    </div>

                    <h1 class="text-6xl lg:text-7xl font-poppins-semibold leading-tight">
                        <span class="text-white">Welcome to</span>
                        <br />
                        <span
                            class="bg-linear-to-r from-primary via-tertiary to-secondary bg-clip-text text-transparent animate-gradient">SIJA</span>
                    </h1>

                    <p class="text-xl text-slate-400 font-poppins-regular leading-relaxed">
                        SMKN 26 Jakarta - Shaping the future of IT professionals. Master technology, build innovations,
                        lead the digital era.
                    </p>

                    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                        <div v-for="(stat, index) in Stats" :key="index"
                            class="text-center p-4 bg-slate-800/50 backdrop-blur-sm rounded-2xl border border-slate-700/50 hover:border-primary/50 transition-all duration-300 transform hover:scale-105">
                            <div class="text-3xl mb-2">{{ stat.icon }}</div>
                            <div
                                class="text-2xl font-poppins-semibold bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">
                                {{ stat.number }}</div>
                            <div class="text-xs text-slate-400 font-poppins-regular">{{ stat.label }}</div>
                        </div>
                    </div>

                    <div class="flex gap-4 pt-4">
                        <a href="#itclub" class="group relative">
                            <div
                                class="absolute inset-0 bg-linear-to-r from-primary to-secondary rounded-xl blur opacity-75 group-hover:opacity-100 transition-all duration-300">
                            </div>
                            <div
                                class="relative px-10 py-4 bg-linear-to-r from-primary to-secondary rounded-xl font-poppins-semibold text-white transform group-hover:scale-105 transition-all duration-300 flex items-center gap-2">
                                <span>Explore Now</span>
                                <svg class="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none"
                                    stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M13 7l5 5m0 0l-5 5m5-5H6" />
                                </svg>
                            </div>
                        </a>
                        <a href="#itclub"
                            class="px-10 py-4 border-2 border-primary text-primary rounded-xl font-poppins-semibold hover:bg-primary hover:text-white transition-all duration-300 transform hover:scale-105">
                            Join IT Club
                        </a>
                    </div>
                </div>

                <div class="relative">
                    <div
                        class="absolute inset-0 bg-linear-to-r from-primary to-secondary rounded-[3rem] blur-3xl opacity-30 animate-pulse">
                    </div>
                    <div
                        class="relative bg-linear-to-br from-slate-800 to-slate-900 rounded-[3rem] p-8 border border-slate-700 shadow-2xl transform hover:scale-105 transition-all duration-500">
                        <img :src="sija" alt="">
                    </div>
                    <div
                        class="absolute -bottom-8 -right-8 w-32 h-32 bg-linear-to-br from-primary to-tertiary rounded-3xl shadow-2xl transform rotate-12 animate-spin-slow">
                    </div>
                    <div
                        class="absolute -top-8 -left-8 w-24 h-24 bg-linear-to-br from-secondary to-fourthiary rounded-2xl shadow-2xl transform -rotate-12 animate-bounce-slow">
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
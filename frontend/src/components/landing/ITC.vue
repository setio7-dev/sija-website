<script setup lang="ts">
import { ItClubDivisions } from '../../data/LandingPageData';

</script>

<template>
    <section id="itclub" class="relative py-32 px-6 overflow-hidden">
        <div class="absolute inset-0 bg-linear-to-br from-slate-900 via-slate-800 to-slate-900"></div>

        <div class="container mx-auto relative z-10">
            <div class="text-center mb-20 animate-fadeInUp">
                <div
                    class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Information
                        Technology Club</span>
                </div>
                <h2 class="text-5xl lg:text-6xl font-poppins-semibold mb-6">
                    <span class="text-white">6 Divisi</span>
                    <br />
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">IT Club
                        SIJA</span>
                </h2>
                <p class="text-xl text-slate-400 font-poppins-regular max-w-3xl mx-auto">
                    Komunitas siswa yang passionate dalam teknologi, belajar bersama dan mengembangkan project nyata
                </p>
            </div>

            <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
                <div v-for="(division, index) in ItClubDivisions" :key="index" class="group relative animate-fadeInUp"
                    :style="{ animationDelay: `${index * 0.1}s` }">
                    <div
                        :class="['absolute inset-0 bg-linear-to-br', division.gradient, 'rounded-3xl blur-xl opacity-0 group-hover:opacity-50 transition-all duration-500 transform group-hover:scale-110']">
                    </div>

                    <div
                        class="relative bg-slate-800/80 backdrop-blur-xl rounded-3xl border border-slate-700/50 p-8 hover:border-slate-600 transition-all duration-500 transform hover:scale-105 hover:-translate-y-2">
                        <div
                            :class="['w-20 h-20 rounded-2xl bg-linear-to-br', division.gradient, 'flex items-center justify-center text-4xl mb-6 transform group-hover:rotate-12 group-hover:scale-110 transition-all duration-500 shadow-2xl']">
                            {{ division.icon }}
                        </div>

                        <h3
                            class="text-2xl font-poppins-semibold text-white mb-3 group-hover:text-transparent group-hover:bg-clip-text group-hover:bg-linear-to-r group-hover:from-primary group-hover:to-secondary transition-all duration-300">
                            {{ division.name }}
                        </h3>

                        <p class="text-slate-400 font-poppins-regular mb-6 leading-relaxed">
                            {{ division.description }}
                        </p>

                        <div class="flex items-center justify-between pt-4 border-t border-slate-700/50">
                            <div class="flex items-center gap-2 text-slate-400">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                                </svg>
                                <span class="text-sm font-poppins-medium">{{ division.members }} Members</span>
                            </div>
                            <div class="flex items-center gap-2 text-slate-400">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                                </svg>
                                <span class="text-sm font-poppins-medium">{{ division.projects }} Projects</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
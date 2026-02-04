<script setup lang="ts">
import { Achievements } from '../../data/LandingPageData';

</script>

<template>
    <section id="prestasi" class="relative py-32 px-6 overflow-hidden">
        <div class="absolute inset-0 bg-linear-to-b from-slate-950 via-slate-900 to-slate-950"></div>

        <div class="container mx-auto relative z-10">
            <div class="text-center mb-20 animate-fadeInUp">
                <div
                    class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Hall of
                        Fame</span>
                </div>
                <h2 class="text-5xl lg:text-6xl font-poppins-semibold mb-6">
                    <span class="text-white">Prestasi</span>
                    <br />
                    <span
                        class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Membanggakan</span>
                </h2>
                <p class="text-xl text-slate-400 font-poppins-regular max-w-3xl mx-auto">
                    Bukti nyata kompetensi siswa SIJA di berbagai kompetisi tingkat nasional dan internasional
                </p>
            </div>

            <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
                <div v-for="(achievement, index) in Achievements" :key="index"
                    class="group relative bg-slate-800/50 backdrop-blur-sm rounded-3xl border border-slate-700/50 overflow-hidden hover:border-primary/50 transition-all duration-500 transform hover:scale-105 hover:-translate-y-2 animate-fadeInUp"
                    :style="{ animationDelay: `${index * 0.1}s` }">
                    <div
                        class="absolute inset-0 bg-linear-to-br from-primary/0 to-secondary/0 group-hover:from-primary/10 group-hover:to-secondary/10 transition-all duration-500">
                    </div>

                    <div class="relative p-8">
                        <div class="flex items-start justify-between mb-6">
                            <div class="text-6xl animate-bounce-slow">{{ achievement.medal }}</div>
                            <div class="flex flex-col gap-2">
                                <span
                                    class="px-4 py-2 bg-primary/20 border border-primary/30 rounded-full text-xs font-poppins-semibold text-primary">
                                    {{ achievement.year }}
                                </span>
                                <span
                                    class="px-4 py-2 bg-secondary/20 border border-secondary/30 rounded-full text-xs font-poppins-semibold text-secondary text-center">
                                    {{ achievement.category }}
                                </span>
                            </div>
                        </div>

                        <h3
                            class="text-xl font-poppins-semibold text-white mb-4 leading-tight group-hover:text-transparent group-hover:bg-clip-text group-hover:bg-linear-to-r group-hover:from-primary group-hover:to-secondary transition-all duration-300">
                            {{ achievement.title }}
                        </h3>

                        <div class="space-y-3 text-sm text-slate-400 font-poppins-regular mb-6">
                            <div class="flex items-center gap-2">
                                <svg class="w-5 h-5 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                                </svg>
                                <span>{{ achievement.organizer }}</span>
                            </div>
                            <div class="flex items-center gap-2">
                                <svg class="w-5 h-5 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                                </svg>
                                <span>{{ achievement.participants }} Peserta</span>
                            </div>
                        </div>

                        <div class="pt-4 border-t border-slate-700/50">
                            <button
                                class="w-full py-3 bg-slate-700/50 hover:bg-linear-to-r hover:from-primary hover:to-secondary rounded-xl font-poppins-semibold text-white transition-all duration-300 group-hover:shadow-lg group-hover:shadow-primary/50">
                                Lihat Detail
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
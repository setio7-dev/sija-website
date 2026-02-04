<script setup lang="ts">
import { Modules } from '../../data/LandingPageData';
</script>

<template>
    <section id="tentang" class="relative py-32 px-6 bg-linear-to-b from-slate-950 to-slate-900" data-aos="fade-up" data-aos-duration="1000">
        <div class="container mx-auto relative z-10">
            <div class="text-center mb-20 animate-fadeInUp">
                <div
                    class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Kurikulum
                        Pembelajaran</span>
                </div>
                <h2 class="text-5xl lg:text-6xl font-poppins-semibold mb-6">
                    <span class="text-white">Mata Pelajaran</span>
                    <br />
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Unggulan</span>
                </h2>
                <p class="text-xl text-slate-400 font-poppins-regular max-w-3xl mx-auto">
                    Materi pembelajaran yang dirancang untuk membentuk IT professional yang siap menghadapi tantangan
                    industri
                </p>
            </div>

            <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
                <div v-for="(module, index) in Modules" :key="index"
                    class="group relative bg-slate-800/50 backdrop-blur-sm rounded-3xl border border-slate-700/50 overflow-hidden hover:border-primary/50 transition-all duration-500 transform hover:scale-105 hover:-translate-y-2 animate-fadeInUp"
                    :style="{ animationDelay: `${index * 0.1}s` }">
                    <div
                        class="absolute inset-0 bg-linear-to-br from-primary/0 to-secondary/0 group-hover:from-primary/10 group-hover:to-secondary/10 transition-all duration-500">
                    </div>

                    <div class="relative aspect-video overflow-hidden">
                        <img :src="module.image" :alt="module.title"
                            class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-500" />
                        <div class="absolute inset-0 bg-linear-to-t from-slate-900 to-transparent"></div>
                        <div class="absolute top-4 right-4 text-5xl animate-float">{{ module.icon }}</div>
                        <div class="absolute bottom-4 left-4">
                            <span :class="[
                                'px-4 py-2 rounded-full text-xs font-poppins-semibold backdrop-blur-sm',
                                module.level === 'Beginner' ? 'bg-emerald-500/30 border border-emerald-500/50 text-emerald-400' :
                                    module.level === 'Intermediate' ? 'bg-blue-500/30 border border-blue-500/50 text-blue-400' :
                                        module.level === 'Advanced' ? 'bg-purple-500/30 border border-purple-500/50 text-purple-400' :
                                            'bg-red-500/30 border border-red-500/50 text-red-400'
                            ]">
                                {{ module.level }}
                            </span>
                        </div>
                    </div>

                    <div class="relative p-6 space-y-4">
                        <h3
                            class="text-2xl font-poppins-semibold text-white group-hover:text-transparent group-hover:bg-clip-text group-hover:bg-linear-to-r group-hover:from-primary group-hover:to-secondary transition-all duration-300">
                            {{ module.title }}
                        </h3>

                        <div class="flex items-center justify-between text-sm">
                            <div class="flex items-center gap-2 text-slate-400">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                                </svg>
                                <span class="font-poppins-medium">{{ module.duration }}</span>
                            </div>
                            <div class="flex items-center gap-2 text-slate-400">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                                </svg>
                                <span class="font-poppins-medium">{{ module.students }} Siswa</span>
                            </div>
                        </div>

                        <button
                            class="w-full py-3 bg-slate-700/50 hover:bg-linear-to-r hover:from-primary hover:to-secondary rounded-xl font-poppins-semibold text-white transition-all duration-500 group-hover:shadow-lg group-hover:shadow-primary/50">
                            Lihat Detail
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
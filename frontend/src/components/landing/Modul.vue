<template>
    <section id="modul" class="relative py-32 px-6 bg-linear-to-b from-slate-900 to-slate-950">
        <div class="container mx-auto relative z-10">
            <div class="text-center mb-20 animate-fadeInUp">
                <div
                    class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                    <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Learning
                        Path</span>
                </div>
                <h2 class="text-5xl lg:text-6xl font-poppins-semibold mb-6">
                    <span class="text-white">Modul</span>
                    <br />
                    <span
                        class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Pembelajaran</span>
                </h2>
                <p class="text-xl text-slate-400 font-poppins-regular max-w-3xl mx-auto">
                    Tersedia lebih dari 50+ modul pembelajaran yang dapat diakses kapan saja
                </p>
            </div>

            <div class="grid lg:grid-cols-2 gap-8 max-w-6xl mx-auto">
                <div
                    class="group relative bg-linear-to-br from-secondary/20 to-fourthiary/20 backdrop-blur-sm rounded-3xl border border-secondary/30 p-10 hover:border-secondary/50 transition-all duration-500 transform hover:scale-105 animate-fadeInUp">
                    <div
                        class="absolute inset-0 bg-linear-to-br from-secondary/0 to-fourthiary/0 group-hover:from-secondary/10 group-hover:to-fourthiary/10 rounded-3xl transition-all duration-500">
                    </div>

                    <div class="relative">
                        <div class="flex items-center gap-4 mb-8">
                            <div
                                class="w-16 h-16 bg-linear-to-br from-secondary to-fourthiary rounded-2xl flex items-center justify-center text-3xl transform group-hover:rotate-12 transition-transform duration-500 shadow-2xl">
                                📚
                            </div>
                            <div>
                                <h3 class="text-3xl font-poppins-semibold text-white">Semester 1 - 4</h3>
                                <p class="text-fourthiary font-poppins-medium">Foundation & Core Skills</p>
                            </div>
                        </div>

                        <ul class="space-y-4 font-poppins-regular text-slate-200">
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-secondary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-secondary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Dasar Jaringan Komputer</p>
                                    <p class="text-sm text-slate-400">OSI Layer, TCP/IP, Subnetting, VLAN</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-secondary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-secondary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Pemrograman Dasar & Web</p>
                                    <p class="text-sm text-slate-400">HTML, CSS, JavaScript, PHP, MySQL</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-secondary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-secondary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Database Management</p>
                                    <p class="text-sm text-slate-400">SQL, NoSQL, Database Design, ERD</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-secondary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-secondary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Sistem Operasi & Linux</p>
                                    <p class="text-sm text-slate-400">Ubuntu, CentOS, CLI, Server Administration</p>
                                </div>
                            </li>
                        </ul>

                        <button
                            class="mt-8 w-full py-4 bg-secondary/30 hover:bg-linear-to-r hover:from-secondary hover:to-fourthiary rounded-xl font-poppins-semibold text-white transition-all duration-300 border border-secondary/50 hover:border-transparent group-hover:shadow-lg group-hover:shadow-secondary/50">
                            Lihat Modul Lengkap
                        </button>
                    </div>
                </div>

                <div class="group relative bg-linear-to-br from-primary/20 to-tertiary/20 backdrop-blur-sm rounded-3xl border border-primary/30 p-10 hover:border-primary/50 transition-all duration-500 transform hover:scale-105 animate-fadeInUp"
                    style="animation-delay: 0.2s">
                    <div
                        class="absolute inset-0 bg-linear-to-br from-primary/0 to-tertiary/0 group-hover:from-primary/10 group-hover:to-tertiary/10 rounded-3xl transition-all duration-500">
                    </div>

                    <div class="relative">
                        <div class="flex items-center gap-4 mb-8">
                            <div
                                class="w-16 h-16 bg-linear-to-br from-primary to-tertiary rounded-2xl flex items-center justify-center text-3xl transform group-hover:rotate-12 transition-transform duration-500 shadow-2xl">
                                🚀
                            </div>
                            <div>
                                <h3 class="text-3xl font-poppins-semibold text-white">Semester 5 - 6</h3>
                                <p class="text-tertiary font-poppins-medium">Advanced & Industry Practice</p>
                            </div>
                        </div>

                        <ul class="space-y-4 font-poppins-regular text-slate-200">
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-primary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-primary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Advanced Networking & Security</p>
                                    <p class="text-sm text-slate-400">Routing Protocols, Firewall, VPN, IDS/IPS</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-primary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-primary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Mobile App Development</p>
                                    <p class="text-sm text-slate-400">Android Native, Flutter, React Native, API
                                        Integration</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-primary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-primary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Cloud Computing & DevOps</p>
                                    <p class="text-sm text-slate-400">AWS, Azure, Docker, Kubernetes, CI/CD</p>
                                </div>
                            </li>
                            <li class="flex items-start gap-4 group/item">
                                <div
                                    class="w-8 h-8 bg-primary/30 rounded-lg flex items-center justify-center shrink-0 group-hover/item:bg-primary transition-colors">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M5 13l4 4L19 7" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="font-poppins-semibold text-white">Praktik Kerja Industri</p>
                                    <p class="text-sm text-slate-400">6 Bulan Magang di Perusahaan IT Terkemuka</p>
                                </div>
                            </li>
                        </ul>

                        <button
                            class="mt-8 w-full py-4 bg-primary/30 hover:bg-linear-to-r hover:from-primary hover:to-tertiary rounded-xl font-poppins-semibold text-white transition-all duration-300 border border-primary/50 hover:border-transparent group-hover:shadow-lg group-hover:shadow-primary/50">
                            Lihat Modul Lengkap
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
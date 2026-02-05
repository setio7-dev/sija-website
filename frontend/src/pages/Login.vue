<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const nis = ref('');
const password = ref('');
const showPassword = ref(false);
const isLoading = ref(false);

const handleLogin = async () => {
  isLoading.value = true;
  
  setTimeout(() => {
    console.log('Login with:', { nis: nis.value, password: password.value });
    isLoading.value = false;
  }, 1500);
};

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value;
};

const navigate = useRouter();
</script>

<template>
    <section class="relative min-h-screen flex items-center justify-center px-6 overflow-hidden">
        <div class="container mx-auto relative z-10" data-aos="fade-up" data-aos-duration="1000">
            <div class="max-w-xl mx-auto">
                <div class="text-center mb-12 animate-fadeInUp">
                    <div class="inline-block px-6 py-2 bg-linear-to-r from-primary/20 to-secondary/20 backdrop-blur-sm border border-primary/30 rounded-full text-sm font-poppins-medium mb-6">
                        <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">SIJA Portal</span>
                    </div>
                    <h1 class="text-5xl lg:text-6xl font-poppins-semibold mb-4">
                        <span class="text-white">Selamat </span>
                        <span class="bg-linear-to-r from-primary to-secondary bg-clip-text text-transparent">Datang</span>
                    </h1>
                    <p class="text-lg text-slate-400 font-poppins-regular">
                        Masuk ke akun Anda untuk melanjutkan
                    </p>
                </div>

                <div class="relative bg-slate-800/50 backdrop-blur-sm rounded-3xl border border-slate-700/50 p-8 animate-fadeInUp" style="animation-delay: 0.2s;">   
                    <form @submit.prevent="handleLogin" class="relative space-y-6">
                        <div>
                            <label for="nis" class="block text-sm font-poppins-semibold text-slate-300 mb-2">
                                Nomor Induk Siswa (NIS)
                            </label>
                            <div class="relative">
                                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                                    <svg class="w-5 h-5 text-slate-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                    </svg>
                                </div>
                                <input
                                    id="nis"
                                    v-model="nis"
                                    type="text"
                                    required
                                    placeholder="Masukkan NIS Anda"
                                    class="w-full pl-12 pr-4 py-4 bg-slate-900/50 border border-slate-700/50 rounded-xl text-white placeholder-slate-500 font-poppins-regular focus:outline-none focus:border-primary/50 focus:ring-2 focus:ring-primary/20 transition-all duration-300"
                                />
                            </div>
                        </div>

                        <div>
                            <label for="password" class="block text-sm font-poppins-semibold text-slate-300 mb-2">
                                Kata Sandi
                            </label>
                            <div class="relative">
                                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                                    <svg class="w-5 h-5 text-slate-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                                    </svg>
                                </div>
                                <input
                                    id="password"
                                    v-model="password"
                                    :type="showPassword ? 'text' : 'password'"
                                    required
                                    placeholder="Masukkan kata sandi Anda"
                                    class="w-full pl-12 pr-12 py-4 bg-slate-900/50 border border-slate-700/50 rounded-xl text-white placeholder-slate-500 font-poppins-regular focus:outline-none focus:border-primary/50 focus:ring-2 focus:ring-primary/20 transition-all duration-300"
                                />
                                <button
                                    type="button"
                                    @click="togglePasswordVisibility"
                                    class="absolute inset-y-0 right-0 pr-4 flex items-center text-slate-500 hover:text-primary transition-colors duration-300"
                                >
                                    <svg v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                                    </svg>
                                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
                                    </svg>
                                </button>
                            </div>
                        </div>

                        <button
                            type="submit"
                            :disabled="isLoading"
                            class="w-full py-4 bg-linear-to-r from-primary to-secondary hover:from-primary/80 hover:to-secondary/80 rounded-xl font-poppins-semibold text-white transition-all duration-300 shadow-lg hover:shadow-primary/50 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
                        >
                            <svg v-if="isLoading" class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            <span>{{ isLoading ? 'Memproses...' : 'Masuk' }}</span>
                        </button>
                    </form>

                    <div class="mt-8 pt-6 border-t border-slate-700/50">
                        <p class="text-center text-sm text-slate-400 font-poppins-regular">
                            Tidak memiliki akun? 
                            <span @click="() => navigate.push('/')" class="font-poppins-semibold cursor-pointer text-primary hover:text-secondary transition-colors duration-300">
                                Silahkan kembali
                            </span>
                        </p>
                    </div>
                </div>

                <div class="mt-8 text-center">
                    <p class="text-sm text-slate-500 font-poppins-regular">
                        © 2026 SIJA. All rights reserved.
                    </p>
                </div>
            </div>
        </div>
    </section>
</template>
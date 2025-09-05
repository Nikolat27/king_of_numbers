<template>
  <div class="min-h-screen bg-bg-primary dark:bg-deep-navy transition-colors duration-300">
    <!-- Header -->
    <header class="bg-bg-secondary dark:bg-gray-800 shadow-lg border-b border-light-gray dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <!-- Logo/Title -->
          <div class="flex items-center">
            <h1 class="text-2xl font-bold text-gold">♔ King of Numbers</h1>
          </div>

          <!-- Navigation & Theme Toggle -->
          <div class="flex items-center space-x-4">
            <!-- Theme Toggle Button -->
            <button
              @click="toggleTheme"
              class="p-2 rounded-lg bg-light-gray dark:bg-gray-600 hover:bg-gray-300 dark:hover:bg-gray-500
                     transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gold"
              aria-label="Toggle theme"
            >
              <svg
                v-if="colorMode.preference === 'light'"
                class="w-5 h-5 text-deep-navy"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
              </svg>
              <svg
                v-else
                class="w-5 h-5 text-soft-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-bg-secondary dark:bg-gray-800 border-t border-light-gray dark:border-gray-700 mt-auto">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="text-center text-text-secondary text-sm">
          <p>&copy; 2024 King of Numbers. Built with Nuxt 3 & Tailwind CSS.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
// Theme management
const colorMode = useColorMode()

const toggleTheme = () => {
  colorMode.preference = colorMode.preference === 'light' ? 'dark' : 'light'
}
</script>

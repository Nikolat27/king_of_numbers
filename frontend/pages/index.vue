<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <section class="relative py-20 px-4 sm:px-6 lg:px-8">
      <div class="max-w-6xl mx-auto text-center">
        <!-- Main Title -->
        <div class="mb-8">
          <h1 class="text-5xl md:text-8xl font-bold text-gold mb-6 drop-shadow-2xl animate-pulse">
            ♔ KING OF NUMBERS ♔
          </h1>
          <div class="w-24 h-1 bg-gold mx-auto mb-6 rounded-full"></div>
          <p class="text-xl md:text-2xl text-text-secondary dark:text-soft-white max-w-3xl mx-auto leading-relaxed">
            Inspired by Alice in Borderland's King of Diamonds.
            <br class="hidden md:block">
            Choose wisely, calculate perfectly, or face elimination.
          </p>
        </div>

        <!-- CTA Button -->
        <div class="mb-16">
          <NuxtLink
            to="/game"
            class="inline-block bg-gold hover:bg-yellow-500 text-deep-navy font-bold text-xl px-12 py-4 rounded-xl
                     transform hover:scale-105 transition-all duration-300 shadow-2xl hover:shadow-gold/50
                     border-4 border-gold hover:border-yellow-500"
          >
            🎯 ENTER THE GAME
          </NuxtLink>
        </div>

        <!-- Game Stats Preview -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto mb-16">
          <div class="card text-center transform hover:scale-105 transition-transform duration-300">
            <div class="text-4xl mb-2">🎲</div>
            <h3 class="text-lg font-bold text-text-primary mb-2">Strategic Choice</h3>
            <p class="text-text-secondary text-sm">Numbers 0-100 determine your fate</p>
          </div>
          <div class="card text-center transform hover:scale-105 transition-transform duration-300">
            <div class="text-4xl mb-2">⚖️</div>
            <h3 class="text-lg font-bold text-text-primary mb-2">Mathematical Precision</h3>
            <p class="text-text-secondary text-sm">Average × 0.8 = target number</p>
          </div>
          <div class="card text-center transform hover:scale-105 transition-transform duration-300">
            <div class="text-4xl mb-2">🏆</div>
            <h3 class="text-lg font-bold text-text-primary mb-2">Survival of the Smartest</h3>
            <p class="text-text-secondary text-sm">Closest number wins the round</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Game Rules Section -->
    <section class="py-16 px-4 sm:px-6 lg:px-8 bg-bg-secondary-light dark:bg-bg-secondary-dark">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-3xl md:text-4xl font-bold text-center text-text-primary mb-12">
          🎯 GAME MECHANICS
        </h2>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <!-- Basic Rules -->
          <div class="card">
            <h3 class="text-xl font-bold text-gold mb-4 flex items-center">
              <span class="text-2xl mr-2">📋</span> Basic Rules
            </h3>
            <ul class="space-y-3 text-text-secondary">
              <li class="flex items-start">
                <span class="text-gold mr-2 mt-1">•</span>
                <span>Choose a number between <strong>0 and 100</strong></span>
              </li>
              <li class="flex items-start">
                <span class="text-gold mr-2 mt-1">•</span>
                <span>Target = <strong>Average of all numbers × 0.8</strong></span>
              </li>
              <li class="flex items-start">
                <span class="text-gold mr-2 mt-1">•</span>
                <span><strong>Closest number to target wins</strong> the round</span>
              </li>
              <li class="flex items-start">
                <span class="text-gold mr-2 mt-1">•</span>
                <span>Four other players <strong>lose 1 point each</strong></span>
              </li>
            </ul>
          </div>

          <!-- Special Rules -->
          <div class="card">
            <h3 class="text-xl font-bold text-gold mb-4 flex items-center">
              <span class="text-2xl mr-2">⚡</span> Dynamic Rules
            </h3>
            <div class="space-y-4">
              <div class="border-l-4 border-gold pl-4">
                <h4 class="font-semibold text-alert-red">1 Player Eliminated</h4>
                <p class="text-sm text-text-secondary">Duplicate numbers become invalid</p>
              </div>
              <div class="border-l-4 border-gold pl-4">
                <h4 class="font-semibold text-alert-red">2 Players Eliminated</h4>
                <p class="text-sm text-text-secondary">Exact matches cause others to lose 2 points</p>
              </div>
              <div class="border-l-4 border-gold pl-4">
                <h4 class="font-semibold text-alert-red">3 Players Eliminated</h4>
                <p class="text-sm text-text-secondary">0 vs 100 becomes a winning combination</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Win/Lose Conditions -->
        <div class="card mt-8 text-center">
          <h3 class="text-xl font-bold text-text-primary mb-4">🏆 GAME END CONDITIONS</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="p-4 bg-green-100 dark:bg-green-900 rounded-lg">
              <h4 class="font-bold text-green-800 dark:text-green-200 mb-2">GAME CLEAR</h4>
              <p class="text-green-700 dark:text-green-300">Last remaining player wins!</p>
            </div>
            <div class="p-4 bg-red-100 dark:bg-red-900 rounded-lg">
              <h4 class="font-bold text-red-800 dark:text-red-200 mb-2">GAME OVER</h4>
              <p class="text-red-700 dark:text-red-300">Reach -10 points = Acid Death</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-16 px-4 sm:px-6 lg:px-8">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-text-primary mb-6">
          Ready to Test Your Mathematical Survival Skills?
        </h2>
        <p class="text-lg text-text-secondary mb-8 max-w-2xl mx-auto">
          Every choice matters. Every calculation counts. Will you be the King of Numbers,
          or will you fall victim to the deadly game?
        </p>

        <div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
          <NuxtLink
            to="/game"
            class="btn-primary text-lg px-8 py-4 transform hover:scale-105 transition-transform duration-300"
          >
            🎲 START YOUR JOURNEY
          </NuxtLink>

          <button
            @click="toggleTheme"
            class="btn-secondary text-lg px-6 py-4"
          >
            {{ isDark ? '☀️ Light Mode' : '🌙 Dark Mode' }}
          </button>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-8 px-4 sm:px-6 lg:px-8 border-t border-light-gray dark:border-gray-600">
      <div class="max-w-4xl mx-auto text-center">
        <p class="text-text-secondary">
          Inspired by <strong>Alice in Borderland</strong> • Built with Vue.js & Nuxt
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
// Theme management
const colorMode = useColorMode()
const isDark = computed(() => colorMode.preference === 'dark')

const toggleTheme = () => {
  colorMode.preference = isDark.value ? 'light' : 'dark'
}
</script>

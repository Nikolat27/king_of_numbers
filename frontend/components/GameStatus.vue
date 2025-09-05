<template>
  <div class="grid grid-cols-2 md:grid-cols-3 gap-3 md:gap-6 mb-6 md:mb-8">
    <!-- Round -->
    <div class="card text-center p-3 md:p-6 select-none cursor-default">
      <h3 class="text-xs md:text-sm font-medium text-text-secondary mb-1 md:mb-2">Round</h3>
      <div class="text-2xl md:text-3xl font-bold text-gold">{{ currentRound }}</div>
    </div>

    <!-- Score -->
    <div class="card text-center p-3 md:p-6 select-none cursor-default">
      <h3 class="text-xs md:text-sm font-medium text-text-secondary mb-1 md:mb-2">Score</h3>
      <div class="text-2xl md:text-3xl font-bold" :class="score < 0 ? 'text-alert-red' : 'text-accent-cyan'">
        {{ score }}
      </div>
    </div>

    <!-- Mobile-only eliminated players -->
    <div class="card text-center p-3 md:hidden select-none cursor-default">
      <h3 class="text-xs font-medium text-text-secondary mb-1">Eliminated</h3>
      <div class="text-xl font-bold text-alert-red">{{ eliminatedPlayers }}</div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  currentRound: {
    type: Number,
    required: true
  },
  score: {
    type: Number,
    required: true
  },
  eliminatedPlayers: {
    type: Number,
    required: true
  }
})
</script>

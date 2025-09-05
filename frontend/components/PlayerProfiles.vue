<template>
  <div class="grid grid-cols-5 gap-2 md:gap-4 mb-6 md:mb-8">
    <div
      v-for="(player, index) in players"
      :key="index"
      class="card text-center transform hover:scale-105 transition-transform duration-300 p-2 md:p-4 select-none cursor-default"
      :class="index === 0 ? 'ring-2 ring-gold' : ''"
    >
      <div
        class="w-12 h-12 md:w-16 md:h-16 mx-auto mb-2 md:mb-3 rounded-full bg-gradient-to-br from-gold to-yellow-600 flex items-center justify-center text-lg md:text-2xl font-bold text-deep-navy">
        {{ player.avatar }}
      </div>
      <h4 class="font-bold text-text-primary text-xs md:text-sm mb-1 truncate">{{ player.name }}</h4>
      <div class="text-sm md:text-base text-text-secondary font-semibold">
        <span :class="player.score < 0 ? 'text-alert-red' : 'text-accent-cyan'">{{ player.score }}</span>
      </div>
      <div v-if="index === 0" class="text-xs text-gold font-medium mt-1">YOU</div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  players: {
    type: Array,
    required: true
  }
})
</script>

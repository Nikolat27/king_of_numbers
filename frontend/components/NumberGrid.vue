<template>
  <div class="max-w-2xl mx-auto px-2">
    <div class="grid grid-cols-10 gap-1 md:gap-1 mb-4 md:mb-6">
      <button
        v-for="number in 101"
        :key="number - 1"
        @click="$emit('number-selected', number - 1)"
        class="aspect-square text-xs md:text-sm font-bold rounded-md md:rounded-lg border-2 transition-all duration-200
               min-h-[32px] md:min-h-[40px] touch-manipulation
               hover:scale-110 hover:shadow-lg active:scale-95"
        :class="selectedNumber === (number - 1)
          ? 'bg-gold text-white border-gold shadow-lg scale-110'
          : 'bg-bg-secondary-light dark:bg-bg-secondary-dark text-text-primary border-light-gray dark:border-gray-600 hover:border-gold'"
      >
        {{ number - 1 }}
      </button>
    </div>

    <!-- Selected Number Display -->
    <div v-if="selectedNumber !== null" class="mb-4 md:mb-6">
      <p class="text-base md:text-lg text-text-secondary">Selected: <span class="font-bold text-gold text-lg md:text-xl">{{
        selectedNumber }}</span></p>
    </div>

    <!-- Submit Button -->
    <button
      @click="$emit('submit-number')"
      class="btn-primary w-full md:w-auto text-base md:text-lg py-3 md:py-2 px-6 md:px-4 min-h-[48px] md:min-h-[40px] touch-manipulation"
      :disabled="selectedNumber === null"
    >
      Confirm Selection
    </button>
  </div>
</template>

<script setup>
defineProps({
  selectedNumber: {
    type: Number,
    default: null
  }
})

defineEmits(['number-selected', 'submit-number'])
</script>

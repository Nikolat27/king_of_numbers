<template>
  <!-- Results Modal -->
  <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-75 backdrop-blur-md flex items-center justify-center z-50 p-4 transition-all duration-300" :class="modalStep < 5 ? 'cursor-none' : ''" @click.self="$event.stopPropagation()">
    <div class="bg-bg-primary dark:bg-deep-navy rounded-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto shadow-2xl border border-light-gray dark:border-gray-600" ref="modalContent">
      <div class="p-6 md:p-8">
        <div class="text-center mb-6">
          <h3 class="text-2xl md:text-3xl font-bold text-gold mb-2">🎭 Round {{ currentRound }} Results</h3>
          <div class="w-16 h-1 bg-gold mx-auto rounded-full"></div>
        </div>

        <!-- Step 1: Reveal Player Choices -->
        <div v-if="modalStep >= 1" class="mb-6">
          <h4 class="text-lg md:text-xl font-bold text-white mb-4 flex items-center">
            <span class="text-2xl mr-2">👥</span> Player Choices Revealed
          </h4>
          <div class="grid grid-cols-1 md:grid-cols-5 gap-3 mb-4">
            <div
              v-for="(choice, index) in revealedChoices"
              :key="index"
              class="card text-center p-3 transition-all duration-500"
              :class="index === 0 ? 'ring-2 ring-gold' : ''"
            >
              <div class="text-2xl mb-2">{{ players[index].avatar }}</div>
              <div class="font-bold text-text-primary text-sm">{{ players[index].name }}</div>
              <div class="text-xl md:text-2xl font-bold text-gold mt-1">{{ choice }}</div>
            </div>
          </div>
        </div>

        <!-- Step 2: Combine Numbers -->
        <div v-if="modalStep >= 2" class="mb-6">
          <h4 class="text-lg md:text-xl font-bold text-white mb-4 flex items-center">
            <span class="text-2xl mr-2">🔢</span> All Numbers Combined
          </h4>
          <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-4 rounded-lg">
            <p class="text-text-secondary text-center text-lg">
              <span class="font-bold text-gold">{{ allChoices.join(' + ') }}</span>
              <span class="mx-2">=</span>
              <span class="font-bold text-accent-cyan text-xl">{{ sumOfChoices }}</span>
            </p>
          </div>
        </div>

        <!-- Step 3: Calculate Average -->
        <div v-if="modalStep >= 3" class="mb-6">
          <h4 class="text-lg md:text-xl font-bold text-white mb-4 flex items-center">
            <span class="text-2xl mr-2">📊</span> Calculate Average
          </h4>
          <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-4 rounded-lg">
            <div class="text-center">
              <p class="text-text-secondary text-lg mb-2">
                <span class="font-bold text-gold">{{ sumOfChoices }}</span>
                <span class="mx-2">÷</span>
                <span class="font-bold text-gold">5</span>
                <span class="mx-2">=</span>
                <span class="font-bold text-accent-cyan text-xl">{{ average.toFixed(2) }}</span>
              </p>
              <p class="text-sm text-text-secondary">Average of all 5 numbers</p>
            </div>
          </div>
        </div>

        <!-- Step 4: Apply Formula -->
        <div v-if="modalStep >= 4" class="mb-6">
          <h4 class="text-lg md:text-xl font-bold text-white mb-4 flex items-center">
            <span class="text-2xl mr-2">🎯</span> Apply Target Formula
          </h4>
          <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-4 rounded-lg">
            <div class="text-center">
              <p class="text-text-secondary text-lg mb-2">
                <span class="font-bold text-accent-cyan">{{ average.toFixed(2) }}</span>
                <span class="mx-2">×</span>
                <span class="font-bold text-gold">0.8</span>
                <span class="mx-2">=</span>
                <span class="font-bold text-green-600 text-xl">{{ target.toFixed(2) }}</span>
              </p>
              <p class="text-sm text-text-secondary">Target = Average × 0.8</p>
            </div>
          </div>
        </div>

        <!-- Step 5: Reveal Winner -->
        <div v-if="modalStep >= 5" class="mb-6">
          <h4 class="text-lg md:text-xl font-bold text-white mb-4 flex items-center">
            <span class="text-2xl mr-2">🏆</span> Winner Revealed
          </h4>
          <div class="text-center">
            <div class="text-4xl md:text-5xl mb-4">{{ isWinner ? '🎉' : '💔' }}</div>
            <p class="text-xl md:text-2xl font-bold mb-2" :class="isWinner ? 'text-green-600' : 'text-alert-red'">
              {{ isWinner ? 'You Won!' : 'You Lost!' }}
            </p>
            <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-4 rounded-lg">
              <p class="text-text-secondary mb-2">Target: <span class="font-bold text-green-600">{{ target.toFixed(2) }}</span></p>
              <p class="text-text-secondary mb-2">Winner's Choice: <span class="font-bold text-accent-cyan">{{ winnerChoice }}</span></p>
              <p class="text-text-secondary">Winner's Distance: <span class="font-bold text-green-600">{{ winnerDistance.toFixed(2) }}</span></p>
            </div>
          </div>
        </div>

        <!-- Continue Button -->
        <div v-if="modalStep >= 5" class="text-center">
          <button
            @click="$emit('continue')"
            class="btn-primary min-h-[48px] touch-manipulation text-base md:text-lg py-3 px-6"
          >
            Continue to Next Round
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const props = defineProps({
  showModal: {
    type: Boolean,
    default: false
  },
  modalStep: {
    type: Number,
    default: 0
  },
  currentRound: {
    type: Number,
    required: true
  },
  players: {
    type: Array,
    required: true
  },
  allChoices: {
    type: Array,
    required: true
  },
  revealedChoices: {
    type: Array,
    default: () => []
  },
  sumOfChoices: {
    type: Number,
    required: true
  },
  average: {
    type: Number,
    required: true
  },
  target: {
    type: Number,
    required: true
  },
  winnerChoice: {
    type: [Number, String],
    required: true
  },
  winnerDistance: {
    type: Number,
    required: true
  },
  isWinner: {
    type: Boolean,
    required: true
  }
})

const emit = defineEmits(['continue'])

const modalContent = ref(null)

// Auto-scroll function
const scrollToBottom = () => {
  if (modalContent.value) {
    setTimeout(() => {
      modalContent.value.scrollTop = modalContent.value.scrollHeight
    }, 100)
  }
}

// Watch for modal step changes to auto-scroll
watch(() => props.modalStep, (newStep) => {
  if (newStep >= 1) {
    scrollToBottom()
  }
})
</script>

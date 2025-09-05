<template>
  <div class="min-h-screen py-4 md:py-8">
    <div class="max-w-4xl mx-auto px-2 sm:px-4 md:px-6 lg:px-8">
      <!-- Game Header -->
      <div class="text-center mb-6 md:mb-8 px-4">
        <NuxtLink to="/" class="inline-block">
          <h1 class="text-3xl md:text-6xl font-bold text-gold mb-3 md:mb-4 drop-shadow-lg cursor-pointer hover:text-yellow-400 transition-colors duration-300">
            ♔ King of Numbers ♔
          </h1>
        </NuxtLink>
        <p class="text-base md:text-lg text-text-secondary dark:text-soft-white max-w-2xl mx-auto leading-relaxed">
          Choose your number wisely. Survival depends on your mathematical strategy!
        </p>
      </div>

      <!-- Player Profiles Component -->
      <PlayerProfiles :players="players" />

      <!-- Game Status Component -->
      <GameStatus
        :current-round="gameState.currentRound"
        :score="gameState.score"
        :eliminated-players="gameState.eliminatedPlayers"
      />

      <!-- Game Area -->
      <div class="card mb-8">
        <!-- Current Rules Display -->
        <div v-if="gameState.eliminatedPlayers > 0"
          class="mb-4 md:mb-6 p-3 md:p-4 bg-alert-red/10 border border-alert-red/20 rounded-lg mx-2 md:mx-0">
          <h4 class="font-bold text-alert-red mb-2 text-sm md:text-base">⚠️ ACTIVE RULE CHANGES</h4>
          <div v-if="gameState.eliminatedPlayers >= 1" class="text-xs md:text-sm text-alert-red mb-1">
            • Duplicate numbers are invalid
          </div>
          <div v-if="gameState.eliminatedPlayers >= 2" class="text-xs md:text-sm text-alert-red mb-1">
            • Exact matches cause others to lose 2 points
          </div>
          <div v-if="gameState.eliminatedPlayers >= 3" class="text-xs md:text-sm text-alert-red mb-1">
            • 0 vs 100 is a winning combination
          </div>
        </div>

        <!-- Input Phase -->
        <div v-if="gameState.gamePhase === 'input'" class="text-center">
          <h3 class="text-xl md:text-2xl font-bold text-text-primary mb-4 md:mb-6 px-2">Choose Your Number (0-100)</h3>
          <p class="text-sm md:text-base text-text-secondary mb-4 px-2">Click on any number to select it</p>

          <!-- Number Grid Component -->
          <NumberGrid
            :selected-number="selectedNumber"
            @number-selected="selectNumber"
            @submit-number="submitNumber"
          />
        </div>

        <!-- Waiting Phase -->
        <div v-else-if="gameState.gamePhase === 'waiting'" class="text-center px-4">
          <div class="text-5xl md:text-6xl mb-4">⏳</div>
          <h3 class="text-xl md:text-2xl font-bold text-text-primary mb-4">Waiting for Other Players</h3>
          <p class="text-sm md:text-base text-text-secondary">All 5 players must submit their numbers...</p>
        </div>

        <!-- Results Modal -->
        <div v-if="showResultsModal" class="fixed inset-0 bg-black bg-opacity-75 backdrop-blur-md flex items-center justify-center z-50 p-4 transition-all duration-300" :class="modalStep < 5 ? 'cursor-none' : ''" @click.self="$event.stopPropagation()">
          <div class="bg-bg-primary dark:bg-deep-navy rounded-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto shadow-2xl border border-light-gray dark:border-gray-600" ref="modalContent">
            <div class="p-6 md:p-8">
              <div class="text-center mb-6">
                <h3 class="text-2xl md:text-3xl font-bold text-gold mb-2">🎭 Round {{ gameState.currentRound }} Results</h3>
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
                    <span class="font-bold text-gold">{{ gameState.allChoices.join(' + ') }}</span>
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
                      <span class="font-bold text-accent-cyan text-xl">{{ gameState.average.toFixed(2) }}</span>
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
                      <span class="font-bold text-accent-cyan">{{ gameState.average.toFixed(2) }}</span>
                      <span class="mx-2">×</span>
                      <span class="font-bold text-gold">0.8</span>
                      <span class="mx-2">=</span>
                      <span class="font-bold text-green-600 text-xl">{{ gameState.target.toFixed(2) }}</span>
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
                  <div class="text-4xl md:text-5xl mb-4">{{ gameState.isWinner ? '🎉' : '💔' }}</div>
                  <p class="text-xl md:text-2xl font-bold mb-2" :class="gameState.isWinner ? 'text-green-600' : 'text-alert-red'">
                    {{ gameState.isWinner ? 'You Won!' : 'You Lost!' }}
                  </p>
                  <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-4 rounded-lg">
                    <p class="text-text-secondary mb-2">Target: <span class="font-bold text-green-600">{{ gameState.target.toFixed(2) }}</span></p>
                    <p class="text-text-secondary mb-2">Winner's Choice: <span class="font-bold text-accent-cyan">{{ winnerChoice }}</span></p>
                    <p class="text-text-secondary">Winner's Distance: <span class="font-bold text-green-600">{{ winnerDistance.toFixed(2) }}</span></p>
                  </div>
                </div>
              </div>

              <!-- Continue Button -->
              <div v-if="modalStep >= 5" class="text-center">
                <button
                  @click="closeResultsModal"
                  class="btn-primary min-h-[48px] touch-manipulation text-base md:text-lg py-3 px-6"
                >
                  Continue to Next Round
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Results Phase -->
        <div v-else-if="gameState.gamePhase === 'results'" class="text-center">
          <h3 class="text-xl md:text-2xl font-bold text-text-primary mb-4 md:mb-6 px-2">Round Results</h3>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6 mb-4 md:mb-6 px-2">
            <div class="card p-3 md:p-6">
              <h4 class="font-bold text-gold mb-3 text-sm md:text-base">📊 Mathematical Analysis</h4>
              <div class="space-y-2">
                <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-2 rounded">
                  <p class="text-text-secondary text-xs md:text-sm font-medium">All Numbers: {{ allPlayerChoices.join(', ') }}</p>
                </div>
                <div class="grid grid-cols-2 gap-2 text-xs md:text-sm">
                  <div class="text-text-secondary">Sum: <span class="font-bold text-gold">{{ sumOfChoices }}</span></div>
                  <div class="text-text-secondary">Count: <span class="font-bold text-gold">5</span></div>
                  <div class="text-text-secondary">Average: <span class="font-bold text-accent-cyan">{{ gameState.average.toFixed(2) }}</span></div>
                  <div class="text-text-secondary">Target: <span class="font-bold text-accent-cyan">{{ gameState.target.toFixed(2) }}</span></div>
                </div>
                <div class="border-t border-light-gray dark:border-gray-600 pt-2 mt-2">
                  <p class="text-text-secondary text-xs md:text-sm">Your Choice: <span class="font-bold text-gold">{{ gameState.playerChoice }}</span></p>
                  <p class="text-text-secondary text-xs md:text-sm">Your Distance: <span class="font-bold" :class="gameState.distance < 1 ? 'text-green-600' : 'text-alert-red'">{{ gameState.distance.toFixed(2) }}</span></p>
                  <p class="text-text-secondary text-xs md:text-sm">Winner's Choice: <span class="font-bold text-accent-cyan">{{ winnerChoice }}</span></p>
                  <p class="text-text-secondary text-xs md:text-sm">Winner's Distance: <span class="font-bold text-green-600">{{ winnerDistance.toFixed(2) }}</span></p>
                </div>
              </div>
            </div>

            <div class="card p-3 md:p-6">
              <h4 class="font-bold text-gold mb-3 text-sm md:text-base">🏆 Round Results</h4>
              <div class="text-center">
                <div class="text-3xl md:text-4xl mb-3">{{ gameState.isWinner ? '🎉' : '💔' }}</div>
                <p class="text-base md:text-lg font-bold mb-2" :class="gameState.isWinner ? 'text-green-600' : 'text-alert-red'">
                  {{ gameState.isWinner ? 'Victory!' : 'Defeat!' }}
                </p>
                <div class="bg-bg-secondary-light dark:bg-bg-secondary-dark p-3 rounded mb-3">
                  <p class="text-text-secondary text-xs md:text-sm mb-1">{{ gameState.roundResult }}</p>
                  <p class="text-text-secondary text-xs md:text-sm">Score Change: <span :class="gameState.scoreChange < 0 ? 'text-alert-red' : 'text-accent-cyan'">{{ gameState.scoreChange }}</span></p>
                </div>
                <div class="text-xs md:text-sm text-text-secondary">
                  <p>Round {{ gameState.currentRound }} Complete</p>
                  <p>Ready for next challenge?</p>
                </div>
              </div>
            </div>
          </div>

          <button @click="nextRound" class="btn-primary min-h-[48px] touch-manipulation text-base md:text-lg py-3 md:py-2 px-6 md:px-4">
            Next Round
          </button>
        </div>

        <!-- Results Modal Component -->
        <ResultsModal
          :show-modal="showResultsModal"
          :modal-step="modalStep"
          :current-round="gameState.currentRound"
          :players="players"
          :all-choices="allPlayerChoices"
          :revealed-choices="revealedChoices"
          :sum-of-choices="sumOfChoices"
          :average="gameState.average"
          :target="gameState.target"
          :winner-choice="winnerChoice"
          :winner-distance="winnerDistance"
          :is-winner="gameState.isWinner"
          @continue="closeResultsModal"
        />

        <!-- Game Over -->
        <div v-if="gameState.gamePhase === 'gameOver'" class="text-center px-4">
          <div class="text-5xl md:text-6xl mb-4">{{ gameState.score >= 0 ? '🏆' : '💀' }}</div>
          <h3 class="text-2xl md:text-3xl font-bold mb-4" :class="gameState.score >= 0 ? 'text-green-600' : 'text-alert-red'">
            {{ gameState.score >= 0 ? 'GAME CLEAR!' : 'GAME OVER' }}
          </h3>
          <p class="text-lg md:text-xl text-text-secondary mb-6">{{ gameState.gameOverMessage }}</p>

          <div class="flex flex-col sm:flex-row gap-3 md:gap-4 justify-center items-center">
            <button @click="resetGame" class="btn-primary w-full sm:w-auto min-h-[48px] touch-manipulation">
              Play Again
            </button>
            <NuxtLink to="/" class="btn-secondary w-full sm:w-auto min-h-[48px] touch-manipulation text-center">
              Back to Home
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Navigation -->
      <div class="text-center mt-6 md:mt-8 px-4">
        <NuxtLink to="/" class="inline-block text-text-secondary hover:text-gold transition-colors text-sm md:text-base py-2 px-4 min-h-[44px] touch-manipulation rounded-lg hover:bg-bg-secondary-light dark:hover:bg-bg-secondary-dark">
          ← Back to Home
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
// Use the game logic composable
const {
  gameState,
  showResultsModal,
  modalStep,
  revealedChoices,
  selectedNumber,
  players,
  allPlayerChoices,
  sumOfChoices,
  winnerChoice,
  winnerDistance,
  selectNumber,
  submitNumber,
  nextRound,
  resetGame,
  closeResultsModal
} = useGameLogic()

// All game logic is now handled by the useGameLogic composable above
</script>

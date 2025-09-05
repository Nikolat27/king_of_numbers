<template>
  <div class="min-h-screen py-4 md:py-8">
    <div class="max-w-4xl mx-auto px-2 sm:px-4 md:px-6 lg:px-8">
      <!-- Game Header -->
      <div class="text-center mb-6 md:mb-8 px-4">
        <h1 class="text-3xl md:text-6xl font-bold text-gold mb-3 md:mb-4 drop-shadow-lg">
          ♔ King of Numbers ♔
        </h1>
        <p class="text-base md:text-lg text-text-secondary dark:text-soft-white max-w-2xl mx-auto leading-relaxed">
          Choose your number wisely. Survival depends on your mathematical strategy!
        </p>
      </div>

      <!-- Player Profiles -->
      <div class="grid grid-cols-5 gap-2 md:gap-4 mb-6 md:mb-8">
        <div v-for="(player, index) in players" :key="index"
          class="card text-center transform hover:scale-105 transition-transform duration-300 p-2 md:p-4"
          :class="index === 0 ? 'ring-2 ring-gold' : ''">
          <div
            class="w-12 h-12 md:w-16 md:h-16 mx-auto mb-2 md:mb-3 rounded-full bg-gradient-to-br from-gold to-yellow-600 flex items-center justify-center text-lg md:text-2xl font-bold text-deep-navy">
            {{ player.avatar }}
          </div>
          <h4 class="font-bold text-text-primary text-xs md:text-sm mb-1 truncate">{{ player.name }}</h4>
          <div class="text-xs text-text-secondary">
            <span :class="player.score < 0 ? 'text-alert-red' : 'text-accent-cyan'">{{ player.score }}</span>
          </div>
          <div v-if="index === 0" class="text-xs text-gold font-medium mt-1">YOU</div>
        </div>
      </div>

      <!-- Game Status -->
      <div class="grid grid-cols-2 md:grid-cols-3 gap-3 md:gap-6 mb-6 md:mb-8">
        <!-- Round -->
        <div class="card text-center p-3 md:p-6">
          <h3 class="text-xs md:text-sm font-medium text-text-secondary mb-1 md:mb-2">Round</h3>
          <div class="text-2xl md:text-3xl font-bold text-gold">{{ gameState.currentRound }}</div>
        </div>

        <!-- Score -->
        <div class="card text-center p-3 md:p-6">
          <h3 class="text-xs md:text-sm font-medium text-text-secondary mb-1 md:mb-2">Score</h3>
          <div class="text-2xl md:text-3xl font-bold" :class="gameState.score < 0 ? 'text-alert-red' : 'text-accent-cyan'">
            {{ gameState.score }}
          </div>
        </div>

        <!-- Mobile-only eliminated players -->
        <div class="card text-center p-3 md:hidden">
          <h3 class="text-xs font-medium text-text-secondary mb-1">Eliminated</h3>
          <div class="text-xl font-bold text-alert-red">{{ gameState.eliminatedPlayers }}</div>
        </div>
      </div>

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

          <!-- Number Grid -->
          <div class="max-w-2xl mx-auto px-2">
            <div class="grid grid-cols-10 gap-1 md:gap-1 mb-4 md:mb-6">
              <button v-for="number in 101" :key="number - 1" @click="selectNumber(number - 1)" class="aspect-square text-xs md:text-sm font-bold rounded-md md:rounded-lg border-2 transition-all duration-200
                       min-h-[32px] md:min-h-[40px] touch-manipulation
                       hover:scale-110 hover:shadow-lg active:scale-95"
                :class="selectedNumber === (number - 1)
                  ? 'bg-gold text-deep-navy border-gold shadow-lg scale-110'
                  : 'bg-bg-secondary-light dark:bg-bg-secondary-dark text-text-primary border-light-gray dark:border-gray-600 hover:border-gold'">
                {{ number - 1 }}
              </button>
            </div>

            <!-- Selected Number Display -->
            <div v-if="selectedNumber !== null" class="mb-4 md:mb-6">
              <p class="text-base md:text-lg text-text-secondary">Selected: <span class="font-bold text-gold text-lg md:text-xl">{{
                selectedNumber }}</span></p>
            </div>

            <!-- Submit Button -->
            <button @click="submitNumber" class="btn-primary w-full md:w-auto text-base md:text-lg py-3 md:py-2 px-6 md:px-4 min-h-[48px] md:min-h-[40px] touch-manipulation" :disabled="selectedNumber === null">
              Confirm Selection
            </button>
          </div>
        </div>

        <!-- Waiting Phase -->
        <div v-else-if="gameState.gamePhase === 'waiting'" class="text-center px-4">
          <div class="text-5xl md:text-6xl mb-4">⏳</div>
          <h3 class="text-xl md:text-2xl font-bold text-text-primary mb-4">Waiting for Other Players</h3>
          <p class="text-sm md:text-base text-text-secondary">All 5 players must submit their numbers...</p>
        </div>

        <!-- Results Phase -->
        <div v-else-if="gameState.gamePhase === 'results'" class="text-center">
          <h3 class="text-xl md:text-2xl font-bold text-text-primary mb-4 md:mb-6 px-2">Round Results</h3>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6 mb-4 md:mb-6 px-2">
            <div class="card p-3 md:p-6">
              <h4 class="font-bold text-gold mb-2 text-sm md:text-base">📊 Calculations</h4>
              <p class="text-text-secondary text-xs md:text-sm">Average: {{ gameState.average.toFixed(2) }}</p>
              <p class="text-text-secondary text-xs md:text-sm">Target (×0.8): {{ gameState.target.toFixed(2) }}</p>
              <p class="text-text-secondary text-xs md:text-sm">Your choice: {{ gameState.playerChoice }}</p>
              <p class="text-text-secondary text-xs md:text-sm">Distance: {{ gameState.distance.toFixed(2) }}</p>
            </div>

            <div class="card p-3 md:p-6">
              <h4 class="font-bold text-gold mb-2 text-sm md:text-base">🏆 Round Winner</h4>
              <div class="text-3xl md:text-4xl mb-2">{{ gameState.isWinner ? '🎉' : '💔' }}</div>
              <p class="text-base md:text-lg font-bold" :class="gameState.isWinner ? 'text-green-600' : 'text-alert-red'">
                {{ gameState.isWinner ? 'You Won!' : 'You Lost' }}
              </p>
              <p class="text-text-secondary text-xs md:text-sm">{{ gameState.roundResult }}</p>
            </div>
          </div>

          <button @click="nextRound" class="btn-primary min-h-[48px] touch-manipulation text-base md:text-lg py-3 md:py-2 px-6 md:px-4">
            Next Round
          </button>
        </div>

        <!-- Game Over -->
        <div v-else-if="gameState.gamePhase === 'gameOver'" class="text-center px-4">
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
// Game state management
const gameState = reactive({
  currentRound: 1,
  score: 0,
  eliminatedPlayers: 0,
  gamePhase: 'input', // input, waiting, results, gameOver
  playerChoice: null,
  average: 0,
  target: 0,
  distance: 0,
  isWinner: false,
  roundResult: '',
  gameOverMessage: ''
})

// Player profiles
const players = reactive([
  { name: 'You', avatar: '👤', score: 0 },
  { name: 'Player 2', avatar: '🤖', score: 0 },
  { name: 'Player 3', avatar: '🎭', score: 0 },
  { name: 'Player 4', avatar: '🎪', score: 0 },
  { name: 'Player 5', avatar: '🎯', score: 0 }
])

const selectedNumber = ref(null)

// Select a number from the grid
const selectNumber = (number) => {
  selectedNumber.value = number
}

// Simulate other players' choices
const generateOtherPlayersChoices = () => {
  const choices = []
  for (let i = 0; i < 4; i++) {
    choices.push(Math.floor(Math.random() * 101)) // 0-100
  }
  return choices
}

// Calculate winner and apply rules
const calculateRound = (playerChoice, otherChoices) => {
  const allChoices = [playerChoice, ...otherChoices]
  const average = allChoices.reduce((sum, num) => sum + num, 0) / allChoices.length
  const target = average * 0.8

  // Find closest numbers
  const distances = allChoices.map(num => Math.abs(num - target))
  const minDistance = Math.min(...distances)

  // Check for duplicates (Rule 1)
  const duplicates = allChoices.filter((num, index) => allChoices.indexOf(num) !== index)
  const hasDuplicates = duplicates.length > 0

  // Determine winner
  let winnerIndex = -1
  let winnerDistance = Infinity

  allChoices.forEach((choice, index) => {
    const distance = Math.abs(choice - target)

    // Rule 1: Duplicates are invalid
    if (hasDuplicates && duplicates.includes(choice)) {
      return // Skip this choice
    }

    if (distance < winnerDistance) {
      winnerDistance = distance
      winnerIndex = index
    }
  })

  const isPlayerWinner = winnerIndex === 0
  const winnerChoice = allChoices[winnerIndex]

  // Apply score changes
  let scoreChange = 0
  if (!isPlayerWinner) {
    scoreChange = -1 // Loser loses 1 point

    // Rule 2: Exact match causes others to lose 2 points each
    if (gameState.eliminatedPlayers >= 2 && Math.abs(winnerChoice - target) < 0.01) {
      scoreChange = -2 // Exact match causes double penalty
    }
  }

  // Rule 3: Special 0 vs 100 rule
  if (gameState.eliminatedPlayers >= 3 && playerChoice === 0 && winnerChoice === 100) {
    scoreChange = 0 // Player doesn't lose if they chose 0 and someone chose 100
  }

  return {
    average,
    target,
    winnerIndex,
    isPlayerWinner,
    scoreChange,
    distances: distances[0], // Player's distance
    hasDuplicates,
    winnerChoice
  }
}

const submitNumber = () => {
  if (selectedNumber.value === null) return

  const choice = selectedNumber.value
  gameState.playerChoice = choice
  gameState.gamePhase = 'waiting'

  // Simulate waiting for other players
  setTimeout(() => {
    const otherChoices = generateOtherPlayersChoices()
    const result = calculateRound(choice, otherChoices)

    gameState.average = result.average
    gameState.target = result.target
    gameState.distance = result.distances
    gameState.isWinner = result.isPlayerWinner

    // Update scores
    gameState.score += result.scoreChange
    players[0].score = gameState.score // Update player's profile score

    // Update AI player scores (they lose -1 each when player loses)
    if (!result.isPlayerWinner) {
      for (let i = 1; i < players.length; i++) {
        players[i].score -= 1

        // Apply exact match rule to AI players too
        if (gameState.eliminatedPlayers >= 2 && Math.abs(winnerChoice - result.target) < 0.01) {
          players[i].score -= 1 // Additional -1 for exact match
        }
      }
    }

    // Check for death condition
    if (gameState.score <= -10) {
      gameState.gamePhase = 'gameOver'
      gameState.gameOverMessage = `Final Score: ${gameState.score} | Acid death awaits...`
      return
    }

    // Determine round result message
    if (result.isPlayerWinner) {
      gameState.roundResult = `You won! Score unchanged`
    } else {
      gameState.roundResult = `You lost! (${result.scoreChange} points)`
      if (result.hasDuplicates && gameState.eliminatedPlayers >= 1) {
        gameState.roundResult += ' (Duplicate numbers invalidated)'
      }
    }

    gameState.gamePhase = 'results'
  }, 2000)

  selectedNumber.value = null
}

const nextRound = () => {
  gameState.currentRound++
  gameState.gamePhase = 'input'
  gameState.playerChoice = null
}

const resetGame = () => {
  gameState.currentRound = 1
  gameState.score = 0
  gameState.eliminatedPlayers = 0
  gameState.gamePhase = 'input'
  gameState.playerChoice = null
  selectedNumber.value = null

  // Reset all player scores
  players.forEach(player => {
    player.score = 0
  })
}
</script>

<template>
  <div class="min-h-screen py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Game Header -->
      <div class="text-center mb-8">
        <h1 class="text-4xl md:text-6xl font-bold text-gold mb-4 drop-shadow-lg">
          ♔ King of Numbers ♔
        </h1>
        <p class="text-lg text-text-secondary dark:text-soft-white max-w-2xl mx-auto">
          Choose your number wisely. Survival depends on your mathematical strategy!
        </p>
      </div>

      <!-- Game Status -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <!-- Round -->
        <div class="card text-center">
          <h3 class="text-sm font-medium text-text-secondary mb-2">Round</h3>
          <div class="text-3xl font-bold text-gold">{{ gameState.currentRound }}</div>
        </div>

        <!-- Score -->
        <div class="card text-center">
          <h3 class="text-sm font-medium text-text-secondary mb-2">Score</h3>
          <div class="score-display">{{ gameState.score }}</div>
        </div>

        <!-- Lives/Points -->
        <div class="card text-center">
          <h3 class="text-sm font-medium text-text-secondary mb-2">Points</h3>
          <div class="text-3xl font-bold" :class="gameState.points < 0 ? 'text-alert-red' : 'text-accent-cyan'">
            {{ gameState.points }}
          </div>
        </div>
      </div>

      <!-- Game Area -->
      <div class="card mb-8">
        <!-- Current Rules Display -->
        <div v-if="gameState.eliminatedPlayers > 0" class="mb-6 p-4 bg-alert-red/10 border border-alert-red/20 rounded-lg">
          <h4 class="font-bold text-alert-red mb-2">⚠️ ACTIVE RULE CHANGES</h4>
          <div v-if="gameState.eliminatedPlayers >= 1" class="text-sm text-alert-red mb-1">
            • Duplicate numbers are invalid
          </div>
          <div v-if="gameState.eliminatedPlayers >= 2" class="text-sm text-alert-red mb-1">
            • Exact matches cause others to lose 2 points
          </div>
          <div v-if="gameState.eliminatedPlayers >= 3" class="text-sm text-alert-red mb-1">
            • 0 vs 100 is a winning combination
          </div>
        </div>

        <!-- Input Phase -->
        <div v-if="gameState.gamePhase === 'input'" class="text-center">
          <h3 class="text-2xl font-bold text-text-primary mb-6">Choose Your Number (0-100)</h3>
          <form @submit.prevent="submitNumber" class="max-w-md mx-auto space-y-4">
            <input
              v-model="playerNumber"
              type="number"
              placeholder="Enter 0-100"
              class="input-field w-full text-center text-2xl"
              min="0"
              max="100"
              required
            />
            <button
              type="submit"
              class="btn-primary w-full"
              :disabled="!playerNumber && playerNumber !== 0"
            >
              Submit Number
            </button>
          </form>
        </div>

        <!-- Waiting Phase -->
        <div v-else-if="gameState.gamePhase === 'waiting'" class="text-center">
          <div class="text-6xl mb-4">⏳</div>
          <h3 class="text-2xl font-bold text-text-primary mb-4">Waiting for Other Players</h3>
          <p class="text-text-secondary">All 5 players must submit their numbers...</p>
        </div>

        <!-- Results Phase -->
        <div v-else-if="gameState.gamePhase === 'results'" class="text-center">
          <h3 class="text-2xl font-bold text-text-primary mb-6">Round Results</h3>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div class="card">
              <h4 class="font-bold text-gold mb-2">📊 Calculations</h4>
              <p class="text-text-secondary">Average: {{ gameState.average.toFixed(2) }}</p>
              <p class="text-text-secondary">Target (×0.8): {{ gameState.target.toFixed(2) }}</p>
              <p class="text-text-secondary">Your choice: {{ gameState.playerChoice }}</p>
              <p class="text-text-secondary">Distance: {{ gameState.distance.toFixed(2) }}</p>
            </div>

            <div class="card">
              <h4 class="font-bold text-gold mb-2">🏆 Round Winner</h4>
              <div class="text-4xl mb-2">{{ gameState.isWinner ? '🎉' : '💔' }}</div>
              <p class="text-lg font-bold" :class="gameState.isWinner ? 'text-green-600' : 'text-alert-red'">
                {{ gameState.isWinner ? 'You Won!' : 'You Lost' }}
              </p>
              <p class="text-text-secondary">{{ gameState.roundResult }}</p>
            </div>
          </div>

          <button
            @click="nextRound"
            class="btn-primary"
          >
            Next Round
          </button>
        </div>

        <!-- Game Over -->
        <div v-else-if="gameState.gamePhase === 'gameOver'" class="text-center">
          <div class="text-6xl mb-4">{{ gameState.points >= 0 ? '🏆' : '💀' }}</div>
          <h3 class="text-3xl font-bold mb-4" :class="gameState.points >= 0 ? 'text-green-600' : 'text-alert-red'">
            {{ gameState.points >= 0 ? 'GAME CLEAR!' : 'GAME OVER' }}
          </h3>
          <p class="text-xl text-text-secondary mb-6">{{ gameState.gameOverMessage }}</p>

          <div class="flex gap-4 justify-center">
            <button @click="resetGame" class="btn-primary">
              Play Again
            </button>
            <NuxtLink to="/" class="btn-secondary">
              Back to Home
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Navigation -->
      <div class="text-center">
        <NuxtLink to="/" class="text-text-secondary hover:text-gold transition-colors">
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
  points: 0,
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

const playerNumber = ref('')

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

  // Apply point changes
  let pointsChange = 0
  if (isPlayerWinner) {
    pointsChange = 4 // Winner gains 4 points
  } else {
    pointsChange = -1 // Loser loses 1 point

    // Rule 2: Exact match causes others to lose 2 points
    if (gameState.eliminatedPlayers >= 2 && Math.abs(winnerChoice - target) < 0.01) {
      pointsChange = -2
    }
  }

  // Rule 3: Special 0 vs 100 rule
  if (gameState.eliminatedPlayers >= 3 && playerChoice === 0 && winnerChoice === 100) {
    pointsChange = 4 // Player wins if they chose 0 and someone chose 100
  }

  return {
    average,
    target,
    winnerIndex,
    isPlayerWinner,
    pointsChange,
    distances: distances[0], // Player's distance
    hasDuplicates,
    winnerChoice
  }
}

const submitNumber = () => {
  if (playerNumber.value === '' || playerNumber.value === null) return

  const choice = parseInt(playerNumber.value)
  if (choice < 0 || choice > 100) return

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

    // Update points
    gameState.points += result.pointsChange

    // Check for eliminations
    if (gameState.points <= -10) {
      gameState.gamePhase = 'gameOver'
      gameState.gameOverMessage = `Final Score: ${gameState.score} | Points: ${gameState.points} | Acid death awaits...`
      return
    }

    // Determine round result message
    if (result.isPlayerWinner) {
      gameState.roundResult = `You won! (+${result.pointsChange} points)`
      gameState.score += 100 // Bonus score for winning
    } else {
      gameState.roundResult = `You lost! (${result.pointsChange} points)`
      if (result.hasDuplicates && gameState.eliminatedPlayers >= 1) {
        gameState.roundResult += ' (Duplicate numbers invalidated)'
      }
    }

    gameState.gamePhase = 'results'
  }, 2000)

  playerNumber.value = ''
}

const nextRound = () => {
  gameState.currentRound++
  gameState.gamePhase = 'input'
  gameState.playerChoice = null
}

const resetGame = () => {
  gameState.currentRound = 1
  gameState.score = 0
  gameState.points = 0
  gameState.eliminatedPlayers = 0
  gameState.gamePhase = 'input'
  gameState.playerChoice = null
  playerNumber.value = ''
}
</script>

import { ref, reactive, computed } from 'vue'

export const useGameLogic = () => {
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
    gameOverMessage: '',
    allChoices: [], // Store all player choices for this round
    scoreChange: 0 // Score change for current round
  })

  // Modal state
  const showResultsModal = ref(false)
  const modalStep = ref(0)
  const revealedChoices = ref([])
  const selectedNumber = ref(null)

  // Player profiles
  const players = reactive([
    { name: 'You', avatar: '👤', score: 0 },
    { name: 'Player 2', avatar: '🤖', score: 0 },
    { name: 'Player 3', avatar: '🎭', score: 0 },
    { name: 'Player 4', avatar: '🎪', score: 0 },
    { name: 'Player 5', avatar: '🎯', score: 0 }
  ])

  // Computed properties for round results
  const allPlayerChoices = computed(() => gameState.allChoices)
  const sumOfChoices = computed(() => gameState.allChoices.reduce((sum, num) => sum + num, 0))
  const winnerChoice = computed(() => {
    if (gameState.allChoices.length === 0) return '?'
    const target = gameState.target
    const distances = gameState.allChoices.map(num => Math.abs(num - target))
    const minDistance = Math.min(...distances)
    const winnerIndex = distances.indexOf(minDistance)
    return gameState.allChoices[winnerIndex]
  })
  const winnerDistance = computed(() => {
    if (gameState.allChoices.length === 0) return 0
    const target = gameState.target
    const distances = gameState.allChoices.map(num => Math.abs(num - target))
    return Math.min(...distances)
  })

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

  // Modal control functions
  const startResultsModal = () => {
    showResultsModal.value = true
    modalStep.value = 0
    revealedChoices.value = []

    // Disable body scroll
    if (process.client) {
      document.body.style.overflow = 'hidden'
    }

    // Step 1: Reveal choices one by one
    setTimeout(() => {
      modalStep.value = 1
      revealChoicesStepByStep()
    }, 500)
  }

  const revealChoicesStepByStep = () => {
    gameState.allChoices.forEach((choice, index) => {
      setTimeout(() => {
        revealedChoices.value = [...revealedChoices.value, choice]
        scrollToBottom() // Auto-scroll as content increases

        // After all choices are revealed, move to next step
        if (index === gameState.allChoices.length - 1) {
          setTimeout(() => {
            modalStep.value = 2 // Show combination
            scrollToBottom()
            setTimeout(() => {
              modalStep.value = 3 // Show average calculation
              scrollToBottom()
              setTimeout(() => {
                modalStep.value = 4 // Show target formula
                scrollToBottom()
                setTimeout(() => {
                  modalStep.value = 5 // Show winner
                  scrollToBottom()
                }, 1500)
              }, 1500)
            }, 1500)
          }, 1000)
        }
      }, index * 800) // Stagger each reveal by 800ms
    })
  }

  const closeResultsModal = () => {
    showResultsModal.value = false
    modalStep.value = 0
    revealedChoices.value = []

    // Re-enable body scroll
    if (process.client) {
      document.body.style.overflow = ''
    }

    gameState.gamePhase = 'input'
  }

  // Auto-scroll function
  const scrollToBottom = () => {
    // This will be handled by the component
  }

  // Game functions
  const selectNumber = (number) => {
    selectedNumber.value = number
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

      // Store all choices for display
      gameState.allChoices = [choice, ...otherChoices]
      gameState.scoreChange = result.scoreChange

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
          if (gameState.eliminatedPlayers >= 2 && Math.abs(result.winnerChoice - result.target) < 0.01) {
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

      // Start the cinematic results modal
      startResultsModal()
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
    gameState.allChoices = []
    gameState.scoreChange = 0
    selectedNumber.value = null

    // Reset all player scores
    players.forEach(player => {
      player.score = 0
    })
  }

  return {
    // State
    gameState,
    showResultsModal,
    modalStep,
    revealedChoices,
    selectedNumber,
    players,

    // Computed
    allPlayerChoices,
    sumOfChoices,
    winnerChoice,
    winnerDistance,

    // Functions
    selectNumber,
    submitNumber,
    nextRound,
    resetGame,
    closeResultsModal
  }
}

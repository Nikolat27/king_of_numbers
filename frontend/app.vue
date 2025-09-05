<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
// Theme management
const colorMode = useColorMode()

// Set initial theme based on system preference or localStorage
onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme) {
    colorMode.preference = savedTheme
  } else {
    // Default to light mode
    colorMode.preference = 'light'
  }
})

// Watch for theme changes and save to localStorage
watch(() => colorMode.preference, (newTheme) => {
  localStorage.setItem('theme', newTheme)
})
</script>

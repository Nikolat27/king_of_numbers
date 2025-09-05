/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue"
  ],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        sans: ['Roboto', 'sans-serif']
      },
      colors: {
        // Primary Colors
        'deep-navy': '#1E2A38',
        'gold': '#F5C518',
        // Accent Colors
        'accent-cyan': '#4FD1C5',
        'soft-white': '#F8FAFC',
        'light-gray': '#E2E8F0',
        'alert-red': '#F56565',
        // Additional semantic colors for dark/light mode
        'bg-primary': {
          light: '#F8FAFC',
          dark: '#1E2A38'
        },
        'bg-secondary': {
          light: '#E2E8F0',
          dark: '#2D3748'
        },
        'text-primary': {
          light: '#1E2A38',
          dark: '#F8FAFC'
        },
        'text-secondary': {
          light: '#4A5568',
          dark: '#E2E8F0'
        }
      },
      animation: {
        'pulse-gold': 'pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        'bounce-subtle': 'bounce 1s infinite'
      }
    },
  },
  plugins: [],
}

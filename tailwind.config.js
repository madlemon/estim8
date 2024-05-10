const defaultTheme = require('tailwindcss/defaultTheme')
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './web/**/*.templ',
  ],
  theme: {
    fontSize: {
      sm: '0.800rem',
      base: '1rem',
      xl: '1.250rem',
      '2xl': '1.563rem',
      '3xl': '1.954rem',
      '4xl': '2.442rem',
      '5xl': '3.053rem',
    },
    fontFamily: {
      sans: ['ui-sans-serif', 'sans-serif'],
      serif: ['ui-serif', 'serif'],
      mono: ['Lucida Console', 'Courier', 'monospace'],
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}


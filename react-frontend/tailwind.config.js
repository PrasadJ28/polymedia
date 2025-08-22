/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{js,ts,jsx,tsx}'
  ],
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      fontFamily: {
        lobster: ['Lobster', 'cursive'], // 👈 custom font
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}

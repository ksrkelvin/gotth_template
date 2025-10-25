/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/pages/**/*.templ",
    "./views/components/**/*.templ", 
    "./static/**/*.html",
    "./cmd/**/*.go",
    "./internal/**/*.go"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};

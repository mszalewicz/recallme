/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: {"400":"#3984ff", "700":"#0057e5"},
        background: {"200":"#e3e3e3"}
      }
    },
  },
  plugins: [],
};

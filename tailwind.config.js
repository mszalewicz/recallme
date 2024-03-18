/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: {"400":"#3929ff", "700":"#23199e"}
      }
    },
  },
  plugins: [],
};

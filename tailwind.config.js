/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: {"400":"#839bc7", "700":"#687b9e"}
      }
    },
  },
  plugins: [],
};

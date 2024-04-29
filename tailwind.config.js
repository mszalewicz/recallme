/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: { 400: "#000000", 700: "#454747" },
        background: { 200: "#e3e3e3" },
      },
    },
  },
  plugins: [],
};

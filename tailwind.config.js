/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: {"400":"#eb4034", "700":"#942821"}
      }
    },
  },
  plugins: [],
};

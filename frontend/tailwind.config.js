/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        primary: ["Inter", "sans-serif"],
      },
      backgroundImage: {
        login: "url('/src/assets/bg_pattern.png')",
      },
      colors: {
        gray : {
          100: "#EEEEEE",
          200: "#ACACAC",
          300: "#9E9E9E",
          400: "#3D3D3D",
          500: "#292929",
          600: "#171717",
        },
        orange: "#F97F50",
        purple: "#3E47D4",
        red: "#D86045",
        green: "#559137",
      }
    },
  },
  plugins: [],
}


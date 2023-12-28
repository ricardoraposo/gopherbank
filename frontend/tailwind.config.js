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
        bg: "#171717",
        gray: "#292929",
        "mid-gray": "#3D3D3D",
        grayish: "#ACACAC",
        orange: "#F97F50",
        "purple": "#3E47D4",
      }
    },
  },
  plugins: [],
}


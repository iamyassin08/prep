// tailwind.config.js
module.exports = {
  darkMode: 'class',
  content: [
    "./index.html",
    './src/**/*.{vue,js,ts,jsx,tsx}', 
    'node_modules/preline/dist/*.js',
    "./node_modules/preline/preline.js" 
  ],
  theme: {
    extend: {
     
    },
  },
  plugins: [
    require('preline/plugin'), 
  ],
}

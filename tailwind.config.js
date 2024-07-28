/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.{html,js,tmpl}"],
  theme: {
    extend: {
      colors: {
        'tTeal': '#1A3636',
        'tDarkSage': '#40534C',
        'tSage': '#677D6A',
        'tBeige': '#D6BD98',
        'tCharcoal': '#0C0C0C', 
        'tBurntUmber': '#481E14', 
        'tMaroon': '#9B3922',
        'tOrange': '#F2613F'
      }
    },
  },
  plugins: [],
}

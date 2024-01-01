/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/pages/*.tmpl', // Adjust the extension to match your template files
    './templates/**/*.tmpl' // If you have nested directories with template files
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

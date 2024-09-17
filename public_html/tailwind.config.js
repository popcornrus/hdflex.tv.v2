/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  safelist: [
    {
      pattern: /grid-cols-([0-9])/,
    }
  ],
  theme: {
    extend: {
      maxWidth: {
        'container': '90%',
      },
    },
  },
  plugins: [],
}


/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    darkMode: 'class',
    theme: {
        extend: {
            colors: {
                ide: {
                    bg: '#1e1e1e',
                    panel: '#252526',
                    border: '#333333',
                    brand: '#007acc'
                }
            }
        },
    },
    plugins: [],
}

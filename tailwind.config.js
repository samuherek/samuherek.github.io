const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/pages/**/*.{js,jsx,ts,tsx}",
        "./src/components/**/*.{js,jsx,ts,tsx}",
    ],
    theme: {
        extend: {
            fontFamily: {
                "mono": ['"Space Mono"', 'monospace', ...defaultTheme.fontFamily.sans],
            },
            colors: {
                'sam': {
                    '50': '#effafc',
                    '100': '#d6f0f7',
                    '200': '#b3e2ee',
                    '300': '#7ecce2',
                    '400': '#42adce',
                    '500': '#299bc2',
                    '600': '#227498',
                    '700': '#225e7c',
                    '800': '#244f66',
                    '900': '#224357',
                    '950': '#112a3b',
                },
                'uh': {
                    '50': '#f6f6f6',
                    '100': '#e7e7e7',
                    '200': '#d1d1d1',
                    '300': '#b0b0b0',
                    '400': '#888888',
                    '500': '#6d6d6d',
                    '600': '#5d5d5d',
                    '700': '#4f4f4f',
                    '800': '#454545',
                    '900': '#3d3d3d',
                    '950': '#0a0a0a',
                },
            }
        },
    },
    plugins: [],
}


import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			backgroundColor: {
				primary: "#111111",
			},
			colors: {
				secondary: "#AAAAAA",
			},
		},
	},

	plugins: [typography]
} satisfies Config;

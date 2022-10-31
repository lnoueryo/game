import type { PlaywrightTestConfig } from '@playwright/test';

const host = process.env.BFF_HOST || 'localhost';
const port = process.env.BFF_PORT || '3000';

const config: PlaywrightTestConfig = {
	testMatch: '/tests/**/*.ts',
    use: {
        baseURL: 'http://' + host + ':' + port,
        // baseURL: process.env.BASE_URL || 'http://localhost:3000',
        // viewport: { width: 1280, height: 720 },
        browserName: 'chromium',
        headless: true,
        ignoreHTTPSErrors: true,
        actionTimeout: 10_000
    },
	// webServer: {
	// 	command: 'npm run build && npm run preview',aa
	// 	port: 4173
	// }
};

console.log(config.use?.baseURL)

export default config;

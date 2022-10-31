import { expect, test } from '@playwright/test';

test('about page has expected h1', async ({ page }) => {
	await page.goto('/login');
	await page.locator('input[name="username"]').fill('test');
	await page.locator('input[name="password"]').fill('123');
	// await page.locator('text=Login').click();
	await page.locator('button:has-text("Login")').click();
	await expect(page).toHaveURL('http://localhost:3000');
	// expect(await page.textContent('h1')).toBe('About this app');
});
// npx playwright install
// npm install playwright-watch --save-dev
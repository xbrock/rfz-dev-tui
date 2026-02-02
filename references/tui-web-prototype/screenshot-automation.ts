import { chromium } from 'playwright';
import type { Browser, Page } from 'playwright';

import path from 'path';
import fs from 'fs';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const SCREENSHOT_DIR = path.join(__dirname, '..', 'prototype-screenshots');
const BASE_URL = 'http://localhost:3000';

async function waitForServer(url: string, timeoutMs: number = 120000, intervalMs: number = 500): Promise<void> {
  const start = Date.now();
  const deadline = start + timeoutMs;
  while (Date.now() < deadline) {
    try {
      // Use HEAD first to be lightweight; fall back to GET if HEAD not allowed
      const res = await fetch(url, { method: 'HEAD' });
      if (res && (res.ok || (res.status >= 200 && res.status < 500))) {
        return;
      }
    } catch (err) {
      // ignore and retry
    }
    await delay(intervalMs);
  }
  throw new Error(`Timed out waiting for server at ${url} after ${timeoutMs} ms`);
}

// Ensure screenshot directory exists
if (!fs.existsSync(SCREENSHOT_DIR)) {
  fs.mkdirSync(SCREENSHOT_DIR, { recursive: true });
}

async function delay(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function screenshot(page: Page, name: string): Promise<void> {
  const filename = `${name.replace(/[^a-zA-Z0-9-_]/g, '_')}.png`;
  await page.screenshot({
    path: path.join(SCREENSHOT_DIR, filename),
    fullPage: true
  });
  console.log(`  Captured: ${filename}`);
}

async function pressKey(page: Page, key: string, times: number = 1): Promise<void> {
  for (let i = 0; i < times; i++) {
    await page.keyboard.press(key);
    await delay(100);
  }
}

async function captureAllStates(browser: Browser): Promise<void> {
  const page = await browser.newPage();
  await page.setViewportSize({ width: 1440, height: 900 });

  console.log('Starting screenshot automation...\n');

  // Navigate to app
  await page.goto(BASE_URL);
  await delay(500);

  // ==================== WELCOME SCREEN ====================
  console.log('=== Welcome Screen ===');
  await screenshot(page, '01-welcome-default');

  // ==================== NAVIGATION STATES ====================
  console.log('\n=== Navigation States ===');

  // Navigate through nav items
  await pressKey(page, 'j'); // Move to Build
  await delay(200);
  await screenshot(page, '02-nav-build-focused');

  await pressKey(page, 'j'); // Move to Logs
  await delay(200);
  await screenshot(page, '03-nav-logs-focused');

  await pressKey(page, 'j'); // Move to Discover
  await delay(200);
  await screenshot(page, '04-nav-discover-focused');

  await pressKey(page, 'j'); // Move to Config
  await delay(200);
  await screenshot(page, '05-nav-config-focused');

  await pressKey(page, 'j'); // Move to Exit
  await delay(200);
  await screenshot(page, '06-nav-exit-focused');

  // ==================== BUILD SCREEN ====================
  console.log('\n=== Build Screen ===');

  // Press 1 for Build screen
  await pressKey(page, '1');
  await delay(300);
  await screenshot(page, '10-build-empty-selection');

  // Navigate through component list
  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '11-build-list-navigation');

  // Select some components with space
  await pressKey(page, 'k', 3); // Back to top
  await delay(100);
  await pressKey(page, ' '); // Select first
  await delay(100);
  await pressKey(page, 'j');
  await pressKey(page, ' '); // Select second
  await delay(100);
  await pressKey(page, 'j');
  await pressKey(page, ' '); // Select third
  await delay(200);
  await screenshot(page, '12-build-partial-selection');

  // Select all
  await pressKey(page, 'a');
  await delay(200);
  await screenshot(page, '13-build-all-selected');

  // Clear all
  await pressKey(page, 'n');
  await delay(200);
  await screenshot(page, '14-build-cleared');

  // Select a few for build demo
  await pressKey(page, ' ');
  await pressKey(page, 'j');
  await pressKey(page, ' ');
  await pressKey(page, 'j');
  await pressKey(page, ' ');
  await delay(200);

  // Switch to actions area
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '15-build-actions-focused');

  // Navigate action buttons
  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '16-build-action-select-all');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '17-build-action-clear');

  // Back to build button
  await pressKey(page, 'h', 2);
  await delay(200);

  // ==================== BUILD CONFIG MODAL ====================
  console.log('\n=== Build Config Modal ===');

  // Open config modal
  await pressKey(page, 'Enter');
  await delay(300);
  await screenshot(page, '20-config-modal-goals-focus');

  // Navigate goals
  await pressKey(page, 'h');
  await delay(200);
  await screenshot(page, '21-config-modal-goal-clean');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '22-config-modal-goal-install');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '23-config-modal-goal-package');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '24-config-modal-goal-clean-install');

  // Tab to profiles
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '25-config-modal-profiles-focus');

  // Toggle profile
  await pressKey(page, ' ');
  await delay(200);
  await screenshot(page, '26-config-modal-profile-toggled');

  // Navigate to second profile
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '27-config-modal-profile-second');

  // Toggle second profile
  await pressKey(page, ' ');
  await delay(200);
  await screenshot(page, '28-config-modal-both-profiles');

  // Tab to port selection
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '29-config-modal-port-focus');

  // Switch port
  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '30-config-modal-port-11091');

  // Tab to skip tests
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '31-config-modal-skip-tests-focus');

  // Toggle skip tests
  await pressKey(page, ' ');
  await delay(200);
  await screenshot(page, '32-config-modal-skip-tests-off');

  // Toggle back
  await pressKey(page, ' ');
  await delay(200);

  // Tab to actions
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '33-config-modal-actions-cancel');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '34-config-modal-actions-start');

  // ==================== BUILD EXECUTION ====================
  console.log('\n=== Build Execution ===');

  // Start build
  await pressKey(page, 'Enter');
  await delay(500);
  await screenshot(page, '40-build-execution-starting');

  // Wait for some progress
  await delay(1500);
  await screenshot(page, '41-build-execution-running');

  // Navigate through build list
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '42-build-execution-second-focused');

  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '43-build-execution-third-focused');

  // Wait for more progress
  await delay(2000);
  await screenshot(page, '44-build-execution-progress');

  // Switch to actions
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '45-build-execution-actions-focus');

  // View logs (press l)
  await pressKey(page, 'l');
  await delay(300);
  await screenshot(page, '46-logs-modal-default');

  // Scroll logs
  await pressKey(page, 'j', 5);
  await delay(200);
  await screenshot(page, '47-logs-modal-scrolled');

  // Toggle error filter
  await pressKey(page, 'e');
  await delay(200);
  await screenshot(page, '48-logs-modal-error-filter');

  // Toggle back
  await pressKey(page, 'e');
  await delay(200);

  // Close logs modal
  await pressKey(page, 'Escape');
  await delay(300);

  // Wait for build to complete
  console.log('  Waiting for build completion...');
  await delay(8000);
  await screenshot(page, '49-build-execution-complete');

  // Back to build screen
  await pressKey(page, 'Escape');
  await delay(300);

  // ==================== VIEW LOGS SCREEN ====================
  console.log('\n=== View Logs Screen ===');

  // Navigate to Logs screen
  await pressKey(page, 'Escape'); // Back to nav
  await delay(200);
  await pressKey(page, '2');
  await delay(300);
  await screenshot(page, '50-logs-screen-default');

  // Navigate components
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '51-logs-screen-component-2');

  await pressKey(page, 'j', 2);
  await delay(200);
  await screenshot(page, '52-logs-screen-component-4');

  // Tab to logs panel
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '53-logs-screen-logs-focus');

  // Scroll logs
  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '54-logs-screen-scrolled');

  // Toggle follow mode
  await pressKey(page, ' ');
  await delay(200);
  await screenshot(page, '55-logs-screen-follow-off');

  // Toggle back
  await pressKey(page, ' ');
  await delay(200);
  await screenshot(page, '56-logs-screen-follow-on');

  // Tab to filters
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '57-logs-screen-filters-focus');

  // Change filter levels
  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '58-logs-screen-filter-info');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '59-logs-screen-filter-warn');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '60-logs-screen-filter-error');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '61-logs-screen-filter-debug');

  // ==================== DISCOVER SCREEN ====================
  console.log('\n=== Discover Screen ===');

  // Navigate to Discover
  await pressKey(page, 'Escape');
  await delay(200);
  await pressKey(page, '3');
  await delay(300);
  await screenshot(page, '70-discover-default');

  // Navigate list
  await pressKey(page, 'j', 2);
  await delay(200);
  await screenshot(page, '71-discover-row-3');

  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '72-discover-row-6');

  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '73-discover-row-9');

  // Switch to actions
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '74-discover-actions-focus');

  // Navigate action buttons
  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '75-discover-action-checkout');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '76-discover-action-update');

  await pressKey(page, 'l');
  await delay(200);
  await screenshot(page, '77-discover-action-status');

  // ==================== CONFIGURATION SCREEN ====================
  console.log('\n=== Configuration Screen ===');

  // Navigate to Configuration
  await pressKey(page, 'Escape');
  await delay(200);
  await pressKey(page, '4');
  await delay(300);
  await screenshot(page, '80-config-screen-default');

  // Navigate sections
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '81-config-section-registry');

  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '82-config-section-detected');

  // Back to first section
  await pressKey(page, 'k', 2);
  await delay(200);

  // Enter details (scan config)
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '83-config-scan-paths-focus');

  // Navigate paths
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '84-config-scan-path-2');

  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '85-config-scan-path-3');

  // Tab to behavior
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '86-config-scan-behavior');

  // Navigate behavior items
  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '87-config-behavior-depth');

  await pressKey(page, 'j');
  await delay(200);
  await screenshot(page, '88-config-behavior-exclude');

  // Go back to sections
  await pressKey(page, 'Escape');
  await delay(200);

  // Navigate to Component Registry
  await pressKey(page, 'j');
  await delay(200);
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '89-config-registry-focus');

  // Navigate registry items
  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '90-config-registry-row-4');

  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '91-config-registry-row-7');

  // Go back to sections
  await pressKey(page, 'Escape');
  await delay(200);

  // Navigate to Detected Components
  await pressKey(page, 'j');
  await delay(200);
  await pressKey(page, 'Tab');
  await delay(200);
  await screenshot(page, '92-config-detected-focus');

  // Navigate detected items
  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '93-config-detected-row-4');

  await pressKey(page, 'j', 3);
  await delay(200);
  await screenshot(page, '94-config-detected-row-7');

  // ==================== DESIGN SYSTEM PAGE ====================
  console.log('\n=== Design System Page ===');

  await page.goto(`${BASE_URL}/design-system`);
  await delay(500);
  await screenshot(page, '95-design-system-top');

  // Scroll down
  await page.evaluate(() => window.scrollBy(0, 800));
  await delay(300);
  await screenshot(page, '96-design-system-middle');

  await page.evaluate(() => window.scrollBy(0, 800));
  await delay(300);
  await screenshot(page, '97-design-system-bottom');

  console.log('\n=== Screenshot automation complete! ===');
  console.log(`Screenshots saved to: ${SCREENSHOT_DIR}`);

  await page.close();
}

async function main(): Promise<void> {
  const waitTimeout = Number(process.env.WAIT_TIMEOUT_MS) || 120000;
  console.log(`Waiting for server at ${BASE_URL} (timeout ${waitTimeout} ms)...`);
  await waitForServer(BASE_URL, waitTimeout);

  console.log('Launching browser...');
  const browser = await chromium.launch({
    headless: true,
    args: ['--no-sandbox']
  });

  try {
    await captureAllStates(browser);
  } catch (error) {
    console.error('Error during screenshot capture:', error);
    throw error;
  } finally {
    await browser.close();
  }
}

main().catch(console.error);

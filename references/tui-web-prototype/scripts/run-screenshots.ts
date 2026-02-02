import { spawn } from 'child_process';

const START_CMD = process.env.START_CMD || 'pnpm run dev';
const SCREENSHOT_CMD = process.env.SCREENSHOT_CMD || 'pnpm run screenshot';
const BASE_URL = process.env.BASE_URL || 'http://localhost:3000';
const WAIT_TIMEOUT_MS = Number(process.env.WAIT_TIMEOUT_MS) || 120000;
const POLL_INTERVAL_MS = Number(process.env.POLL_INTERVAL_MS) || 500;

function delay(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function waitForServer(url: string, timeoutMs: number, intervalMs: number) {
  const start = Date.now();
  const deadline = start + timeoutMs;
  while (Date.now() < deadline) {
    try {
      const res = await fetch(url, { method: 'HEAD' });
      if (res && (res.ok || (res.status >= 200 && res.status < 500))) return;
    } catch (err) {
      // ignore
    }
    await delay(intervalMs);
  }
  throw new Error(`Timed out waiting for server at ${url} after ${timeoutMs} ms`);
}

function spawnShell(cmd: string) {
  // Use shell spawn so complex commands like "pnpm run start" work
  const child = spawn(cmd, { shell: true, stdio: 'inherit' });
  return child;
}

async function run(): Promise<number> {
  if (process.env.BUILD_BEFORE_START === '1' || process.env.BUILD_BEFORE_START === 'true') {
    console.log('Running build step before starting server...');
    // run build synchronously and inherit stdio so user sees output
    const builder = spawnShell('pnpm run build');
    const buildExit: number = await new Promise(resolve => {
      builder.on('exit', (code) => resolve(typeof code === 'number' ? code : 1));
      builder.on('error', () => resolve(1));
    });
    if (buildExit !== 0) {
      console.error('Build failed, aborting.');
      return buildExit;
    }
  }

  console.log(`Starting server with: ${START_CMD}`);
  const server = spawnShell(START_CMD);

  // If server exits early, abort
  let serverExited = false;
  server.on('exit', (code, signal) => {
    serverExited = true;
    console.log(`Server process exited with code=${code} signal=${signal}`);
  });

  try {
    console.log(`Waiting for ${BASE_URL} (timeout ${WAIT_TIMEOUT_MS} ms)...`);
    await waitForServer(BASE_URL, WAIT_TIMEOUT_MS, POLL_INTERVAL_MS);
  } catch (err) {
    console.error('Server did not become ready:', err.message || err);
    // ensure server is killed if we spawned it
    if (!serverExited) {
      console.log('Shutting down server process...');
      server.kill('SIGINT');
      await delay(1000);
      if (!server.killed) server.kill('SIGKILL');
    }
    return 2;
  }

  console.log('Server is ready — running screenshots...');

  // Run screenshot command
  const runner = spawn(SCREENSHOT_CMD, { shell: true, stdio: 'inherit' });
  const exitCode: number = await new Promise(resolve => {
    runner.on('exit', (code) => resolve(typeof code === 'number' ? code : 1));
    runner.on('error', () => resolve(1));
  });

  console.log('Screenshot process finished with code', exitCode);

  // Shut down server
  if (!serverExited) {
    console.log('Stopping server...');
    server.kill('SIGINT');
    await delay(1000);
    if (!server.killed) server.kill('SIGKILL');
  }

  return exitCode;
}

run().then(code => process.exit(code)).catch(err => {
  console.error('Runner failed:', err);
  process.exit(1);
});

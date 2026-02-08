/**
 * TYPESCRIPT BROWSER APIS (Lessons 1-10)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Read one lesson block at a time
 * 3) For runtime, compile to JS and run in a browser context
 *
 * Extra context:
 * - lessons/notes/86-ts-browser-apis-principles.md
 * - lessons/notes/87-ts-browser-api-history.md
 * - lessons/notes/82-ts-tooling-runtime.md
 */

// -----------------------------------------------------------------------------
// LESSON 1: localStorage set/get
// Why this matters: apps often persist small preferences.

function saveTheme(theme: "light" | "dark"): void {
  localStorage.setItem("theme", theme);
}

function loadTheme(): "light" | "dark" {
  const raw = localStorage.getItem("theme");
  return raw === "dark" ? "dark" : "light";
}

saveTheme("dark");
console.log("Lesson 1:", loadTheme());

// -----------------------------------------------------------------------------
// LESSON 2: JSON with localStorage
// Why this matters: storage values are strings, not objects.

type Preferences = {
  language: "en" | "es";
  compactMode: boolean;
};

function savePrefs(prefs: Preferences): void {
  localStorage.setItem("prefs", JSON.stringify(prefs));
}

function loadPrefs(): Preferences {
  const raw = localStorage.getItem("prefs");
  if (!raw) return { language: "en", compactMode: false };
  try {
    const parsed = JSON.parse(raw) as Partial<Preferences>;
    return {
      language: parsed.language === "es" ? "es" : "en",
      compactMode: parsed.compactMode === true
    };
  } catch {
    return { language: "en", compactMode: false };
  }
}

savePrefs({ language: "es", compactMode: true });
console.log("Lesson 2:", loadPrefs());

// -----------------------------------------------------------------------------
// LESSON 3: URL and URLSearchParams
// Why this matters: query strings drive filtering and deep linking.

function parseFilter(urlText: string): string {
  const url = new URL(urlText);
  return url.searchParams.get("filter") ?? "all";
}

console.log("Lesson 3:", parseFilter("https://example.com/?filter=active"));

// -----------------------------------------------------------------------------
// LESSON 4: Build URLs safely
// Why this matters: avoid broken string concatenation.

function buildSearchUrl(base: string, term: string): string {
  const url = new URL(base);
  url.searchParams.set("q", term);
  return url.toString();
}

console.log("Lesson 4:", buildSearchUrl("https://example.com/search", "typescript"));

// -----------------------------------------------------------------------------
// LESSON 5: Timers with clearTimeout
// Why this matters: debounce-like behavior reduces noisy updates.

let pendingTimer: number | undefined;

function schedulePreview(callback: () => void): void {
  if (pendingTimer !== undefined) {
    window.clearTimeout(pendingTimer);
  }
  pendingTimer = window.setTimeout(() => {
    callback();
  }, 300);
}

schedulePreview(() => console.log("Lesson 5: preview update"));

// -----------------------------------------------------------------------------
// LESSON 6: Clipboard API (with permission-safe fallback)
// Why this matters: copy-to-clipboard is common UX.

async function copyText(text: string): Promise<boolean> {
  try {
    await navigator.clipboard.writeText(text);
    return true;
  } catch {
    return false;
  }
}

copyText("hello").then((ok) => console.log("Lesson 6:", ok));

// -----------------------------------------------------------------------------
// LESSON 7: Fetch with typed response shape
// Why this matters: API contracts should be explicit.

type HealthResponse = {
  ok: boolean;
  service: string;
};

async function fetchHealth(): Promise<HealthResponse | null> {
  try {
    // Offline-safe data URL for predictable lesson behavior.
    const res = await fetch("data:application/json,%7B%22ok%22%3Atrue%2C%22service%22%3A%22demo%22%7D");
    if (!res.ok) return null;
    const data = (await res.json()) as Partial<HealthResponse>;
    if (typeof data.ok !== "boolean") return null;
    if (typeof data.service !== "string") return null;
    return { ok: data.ok, service: data.service };
  } catch {
    return null;
  }
}

fetchHealth().then((result) => console.log("Lesson 7:", result));

// -----------------------------------------------------------------------------
// LESSON 8: AbortController for request cancellation
// Why this matters: user navigation should cancel stale requests.

async function fetchWithAbort(signal: AbortSignal): Promise<boolean> {
  try {
    await fetch("data:text/plain,ok", { signal });
    return true;
  } catch {
    return false;
  }
}

const controller = new AbortController();
fetchWithAbort(controller.signal).then((ok) => console.log("Lesson 8:", ok));
controller.abort();

// -----------------------------------------------------------------------------
// LESSON 9: Browser online/offline events
// Why this matters: apps should react to connectivity changes.

function handleOnline(): void {
  console.log("Lesson 9: online");
}

function handleOffline(): void {
  console.log("Lesson 9: offline");
}

window.addEventListener("online", handleOnline);
window.addEventListener("offline", handleOffline);

// -----------------------------------------------------------------------------
// LESSON 10: Typed custom event payload
// Why this matters: event-based architecture scales UI communication.

type ToastDetail = {
  message: string;
  level: "info" | "error";
};

const toastEvent = new CustomEvent<ToastDetail>("toast", {
  detail: { message: "Saved", level: "info" }
});

window.addEventListener("toast", (event: Event) => {
  const typed = event as CustomEvent<ToastDetail>;
  console.log("Lesson 10:", typed.detail.message, typed.detail.level);
});

window.dispatchEvent(toastEvent);

// End of TypeScript Browser APIs 1-10

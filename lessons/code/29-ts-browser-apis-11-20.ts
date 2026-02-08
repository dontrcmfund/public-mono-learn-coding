/**
 * TYPESCRIPT BROWSER APIS (Lessons 11-20)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Read one lesson block at a time
 * 3) Change one rule, predict output, then run
 *
 * Extra context:
 * - lessons/notes/86-ts-browser-apis-principles.md
 * - lessons/notes/87-ts-browser-api-history.md
 * - lessons/notes/88-browser-api-permissions-gotchas.md
 */

// -----------------------------------------------------------------------------
// LESSON 11: sessionStorage basics
// Why this matters: store temporary per-tab state.

function saveDraft(text: string): void {
  sessionStorage.setItem("draft", text);
}

function loadDraft(): string {
  return sessionStorage.getItem("draft") ?? "";
}

saveDraft("draft note");
console.log("Lesson 11:", loadDraft());

// -----------------------------------------------------------------------------
// LESSON 12: history pushState
// Why this matters: update URL without full reload.

function setViewParam(view: string): void {
  const url = new URL(window.location.href);
  url.searchParams.set("view", view);
  history.pushState({ view }, "", url.toString());
}

setViewParam("list");
console.log("Lesson 12:", window.location.href);

// -----------------------------------------------------------------------------
// LESSON 13: history popstate event
// Why this matters: react when user navigates browser history.

window.addEventListener("popstate", (event: PopStateEvent) => {
  console.log("Lesson 13:", event.state);
});

// -----------------------------------------------------------------------------
// LESSON 14: location helpers
// Why this matters: route and query behavior starts from URL parts.

function currentPath(): string {
  return window.location.pathname;
}

console.log("Lesson 14:", currentPath());

// -----------------------------------------------------------------------------
// LESSON 15: FormData parsing
// Why this matters: forms are a core web input system.

function formDataToObject(form: HTMLFormElement): Record<string, string> {
  const data = new FormData(form);
  const out: Record<string, string> = {};
  for (const [key, value] of data.entries()) {
    out[key] = String(value);
  }
  return out;
}

const form = document.createElement("form");
const input = document.createElement("input");
input.name = "email";
input.value = "dev@example.com";
form.appendChild(input);
console.log("Lesson 15:", formDataToObject(form));

// -----------------------------------------------------------------------------
// LESSON 16: querySelector with typed narrowing
// Why this matters: DOM queries can return null.

const maybeButton = document.querySelector("button");
if (maybeButton instanceof HTMLButtonElement) {
  console.log("Lesson 16: found button with text", maybeButton.textContent);
} else {
  console.log("Lesson 16: no button found");
}

// -----------------------------------------------------------------------------
// LESSON 17: requestAnimationFrame basics
// Why this matters: smooth UI updates should sync with paint cycles.

let frameCount = 0;
function animate(): void {
  frameCount += 1;
  if (frameCount <= 2) {
    console.log("Lesson 17 frame:", frameCount);
    requestAnimationFrame(animate);
  }
}
requestAnimationFrame(animate);

// -----------------------------------------------------------------------------
// LESSON 18: IntersectionObserver pattern (guarded)
// Why this matters: lazy loading and visibility logic use this API.

if ("IntersectionObserver" in window) {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        console.log("Lesson 18: element visible");
      }
    });
  });
  const el = document.createElement("div");
  document.body.appendChild(el);
  observer.observe(el);
  // Best practice: disconnect observers when no longer needed.
  setTimeout(() => observer.disconnect(), 500);
}

// -----------------------------------------------------------------------------
// LESSON 19: navigator language
// Why this matters: localization starts from user locale.

function getPreferredLanguage(): string {
  return navigator.language;
}

console.log("Lesson 19:", getPreferredLanguage());

// -----------------------------------------------------------------------------
// LESSON 20: simple API capability check helper
// Why this matters: feature detection avoids runtime surprises.

function hasClipboardSupport(): boolean {
  return typeof navigator.clipboard !== "undefined";
}

console.log("Lesson 20:", hasClipboardSupport());

// End of TypeScript Browser APIs 11-20

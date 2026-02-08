/**
 * NODE BASICS (Lessons 31–40)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/22-node-basics-31-40.js
 */

// -----------------------------------------------------------------------------
// LESSON 31: Reading user input from stdin
// stdin lets users type into your program.
// Why this matters: it enables interactive scripts.
process.stdin.setEncoding("utf8");
process.stdin.on("data", (chunk) => {
  const text = chunk.trim();
  if (text) {
    console.log(`Lesson 31: You typed: ${text}`);
    process.stdin.pause(); // stop listening after one input
  }
});
console.log("Lesson 31: Type something and press Enter");

// -----------------------------------------------------------------------------
// LESSON 32: Simple prompt loop (concept)
// You can keep listening for more input.
// Why this matters: it allows multi-step CLIs.
// (We pause above to keep this beginner-friendly.)

// -----------------------------------------------------------------------------
// LESSON 33: Fetching data (Node 18+)
// fetch gets data from a URL.
// Why this matters: APIs are everywhere.
// More details: lessons/notes/77-node-fetch-gotchas.md
async function fetchExample() {
  try {
    // Use a data URL so this lesson works without internet access.
    const res = await fetch("data:text/plain,hello");
    console.log("Lesson 33 status:", res.status, "ok:", res.ok);
  } catch (err) {
    console.log("Lesson 33 error:", err.message);
  }
}
fetchExample();

// -----------------------------------------------------------------------------
// LESSON 34: Reading JSON from a response
// res.json() parses JSON automatically.
// Why this matters: most APIs return JSON.
// More details: lessons/notes/77-node-fetch-gotchas.md
async function fetchJson() {
  try {
    // Also offline-safe: parse JSON from an in-memory data URL.
    const res = await fetch("data:application/json,%7B%22ok%22%3Atrue%7D");
    if (!res.ok) {
      console.log("Lesson 34: HTTP error", res.status);
      return;
    }
    const data = await res.json();
    console.log("Lesson 34:", typeof data, data.ok);
  } catch (err) {
    console.log("Lesson 34 error:", err.message);
  }
}
fetchJson();

// -----------------------------------------------------------------------------
// LESSON 35: Basic error-first callback pattern
// Some Node APIs use (err, result) callbacks.
// Why this matters: you will see this in older code.
const fs = require("fs");
fs.readFile(__filename, "utf8", (err, text) => {
  if (err) {
    console.log("Lesson 35 error:", err.message);
    return;
  }
  console.log("Lesson 35 length:", text.length);
});

// -----------------------------------------------------------------------------
// LESSON 36: Streams (concept)
// Streams process data in chunks.
// Why this matters: large files are too big to load at once.
const stream = fs.createReadStream(__filename, { encoding: "utf8" });
let total = 0;
stream.on("data", (chunk) => {
  total += chunk.length;
});
stream.on("end", () => {
  console.log("Lesson 36 bytes:", total);
});

// -----------------------------------------------------------------------------
// LESSON 37: Simple progress logging
// Log progress during long tasks.
// Why this matters: feedback reduces confusion.
let progress = 0;
const progressTimer = setInterval(() => {
  progress += 20;
  console.log(`Lesson 37: progress ${progress}%`);
  if (progress >= 100) {
    clearInterval(progressTimer);
  }
}, 200);

// -----------------------------------------------------------------------------
// LESSON 38: Working with paths safely
// path.dirname and path.extname help analyze file paths.
// Why this matters: it avoids string hacks.
const path = require("path");
console.log("Lesson 38 dir:", path.dirname(__filename));
console.log("Lesson 38 ext:", path.extname(__filename));

// -----------------------------------------------------------------------------
// LESSON 39: Using require with JSON
// You can require JSON files directly.
// Why this matters: config files are often JSON.
const tempJsonPath = path.join(__dirname, "tmp-config.json");
fs.writeFileSync(tempJsonPath, JSON.stringify({ theme: "light" }));
const config = require(tempJsonPath);
console.log("Lesson 39:", config.theme);

// -----------------------------------------------------------------------------
// LESSON 40: Clean up temporary files
// Remove temp files to keep things tidy.
// Why this matters: scripts should clean up after themselves.
fs.unlinkSync(tempJsonPath);
console.log("Lesson 40: cleaned up temp file");

// End of Node Basics 31–40

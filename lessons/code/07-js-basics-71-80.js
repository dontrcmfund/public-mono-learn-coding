/**
 * JS BASICS (Lessons 71–80 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/07-js-basics-71-80.js
 */

// -----------------------------------------------------------------------------
// LESSON 71: Common array methods recap (map/filter/find)
// These help you transform, filter, and locate data.
// Why this matters: most data work is list work.
const nums = [1, 2, 3, 4];
console.log("Lesson 71 map:", nums.map((n) => n * 10));
console.log("Lesson 71 filter:", nums.filter((n) => n % 2 === 0));
console.log("Lesson 71 find:", nums.find((n) => n > 2));

// -----------------------------------------------------------------------------
// LESSON 72: reduce recap (summaries)
// reduce can add, count, or summarize a list.
// Why this matters: it turns many values into one.
const total = nums.reduce((acc, n) => acc + n, 0);
console.log("Lesson 72:", total);

// -----------------------------------------------------------------------------
// LESSON 73: Object entries (looping key/value pairs)
// Object.entries gives you [key, value] pairs.
// Why this matters: it makes objects easy to loop.
const settings = { volume: 3, theme: "dark" };
for (const [key, value] of Object.entries(settings)) {
  console.log("Lesson 73:", key, value);
}

// -----------------------------------------------------------------------------
// LESSON 74: Early return in loops
// You can stop a loop with break.
// Why this matters: it saves time once you find what you need.
for (const n of nums) {
  if (n === 3) {
    console.log("Lesson 74: found 3, stopping");
    break;
  }
}

// -----------------------------------------------------------------------------
// LESSON 75: continue (skip one loop step)
// continue skips to the next item.
// Why this matters: it avoids nested if blocks.
for (const n of nums) {
  if (n % 2 !== 0) {
    continue; // skip odd numbers
  }
  console.log("Lesson 75 even:", n);
}

// -----------------------------------------------------------------------------
// LESSON 76: try/catch with custom errors
// You can throw your own error when something is wrong.
// Why this matters: it makes bugs clearer.
function requireName(name) {
  if (!name) {
    throw new Error("Name is required");
  }
  return name;
}
try {
  requireName("");
} catch (err) {
  console.log("Lesson 76:", err.message);
}

// -----------------------------------------------------------------------------
// LESSON 77: Basic input validation pattern
// Check input early before using it.
// Why this matters: it prevents bad data from spreading.
function isPositiveNumber(value) {
  return typeof value === "number" && value > 0;
}
console.log("Lesson 77:", isPositiveNumber(5), isPositiveNumber(-2));

// -----------------------------------------------------------------------------
// LESSON 78: String trimming and checking
// trim removes extra spaces.
// Why this matters: user input often has extra whitespace.
const raw = "  hello  ";
const clean = raw.trim();
console.log("Lesson 78:", clean, clean.length);

// -----------------------------------------------------------------------------
// LESSON 79: Safe number conversion
// Number() converts, Number.isNaN checks for failure.
// Why this matters: user input is often text.
const input = "42";
const parsed = Number(input);
console.log("Lesson 79:", parsed, Number.isNaN(parsed));

// -----------------------------------------------------------------------------
// LESSON 80: Basic logging with labels
// Clear labels make debugging easier.
// Why this matters: it reduces confusion when reading output.
const result = { ok: true, count: 3 };
console.log("Lesson 80 result:", result);

// End of Lessons 71–80

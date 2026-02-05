/**
 * JS BASICS (Lessons 61–70 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/06-js-basics-61-70.js
 */

// -----------------------------------------------------------------------------
// LESSON 61: Object and array references (shared vs copied)
// Variables can point to the same object.
// Why this matters: changes in one place can affect another.
const original = { count: 1 };
const alias = original; // both point to the same object
alias.count = 2;
console.log("Lesson 61:", original.count); // 2

// -----------------------------------------------------------------------------
// LESSON 62: Shallow vs deep copy (basic idea)
// Spread makes a shallow copy, not a deep one.
// Why this matters: nested objects can still be linked.
// More details: lessons/notes/68-js-shallow-vs-deep.md
const nested = { info: { level: 1 } };
const shallow = { ...nested };
shallow.info.level = 2;
console.log("Lesson 62:", nested.info.level); // 2

// -----------------------------------------------------------------------------
// LESSON 63: Rest parameters (collect many args)
// Rest gathers extra arguments into an array.
// Why this matters: functions can accept flexible input.
function sumAll(...nums) {
  return nums.reduce((acc, n) => acc + n, 0);
}
console.log("Lesson 63:", sumAll(1, 2, 3));

// -----------------------------------------------------------------------------
// LESSON 64: Default values with || and ??
// || uses a fallback for any falsy value.
// ?? uses a fallback only for null/undefined.
// Why this matters: 0 and "" are real values.
const points = 0;
console.log("Lesson 64 (||):", points || 10); // gives 10
console.log("Lesson 64 (??):", points ?? 10); // keeps 0

// -----------------------------------------------------------------------------
// LESSON 65: Optional chaining with functions
// You can safely call a function if it exists.
// Why this matters: missing functions should not crash.
const maybe = {};
maybe.run?.(); // nothing happens, no crash
console.log("Lesson 65: optional call did not crash");

// -----------------------------------------------------------------------------
// LESSON 66: Destructuring with defaults
// You can provide default values while destructuring.
// Why this matters: missing data will not break your code.
const config = { theme: "light" };
const { theme, mode = "simple" } = config;
console.log("Lesson 66:", theme, mode);

// -----------------------------------------------------------------------------
// LESSON 67: Ternary operator (short if/else)
// condition ? valueIfTrue : valueIfFalse
// Why this matters: it keeps simple choices compact.
const age = 18;
const access = age >= 18 ? "allowed" : "denied";
console.log("Lesson 67:", access);

// -----------------------------------------------------------------------------
// LESSON 68: Map (object-like key/value store)
// Map can use any type as a key.
// Why this matters: it avoids key collisions on plain objects.
const visits = new Map();
visits.set("home", 1);
visits.set("about", 2);
console.log("Lesson 68:", visits.get("home"));

// -----------------------------------------------------------------------------
// LESSON 69: Set (unique values only)
// Set stores unique values.
// Why this matters: it removes duplicates easily.
const unique = new Set([1, 2, 2, 3]);
console.log("Lesson 69:", Array.from(unique));

// -----------------------------------------------------------------------------
// LESSON 70: Basic modules (export/import idea)
// Modules let you split code across files.
// Why this matters: large programs stay organized.
// We will practice real imports later.
console.log("Lesson 70: Modules let you split code into files.");

// End of Lessons 61–70

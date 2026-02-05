/**
 * JS BASICS (Lessons 31–40 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/03-js-basics-31-40.js
 */

// -----------------------------------------------------------------------------
// LESSON 31: const vs let (when to change values)
// const means “do not reassign this.” let means “I might change it.”
// Why this matters: it prevents accidental changes.
// More details: lessons/notes/64-js-const-vs-let.md
const siteName = "LearnJS";
let visits = 0;
visits = visits + 1;
console.log("Lesson 31:", siteName, visits);

// -----------------------------------------------------------------------------
// LESSON 32: Array indexes and last item
// Indexes start at 0, and last index is length - 1.
// Why this matters: off‑by‑one errors are common.
const numbers = [10, 20, 30, 40];
const lastIndex = numbers.length - 1;
console.log("Lesson 32 last:", numbers[lastIndex]);

// -----------------------------------------------------------------------------
// LESSON 33: Slicing arrays (getting a portion)
// slice gives you a new array without changing the original.
// Why this matters: it helps you avoid unwanted changes.
const letters = ["a", "b", "c", "d"];
const firstTwo = letters.slice(0, 2);
console.log("Lesson 33:", firstTwo, letters);

// -----------------------------------------------------------------------------
// LESSON 34: Spreading arrays (copying and combining)
// The ... spread operator expands an array.
// Why this matters: it helps you create new arrays safely.
const base = [1, 2];
const combined = [...base, 3, 4];
console.log("Lesson 34:", combined);

// -----------------------------------------------------------------------------
// LESSON 35: Object destructuring (pull values out)
// Destructuring is a shortcut to grab values by name.
// Why this matters: it makes code cleaner when you need a few fields.
// More details: lessons/notes/65-js-object-destructuring.md
const user = { name: "Kai", role: "admin" };
const { name, role } = user;
console.log("Lesson 35:", name, role);

// -----------------------------------------------------------------------------
// LESSON 36: String splitting (turn text into a list)
// split turns a string into an array.
// Why this matters: many inputs are plain text.
const sentence = "red,green,blue";
const parts = sentence.split(",");
console.log("Lesson 36:", parts);

// -----------------------------------------------------------------------------
// LESSON 37: Joining arrays (turn list into text)
// join turns an array into a string.
// Why this matters: you often need to display lists.
const joined = parts.join(" | ");
console.log("Lesson 37:", joined);

// -----------------------------------------------------------------------------
// LESSON 38: Math with Math
// Math has helpful tools like rounding.
// Why this matters: real programs need clean numbers.
const price = 9.99;
console.log("Lesson 38 round:", Math.round(price));
console.log("Lesson 38 floor:", Math.floor(price));
console.log("Lesson 38 ceil:", Math.ceil(price));

// -----------------------------------------------------------------------------
// LESSON 39: Random numbers (simple use)
// Math.random() gives a number between 0 and 1.
// Why this matters: randomness is useful in games and tests.
const randomValue = Math.random();
console.log("Lesson 39:", randomValue);

// -----------------------------------------------------------------------------
// LESSON 40: Early return (stop a function)
// return can end a function early.
// Why this matters: it keeps logic simple and avoids extra work.
function printIfPositive(n) {
  if (n <= 0) {
    return; // stop early
  }
  console.log("Lesson 40:", n);
}
printIfPositive(5);
printIfPositive(-1);

// End of Lessons 31–40

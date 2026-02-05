/**
 * JS BASICS (Lessons 91–100 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/09-js-basics-91-100.js
 */

// -----------------------------------------------------------------------------
// LESSON 91: Converting between strings and numbers
// String() and Number() convert values safely.
// Why this matters: user input is usually text.
const num = 123;
const asText = String(num);
const backToNum = Number(asText);
console.log("Lesson 91:", asText, backToNum);

// -----------------------------------------------------------------------------
// LESSON 92: Simple rounding with toFixed
// toFixed returns a string with set decimals.
// Why this matters: money often needs 2 decimals.
const price = 3.456;
console.log("Lesson 92:", price.toFixed(2));

// -----------------------------------------------------------------------------
// LESSON 93: Template strings with expressions
// You can insert calculations inside ${ }.
// Why this matters: it keeps output readable.
const x = 7;
const y = 8;
console.log(`Lesson 93: ${x} + ${y} = ${x + y}`);

// -----------------------------------------------------------------------------
// LESSON 94: Array join with different separators
// join can build strings for display.
// Why this matters: UI often needs a single string.
const items = ["apples", "oranges", "bananas"];
console.log("Lesson 94:", items.join(", "));

// -----------------------------------------------------------------------------
// LESSON 95: Simple sorting with localeCompare (strings)
// localeCompare handles alphabetical order better.
// Why this matters: it sorts text more naturally.
const words = ["banana", "apple", "cherry"];
words.sort((a, b) => a.localeCompare(b));
console.log("Lesson 95:", words);

// -----------------------------------------------------------------------------
// LESSON 96: Object shorthand
// When key and variable name match, you can shorten.
// Why this matters: it reduces clutter.
const city = "Denver";
const state = "CO";
const place = { city, state };
console.log("Lesson 96:", place);

// -----------------------------------------------------------------------------
// LESSON 97: Destructuring in function parameters
// You can pull values out directly in the parameter list.
// Why this matters: it makes function inputs clear.
function printUser({ name, role }) {
  console.log(`Lesson 97: ${name} (${role})`);
}
printUser({ name: "Mia", role: "student" });

// -----------------------------------------------------------------------------
// LESSON 98: Rest and spread recap
// Rest collects; spread expands.
// Why this matters: it helps manage flexible data.
const nums = [1, 2, 3];
function sum(...values) {
  return values.reduce((acc, v) => acc + v, 0);
}
console.log("Lesson 98:", sum(...nums));

// -----------------------------------------------------------------------------
// LESSON 99: Date formatting basics
// toLocaleString formats dates for humans.
// Why this matters: raw dates are hard to read.
const now = new Date();
console.log("Lesson 99:", now.toLocaleString());

// -----------------------------------------------------------------------------
// LESSON 100: Tiny recap (confidence check)
// You now know: values, variables, functions, arrays, objects, and more.
// Why this matters: these are the building blocks of real programs.
console.log("Lesson 100: You have the core building blocks of JavaScript.");

// End of Lessons 91–100

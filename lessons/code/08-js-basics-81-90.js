/**
 * JS BASICS (Lessons 81–90 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/08-js-basics-81-90.js
 */

// -----------------------------------------------------------------------------
// LESSON 81: Null vs undefined
// undefined = not assigned; null = intentionally empty.
// Why this matters: it helps you understand missing data.
let notSet;
const intentionallyEmpty = null;
console.log("Lesson 81:", notSet, intentionallyEmpty);

// -----------------------------------------------------------------------------
// LESSON 82: Checking for existence
// Use == null to check for null or undefined together.
// Why this matters: it simplifies “missing value” checks.
function hasValue(value) {
  return value != null; // true unless null or undefined
}
console.log("Lesson 82:", hasValue(0), hasValue(null), hasValue(undefined));

// -----------------------------------------------------------------------------
// LESSON 83: Array includes vs indexOf
// includes is clearer and returns true/false.
// Why this matters: clarity prevents mistakes.
const pets = ["cat", "dog"];
console.log("Lesson 83:", pets.includes("dog"), pets.indexOf("dog"));

// -----------------------------------------------------------------------------
// LESSON 84: String includes and startsWith
// These help you search text.
// Why this matters: text search is common.
const phrase = "learn javascript";
console.log("Lesson 84:", phrase.includes("java"), phrase.startsWith("learn"));

// -----------------------------------------------------------------------------
// LESSON 85: Array every vs some
// every checks if all items pass; some checks if any pass.
// Why this matters: it makes list checks easy.
const ages = [16, 22, 19];
console.log("Lesson 85:", ages.every((a) => a >= 18), ages.some((a) => a >= 18));

// -----------------------------------------------------------------------------
// LESSON 86: More on functions (return vs console.log)
// return gives a value back; console.log just prints.
// Why this matters: returning is needed to reuse results.
function multiply(a, b) {
  return a * b;
}
console.log("Lesson 86:", multiply(3, 4));

// -----------------------------------------------------------------------------
// LESSON 87: Guard clauses (early exits)
// Guard clauses keep functions simple.
// Why this matters: fewer nested if blocks.
function safeDivide(a, b) {
  if (b === 0) {
    return "Cannot divide by zero";
  }
  return a / b;
}
console.log("Lesson 87:", safeDivide(10, 2), safeDivide(10, 0));

// -----------------------------------------------------------------------------
// LESSON 88: Object hasOwnProperty (own vs inherited)
// hasOwnProperty checks if a key belongs to the object itself.
// Why this matters: it avoids surprises from prototypes.
const bag = { apple: 1 };
console.log("Lesson 88:", bag.hasOwnProperty("apple"));

// -----------------------------------------------------------------------------
// LESSON 89: Basic string replace
// replace swaps one part of a string for another.
// Why this matters: text cleanup is common.
const originalText = "Hello, world";
console.log("Lesson 89:", originalText.replace("world", "JS"));

// -----------------------------------------------------------------------------
// LESSON 90: Simple formatting with padStart
// padStart adds padding to the start of a string.
// Why this matters: it helps format numbers like "05".
const minute = "5";
console.log("Lesson 90:", minute.padStart(2, "0"));

// End of Lessons 81–90

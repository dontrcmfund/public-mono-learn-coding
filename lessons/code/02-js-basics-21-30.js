/**
 * JS BASICS (Lessons 21–30 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/02-js-basics-21-30.js
 */

// -----------------------------------------------------------------------------
// LESSON 21: else if (more than two choices)
// Use else if when there are multiple possibilities.
// Why this matters: real decisions often have more than two options.
const score = 82;
if (score >= 90) {
  console.log("Lesson 21: Grade A");
} else if (score >= 80) {
  console.log("Lesson 21: Grade B");
} else {
  console.log("Lesson 21: Grade C or below");
}

// -----------------------------------------------------------------------------
// LESSON 22: === vs == (strict vs loose equality)
// === checks value AND type. == can do strange conversions.
// Why this matters: strict checks prevent surprises.
// More details: lessons/notes/61-js-equality-gotchas.md
console.log("Lesson 22:", 5 === "5", 5 == "5");

// -----------------------------------------------------------------------------
// LESSON 23: Truthy and falsy values
// Some values act like true or false in an if statement.
// Why this matters: it helps you understand bugs in conditions.
// More details: lessons/notes/62-js-truthy-falsy.md
const emptyText = "";
const someText = "hi";
if (emptyText) {
  console.log("Lesson 23: This will not run.");
}
if (someText) {
  console.log("Lesson 23: This will run.");
}

// -----------------------------------------------------------------------------
// LESSON 24: for...of loops (easy list looping)
// for...of gives you each item in a list.
// Why this matters: it is simpler than counting indexes.
const fruits = ["apple", "banana", "orange"];
for (const fruit of fruits) {
  console.log("Lesson 24:", fruit);
}

// -----------------------------------------------------------------------------
// LESSON 25: Array includes (is it in the list?)
// includes checks whether a list has a value.
// Why this matters: you can quickly check membership.
const pets = ["cat", "dog"];
console.log("Lesson 25:", pets.includes("dog"), pets.includes("fish"));

// -----------------------------------------------------------------------------
// LESSON 26: Object keys and values
// You can list the keys or values of an object.
// Why this matters: it helps you inspect data.
const profile = { name: "Ari", level: 3 };
console.log("Lesson 26 keys:", Object.keys(profile));
console.log("Lesson 26 values:", Object.values(profile));

// -----------------------------------------------------------------------------
// LESSON 27: Default function parameters
// You can set a default value if none is provided.
// Why this matters: your function stays safe and predictable.
function greet(name = "friend") {
  return `Hello, ${name}!`;
}
console.log("Lesson 27:", greet(), greet("Lee"));

// -----------------------------------------------------------------------------
// LESSON 28: String length and simple methods
// Strings have properties and methods.
// Why this matters: text processing is common.
const phrase = "hello world";
console.log("Lesson 28 length:", phrase.length);
console.log("Lesson 28 upper:", phrase.toUpperCase());

// -----------------------------------------------------------------------------
// LESSON 29: typeof (what kind of value is this?)
// typeof tells you the type of a value.
// Why this matters: it helps debug strange behavior.
console.log("Lesson 29:", typeof 123, typeof "abc", typeof true);

// -----------------------------------------------------------------------------
// LESSON 30: NaN and checking numbers
// NaN means “not a number.” It shows up when math fails.
// Why this matters: it helps you catch bad input.
// More details: lessons/notes/63-js-nan-and-typeof.md
const notNumber = Number("abc");
console.log("Lesson 30:", notNumber, Number.isNaN(notNumber));

// End of Lessons 21–30

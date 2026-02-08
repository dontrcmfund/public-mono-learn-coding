/**
 * JS BASICS (Lessons 1–20 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/01-js-basics-1-20.js
 *
 * Study for retention (not speed):
 * 1) Read one lesson block
 * 2) Say the "why" out loud
 * 3) Change one line and predict output before running
 *
 * Extra context:
 * - lessons/notes/60-what-is-javascript.md
 * - lessons/notes/78-how-to-study-this-curriculum.md
 * - lessons/notes/79-history-and-etymology-map.md
 */

// -----------------------------------------------------------------------------
// LESSON 1: What is a program?
// A program is a list of instructions for the computer.
// This line is an instruction that prints text to the screen.
// Why this matters: You are learning to write clear instructions.
console.log("Lesson 1: A program is a list of instructions.");

// -----------------------------------------------------------------------------
// LESSON 2: Comments are for humans.
// Comments help you (and future you) remember WHY something exists.
// The computer ignores comments.
// Why this matters: Code is read more often than it is written.
console.log("Lesson 2: Comments are for humans, not the computer.");

// -----------------------------------------------------------------------------
// LESSON 3: Values are pieces of information.
// In JavaScript, a value can be text, a number, or true/false.
// Why this matters: Programs are about moving and changing values.
const myName = "Alex"; // text value (a string)
const myAge = 20;      // number value
const isLearning = true; // boolean value (true/false)
console.log("Lesson 3 values:", myName, myAge, isLearning);

// -----------------------------------------------------------------------------
// LESSON 4: Variables are labeled boxes.
// A variable stores a value so you can use it later.
// Why this matters: You reuse information without rewriting it.
let score = 0; // start at zero
score = score + 1; // increase by 1
console.log("Lesson 4 score:", score);

// -----------------------------------------------------------------------------
// LESSON 5: Simple math is just instruction.
// JavaScript can do math like a calculator.
// Why this matters: Many programs are just math + decisions.
const total = 10 + 5;
const difference = 10 - 5;
const product = 10 * 5;
const quotient = 10 / 5;
console.log("Lesson 5 math:", total, difference, product, quotient);

// -----------------------------------------------------------------------------
// LESSON 6: Strings and combining text.
// A string is text wrapped in quotes.
// You can combine strings with the + operator.
// Why this matters: Most programs display or build text.
const firstName = "Sam";
const greeting = "Hello, " + firstName + "!";
console.log("Lesson 6:", greeting);

// -----------------------------------------------------------------------------
// LESSON 7: Arrays are ordered lists.
// Arrays store multiple values in one place.
// Why this matters: Lists show up everywhere (names, tasks, items).
const colors = ["red", "green", "blue"];
console.log("Lesson 7:", colors);
console.log("First color:", colors[0]); // arrays start at index 0

// -----------------------------------------------------------------------------
// LESSON 8: Objects are labeled groups.
// Objects store related information using key -> value pairs.
// Why this matters: Real things have properties (name, age, etc.).
const person = {
  name: "Jordan",
  age: 25,
  isStudent: true
};
console.log("Lesson 8:", person);
console.log("Person name:", person.name);

// -----------------------------------------------------------------------------
// LESSON 9: Functions are reusable steps.
// A function is a named recipe you can reuse.
// Why this matters: You avoid repeating yourself.
function sayHello(name) {
  return "Hello, " + name + "!";
}
console.log("Lesson 9:", sayHello("Avery"));

// -----------------------------------------------------------------------------
// LESSON 10: If statements are decisions.
// Code can make choices based on true/false.
// Why this matters: Programs react to different situations.
const temperature = 70;
if (temperature > 75) {
  console.log("Lesson 10: It is warm.");
} else {
  console.log("Lesson 10: It is cool.");
}

// -----------------------------------------------------------------------------
// LESSON 11: Comparison operators (true/false results)
// These help you ask questions like: is this bigger? is it equal?
// Why this matters: decisions need yes/no answers.
const a = 5;
const b = 10;
console.log("Lesson 11:", a < b, a === b, a !== b);

// -----------------------------------------------------------------------------
// LESSON 12: Logical operators (and/or/not)
// These combine multiple true/false answers.
// Why this matters: real decisions often have more than one condition.
const hasTicket = true;
const hasID = false;
console.log("Lesson 12 (AND):", hasTicket && hasID); // both must be true
console.log("Lesson 12 (OR):", hasTicket || hasID);  // either can be true
console.log("Lesson 12 (NOT):", !hasID);            // flips true/false

// -----------------------------------------------------------------------------
// LESSON 13: Loops (repeat work)
// A loop runs the same instructions multiple times.
// Why this matters: repetition is common (lists, steps, tasks).
for (let i = 1; i <= 3; i++) {
  console.log("Lesson 13 loop count:", i);
}

// -----------------------------------------------------------------------------
// LESSON 14: While loops (repeat until a condition changes)
// A while loop keeps running while something is true.
// Why this matters: some tasks repeat until they are done.
let count = 0;
while (count < 2) {
  console.log("Lesson 14 while:", count);
  count = count + 1;
}

// -----------------------------------------------------------------------------
// LESSON 15: Scope (where variables live)
// Variables only exist in certain places.
// Why this matters: it prevents name conflicts and confusion.
// If you want more background, see lessons/notes/60-what-is-javascript.md
const outer = "outside";
if (true) {
  const inner = "inside";
  console.log("Lesson 15:", outer, inner);
}
// console.log(inner); // This would cause an error (inner is out of scope)

// -----------------------------------------------------------------------------
// LESSON 16: Array length and adding items
// Arrays know how many items they contain.
// Why this matters: you can grow lists as you learn more.
const pets = ["cat", "dog"];
console.log("Lesson 16 length:", pets.length);
pets.push("bird"); // add to the end
console.log("Lesson 16 after push:", pets);

// -----------------------------------------------------------------------------
// LESSON 17: Updating object values
// Objects can change over time.
// Why this matters: real data changes (age, status, score).
const profile = { name: "Riley", level: 1 };
profile.level = 2; // update an existing value
console.log("Lesson 17:", profile);

// -----------------------------------------------------------------------------
// LESSON 18: Function parameters and return values
// Parameters are input; return values are output.
// Why this matters: functions are tiny machines that transform data.
function double(number) {
  return number * 2;
}
console.log("Lesson 18:", double(6));

// -----------------------------------------------------------------------------
// LESSON 19: Template strings (easier text building)
// Backticks let you insert values into text.
// Why this matters: it makes output clearer and less error‑prone.
const city = "Seattle";
const message = `Lesson 19: I live in ${city}.`;
console.log(message);

// -----------------------------------------------------------------------------
// LESSON 20: Errors are clues, not failures
// Errors are messages that help you fix problems.
// Why this matters: learning to read errors builds confidence.
// Try to imagine the message below without running it.
// If you want more context on JS history, see lessons/notes/60-what-is-javascript.md
console.log("Lesson 20: Errors are clues that guide you to a fix.");

// End of Lessons 1–20

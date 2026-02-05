/**
 * JS BASICS (Lessons 51–60 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/05-js-basics-51-60.js
 */

// -----------------------------------------------------------------------------
// LESSON 51: Function declarations vs arrow functions (same idea)
// Both create functions. Arrow functions are shorter.
// Why this matters: you will see both styles in real code.
function add(a, b) {
  return a + b;
}
const addArrow = (a, b) => a + b;
console.log("Lesson 51:", add(2, 3), addArrow(2, 3));

// -----------------------------------------------------------------------------
// LESSON 52: Functions as values (callbacks)
// You can pass a function into another function.
// Why this matters: it lets you customize behavior.
function runLater(action) {
  action();
}
runLater(() => {
  console.log("Lesson 52: This is a callback.");
});

// -----------------------------------------------------------------------------
// LESSON 53: Arrays of objects (real‑life data)
// Many real programs use lists of objects.
// Why this matters: it is how apps store people, tasks, items.
const students = [
  { name: "Ada", score: 95 },
  { name: "Bo", score: 72 },
  { name: "Cam", score: 88 }
];
console.log("Lesson 53:", students[0].name, students[0].score);

// -----------------------------------------------------------------------------
// LESSON 54: Combining map + filter on real data
// Filter removes items; map transforms the rest.
// Why this matters: it is a common data workflow.
const passingNames = students
  .filter((s) => s.score >= 80)
  .map((s) => s.name);
console.log("Lesson 54:", passingNames);

// -----------------------------------------------------------------------------
// LESSON 55: Object methods (functions inside objects)
// An object can store a function as a property.
// Why this matters: it groups behavior with data.
const calculator = {
  add(x, y) {
    return x + y;
  }
};
console.log("Lesson 55:", calculator.add(4, 5));

// -----------------------------------------------------------------------------
// LESSON 56: JSON (data you can save or send)
// JSON is a text format for data.
// Why this matters: it is how apps share data.
// More details: lessons/notes/66-js-json-and-try-catch.md
const user = { name: "Dee", level: 2 };
const jsonText = JSON.stringify(user);
const parsed = JSON.parse(jsonText);
console.log("Lesson 56:", jsonText, parsed.name);

// -----------------------------------------------------------------------------
// LESSON 57: try/catch (handling errors safely)
// try/catch lets you handle errors without crashing.
// Why this matters: it keeps programs stable.
// More details: lessons/notes/66-js-json-and-try-catch.md
try {
  JSON.parse("not valid json");
} catch (err) {
  console.log("Lesson 57: Caught error safely.");
}

// -----------------------------------------------------------------------------
// LESSON 58: Dates (basic use)
// Dates represent time.
// Why this matters: timestamps and scheduling are common.
const now = new Date();
console.log("Lesson 58:", now.toISOString());

// -----------------------------------------------------------------------------
// LESSON 59: Timers (run later)
// setTimeout runs code after a delay.
// Why this matters: many apps wait for things to happen.
// More details: lessons/notes/67-js-timers.md
setTimeout(() => {
  console.log("Lesson 59: This ran later.");
}, 500);

// -----------------------------------------------------------------------------
// LESSON 60: switch (many choices)
// switch is another way to handle multiple options.
// Why this matters: it can be cleaner than many else-if blocks.
const day = "mon";
switch (day) {
  case "mon":
    console.log("Lesson 60: Start of week");
    break;
  case "fri":
    console.log("Lesson 60: Almost weekend");
    break;
  default:
    console.log("Lesson 60: Regular day");
}

// End of Lessons 51–60

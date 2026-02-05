/**
 * NODE BASICS (Lessons 11–20)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/20-node-basics-11-20.js
 */

// -----------------------------------------------------------------------------
// LESSON 11: Environment variables
// process.env holds environment settings.
// Why this matters: it keeps secrets and settings outside code.
console.log("Lesson 11:", process.env.NODE_ENV || "(not set)");

// -----------------------------------------------------------------------------
// LESSON 12: Exit codes
// process.exitCode tells the OS if the script succeeded.
// Why this matters: other tools can detect failure.
process.exitCode = 0;

// -----------------------------------------------------------------------------
// LESSON 13: Reading a file async
// Async avoids blocking the program.
// Why this matters: it scales better for big tasks.
// More details: lessons/notes/73-js-async-promises.md
const fs = require("fs");
fs.readFile(__filename, "utf8", (err, text) => {
  if (err) {
    console.log("Lesson 13 error:", err.message);
    return;
  }
  console.log("Lesson 13: read this file, length", text.length);
});

// -----------------------------------------------------------------------------
// LESSON 14: Promises (basic idea)
// Promises represent future results.
// Why this matters: async code needs structure.
// More details: lessons/notes/73-js-async-promises.md
const fsPromises = require("fs/promises");
fsPromises.readFile(__filename, "utf8")
  .then((text) => {
    console.log("Lesson 14: promise read length", text.length);
  })
  .catch((err) => {
    console.log("Lesson 14 error:", err.message);
  });

// -----------------------------------------------------------------------------
// LESSON 15: Async/await (cleaner async)
// async/await is a cleaner way to write async code.
// Why this matters: it reads like normal code.
// More details: lessons/notes/73-js-async-promises.md
async function readSelf() {
  try {
    const text = await fsPromises.readFile(__filename, "utf8");
    console.log("Lesson 15: await read length", text.length);
  } catch (err) {
    console.log("Lesson 15 error:", err.message);
  }
}
readSelf();

// -----------------------------------------------------------------------------
// LESSON 16: Simple timers in Node
// setTimeout works in Node too.
// Why this matters: you can schedule tasks.
setTimeout(() => {
  console.log("Lesson 16: timer fired");
}, 300);

// -----------------------------------------------------------------------------
// LESSON 17: setInterval and clearInterval
// Repeat work until you stop it.
// Why this matters: it’s useful for polling.
let ticks = 0;
const intervalId = setInterval(() => {
  ticks += 1;
  console.log("Lesson 17: tick", ticks);
  if (ticks >= 2) {
    clearInterval(intervalId);
  }
}, 200);

// -----------------------------------------------------------------------------
// LESSON 18: Path basics
// path.resolve gives an absolute path.
// Why this matters: absolute paths are reliable.
const path = require("path");
console.log("Lesson 18:", path.resolve("."));

// -----------------------------------------------------------------------------
// LESSON 19: Checking if a file exists
// fs.existsSync returns true/false.
// Why this matters: it avoids crashes.
console.log("Lesson 19:", fs.existsSync(__filename));

// -----------------------------------------------------------------------------
// LESSON 20: Simple CLI output formatting
// Add labels to keep output clear.
// Why this matters: clarity reduces confusion.
console.log("Lesson 20: Done with Node basics 11–20");

// End of Node Basics 11–20

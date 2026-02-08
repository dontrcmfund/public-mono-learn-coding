/**
 * NODE BASICS (Lessons 1–10)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/19-node-basics-1-10.js
 *
 * This file creates temporary output files while teaching file IO.
 * Reading files and writing files is expected here.
 *
 * Extra context:
 * - lessons/notes/71-what-is-node.md
 * - lessons/notes/76-node-sync-vs-async.md
 * - lessons/notes/78-how-to-study-this-curriculum.md
 * - lessons/notes/79-history-and-etymology-map.md
 */

// -----------------------------------------------------------------------------
// LESSON 1: Node runs JavaScript on your computer
// console.log works here just like in the browser.
// Why this matters: you can write tools without a webpage.
console.log("Lesson 1: Hello from Node.js");

// -----------------------------------------------------------------------------
// LESSON 2: process.argv (reading arguments)
// Node gives you command-line arguments.
// Why this matters: you can build reusable scripts.
console.log("Lesson 2 argv:", process.argv.slice(2));

// -----------------------------------------------------------------------------
// LESSON 3: Built-in modules (fs, path)
// Node includes modules for files and paths.
// Why this matters: you can read/write files.
const fs = require("fs");
const path = require("path");

// -----------------------------------------------------------------------------
// LESSON 4: __dirname (where this file lives)
// __dirname is the folder of the current file.
// Why this matters: it helps build reliable paths.
console.log("Lesson 4 __dirname:", __dirname);

// -----------------------------------------------------------------------------
// LESSON 5: path.join (safe file paths)
// join builds paths that work across systems.
// Why this matters: it avoids broken paths.
const outPath = path.join(__dirname, "tmp-output.txt");
console.log("Lesson 5 outPath:", outPath);

// -----------------------------------------------------------------------------
// LESSON 6: Writing a file (sync)
// Write text to a file.
// Why this matters: scripts can save results.
// More details: lessons/notes/76-node-sync-vs-async.md
fs.writeFileSync(outPath, "Lesson 6: File created by Node\n");
console.log("Lesson 6: wrote file");

// -----------------------------------------------------------------------------
// LESSON 7: Reading a file (sync)
// Read text from a file.
// Why this matters: scripts can read data.
// More details: lessons/notes/76-node-sync-vs-async.md
const fileText = fs.readFileSync(outPath, "utf8");
console.log("Lesson 7:", fileText.trim());

// -----------------------------------------------------------------------------
// LESSON 8: Using a local module
// require can load your own files.
// Why this matters: you can split code into parts.
const math = require("./utils/math.js");
console.log("Lesson 8:", math.add(2, 3));

// -----------------------------------------------------------------------------
// LESSON 9: JSON in Node
// JSON is just text; Node can write it to a file.
// Why this matters: many tools use JSON configs.
const jsonPath = path.join(__dirname, "tmp-data.json");
const data = { ok: true, count: 3 };
fs.writeFileSync(jsonPath, JSON.stringify(data, null, 2));
console.log("Lesson 9: wrote JSON");

// -----------------------------------------------------------------------------
// LESSON 10: Safe file reading with try/catch
// Files might not exist; handle errors.
// Why this matters: scripts should fail gracefully.
try {
  const missing = fs.readFileSync(path.join(__dirname, "missing.txt"), "utf8");
  console.log(missing);
} catch (err) {
  console.log("Lesson 10: file not found (handled)");
}

// End of Node Basics 1–10

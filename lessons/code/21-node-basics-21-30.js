/**
 * NODE BASICS (Lessons 21–30)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/21-node-basics-21-30.js
 */

// -----------------------------------------------------------------------------
// LESSON 21: Reading a directory
// fs.readdirSync lists files in a folder.
// Why this matters: scripts often scan folders.
const fs = require("fs");
const path = require("path");
const dirItems = fs.readdirSync(__dirname);
console.log("Lesson 21:", dirItems.slice(0, 5));

// -----------------------------------------------------------------------------
// LESSON 22: Filtering files by extension
// You can filter for specific file types.
// Why this matters: it helps find what you need.
const jsFiles = dirItems.filter((name) => name.endsWith(".js"));
console.log("Lesson 22:", jsFiles);

// -----------------------------------------------------------------------------
// LESSON 23: Simple file stats
// fs.statSync tells you file info.
// Why this matters: you can check size or type.
const thisFile = path.join(__dirname, path.basename(__filename));
const stats = fs.statSync(thisFile);
console.log("Lesson 23 size:", stats.size);

// -----------------------------------------------------------------------------
// LESSON 24: Creating a folder if missing
// fs.existsSync + mkdirSync avoids errors.
// Why this matters: scripts can prepare folders.
const outputDir = path.join(__dirname, "tmp");
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir);
}
console.log("Lesson 24:", outputDir);

// -----------------------------------------------------------------------------
// LESSON 25: Writing multiple files
// Scripts can generate files.
// Why this matters: automation saves time.
for (let i = 1; i <= 3; i++) {
  const filePath = path.join(outputDir, `file-${i}.txt`);
  fs.writeFileSync(filePath, `File ${i}\n`);
}
console.log("Lesson 25: wrote 3 files");

// -----------------------------------------------------------------------------
// LESSON 26: Reading JSON from a file
// Read and parse JSON safely.
// Why this matters: configs and data are often JSON.
const jsonPath = path.join(outputDir, "data.json");
fs.writeFileSync(jsonPath, JSON.stringify({ ok: true }, null, 2));
const jsonText = fs.readFileSync(jsonPath, "utf8");
const json = JSON.parse(jsonText);
console.log("Lesson 26:", json.ok);

// -----------------------------------------------------------------------------
// LESSON 27: Basic CLI input from argv
// Use process.argv to accept input.
// Why this matters: scripts can be reused with different input.
const nameArg = process.argv[2] || "friend";
console.log(`Lesson 27: Hello, ${nameArg}`);

// -----------------------------------------------------------------------------
// LESSON 28: Simple HTTP server
// Node can start a server that returns text.
// Why this matters: servers are a core use of Node.
// More details: lessons/notes/72-node-http-server-gotchas.md
const http = require("http");
const server = http.createServer((req, res) => {
  res.writeHead(200, { "Content-Type": "text/plain" });
  res.end("Lesson 28: Hello from a Node server\n");
});
server.listen(3000, () => {
  console.log("Lesson 28: Server running at http://localhost:3000");
});

// -----------------------------------------------------------------------------
// LESSON 29: Stopping the server
// Call server.close() to stop it.
// Why this matters: you control the lifecycle.
setTimeout(() => {
  server.close(() => {
    console.log("Lesson 29: Server stopped");
  });
}, 2000);

// -----------------------------------------------------------------------------
// LESSON 30: Putting it together
// Node can read files, process data, and serve results.
// Why this matters: this is the foundation of backend work.
console.log("Lesson 30: Node can read files, process data, and serve responses.");

// End of Node Basics 21–30

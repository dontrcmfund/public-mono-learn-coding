/**
 * NODE BASICS (Lessons 41–50)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/23-node-basics-41-50.js
 */

// -----------------------------------------------------------------------------
// LESSON 41: fs.promises (modern async)
// fs.promises gives promise-based file operations.
// Why this matters: it works well with async/await.
const fs = require("fs/promises");
const path = require("path");

// -----------------------------------------------------------------------------
// LESSON 42: Write and read with async/await
// Write a file, then read it back.
// Why this matters: it is a common script pattern.
async function writeAndRead() {
  const filePath = path.join(__dirname, "tmp-async.txt");
  await fs.writeFile(filePath, "Lesson 42: async write\n");
  const text = await fs.readFile(filePath, "utf8");
  console.log("Lesson 42:", text.trim());
  await fs.unlink(filePath);
}
writeAndRead();

// -----------------------------------------------------------------------------
// LESSON 43: List directory with fs.promises
// readdir returns file names.
// Why this matters: many scripts scan folders.
async function listDir() {
  const items = await fs.readdir(__dirname);
  console.log("Lesson 43:", items.slice(0, 3));
}
listDir();

// -----------------------------------------------------------------------------
// LESSON 44: Simple file copy
// Copy a file with copyFile.
// Why this matters: backups are common.
async function copySelf() {
  const src = __filename;
  const dest = path.join(__dirname, "tmp-copy.js");
  await fs.copyFile(src, dest);
  console.log("Lesson 44: copied file");
  await fs.unlink(dest);
}
copySelf();

// -----------------------------------------------------------------------------
// LESSON 45: Create and remove folders
// mkdir and rmdir (with recursive) manage folders.
// Why this matters: scripts set up clean workspaces.
async function tempFolder() {
  const dir = path.join(__dirname, "tmp-dir");
  await fs.mkdir(dir, { recursive: true });
  console.log("Lesson 45: created", dir);
  await fs.rmdir(dir);
}
tempFolder();

// -----------------------------------------------------------------------------
// LESSON 46: Simple file existence check
// fs.access tells you if a file is reachable.
// Why this matters: you can avoid crashes.
async function checkAccess() {
  try {
    await fs.access(__filename);
    console.log("Lesson 46: file is accessible");
  } catch {
    console.log("Lesson 46: file not accessible");
  }
}
checkAccess();

// -----------------------------------------------------------------------------
// LESSON 47: Reading JSON safely (async)
// Wrap JSON.parse in try/catch.
// Why this matters: JSON can be invalid.
async function readJsonSafely() {
  const filePath = path.join(__dirname, "tmp.json");
  await fs.writeFile(filePath, JSON.stringify({ ok: true }));
  const text = await fs.readFile(filePath, "utf8");
  try {
    const data = JSON.parse(text);
    console.log("Lesson 47:", data.ok);
  } catch (err) {
    console.log("Lesson 47 error:", err.message);
  } finally {
    await fs.unlink(filePath);
  }
}
readJsonSafely();

// -----------------------------------------------------------------------------
// LESSON 48: Simple logger function
// Wrap console.log for consistent output.
// Why this matters: it keeps logs readable.
function log(label, value) {
  console.log(`${label}:`, value);
}
log("Lesson 48", "Logging is consistent");

// -----------------------------------------------------------------------------
// LESSON 49: Simple configuration pattern
// Load config from environment or default.
// Why this matters: scripts need flexible settings.
const PORT = process.env.PORT || 3000;
log("Lesson 49 PORT", PORT);

// -----------------------------------------------------------------------------
// LESSON 50: Summary
// Node lets you work with files, folders, and the OS.
// Why this matters: it’s the foundation for backend work.
log("Lesson 50", "Node is for scripts and servers");

// End of Node Basics 41–50

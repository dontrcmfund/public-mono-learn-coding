/**
 * JS BASICS (Lessons 41–50 in one file)
 *
 * How to run this file (later, when ready):
 * 1) Open a terminal in the repo
 * 2) Run: node lessons/code/04-js-basics-41-50.js
 */

// -----------------------------------------------------------------------------
// LESSON 41: Array map (transform every item)
// map creates a new array by transforming each item.
// Why this matters: it is a clean way to change data.
const prices = [5, 10, 15];
const doubledPrices = prices.map((p) => p * 2);
console.log("Lesson 41:", doubledPrices);

// -----------------------------------------------------------------------------
// LESSON 42: Array filter (keep what matches)
// filter keeps only the items that pass a test.
// Why this matters: it helps you remove unwanted data.
const numbers = [1, 2, 3, 4, 5];
const evens = numbers.filter((n) => n % 2 === 0);
console.log("Lesson 42:", evens);

// -----------------------------------------------------------------------------
// LESSON 43: Array find (first match)
// find returns the first item that matches a test.
// Why this matters: it helps you locate one thing quickly.
const names = ["Ava", "Ben", "Cara"];
const found = names.find((n) => n.startsWith("C"));
console.log("Lesson 43:", found);

// -----------------------------------------------------------------------------
// LESSON 44: Array reduce (combine into one value)
// reduce combines all items into a single result.
// Why this matters: it helps you total, sum, or build a summary.
const totals = [2, 4, 6];
const sum = totals.reduce((acc, n) => acc + n, 0);
console.log("Lesson 44:", sum);

// -----------------------------------------------------------------------------
// LESSON 45: Mutating vs non‑mutating methods
// Some methods change the original array, some do not.
// Why this matters: unexpected changes are a common bug.
const original = [1, 2, 3];
const sliced = original.slice(0, 2); // non‑mutating
original.push(4); // mutating
console.log("Lesson 45:", original, sliced);

// -----------------------------------------------------------------------------
// LESSON 46: forEach (do something for each item)
// forEach runs a function for every item, but returns nothing.
// Why this matters: it is good for side‑effects like logging.
const tasks = ["wash", "cook", "study"];
tasks.forEach((task) => {
  console.log("Lesson 46:", task);
});

// -----------------------------------------------------------------------------
// LESSON 47: Sorting arrays (careful!)
// sort changes the original array and sorts as strings by default.
// Why this matters: it can surprise you with numbers.
const nums = [10, 2, 30];
nums.sort();
console.log("Lesson 47:", nums); // "10" comes before "2" as strings

// -----------------------------------------------------------------------------
// LESSON 48: Sorting numbers with a compare function
// Provide a compare function to sort numbers correctly.
// Why this matters: numbers need numeric sorting, not string sorting.
const nums2 = [10, 2, 30];
nums2.sort((a, b) => a - b);
console.log("Lesson 48:", nums2);

// -----------------------------------------------------------------------------
// LESSON 49: Object cloning (shallow copy)
// Object spread makes a shallow copy.
// Why this matters: you can avoid accidental mutation.
const settings = { theme: "light", volume: 5 };
const settingsCopy = { ...settings };
settingsCopy.volume = 7;
console.log("Lesson 49:", settings, settingsCopy);

// -----------------------------------------------------------------------------
// LESSON 50: Optional chaining (safe access)
// Optional chaining avoids errors when a value might be missing.
// Why this matters: data can be incomplete.
const data = { user: { name: "Noa" } };
console.log("Lesson 50:", data.user?.name, data.profile?.name);

// End of Lessons 41–50

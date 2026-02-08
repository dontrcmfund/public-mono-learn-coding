/**
 * TYPESCRIPT BRIDGE (Lessons 1-10)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Read one lesson block at a time
 * 3) Change one line and predict editor feedback before running
 *
 * Extra context:
 * - lessons/notes/80-what-is-typescript.md
 * - lessons/notes/81-ts-common-gotchas.md
 * - lessons/notes/78-how-to-study-this-curriculum.md
 */

// -----------------------------------------------------------------------------
// LESSON 1: Type annotations on variables
// Why this matters: you make value expectations explicit.
const learnerName: string = "Alex";
const attempts: number = 3;
const isActive: boolean = true;
console.log("Lesson 1:", learnerName, attempts, isActive);

// -----------------------------------------------------------------------------
// LESSON 2: Type inference
// TypeScript can infer many types without annotations.
// Why this matters: cleaner code with safety.
const city = "Seattle"; // inferred as string
const year = 2026; // inferred as number
console.log("Lesson 2:", city, year);

// -----------------------------------------------------------------------------
// LESSON 3: Function parameter and return types
// Why this matters: function contracts become clear.
function add(a: number, b: number): number {
  return a + b;
}
console.log("Lesson 3:", add(2, 5));

// -----------------------------------------------------------------------------
// LESSON 4: Union types
// A value can be one of several types.
// Why this matters: real data is often flexible.
function printId(id: string | number): void {
  console.log("Lesson 4 ID:", id);
}
printId("abc");
printId(42);

// -----------------------------------------------------------------------------
// LESSON 5: Type narrowing
// Narrowing means proving which union member you have.
// Why this matters: safe operations need known types.
function normalize(input: string | number): string {
  if (typeof input === "number") {
    return input.toString();
  }
  return input.trim();
}
console.log("Lesson 5:", normalize("  hi  "), normalize(99));

// -----------------------------------------------------------------------------
// LESSON 6: Type aliases
// Alias gives a reusable name to a type shape.
// Why this matters: clearer communication in teams.
type Student = {
  name: string;
  level: number;
};
const s1: Student = { name: "Mia", level: 1 };
console.log("Lesson 6:", s1);

// -----------------------------------------------------------------------------
// LESSON 7: Interfaces
// Interface also describes object shape.
// Why this matters: common pattern in TypeScript codebases.
interface Task {
  id: number;
  title: string;
  done: boolean;
}
const t1: Task = { id: 1, title: "Read lesson", done: false };
console.log("Lesson 7:", t1);

// -----------------------------------------------------------------------------
// LESSON 8: Optional properties
// Optional means property may be missing.
// Why this matters: API and form data are often partial.
interface Profile {
  username: string;
  bio?: string;
}
const p1: Profile = { username: "dev_starter" };
console.log("Lesson 8:", p1.bio ?? "(no bio yet)");

// -----------------------------------------------------------------------------
// LESSON 9: Array types
// Two equivalent styles: T[] and Array<T>.
// Why this matters: typed lists reduce surprises.
const scores: number[] = [80, 90, 100];
const labels: Array<string> = ["A", "B", "C"];
console.log("Lesson 9:", scores, labels);

// -----------------------------------------------------------------------------
// LESSON 10: Unknown vs any
// any disables safety; unknown requires checks.
// Why this matters: unknown keeps you honest.
function safeLength(value: unknown): number {
  if (typeof value === "string") {
    return value.length;
  }
  return 0;
}
console.log("Lesson 10:", safeLength("typescript"), safeLength(123));

// End of TypeScript Bridge 1-10

/**
 * TYPESCRIPT BRIDGE (Lessons 11-20)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Read one lesson block at a time
 * 3) Predict type feedback before changing code
 *
 * Extra context:
 * - lessons/notes/81-ts-common-gotchas.md
 * - lessons/notes/82-ts-tooling-runtime.md
 * - lessons/notes/79-history-and-etymology-map.md
 */

// -----------------------------------------------------------------------------
// LESSON 11: Literal types
// Limit values to exact allowed options.
// Why this matters: prevents invalid states.
type Theme = "light" | "dark";
const activeTheme: Theme = "light";
console.log("Lesson 11:", activeTheme);

// -----------------------------------------------------------------------------
// LESSON 12: Tuples
// Tuple is a fixed-length, ordered type.
// Why this matters: position-based data stays predictable.
const point: [number, number] = [10, 20];
console.log("Lesson 12:", point);

// -----------------------------------------------------------------------------
// LESSON 13: Readonly properties
// readonly prevents reassignment after creation.
// Why this matters: protects important values.
interface Course {
  readonly id: string;
  title: string;
}
const c1: Course = { id: "ts-101", title: "TypeScript Basics" };
c1.title = "TypeScript Basics Updated";
console.log("Lesson 13:", c1.id, c1.title);

// -----------------------------------------------------------------------------
// LESSON 14: Record utility type
// Record<K, V> maps keys to a value type.
// Why this matters: concise typed dictionaries.
const scoresByName: Record<string, number> = {
  alex: 90,
  mia: 95
};
console.log("Lesson 14:", scoresByName);

// -----------------------------------------------------------------------------
// LESSON 15: Partial utility type
// Partial<T> makes all properties optional.
// Why this matters: useful for updates/patch objects.
interface Settings {
  language: string;
  notifications: boolean;
}
const patch: Partial<Settings> = { notifications: false };
console.log("Lesson 15:", patch);

// -----------------------------------------------------------------------------
// LESSON 16: Pick and Omit utility types
// Pick keeps fields; Omit removes fields.
// Why this matters: build precise types from existing ones.
interface User {
  id: number;
  name: string;
  email: string;
}
type PublicUser = Pick<User, "id" | "name">;
type UserWithoutEmail = Omit<User, "email">;
const publicUser: PublicUser = { id: 1, name: "Kai" };
const userNoEmail: UserWithoutEmail = { id: 1, name: "Kai" };
console.log("Lesson 16:", publicUser, userNoEmail);

// -----------------------------------------------------------------------------
// LESSON 17: Function type aliases
// Name function signatures for reuse.
// Why this matters: callback contracts stay clear.
type MathOp = (a: number, b: number) => number;
const multiply: MathOp = (a, b) => a * b;
console.log("Lesson 17:", multiply(3, 4));

// -----------------------------------------------------------------------------
// LESSON 18: Generics basics
// Generic lets type be provided later.
// Why this matters: reusable functions stay type-safe.
function identity<T>(value: T): T {
  return value;
}
console.log("Lesson 18:", identity<string>("hello"), identity<number>(42));

// -----------------------------------------------------------------------------
// LESSON 19: Generic constraints
// Constrain T to types that have required properties.
// Why this matters: safe operations inside generic code.
function logLength<T extends { length: number }>(value: T): T {
  console.log("Lesson 19 length:", value.length);
  return value;
}
logLength("typescript");
logLength([1, 2, 3]);

// -----------------------------------------------------------------------------
// LESSON 20: Discriminated unions
// Use a shared "kind" field to narrow safely.
// Why this matters: reliable branching on object variants.
type Circle = { kind: "circle"; radius: number };
type Square = { kind: "square"; size: number };
type Shape = Circle | Square;

function area(shape: Shape): number {
  if (shape.kind === "circle") {
    return Math.PI * shape.radius * shape.radius;
  }
  return shape.size * shape.size;
}

console.log("Lesson 20:", area({ kind: "circle", radius: 2 }), area({ kind: "square", size: 3 }));

// End of TypeScript Bridge 11-20

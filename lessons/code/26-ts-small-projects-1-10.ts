/**
 * TYPESCRIPT SMALL PROJECTS (Lessons 1-10)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Run one lesson block at a time
 * 3) Change one rule, predict the new output, then run
 *
 * Extra context:
 * - lessons/notes/84-ts-small-projects-principles.md
 * - lessons/notes/82-ts-tooling-runtime.md
 */

// -----------------------------------------------------------------------------
// LESSON 1 PROJECT: Counter with typed actions
// Why this matters: many UIs are state + actions.

type CounterAction = "inc" | "dec" | "reset";

function applyCounter(state: number, action: CounterAction): number {
  if (action === "inc") return state + 1;
  if (action === "dec") return state - 1;
  return 0;
}

let counterState = 0;
counterState = applyCounter(counterState, "inc");
counterState = applyCounter(counterState, "inc");
counterState = applyCounter(counterState, "dec");
console.log("Lesson 1:", counterState);

// -----------------------------------------------------------------------------
// LESSON 2 PROJECT: Todo list data model
// Why this matters: clear models reduce confusion as features grow.

type Todo = {
  id: number;
  title: string;
  done: boolean;
};

const todos: Todo[] = [
  { id: 1, title: "Read TypeScript lesson", done: true },
  { id: 2, title: "Practice project", done: false }
];
console.log("Lesson 2:", todos.length);

// -----------------------------------------------------------------------------
// LESSON 3 PROJECT: Mark todo done by id
// Why this matters: updates by id are common in real apps.

function markDone(items: Todo[], id: number): Todo[] {
  return items.map((item) => {
    if (item.id === id) return { ...item, done: true };
    return item;
  });
}

const updatedTodos = markDone(todos, 2);
console.log("Lesson 3:", updatedTodos);

// -----------------------------------------------------------------------------
// LESSON 4 PROJECT: Shopping cart subtotal
// Why this matters: totals and summaries are core product logic.

type CartItem = {
  name: string;
  price: number;
  qty: number;
};

const cart: CartItem[] = [
  { name: "Notebook", price: 5, qty: 2 },
  { name: "Pen", price: 2, qty: 3 }
];

function cartSubtotal(items: CartItem[]): number {
  return items.reduce((sum, item) => sum + item.price * item.qty, 0);
}

console.log("Lesson 4:", cartSubtotal(cart));

// -----------------------------------------------------------------------------
// LESSON 5 PROJECT: Grade calculator with validation
// Why this matters: validation prevents silent bad results.

function averageGrade(scores: number[]): number {
  if (scores.length === 0) return 0;
  const valid = scores.every((s) => s >= 0 && s <= 100);
  if (!valid) return 0;
  const total = scores.reduce((sum, s) => sum + s, 0);
  return total / scores.length;
}

console.log("Lesson 5:", averageGrade([90, 80, 100]));

// -----------------------------------------------------------------------------
// LESSON 6 PROJECT: Budget status function
// Why this matters: business rules are just typed decision logic.

type BudgetStatus = "under" | "over" | "exact";

function getBudgetStatus(limit: number, spent: number): BudgetStatus {
  if (spent < limit) return "under";
  if (spent > limit) return "over";
  return "exact";
}

console.log("Lesson 6:", getBudgetStatus(100, 120));

// -----------------------------------------------------------------------------
// LESSON 7 PROJECT: Search notes (case-insensitive)
// Why this matters: search behavior appears in almost every app.

function searchNotes(notes: string[], query: string): string[] {
  const q = query.toLowerCase().trim();
  return notes.filter((n) => n.toLowerCase().includes(q));
}

console.log("Lesson 7:", searchNotes(["TypeScript", "Docker", "DOM"], "do"));

// -----------------------------------------------------------------------------
// LESSON 8 PROJECT: Safe API-like response parsing
// Why this matters: external data can be partial or malformed.

type UserSummary = {
  username: string;
  points: number;
};

function parseUserSummary(input: unknown): UserSummary | null {
  if (typeof input !== "object" || input === null) return null;
  const maybeUser = input as { username?: unknown; points?: unknown };
  if (typeof maybeUser.username !== "string") return null;
  if (typeof maybeUser.points !== "number") return null;
  return { username: maybeUser.username, points: maybeUser.points };
}

console.log("Lesson 8:", parseUserSummary({ username: "kai", points: 20 }));

// -----------------------------------------------------------------------------
// LESSON 9 PROJECT: Reusable sorter with generics
// Why this matters: generic utilities scale across project features.

function sortByNumberField<T>(items: T[], selector: (item: T) => number): T[] {
  return [...items].sort((a, b) => selector(a) - selector(b));
}

const ranked = sortByNumberField(
  [
    { name: "A", score: 20 },
    { name: "B", score: 10 }
  ],
  (item) => item.score
);
console.log("Lesson 9:", ranked);

// -----------------------------------------------------------------------------
// LESSON 10 PROJECT: Mini leaderboard summary
// Why this matters: combining typed helpers is real project work.

type Player = {
  name: string;
  score: number;
};

function topPlayer(players: Player[]): Player | null {
  if (players.length === 0) return null;
  const sorted = sortByNumberField(players, (p) => -p.score);
  return sorted[0];
}

const players: Player[] = [
  { name: "Mia", score: 50 },
  { name: "Leo", score: 70 },
  { name: "Ana", score: 65 }
];

console.log("Lesson 10:", topPlayer(players));

// End of TypeScript Small Projects 1-10

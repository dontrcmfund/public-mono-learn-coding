/**
 * TYPESCRIPT SMALL PROJECTS (Lessons 11-20)
 *
 * Suggested use:
 * 1) Open this file in VS Code
 * 2) Run one lesson block at a time
 * 3) Change one rule, predict output, then run
 *
 * Extra context:
 * - lessons/notes/84-ts-small-projects-principles.md
 * - lessons/notes/85-ts-small-projects-gotchas.md
 */

// -----------------------------------------------------------------------------
// LESSON 11 PROJECT: Pagination helper
// Why this matters: lists often need page-based views.

function paginate<T>(items: T[], page: number, pageSize: number): T[] {
  const safePage = Math.max(1, page);
  const safePageSize = Math.max(1, pageSize);
  const start = (safePage - 1) * safePageSize;
  return items.slice(start, start + safePageSize);
}

console.log("Lesson 11:", paginate([1, 2, 3, 4, 5], 2, 2));

// -----------------------------------------------------------------------------
// LESSON 12 PROJECT: Form validation result type
// Why this matters: typed result objects make UI messaging easier.

type ValidationResult =
  | { ok: true }
  | { ok: false; errors: string[] };

function validateSignup(email: string, password: string): ValidationResult {
  const errors: string[] = [];
  if (!email.includes("@")) errors.push("Email must include @");
  if (password.length < 8) errors.push("Password must be at least 8 chars");
  if (errors.length > 0) return { ok: false, errors };
  return { ok: true };
}

console.log("Lesson 12:", validateSignup("dev@example.com", "12345678"));

// -----------------------------------------------------------------------------
// LESSON 13 PROJECT: Inventory low-stock report
// Why this matters: status reports are common in admin tools.

type StockItem = {
  sku: string;
  name: string;
  quantity: number;
  reorderAt: number;
};

function lowStock(items: StockItem[]): StockItem[] {
  return items.filter((item) => item.quantity <= item.reorderAt);
}

const inventory: StockItem[] = [
  { sku: "A1", name: "Notebook", quantity: 3, reorderAt: 5 },
  { sku: "B1", name: "Marker", quantity: 12, reorderAt: 5 }
];
console.log("Lesson 13:", lowStock(inventory));

// -----------------------------------------------------------------------------
// LESSON 14 PROJECT: Add unique tags
// Why this matters: uniqueness rules prevent duplicate UI data.

function addUniqueTag(tags: string[], newTag: string): string[] {
  const normalized = newTag.trim().toLowerCase();
  if (!normalized) return tags;
  if (tags.includes(normalized)) return tags;
  return [...tags, normalized];
}

console.log("Lesson 14:", addUniqueTag(["typescript"], "TypeScript"));

// -----------------------------------------------------------------------------
// LESSON 15 PROJECT: Sort users by role then name
// Why this matters: predictable ordering improves usability.

type User = {
  name: string;
  role: "admin" | "editor" | "viewer";
};

const roleRank: Record<User["role"], number> = {
  admin: 1,
  editor: 2,
  viewer: 3
};

function sortUsers(users: User[]): User[] {
  return [...users].sort((a, b) => {
    const roleDiff = roleRank[a.role] - roleRank[b.role];
    if (roleDiff !== 0) return roleDiff;
    return a.name.localeCompare(b.name);
  });
}

console.log(
  "Lesson 15:",
  sortUsers([
    { name: "Mia", role: "viewer" },
    { name: "Ana", role: "admin" },
    { name: "Leo", role: "editor" }
  ])
);

// -----------------------------------------------------------------------------
// LESSON 16 PROJECT: Safe settings merge
// Why this matters: config updates should keep defaults intact.

type Settings = {
  language: "en" | "es";
  darkMode: boolean;
  itemsPerPage: number;
};

function mergeSettings(base: Settings, patch: Partial<Settings>): Settings {
  return { ...base, ...patch };
}

const defaultSettings: Settings = { language: "en", darkMode: false, itemsPerPage: 20 };
console.log("Lesson 16:", mergeSettings(defaultSettings, { darkMode: true }));

// -----------------------------------------------------------------------------
// LESSON 17 PROJECT: Transaction summary
// Why this matters: financial summaries appear in many apps.

type Transaction = {
  id: string;
  type: "income" | "expense";
  amount: number;
};

function summarize(transactions: Transaction[]): { income: number; expense: number; net: number } {
  const income = transactions
    .filter((t) => t.type === "income")
    .reduce((sum, t) => sum + t.amount, 0);
  const expense = transactions
    .filter((t) => t.type === "expense")
    .reduce((sum, t) => sum + t.amount, 0);
  return { income, expense, net: income - expense };
}

console.log(
  "Lesson 17:",
  summarize([
    { id: "1", type: "income", amount: 1200 },
    { id: "2", type: "expense", amount: 300 }
  ])
);

// -----------------------------------------------------------------------------
// LESSON 18 PROJECT: Retry policy helper
// Why this matters: network and IO tasks often need retries.

type RetryDecision = {
  shouldRetry: boolean;
  nextDelayMs: number;
};

function getRetryDecision(attempt: number, maxAttempts: number): RetryDecision {
  if (attempt >= maxAttempts) return { shouldRetry: false, nextDelayMs: 0 };
  const delay = 200 * Math.pow(2, attempt - 1);
  return { shouldRetry: true, nextDelayMs: delay };
}

console.log("Lesson 18:", getRetryDecision(2, 4));

// -----------------------------------------------------------------------------
// LESSON 19 PROJECT: Route parser
// Why this matters: URL path parsing drives many front-end/back-end features.

type ParsedRoute = {
  resource: string;
  id: string | null;
};

function parseRoute(path: string): ParsedRoute {
  const [resource, id] = path.replace(/^\/+/, "").split("/");
  return { resource: resource || "", id: id || null };
}

console.log("Lesson 19:", parseRoute("/users/42"));

// -----------------------------------------------------------------------------
// LESSON 20 PROJECT: End-to-end mini task service
// Why this matters: composing typed helpers is real software design.

type Task = {
  id: number;
  title: string;
  done: boolean;
};

function addTask(list: Task[], title: string): Task[] {
  const nextId = list.length === 0 ? 1 : Math.max(...list.map((t) => t.id)) + 1;
  return [...list, { id: nextId, title: title.trim(), done: false }];
}

function completeTask(list: Task[], id: number): Task[] {
  return list.map((task) => (task.id === id ? { ...task, done: true } : task));
}

let taskState: Task[] = [];
taskState = addTask(taskState, "Write types");
taskState = addTask(taskState, "Run lesson");
taskState = completeTask(taskState, 1);
console.log("Lesson 20:", taskState);

// End of TypeScript Small Projects 11-20

"""
PYTHON DSA BASICS (Lessons 11-20)

Suggested use:
1) Run: python3 lessons/code/50-python-dsa-basics-11-20.py
2) Focus on pattern recognition: sliding window, recursion, sorting, graph traversal
3) Change inputs and compare output behavior

Extra context:
- lessons/notes/121-big-o-first-principles.md
"""

from collections import deque


# -----------------------------------------------------------------------------
# LESSON 11: Sliding window max sum of fixed size
# Why this matters: avoids repeated full-sum work.
def max_sum_k(items: list[int], k: int) -> int | None:
    if k <= 0 or k > len(items):
        return None
    window = sum(items[:k])
    best = window
    for i in range(k, len(items)):
        window += items[i] - items[i - k]
        best = max(best, window)
    return best

print("Lesson 11:", max_sum_k([2, 1, 5, 1, 3, 2], 3))


# -----------------------------------------------------------------------------
# LESSON 12: Merge two sorted lists
# Why this matters: merge logic is a core sorting/composition primitive.
def merge_sorted(a: list[int], b: list[int]) -> list[int]:
    i = j = 0
    out: list[int] = []
    while i < len(a) and j < len(b):
        if a[i] <= b[j]:
            out.append(a[i])
            i += 1
        else:
            out.append(b[j])
            j += 1
    out.extend(a[i:])
    out.extend(b[j:])
    return out

print("Lesson 12:", merge_sorted([1, 4, 8], [2, 3, 7]))


# -----------------------------------------------------------------------------
# LESSON 13: Recursion basics (factorial)
# Why this matters: recursive decomposition appears in many algorithms.
def factorial(n: int) -> int:
    if n <= 1:
        return 1
    return n * factorial(n - 1)

print("Lesson 13:", factorial(5))


# -----------------------------------------------------------------------------
# LESSON 14: Memoization (top-down DP)
# Why this matters: cache removes repeated expensive subproblems.
def fib(n: int, memo: dict[int, int] | None = None) -> int:
    if memo is None:
        memo = {}
    if n in memo:
        return memo[n]
    if n <= 1:
        return n
    memo[n] = fib(n - 1, memo) + fib(n - 2, memo)
    return memo[n]

print("Lesson 14:", fib(10))


# -----------------------------------------------------------------------------
# LESSON 15: In-place reverse list
# Why this matters: pointer swaps are a common space-efficient pattern.
def reverse_in_place(items: list[int]) -> None:
    left, right = 0, len(items) - 1
    while left < right:
        items[left], items[right] = items[right], items[left]
        left += 1
        right -= 1

arr = [1, 2, 3, 4]
reverse_in_place(arr)
print("Lesson 15:", arr)


# -----------------------------------------------------------------------------
# LESSON 16: Check balanced parentheses (stack)
# Why this matters: validation of nested structures appears in parsers.
def is_balanced(text: str) -> bool:
    stack: list[str] = []
    pairs = {")": "(", "]": "[", "}": "{"}
    opens = set(pairs.values())
    for ch in text:
        if ch in opens:
            stack.append(ch)
        elif ch in pairs:
            if not stack or stack.pop() != pairs[ch]:
                return False
    return not stack

print("Lesson 16:", is_balanced("([]{})"), is_balanced("([)]"))


# -----------------------------------------------------------------------------
# LESSON 17: BFS shortest steps in unweighted graph
# Why this matters: BFS finds shortest path length in unweighted graphs.
def shortest_steps(graph: dict[str, list[str]], start: str, target: str) -> int:
    q = deque([(start, 0)])
    seen = {start}
    while q:
        node, dist = q.popleft()
        if node == target:
            return dist
        for nxt in graph.get(node, []):
            if nxt not in seen:
                seen.add(nxt)
                q.append((nxt, dist + 1))
    return -1

sample_graph = {
    "A": ["B", "C"],
    "B": ["D"],
    "C": ["D", "E"],
    "D": ["F"],
    "E": [],
    "F": [],
}
print("Lesson 17:", shortest_steps(sample_graph, "A", "F"))


# -----------------------------------------------------------------------------
# LESSON 18: Greedy interval scheduling (count non-overlap)
# Why this matters: some optimization problems have efficient greedy solutions.
def max_non_overlapping(intervals: list[tuple[int, int]]) -> int:
    sorted_intervals = sorted(intervals, key=lambda x: x[1])
    count = 0
    current_end = float("-inf")
    for start, end in sorted_intervals:
        if start >= current_end:
            count += 1
            current_end = end
    return count

print("Lesson 18:", max_non_overlapping([(1, 3), (2, 4), (3, 5), (0, 7), (8, 9)]))


# -----------------------------------------------------------------------------
# LESSON 19: Top-k via sorting
# Why this matters: ranking and selection are frequent requirements.
def top_k(items: list[int], k: int) -> list[int]:
    return sorted(items, reverse=True)[: max(0, k)]

print("Lesson 19:", top_k([9, 1, 7, 3, 8], 3))


# -----------------------------------------------------------------------------
# LESSON 20: Complexity reflection helper
# Why this matters: naming complexity trains decision-making.
def complexity_notes() -> dict[str, str]:
    return {
        "binary_search": "O(log n)",
        "linear_search": "O(n)",
        "merge_sorted": "O(n + m)",
        "bfs": "O(V + E)",
    }

print("Lesson 20:", complexity_notes())

# End of Python DSA Basics 11-20

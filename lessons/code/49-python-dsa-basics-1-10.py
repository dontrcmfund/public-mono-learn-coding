"""
PYTHON DSA BASICS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/49-python-dsa-basics-1-10.py
2) For each lesson, identify the data structure and algorithm used
3) Modify input size and observe behavior

Extra context:
- lessons/notes/120-what-are-data-structures-and-algorithms.md
- lessons/notes/121-big-o-first-principles.md
"""

from collections import deque


# -----------------------------------------------------------------------------
# LESSON 1: List append and index access
# Why this matters: lists are default sequence structures in Python.
values = [1, 2, 3]
values.append(4)
print("Lesson 1:", values[0], values[-1])


# -----------------------------------------------------------------------------
# LESSON 2: Set membership for fast lookup
# Why this matters: sets are efficient for membership checks.
seen = {"mia", "leo", "ana"}
print("Lesson 2:", "mia" in seen, "kai" in seen)


# -----------------------------------------------------------------------------
# LESSON 3: Dictionary key-value mapping
# Why this matters: maps enable direct access by key.
scores = {"mia": 92, "leo": 78}
scores["ana"] = 88
print("Lesson 3:", scores.get("ana"), len(scores))


# -----------------------------------------------------------------------------
# LESSON 4: Linear search
# Why this matters: baseline scanning algorithm for unsorted data.
def linear_search(items: list[int], target: int) -> int:
    for i, value in enumerate(items):
        if value == target:
            return i
    return -1

print("Lesson 4:", linear_search([5, 8, 2, 9], 2))


# -----------------------------------------------------------------------------
# LESSON 5: Binary search (sorted input)
# Why this matters: uses order to reduce comparisons.
def binary_search(sorted_items: list[int], target: int) -> int:
    left, right = 0, len(sorted_items) - 1
    while left <= right:
        mid = (left + right) // 2
        if sorted_items[mid] == target:
            return mid
        if sorted_items[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1

print("Lesson 5:", binary_search([1, 3, 5, 7, 9], 7))


# -----------------------------------------------------------------------------
# LESSON 6: Stack with list
# Why this matters: LIFO behavior appears in undo/backtracking.
stack: list[str] = []
stack.append("open")
stack.append("edit")
print("Lesson 6 pop:", stack.pop())


# -----------------------------------------------------------------------------
# LESSON 7: Queue with deque
# Why this matters: FIFO behavior appears in scheduling and BFS.
queue = deque(["task1", "task2"])
queue.append("task3")
print("Lesson 7:", queue.popleft())


# -----------------------------------------------------------------------------
# LESSON 8: Frequency counting
# Why this matters: counting categories is common in analytics/logging.
def frequency(items: list[str]) -> dict[str, int]:
    counts: dict[str, int] = {}
    for item in items:
        counts[item] = counts.get(item, 0) + 1
    return counts

print("Lesson 8:", frequency(["a", "b", "a", "c", "a"]))


# -----------------------------------------------------------------------------
# LESSON 9: Two-pointer technique
# Why this matters: efficient scanning on sorted arrays.
def has_pair_with_sum(sorted_items: list[int], target_sum: int) -> bool:
    left, right = 0, len(sorted_items) - 1
    while left < right:
        current = sorted_items[left] + sorted_items[right]
        if current == target_sum:
            return True
        if current < target_sum:
            left += 1
        else:
            right -= 1
    return False

print("Lesson 9:", has_pair_with_sum([1, 2, 4, 7, 11], 9))


# -----------------------------------------------------------------------------
# LESSON 10: Prefix sum pattern
# Why this matters: range-sum queries become faster after preprocessing.
def prefix_sums(items: list[int]) -> list[int]:
    out = [0]
    running = 0
    for x in items:
        running += x
        out.append(running)
    return out

pref = prefix_sums([3, 5, 2, 6])
# Sum from index 1 to 3 => pref[4] - pref[1]
print("Lesson 10:", pref, "range_sum(1..3)=", pref[4] - pref[1])

# End of Python DSA Basics 1-10

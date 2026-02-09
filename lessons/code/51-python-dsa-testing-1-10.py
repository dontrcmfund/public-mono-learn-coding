"""
PYTHON DSA TESTING (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/51-python-dsa-testing-1-10.py
2) Read PASS/FAIL output
3) Break one algorithm intentionally and verify tests catch it

Extra context:
- lessons/notes/122-python-dsa-gotchas.md
- lessons/notes/123-dsa-problem-solving-workflow.md
"""

from collections import deque


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


def linear_search(items: list[int], target: int) -> int:
    for i, value in enumerate(items):
        if value == target:
            return i
    return -1


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


# -----------------------------------------------------------------------------
# LESSON 1: linear search found case
run_test("Lesson 1 linear found", linear_search([5, 8, 2], 2) == 2)

# LESSON 2: linear search missing case
run_test("Lesson 2 linear missing", linear_search([5, 8, 2], 9) == -1)

# LESSON 3: binary search found case
run_test("Lesson 3 binary found", binary_search([1, 3, 5, 7, 9], 7) == 3)

# LESSON 4: binary search missing case
run_test("Lesson 4 binary missing", binary_search([1, 3, 5, 7, 9], 2) == -1)

# LESSON 5: pair sum true case
run_test("Lesson 5 pair sum true", has_pair_with_sum([1, 2, 4, 7, 11], 9) is True)

# LESSON 6: pair sum false case
run_test("Lesson 6 pair sum false", has_pair_with_sum([1, 2, 4, 7, 11], 20) is False)

# LESSON 7: BFS shortest path
sample_graph = {
    "A": ["B", "C"],
    "B": ["D"],
    "C": ["D", "E"],
    "D": ["F"],
    "E": [],
    "F": [],
}
run_test("Lesson 7 BFS shortest", shortest_steps(sample_graph, "A", "F") == 3)

# LESSON 8: BFS unreachable path
run_test("Lesson 8 BFS unreachable", shortest_steps(sample_graph, "E", "A") == -1)

# LESSON 9: empty input edge case
run_test("Lesson 9 empty linear", linear_search([], 1) == -1)

# LESSON 10: completion marker
print("Lesson 10: DSA testing suite complete")

# End of Python DSA Testing 1-10

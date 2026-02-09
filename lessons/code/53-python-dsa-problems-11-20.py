"""
PYTHON DSA PROBLEMS (Lessons 11-20)

Suggested use:
1) Run: python3 lessons/code/53-python-dsa-problems-11-20.py
2) Identify pattern type for each lesson (graph, heap, DP, etc.)
3) Change input size and reason about complexity

Extra context:
- lessons/notes/123-dsa-problem-solving-workflow.md
- lessons/notes/122-python-dsa-gotchas.md
"""

import heapq


# -----------------------------------------------------------------------------
# LESSON 11: K smallest values (heap)
# Why this matters: heap gives efficient partial ranking.
def k_smallest(nums: list[int], k: int) -> list[int]:
    if k <= 0:
        return []
    return heapq.nsmallest(k, nums)

print("Lesson 11:", k_smallest([9, 4, 7, 1, 3], 3))


# -----------------------------------------------------------------------------
# LESSON 12: Top frequent element
# Why this matters: frequency + ranking is common in logs/analytics.
def top_frequent(nums: list[int]) -> int | None:
    if not nums:
        return None
    counts: dict[int, int] = {}
    for n in nums:
        counts[n] = counts.get(n, 0) + 1
    return max(counts, key=counts.get)

print("Lesson 12:", top_frequent([1, 2, 2, 3, 2, 1]))


# -----------------------------------------------------------------------------
# LESSON 13: Coin change (DP)
# Why this matters: classic optimization with overlapping subproblems.
def min_coins(coins: list[int], amount: int) -> int:
    dp = [float("inf")] * (amount + 1)
    dp[0] = 0
    for a in range(1, amount + 1):
        for c in coins:
            if a - c >= 0:
                dp[a] = min(dp[a], dp[a - c] + 1)
    return -1 if dp[amount] == float("inf") else int(dp[amount])

print("Lesson 13:", min_coins([1, 3, 4], 6))


# -----------------------------------------------------------------------------
# LESSON 14: Longest increasing subsequence (O(n^2) DP)
# Why this matters: sequence optimization is a recurring interview pattern.
def lis_length(nums: list[int]) -> int:
    if not nums:
        return 0
    dp = [1] * len(nums)
    for i in range(len(nums)):
        for j in range(i):
            if nums[j] < nums[i]:
                dp[i] = max(dp[i], dp[j] + 1)
    return max(dp)

print("Lesson 14:", lis_length([10, 9, 2, 5, 3, 7, 101, 18]))


# -----------------------------------------------------------------------------
# LESSON 15: Number of islands (DFS on grid)
# Why this matters: graph traversal over matrix data is common.
def count_islands(grid: list[list[str]]) -> int:
    if not grid:
        return 0

    rows, cols = len(grid), len(grid[0])

    def dfs(r: int, c: int) -> None:
        if r < 0 or c < 0 or r >= rows or c >= cols or grid[r][c] != "1":
            return
        grid[r][c] = "0"
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)

    islands = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "1":
                islands += 1
                dfs(r, c)
    return islands

grid_sample = [
    ["1", "1", "0", "0"],
    ["1", "0", "0", "1"],
    ["0", "0", "1", "1"],
]
print("Lesson 15:", count_islands([row[:] for row in grid_sample]))


# -----------------------------------------------------------------------------
# LESSON 16: Cycle detection in linked list (fast/slow pointers)
# Why this matters: pointer speed technique is high-value.
class Node:
    def __init__(self, value: int, next_node: "Node | None" = None):
        self.value = value
        self.next = next_node


def has_cycle(head: Node | None) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False

n1 = Node(1)
n2 = Node(2)
n3 = Node(3)
n1.next = n2
n2.next = n3
n3.next = n2
print("Lesson 16:", has_cycle(n1))


# -----------------------------------------------------------------------------
# LESSON 17: Parentheses generation count (backtracking)
# Why this matters: controlled search with constraints.
def generate_parentheses(n: int) -> list[str]:
    out: list[str] = []

    def backtrack(current: str, open_count: int, close_count: int) -> None:
        if len(current) == 2 * n:
            out.append(current)
            return
        if open_count < n:
            backtrack(current + "(", open_count + 1, close_count)
        if close_count < open_count:
            backtrack(current + ")", open_count, close_count + 1)

    backtrack("", 0, 0)
    return out

print("Lesson 17 count:", len(generate_parentheses(3)))


# -----------------------------------------------------------------------------
# LESSON 18: Meeting rooms (interval overlap)
# Why this matters: scheduling conflict detection is practical.
def min_meeting_rooms(intervals: list[tuple[int, int]]) -> int:
    if not intervals:
        return 0
    starts = sorted(i[0] for i in intervals)
    ends = sorted(i[1] for i in intervals)

    s = e = 0
    rooms = max_rooms = 0
    while s < len(intervals):
        if starts[s] < ends[e]:
            rooms += 1
            max_rooms = max(max_rooms, rooms)
            s += 1
        else:
            rooms -= 1
            e += 1
    return max_rooms

print("Lesson 18:", min_meeting_rooms([(0, 30), (5, 10), (15, 20)]))


# -----------------------------------------------------------------------------
# LESSON 19: Longest common prefix
# Why this matters: string prefix matching appears in search/autocomplete.
def longest_common_prefix(words: list[str]) -> str:
    if not words:
        return ""
    prefix = words[0]
    for word in words[1:]:
        while not word.startswith(prefix):
            prefix = prefix[:-1]
            if not prefix:
                return ""
    return prefix

print("Lesson 19:", longest_common_prefix(["flower", "flow", "flight"]))


# -----------------------------------------------------------------------------
# LESSON 20: Complexity recap map
# Why this matters: mapping pattern to complexity builds fluency.
def complexity_map() -> dict[str, str]:
    return {
        "k_smallest": "O(n log k) or O(n log n) (library strategy dependent)",
        "min_coins": "O(amount * len(coins))",
        "count_islands": "O(rows * cols)",
        "min_meeting_rooms": "O(n log n)",
    }

print("Lesson 20:", complexity_map())

# End of Python DSA Problems 11-20

"""
PYTHON DSA PROBLEMS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/52-python-dsa-problems-1-10.py
2) For each lesson, explain chosen pattern out loud
3) Modify constraints and reason about complexity

Extra context:
- lessons/notes/123-dsa-problem-solving-workflow.md
- lessons/notes/121-big-o-first-principles.md
"""


# -----------------------------------------------------------------------------
# LESSON 1: Two Sum (sorted, two-pointer)
# Why this matters: classic pointer coordination pattern.
def two_sum_sorted(nums: list[int], target: int) -> tuple[int, int] | None:
    left, right = 0, len(nums) - 1
    while left < right:
        total = nums[left] + nums[right]
        if total == target:
            return left, right
        if total < target:
            left += 1
        else:
            right -= 1
    return None

print("Lesson 1:", two_sum_sorted([1, 2, 4, 7, 11], 9))


# -----------------------------------------------------------------------------
# LESSON 2: Valid palindrome (two-pointer)
# Why this matters: boundary movement without extra memory.
def is_palindrome(text: str) -> bool:
    left, right = 0, len(text) - 1
    while left < right:
        if text[left] != text[right]:
            return False
        left += 1
        right -= 1
    return True

print("Lesson 2:", is_palindrome("racecar"), is_palindrome("python"))


# -----------------------------------------------------------------------------
# LESSON 3: Maximum subarray sum (Kadane)
# Why this matters: dynamic programming in one pass.
def max_subarray(nums: list[int]) -> int:
    best = nums[0]
    current = nums[0]
    for n in nums[1:]:
        current = max(n, current + n)
        best = max(best, current)
    return best

print("Lesson 3:", max_subarray([4, -1, 2, 1, -5, 4]))


# -----------------------------------------------------------------------------
# LESSON 4: Product except self (prefix/suffix)
# Why this matters: avoid division and keep linear complexity.
def product_except_self(nums: list[int]) -> list[int]:
    n = len(nums)
    out = [1] * n
    prefix = 1
    for i in range(n):
        out[i] = prefix
        prefix *= nums[i]
    suffix = 1
    for i in range(n - 1, -1, -1):
        out[i] *= suffix
        suffix *= nums[i]
    return out

print("Lesson 4:", product_except_self([1, 2, 3, 4]))


# -----------------------------------------------------------------------------
# LESSON 5: Merge intervals
# Why this matters: sorting + scan is a common range problem pattern.
def merge_intervals(intervals: list[tuple[int, int]]) -> list[tuple[int, int]]:
    if not intervals:
        return []
    intervals = sorted(intervals)
    merged = [intervals[0]]
    for start, end in intervals[1:]:
        last_start, last_end = merged[-1]
        if start <= last_end:
            merged[-1] = (last_start, max(last_end, end))
        else:
            merged.append((start, end))
    return merged

print("Lesson 5:", merge_intervals([(1, 3), (2, 6), (8, 10), (9, 12)]))


# -----------------------------------------------------------------------------
# LESSON 6: Anagram check using frequency map
# Why this matters: map-based counting appears frequently.
def is_anagram(a: str, b: str) -> bool:
    if len(a) != len(b):
        return False
    counts: dict[str, int] = {}
    for ch in a:
        counts[ch] = counts.get(ch, 0) + 1
    for ch in b:
        if ch not in counts:
            return False
        counts[ch] -= 1
        if counts[ch] == 0:
            del counts[ch]
    return not counts

print("Lesson 6:", is_anagram("listen", "silent"))


# -----------------------------------------------------------------------------
# LESSON 7: Longest unique substring (sliding window)
# Why this matters: window resizing on constraints is a high-value pattern.
def longest_unique_substring_len(text: str) -> int:
    seen: dict[str, int] = {}
    left = 0
    best = 0
    for right, ch in enumerate(text):
        if ch in seen and seen[ch] >= left:
            left = seen[ch] + 1
        seen[ch] = right
        best = max(best, right - left + 1)
    return best

print("Lesson 7:", longest_unique_substring_len("abcabcbb"))


# -----------------------------------------------------------------------------
# LESSON 8: Rotate array right by k
# Why this matters: modular indexing and list slicing practice.
def rotate_right(nums: list[int], k: int) -> list[int]:
    if not nums:
        return []
    k %= len(nums)
    return nums[-k:] + nums[:-k]

print("Lesson 8:", rotate_right([1, 2, 3, 4, 5], 2))


# -----------------------------------------------------------------------------
# LESSON 9: Majority element (Boyer-Moore)
# Why this matters: elegant linear-time constant-space voting algorithm.
def majority_element(nums: list[int]) -> int:
    candidate = nums[0]
    count = 0
    for n in nums:
        if count == 0:
            candidate = n
        count += 1 if n == candidate else -1
    return candidate

print("Lesson 9:", majority_element([2, 2, 1, 1, 2, 2, 2]))


# -----------------------------------------------------------------------------
# LESSON 10: Complexity quick summary for problem set
# Why this matters: explicit complexity reflection builds intuition.
def complexity_summary() -> dict[str, str]:
    return {
        "two_sum_sorted": "O(n)",
        "max_subarray": "O(n)",
        "merge_intervals": "O(n log n)",
        "longest_unique_substring_len": "O(n)",
    }

print("Lesson 10:", complexity_summary())

# End of Python DSA Problems 1-10

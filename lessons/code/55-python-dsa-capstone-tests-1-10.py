"""
PYTHON DSA CAPSTONE TESTS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/55-python-dsa-capstone-tests-1-10.py
2) Read PASS/FAIL output
3) Break one helper in 54 file and verify failures
"""


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


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


def is_palindrome(text: str) -> bool:
    left, right = 0, len(text) - 1
    while left < right:
        if text[left] != text[right]:
            return False
        left += 1
        right -= 1
    return True


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


def parse_intervals(raw_items: list[str]) -> list[tuple[int, int]] | None:
    out: list[tuple[int, int]] = []
    for item in raw_items:
        if ":" not in item:
            return None
        left, right = item.split(":", 1)
        try:
            start, end = int(left), int(right)
        except ValueError:
            return None
        if start > end:
            return None
        out.append((start, end))
    return out


# -----------------------------------------------------------------------------
# LESSON 1: two-sum success
run_test("Lesson 1 two-sum found", two_sum_sorted([1, 2, 4, 7, 11], 9) == (1, 3))

# LESSON 2: two-sum not found
run_test("Lesson 2 two-sum missing", two_sum_sorted([1, 2, 4, 7, 11], 100) is None)

# LESSON 3: palindrome true
run_test("Lesson 3 palindrome true", is_palindrome("racecar") is True)

# LESSON 4: palindrome false
run_test("Lesson 4 palindrome false", is_palindrome("python") is False)

# LESSON 5: merge intervals overlap
run_test(
    "Lesson 5 merge overlap",
    merge_intervals([(1, 3), (2, 6), (8, 10)]) == [(1, 6), (8, 10)],
)

# LESSON 6: merge intervals empty
run_test("Lesson 6 merge empty", merge_intervals([]) == [])

# LESSON 7: parse intervals valid
run_test("Lesson 7 parse valid", parse_intervals(["1:3", "4:6"]) == [(1, 3), (4, 6)])

# LESSON 8: parse intervals invalid order
run_test("Lesson 8 parse invalid order", parse_intervals(["5:2"]) is None)

# LESSON 9: parse intervals invalid format
run_test("Lesson 9 parse invalid format", parse_intervals(["bad"]) is None)

# LESSON 10: completion marker
print("Lesson 10: DSA capstone tests complete")

# End of Python DSA Capstone Tests 1-10

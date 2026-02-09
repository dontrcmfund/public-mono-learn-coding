"""
PYTHON DSA CAPSTONE (Lessons 1-10)

Project: Algorithm Toolkit CLI

Suggested use:
1) Run examples:
   python3 lessons/code/54-python-dsa-capstone-1-10.py two-sum --nums 1 2 4 7 11 --target 9
   python3 lessons/code/54-python-dsa-capstone-1-10.py palindrome --text racecar
   python3 lessons/code/54-python-dsa-capstone-1-10.py merge-intervals --intervals 1:3 2:6 8:10

Extra context:
- lessons/notes/124-dsa-capstone-plan.md
- lessons/notes/123-dsa-problem-solving-workflow.md
"""

import argparse


# -----------------------------------------------------------------------------
# LESSON 1: two-sum sorted implementation
# Why this matters: demonstrates two-pointer toolkit command.
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


# -----------------------------------------------------------------------------
# LESSON 2: palindrome checker
# Why this matters: string pointer checks are common utility logic.
def is_palindrome(text: str) -> bool:
    left, right = 0, len(text) - 1
    while left < right:
        if text[left] != text[right]:
            return False
        left += 1
        right -= 1
    return True


# -----------------------------------------------------------------------------
# LESSON 3: merge intervals helper
# Why this matters: interval merge appears in scheduling/report logic.
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


# -----------------------------------------------------------------------------
# LESSON 4: parser for interval args
# Why this matters: CLI input parsing must be explicit and safe.
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
# LESSON 5: subcommand handlers
# Why this matters: clear command boundaries support maintainability.
def handle_two_sum(nums: list[int], target: int) -> int:
    result = two_sum_sorted(nums, target)
    print("Lesson 5:", result)
    return 0


def handle_palindrome(text: str) -> int:
    print("Lesson 5:", is_palindrome(text))
    return 0


def handle_merge(intervals_raw: list[str]) -> int:
    parsed = parse_intervals(intervals_raw)
    if parsed is None:
        print("Lesson 5: invalid intervals. Use start:end with start <= end")
        return 2
    print("Lesson 5:", merge_intervals(parsed))
    return 0


# -----------------------------------------------------------------------------
# LESSON 6: parser builder
# Why this matters: command grammar should be explicit.
def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="DSA toolkit CLI")
    sub = parser.add_subparsers(dest="cmd", required=True)

    two_sum = sub.add_parser("two-sum", help="Two-sum on sorted numbers")
    two_sum.add_argument("--nums", nargs="+", type=int, required=True)
    two_sum.add_argument("--target", type=int, required=True)

    pal = sub.add_parser("palindrome", help="Check palindrome")
    pal.add_argument("--text", required=True)

    merge = sub.add_parser("merge-intervals", help="Merge intervals")
    merge.add_argument("--intervals", nargs="+", required=True)

    return parser


# -----------------------------------------------------------------------------
# LESSON 7: dispatch
# Why this matters: centralized routing simplifies extension.
def dispatch(args: argparse.Namespace) -> int:
    if args.cmd == "two-sum":
        if args.nums != sorted(args.nums):
            print("Lesson 7: nums must be sorted for two-pointer method")
            return 2
        return handle_two_sum(args.nums, args.target)
    if args.cmd == "palindrome":
        return handle_palindrome(args.text)
    if args.cmd == "merge-intervals":
        return handle_merge(args.intervals)
    print("Unknown command")
    return 1


# -----------------------------------------------------------------------------
# LESSON 8: complexity hint output
# Why this matters: tools can teach alongside execution.
def print_complexity_hint(cmd: str) -> None:
    hints = {
        "two-sum": "O(n)",
        "palindrome": "O(n)",
        "merge-intervals": "O(n log n)",
    }
    print("Lesson 8 complexity:", hints.get(cmd, "unknown"))


# -----------------------------------------------------------------------------
# LESSON 9: main orchestration
# Why this matters: parse -> run -> explain flow is consistent.
def main() -> int:
    parser = build_parser()
    args = parser.parse_args()
    code = dispatch(args)
    if code == 0:
        print_complexity_hint(args.cmd)
    return code


# -----------------------------------------------------------------------------
# LESSON 10: entrypoint
# Why this matters: explicit script execution boundary.
if __name__ == "__main__":
    raise SystemExit(main())

# End of Python DSA Capstone 1-10

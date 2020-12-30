package main

type problemMapping map[int]chapter

type chapter struct {
	Number   int
	Name     string
	Problems []problem
}

type problem struct {
	Name   string
	Folder string
}

func isChapterNumberValid(n int) bool {
	_, ok := chapters[n]
	return ok
}

func getAllChapterNumbers() []int {
	chapterKeys := make([]int, 0, len(chapters))
	for k := range chapters {
		chapterKeys = append(chapterKeys, k)
	}

	return chapterKeys
}

var chapters = problemMapping{
	4: {
		Name:   "Chapter 04: Primitive Types",
		Number: 4,
		Problems: []problem{
			{
				Name:   "4.00 Bootcamp: Primitive Types",
				Folder: "count_bits",
			},
			{
				Name:   "4.01 Computing the parity of a word",
				Folder: "parity",
			},
			{
				Name:   "4.02 Swap bits",
				Folder: "swap_bits",
			},
			{
				Name:   "4.03 Reverse bits",
				Folder: "reverse_bits",
			},
			{
				Name:   "4.04 Find a closest integer with the same weight",
				Folder: "closest_int_same_weight",
			},
			{
				Name:   "4.05 Compute product without arithmetical operators",
				Folder: "primitive_multiply",
			},
			{
				Name:   "4.06 Compute quotient without arithmetical operators",
				Folder: "primitive_divide",
			},
			{
				Name:   "4.07 Compute pow(x,y)",
				Folder: "power_xy",
			},
			{
				Name:   "4.08 Reverse digits",
				Folder: "reverse_digits",
			},
			{
				Name:   "4.09 Check if a decimal integer is a palindrome",
				Folder: "is_number_palindromic",
			},
			{
				Name:   "4.10 Generate uniform random numbers",
				Folder: "uniform_random_number",
			},
			{
				Name:   "4.11 Rectangle intersection",
				Folder: "rectangle_intersection",
			},
		},
	},
	5: {
		Name:   "Chapter 05: Arrays",
		Number: 5,
		Problems: []problem{
			{
				Name:   "5.00 Bootcamp: Arrays",
				Folder: "even_odd_array",
			},
			{
				Name:   "5.01 The Dutch national flag problem",
				Folder: "dutch_national_flag",
			},
			{
				Name:   "5.02 Increment an arbitrary-precision integer",
				Folder: "int_as_array_increment",
			},
			{
				Name:   "5.03 Multiply two arbitrary-precision integers",
				Folder: "int_as_array_multiply",
			},
			{
				Name:   "5.04 Advancing through an array",
				Folder: "advance_by_offsets",
			},
			{
				Name:   "5.05 Delete duplicates from a sorted array",
				Folder: "sorted_array_remove_dups",
			},
			{
				Name:   "5.06 Buy and sell a stock once",
				Folder: "buy_and_sell_stock",
			},
			{
				Name:   "5.07 Buy and sell a stock twice",
				Folder: "buy_and_sell_stock_twice",
			},
			{
				Name:   "5.08 Computing an alternation",
				Folder: "alternating_array",
			},
			{
				Name:   "5.09 Enumerate all primes to n",
				Folder: "prime_sieve",
			},
			{
				Name:   "5.10 Permute the elements of an array",
				Folder: "apply_permutation",
			},
			{
				Name:   "5.11 Compute the next permutation",
				Folder: "next_permutation",
			},
			{
				Name:   "5.12 Sample offline data",
				Folder: "offline_sampling",
			},
			{
				Name:   "5.13 Sample online data",
				Folder: "online_sampling",
			},
			{
				Name:   "5.14 Compute a random permutation",
				Folder: "random_permutation",
			},
			{
				Name:   "5.15 Compute a random subset",
				Folder: "random_subset",
			},
			{
				Name:   "5.16 Generate nonuniform random numbers",
				Folder: "nonuniform_random_number",
			},
			{
				Name:   "5.17 The Sudoku checker problem",
				Folder: "is_valid_sudoku",
			},
			{
				Name:   "5.18 Compute the spiral ordering of a 2D array",
				Folder: "spiral_ordering",
			},
			{
				Name:   "5.19 Rotate a 2D array",
				Folder: "matrix_rotation",
			},
			{
				Name:   "5.20 Compute rows in Pascal's Triangle",
				Folder: "pascal_triangle",
			},
		},
	},
	6: {
		Name:   "Chapter 06: Strings",
		Number: 6,
		Problems: []problem{
			{
				Name:   "6.00 Bootcamp: Strings",
				Folder: "is_string_palindromic",
			},
			{
				Name:   "6.01 Interconvert strings and integers",
				Folder: "string_integer_interconversion",
			},
			{
				Name:   "6.02 Base conversion",
				Folder: "convert_base",
			},
			{
				Name:   "6.03 Compute the spreadsheet column encoding",
				Folder: "spreadsheet_encoding",
			},
			{
				Name:   "6.04 Replace and remove",
				Folder: "replace_and_remove",
			},
			{
				Name:   "6.05 Test palindromicity",
				Folder: "is_string_palindromic_punctuation",
			},
			{
				Name:   "6.06 Reverse all the words in a sentence",
				Folder: "reverse_words",
			},
			{
				Name:   "6.07 The look-and-say problem",
				Folder: "look_and_say",
			},
			{
				Name:   "6.08 Convert from Roman to decimal",
				Folder: "roman_to_integer",
			},
			{
				Name:   "6.09 Compute all valid IP addresses",
				Folder: "valid_ip_addresses",
			},
			{
				Name:   "6.10 Write a string sinusoidally",
				Folder: "snake_string",
			},
			{
				Name:   "6.11 Implement run-length encoding",
				Folder: "run_length_compression",
			},
			{
				Name:   "6.12 Find the first occurrence of a substring",
				Folder: "substring_match",
			},
		},
	},
	7: {
		Name:   "Chapter 07: Linked Lists",
		Number: 7,
		Problems: []problem{
			{
				Name:   "7.00 Bootcamp: Delete From List",
				Folder: "delete_from_list",
			},
			{
				Name:   "7.00 Bootcamp: Insert In List",
				Folder: "insert_in_list",
			},
			{
				Name:   "7.00 Bootcamp: Search In List",
				Folder: "search_in_list",
			},
			{
				Name:   "7.01 Merge two sorted lists",
				Folder: "sorted_lists_merge",
			},
			{
				Name:   "7.02 Reverse a single sublist",
				Folder: "reverse_sublist",
			},
			{
				Name:   "7.03 Test for cyclicity",
				Folder: "is_list_cyclic",
			},
			{
				Name:   "7.04 Test for overlapping lists---lists are cycle-free",
				Folder: "do_terminated_lists_overlap",
			},
			{
				Name:   "7.05 Test for overlapping lists---lists may have cycles",
				Folder: "do_lists_overlap",
			},
			{
				Name:   "7.06 Delete a node from a singly linked list",
				Folder: "delete_node_from_list",
			},
			{
				Name:   "7.07 Remove the kth last element from a list",
				Folder: "delete_kth_last_from_list",
			},
			{
				Name:   "7.08 Remove duplicates from a sorted list",
				Folder: "remove_duplicates_from_sorted_list",
			},
			{
				Name:   "7.09 Implement cyclic right shift for singly linked lists",
				Folder: "list_cyclic_right_shift",
			},
			{
				Name:   "7.10 Implement even-odd merge",
				Folder: "even_odd_list_merge",
			},
			{
				Name:   "7.11 Test whether a singly linked list is palindromic",
				Folder: "is_list_palindromic",
			},
			{
				Name:   "7.12 Implement list pivoting",
				Folder: "pivot_list",
			},
			{
				Name:   "7.13 Add list-based integers",
				Folder: "int_as_list_add",
			},
		},
	},
	8: {
		Name:   "Chapter 08: Stacks and Queues",
		Number: 8,
		Problems: []problem{
			{
				Name:   "8.01 Implement a stack with max API",
				Folder: "stack_with_max",
			},
			{
				Name:   "8.02 Evaluate RPN expressions",
				Folder: "evaluate_rpn",
			},
			{
				Name:   "8.03 Is a string well-formed?",
				Folder: "is_valid_parenthesization",
			},
			{
				Name:   "8.04 Normalize pathnames",
				Folder: "directory_path_normalization",
			},
			{
				Name:   "8.05 Compute buildings with a sunset view",
				Folder: "sunset_view",
			},
			{
				Name:   "8.06 Compute binary tree nodes in order of increasing depth",
				Folder: "tree_level_order",
			},
			{
				Name:   "8.07 Implement a circular queue",
				Folder: "circular_queue",
			},
			{
				Name:   "8.08 Implement a queue using stacks",
				Folder: "queue_from_stacks",
			},
			{
				Name:   "8.09 Implement a queue with max API",
				Folder: "queue_with_max",
			},
		},
	},
	9: {
		Name:   "Chapter 09: Binary Trees",
		Number: 9,
		Problems: []problem{
			{
				Name:   "9.01 Test if a binary tree is height-balanced",
				Folder: "is_tree_balanced",
			},
			{
				Name:   "9.02 Test if a binary tree is symmetric",
				Folder: "is_tree_symmetric",
			},
			{
				Name:   "9.03 Compute the lowest common ancestor in a binary tree",
				Folder: "lowest_common_ancestor",
			},
			{
				Name:   "9.04 Compute the LCA when nodes have parent pointers",
				Folder: "lowest_common_ancestor_with_parent",
			},
			{
				Name:   "9.05 Sum the root-to-leaf paths in a binary tree",
				Folder: "sum_root_to_leaf",
			},
			{
				Name:   "9.06 Find a root to leaf path with specified sum",
				Folder: "path_sum",
			},
			{
				Name:   "9.07 Implement an inorder traversal without recursion",
				Folder: "tree_inorder",
			},
			{
				Name:   "9.08 Compute the kth node in an inorder traversal",
				Folder: "kth_node_in_tree",
			},
			{
				Name:   "9.09 Compute the successor",
				Folder: "successor_in_tree",
			},
			{
				Name:   "9.10 Implement an inorder traversal with constant space",
				Folder: "tree_with_parent_inorder",
			},
			{
				Name:   "9.11 Reconstruct a binary tree from traversal data",
				Folder: "tree_from_preorder_inorder",
			},
			{
				Name:   "9.12 Reconstruct a binary tree from a preorder traversal with markers",
				Folder: "tree_from_preorder_with_null",
			},
			{
				Name:   "9.13 Compute the leaves of a binary tree",
				Folder: "tree_connect_leaves",
			},
			{
				Name:   "9.14 Compute the exterior of a binary tree",
				Folder: "tree_exterior",
			},
			{
				Name:   "9.15 Compute the right sibling tree",
				Folder: "tree_right_sibling",
			},
		},
	},
	10: {
		Name:   "Chapter 10: Heaps",
		Number: 10,
		Problems: []problem{
			{
				Name:   "10.01 Merge sorted files",
				Folder: "sorted_arrays_merge",
			},
			{
				Name:   "10.02 Sort an increasing-decreasing array",
				Folder: "sort_increasing_decreasing_array",
			},
			{
				Name:   "10.03 Sort an almost-sorted array",
				Folder: "sort_almost_sorted_array",
			},
			{
				Name:   "10.04 Compute the k closest stars",
				Folder: "k_closest_stars",
			},
			{
				Name:   "10.05 Compute the median of online data",
				Folder: "online_median",
			},
			{
				Name:   "10.06 Compute the k largest elements in a max-heap",
				Folder: "k_largest_in_heap",
			},
		},
	},
	11: {
		Name:   "Chapter 11: Searching",
		Number: 11,
		Problems: []problem{
			{
				Name:   "11.01 Search a sorted array for first occurrence of k",
				Folder: "search_first_key",
			},
			{
				Name:   "11.02 Search a sorted array for entry equal to its index",
				Folder: "search_entry_equal_to_index",
			},
			{
				Name:   "11.03 Search a cyclically sorted array",
				Folder: "search_shifted_sorted_array",
			},
			{
				Name:   "11.04 Compute the integer square root",
				Folder: "int_square_root",
			},
			{
				Name:   "11.05 Compute the real square root",
				Folder: "real_square_root",
			},
			{
				Name:   "11.06 Search in a 2D sorted array",
				Folder: "search_row_col_sorted_matrix",
			},
			{
				Name:   "11.07 Find the min and max simultaneously",
				Folder: "search_for_min_max_in_array",
			},
			{
				Name:   "11.08 Find the kth largest element",
				Folder: "kth_largest_in_array",
			},
			{
				Name:   "11.09 Find the missing IP address",
				Folder: "absent_value_array",
			},
			{
				Name:   "11.10 Find the duplicate and missing elements",
				Folder: "search_for_missing_element",
			},
		},
	},
	12: {
		Name:   "Chapter 12: Hash Tables",
		Number: 12,
		Problems: []problem{
			{
				Name:   "12.00 Bootcamp: Hash Tables",
				Folder: "anagrams",
			},
			{
				Name:   "12.01 Test for palindromic permutations",
				Folder: "is_string_permutable_to_palindrome",
			},
			{
				Name:   "12.02 Is an anonymous letter constructible?",
				Folder: "is_anonymous_letter_constructible",
			},
			{
				Name:   "12.03 Implement an ISBN cache",
				Folder: "lru_cache",
			},
			{
				Name:   "12.04 Compute the LCA, optimizing for close ancestors",
				Folder: "lowest_common_ancestor_close_ancestor",
			},
			{
				Name:   "12.05 Find the nearest repeated entries in an array",
				Folder: "nearest_repeated_entries",
			},
			{
				Name:   "12.06 Find the smallest subarray covering all values",
				Folder: "smallest_subarray_covering_set",
			},
			{
				Name:   "12.07 Find smallest subarray sequentially covering all values",
				Folder: "smallest_subarray_covering_all_values",
			},
			{
				Name:   "12.08 Find the longest subarray with distinct entries",
				Folder: "longest_subarray_with_distinct_values",
			},
			{
				Name:   "12.09 Find the length of a longest contained interval",
				Folder: "longest_contained_interval",
			},
			{
				Name:   "12.10 Compute all string decompositions",
				Folder: "string_decompositions_into_dictionary_words",
			},
			{
				Name:   "12.11 Test the Collatz conjecture",
				Folder: "collatz_checker",
			},
		},
	},
	13: {
		Name:   "Chapter 13: Sorting",
		Number: 13,
		Problems: []problem{
			{
				Name:   "13.01 Compute the intersection of two sorted arrays",
				Folder: "intersect_sorted_arrays",
			},
			{
				Name:   "13.02 Merge two sorted arrays",
				Folder: "two_sorted_arrays_merge",
			},
			{
				Name:   "13.03 Computing the h-index",
				Folder: "h_index",
			},
			{
				Name:   "13.04 Remove first-name duplicates",
				Folder: "remove_duplicates",
			},
			{
				Name:   "13.05 Smallest nonconstructible value",
				Folder: "smallest_nonconstructible_value",
			},
			{
				Name:   "13.06 Render a calendar",
				Folder: "calendar_rendering",
			},
			{
				Name:   "13.07 Merging intervals",
				Folder: "interval_add",
			},
			{
				Name:   "13.08 Compute the union of intervals",
				Folder: "intervals_union",
			},
			{
				Name:   "13.09 Partitioning and sorting an array with many repeated entries",
				Folder: "group_equal_entries",
			},
			{
				Name:   "13.10 Team photo day---1",
				Folder: "is_array_dominated",
			},
			{
				Name:   "13.11 Implement a fast sorting algorithm for lists",
				Folder: "sort_list",
			},
			{
				Name:   "13.12 Compute a salary threshold",
				Folder: "find_salary_threshold",
			},
		},
	},
	14: {
		Name:   "Chapter 14: Binary Search Trees",
		Number: 14,
		Problems: []problem{
			{
				Name:   "14.00 Bootcamp: Binary Search Trees",
				Folder: "search_in_bst",
			},
			{
				Name:   "14.01 Test if a binary tree satisfies the BST property",
				Folder: "is_tree_a_bst",
			},
			{
				Name:   "14.02 Find the first key greater than a given value in a BST",
				Folder: "search_first_greater_value_in_bst",
			},
			{
				Name:   "14.03 Find the k largest elements in a BST",
				Folder: "k_largest_values_in_bst",
			},
			{
				Name:   "14.04 Compute the LCA in a BST",
				Folder: "lowest_common_ancestor_in_bst",
			},
			{
				Name:   "14.05 Reconstruct a BST from traversal data",
				Folder: "bst_from_preorder",
			},
			{
				Name:   "14.06 Find the closest entries in three sorted arrays",
				Folder: "minimum_distance_3_sorted_arrays",
			},
			{
				Name:   "14.07 Enumerate extended integers",
				Folder: "ab_sqrt_2",
			},
			{
				Name:   "14.08 Build a minimum height BST from a sorted array",
				Folder: "bst_from_sorted_array",
			},
			{
				Name:   "14.09 Test if three BST nodes are totally ordered",
				Folder: "descendant_and_ancestor_in_bst",
			},
			{
				Name:   "14.10 The range lookup problem",
				Folder: "range_lookup_in_bst",
			},
			{
				Name:   "14.11 Add credits",
				Folder: "adding_credits",
			},
		},
	},
	15: {
		Name:   "Chapter 15: Recursion",
		Number: 15,
		Problems: []problem{
			{
				Name:   "15.00 Bootcamp: Recursion",
				Folder: "euclidean_gcd",
			},
			{
				Name:   "15.01 The Towers of Hanoi problem",
				Folder: "hanoi",
			},
			{
				Name:   "15.02 Compute all mnemonics for a phone number",
				Folder: "phone_number_mnemonic",
			},
			{
				Name:   "15.03 Generate all nonattacking placements of n-Queens",
				Folder: "n_queens",
			},
			{
				Name:   "15.04 Generate permutations",
				Folder: "permutations",
			},
			{
				Name:   "15.05 Generate the power set",
				Folder: "power_set",
			},
			{
				Name:   "15.06 Generate all subsets of size k",
				Folder: "combinations",
			},
			{
				Name:   "15.07 Generate strings of matched parens",
				Folder: "enumerate_balanced_parentheses",
			},
			{
				Name:   "15.08 Generate palindromic decompositions",
				Folder: "enumerate_palindromic_decompositions",
			},
			{
				Name:   "15.09 Generate binary trees",
				Folder: "enumerate_trees",
			},
			{
				Name:   "15.10 Implement a Sudoku solver",
				Folder: "sudoku_solve",
			},
			{
				Name:   "15.11 Compute a Gray code",
				Folder: "gray_code",
			},
		},
	},
	16: {
		Name:   "Chapter 16: Dynamic Programming",
		Number: 16,
		Problems: []problem{
			{
				Name:   "16.00 Bootcamp: Max Sum Subarray",
				Folder: "max_sum_subarray",
			},
			{
				Name:   "16.00 Bootcamp: Fibonacci",
				Folder: "fibonacci",
			},
			{
				Name:   "16.01 Count the number of score combinations",
				Folder: "number_of_score_combinations",
			},
			{
				Name:   "16.02 Compute the Levenshtein distance",
				Folder: "levenshtein_distance",
			},
			{
				Name:   "16.03 Count the number of ways to traverse a 2D array",
				Folder: "number_of_traversals_matrix",
			},
			{
				Name:   "16.04 Compute the binomial coefficients",
				Folder: "binomial_coefficients",
			},
			{
				Name:   "16.05 Search for a sequence in a 2D array",
				Folder: "is_string_in_matrix",
			},
			{
				Name:   "16.06 The knapsack problem",
				Folder: "knapsack",
			},
			{
				Name:   "16.07 Building a search index for domains",
				Folder: "is_string_decomposable_into_words",
			},
			{
				Name:   "16.08 Find the minimum weight path in a triangle",
				Folder: "minimum_weight_path_in_a_triangle",
			},
			{
				Name:   "16.09 Pick up coins for maximum gain",
				Folder: "picking_up_coins",
			},
			{
				Name:   "16.10 Count the number of moves to climb stairs",
				Folder: "number_of_traversals_staircase",
			},
			{
				Name:   "16.11 The pretty printing problem",
				Folder: "pretty_printing",
			},
			{
				Name:   "16.12 Find the longest nondecreasing subsequence",
				Folder: "longest_nondecreasing_subsequence",
			},
		},
	},
	17: {
		Name:   "Chapter 17: Greedy Algorithms and Invariants",
		Number: 17,
		Problems: []problem{
			{
				Name:   "17.00 Bootcamp: Greedy Algorithms and Invariants",
				Folder: "making_change",
			},
			{
				Name:   "17.01 Compute an optimum assignment of tasks",
				Folder: "task_pairing",
			},
			{
				Name:   "17.02 Schedule to minimize waiting time",
				Folder: "minimum_waiting_time",
			},
			{
				Name:   "17.03 The interval covering problem",
				Folder: "minimum_points_covering_intervals",
			},
			{
				Name:   "17.03 Invariant Bootcamp: Two Sum",
				Folder: "two_sum",
			},
			{
				Name:   "17.04 The 3-sum problem",
				Folder: "three_sum",
			},
			{
				Name:   "17.05 Find the majority element",
				Folder: "majority_element",
			},
			{
				Name:   "17.06 The gasup problem",
				Folder: "refueling_schedule",
			},
			{
				Name:   "17.07 Compute the maximum water trapped by a pair of vertical lines",
				Folder: "max_trapped_water",
			},
			{
				Name:   "17.08 Compute the largest rectangle under the skyline",
				Folder: "largest_rectangle_under_skyline",
			},
		},
	},
	18: {
		Name:   "Chapter 18: Graphs",
		Number: 18,
		Problems: []problem{
			{
				Name:   "18.01 Search a maze",
				Folder: "search_maze",
			},
			{
				Name:   "18.02 Paint a Boolean matrix",
				Folder: "matrix_connected_regions",
			},
			{
				Name:   "18.03 Compute enclosed regions",
				Folder: "matrix_enclosed_regions",
			},
			{
				Name:   "18.04 Deadlock detection",
				Folder: "deadlock_detection",
			},
			{
				Name:   "18.05 Clone a graph",
				Folder: "graph_clone",
			},
			{
				Name:   "18.06 Making wired connections",
				Folder: "is_circuit_wirable",
			},
			{
				Name:   "18.07 Transform one string to another",
				Folder: "string_transformability",
			},
			{
				Name:   "18.08 Team photo day---2",
				Folder: "max_teams_in_photograph",
			},
		},
	},
	24: {
		Name:   "Chapter 24: Honors Class",
		Number: 24,
		Problems: []problem{
			{
				Name:   "24.01 Compute the greatest common divisor",
				Folder: "gcd",
			},
			{
				Name:   "24.02 Find the first missing positive entry",
				Folder: "first_missing_positive_entry",
			},
			{
				Name:   "24.03 Buy and sell a stock at most k times",
				Folder: "buy_and_sell_stock_k_times",
			},
			{
				Name:   "24.04 Compute the maximum product of all entries but one",
				Folder: "max_product_all_but_one",
			},
			{
				Name:   "24.05 Compute the longest contiguous increasing subarray",
				Folder: "longest_increasing_subarray",
			},
			{
				Name:   "24.06 Rotate an array",
				Folder: "rotate_array",
			},
			{
				Name:   "24.07 Identify positions attacked by rooks",
				Folder: "rook_attack",
			},
			{
				Name:   "24.08 Justify text",
				Folder: "left_right_justify_text",
			},
			{
				Name:   "24.09 Implement list zipping",
				Folder: "zip_list",
			},
			{
				Name:   "24.10 Copy a postings list",
				Folder: "copy_posting_list",
			},
			{
				Name:   "24.11 Compute the longest substring with matching parens",
				Folder: "longest_substring_with_matching_parentheses",
			},
			{
				Name:   "24.12 Compute the maximum of a sliding window",
				Folder: "max_of_sliding_window",
			},
			{
				Name:   "24.13 Compute fair bonuses",
				Folder: "bonus",
			},
			{
				Name:   "24.14 Search a sorted array of unknown length",
				Folder: "search_unknown_length_array",
			},
			{
				Name:   "24.15 Search in two sorted arrays",
				Folder: "kth_largest_element_in_two_sorted_arrays",
			},
			{
				Name:   "24.16 Find the kth largest element---large n, small k",
				Folder: "kth_largest_element_in_long_array",
			},
			{
				Name:   "24.17 Find an element that appears only once",
				Folder: "element_appearing_once",
			},
			{
				Name:   "24.18 Find the line through the most points",
				Folder: "line_through_most_points",
			},
			{
				Name:   "24.19 Convert a sorted doubly linked list into a BST",
				Folder: "sorted_list_to_bst",
			},
			{
				Name:   "24.20 Convert a BST to a sorted doubly linked list",
				Folder: "bst_to_sorted_list",
			},
			{
				Name:   "24.21 Merge two BSTs",
				Folder: "bst_merge",
			},
			{
				Name:   "24.22 Implement regular expression matching",
				Folder: "regular_expression",
			},
			{
				Name:   "24.23 Synthesize an expression",
				Folder: "insert_operators_in_string",
			},
			{
				Name:   "24.24 Count inversions",
				Folder: "count_inversions",
			},
			{
				Name:   "24.25 Draw the skyline",
				Folder: "drawing_skyline",
			},
			{
				Name:   "24.26 Measure with defective jugs",
				Folder: "defective_jugs",
			},
			{
				Name:   "24.27 Compute the maximum subarray sum in a circular array",
				Folder: "maximum_subarray_in_circular_array",
			},
			{
				Name:   "24.28 Determine the critical height",
				Folder: "max_safe_height",
			},
			{
				Name:   "24.29 Max Square Submatrix",
				Folder: "max_square_submatrix",
			},
			{
				Name:   "24.29 Max Submatrix",
				Folder: "max_submatrix",
			},
			{
				Name:   "24.30 Implement Huffman coding",
				Folder: "huffman_coding",
			},
			{
				Name:   "24.31 Trapping water",
				Folder: "max_water_trappable",
			},
			{
				Name:   "24.32 The heavy hitter problem",
				Folder: "search_frequent_items",
			},
			{
				Name:   "24.33 Find the longest subarray with sum constraint",
				Folder: "longest_subarray_with_sum_constraint",
			},
			{
				Name:   "24.34 Road network",
				Folder: "road_network",
			},
			{
				Name:   "24.35 Test if arbitrage is possible",
				Folder: "arbitrage",
			},
		},
	},
}

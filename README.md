# Interview Questions

Basic algorithmic interview questions, answered in Go.
For the majority of cases, where an array is requested, I have
made the assumption that we can guarantee that the array is populated
or that the specification of the question are always true, e.g. when 
finding a repeated character in a string, the string is always populated
with characters and there is always a repeated character. This avoids noisy
defensive programming that distracts from the solution.

From https://simpleprogrammer.com/programming-interview-questions/

Array Questions:
1. [How do you find the missing number in a given integer array of 1 to 100?](src/find_missing_integer.go)
1. [How do you find the duplicate number on a given integer array?](src/array_duplicate_number.go)
1. [How do you find the largest and smallest number in an unsorted integer array?](src/min_max_numbers.go)
1. [How do you find all pairs of an integer array whose sum is equal to a given number?](src/number_pair_sum.go)
1. [How do you find duplicate numbers in an array if it contains multiple duplicates?](src/array_duplicate_numbers.go)
1. [How are duplicates removed from a given array?](src/array_remove_duplicates.go)
1. How is an integer array sorted in place using the quicksort algorithm?
1. How do you remove duplicates from an array in place?
1. How do you reverse an array in place in Java?

Linked List Questions:
1. [How do you find the middle element of a singly linked list in one pass?](src/linked_list_middle_element.go)
1. [How do you check if a given linked list contains a cycle? How do you find the starting node of the cycle?](src/linked_list_cycle.go)
1. [How do you reverse a linked list?](src/linked_list_reverse.go)
1. How do you reverse a singly linked list without recursion?
1. How are duplicate nodes removed in an unsorted linked list?
1. How do you find the length of a singly linked list?
1. How do you find the third node from the end in a singly linked list?
1. How do you find the sum of two linked lists using Stack?

String Questions:
1. [How do you print duplicate characters from a string?](src/duplicate_chars.go)
1. [How do you check if two strings are anagrams of each other?](src/string_anagrams.go)
1. [How do you print the first non-repeated character from a string?](src/first_non_repeated_char.go)
1. [How can a given string be reversed using recursion?](src/reverse_string.go)
1. How do you check if a string contains only digits?
1. How are duplicate characters found in a string?
1. How do you count a number of vowels and consonants in a given string?
1. How do you count the occurrence of a given character in a string?
1. How do you find all permutations of a string?
1. How do you reverse words in a given sentence without using any library method?
1. How do you check if two strings are a rotation of each other?
1. How do you check if a given string is a palindrome?

Binary Tree Questions:
1. [How is a binary search tree implemented?](src/binarsearchtree.go)
1. [How do you perform preorder traversal in a given binary tree?](src/bst_preorder_traversal.go)
1. [How do you perform an inorder traversal in a given binary tree?](src/bst_inorder_traversal.go)
1. [How do you implement a postorder traversal algorithm?](src/bst_postorder_traversal.go)
1. [How do you count a number of leaf nodes in a given binary tree?](src/bst_leaf_node_count.go)
1. [How do you perform a binary search in a given binary tree?](src/bst_search.go)

Misc Questions:
1. How is a bubble sort algorithm implemented?
1. How is an iterative quicksort algorithm implemented?
1. How do you implement an insertion sort algorithm?
1. How is a merge sort algorithm implemented?
1. How do you implement a bucket sort algorithm?
1. How do you implement a counting sort algorithm?
1. How is a radix sort algorithm implemented?
1. How do you swap two numbers without using the third variable?
1. How do you check if two rectangles overlap with each other?
1. How do you design a vending machine?
package main

import "fmt"

// Node represents a node of the linked list
type Node struct {
	value int
	next  *Node
}

// Function to swap alternate nodes of the linked list
func swapAlternateNodes(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	// Initialize pointers
	prev := head
	curr := head.next

	// Swap the first pair
	head = curr

	for {
		// Swap the nodes
		next := curr.next
		curr.next = prev

		// Check if next pair exists
		if next == nil || next.next == nil {
			prev.next = next
			break
		}

		// Point the previous node to the next of the pair to swap
		prev.next = next.next

		// Move the pointers two nodes ahead
		prev = next
		curr = prev.next
	}

	return head
}

// Helper function to print the linked list
func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

// Function to insert a new node after a given value in the linked list
func insertAfter(head *Node, afterValue int, newValue int) *Node {
	// Create the new node
	newNode := &Node{value: newValue}

	// Traverse the list to find the node after which to insert
	current := head
	for current != nil && current.value != afterValue {
		current = current.next
	}

	// If the node was found, insert the new node after it
	if current != nil {
		newNode.next = current.next
		current.next = newNode
	}

	return head
}

// Function to delete a node with a specific value in the linked list
func deleteNode(head *Node, value int) *Node {
	// If the list is empty, return nil
	if head == nil {
		return nil
	}

	// If the head node itself holds the value to be deleted
	if head.value == value {
		return head.next
	}

	// Traverse the list to find the node to delete
	current := head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	// If the node was found, adjust pointers to delete it
	if current.next != nil {
		current.next = current.next.next
	}

	return head
}

// Helper function to create a linked list from a slice of values
func createList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{value: values[0]}
	current := head

	for _, value := range values[1:] {
		current.next = &Node{value: value}
		current = current.next
	}

	return head
}

func main() {
	// Create a linked list
	values := []int{1, 2, 3, 4, 5, 6}
	head := createList(values)

	fmt.Println("Original list:")
	printList(head)

	// Swap alternate nodes
	head = swapAlternateNodes(head)

	fmt.Println("List after swapping alternate nodes:")
	printList(head)
}

package avl_test

import (
	"testing"

	"github.com/jrwoodruff1000/avl"
)

type test_int_key struct {
	key     int
	payload string
}

type test_large_int_key struct {
	key     int
	payload int
}

type test_string_key struct {
	key     string
	payload int
}

var int_tests = []test_int_key{
	{1, "one"},
	{2, "two"},
	{3, "three"},
}

var string_tests = []test_string_key{
	{"one", 1},
	{"two", 2},
	{"three", 3},
}

var large_test = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49}

func TestNewAvl(t *testing.T) {
	// Also tests Get_metadata()
	myAVL := avl.NewAvl[int]("Test AVL")
	nm := myAVL.Get_metadata()
	if nm != "Test AVL" {
		t.Error("Expected 'Test AVL', got ", nm)
	}
}

func Test_Add_node_int(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Add Int Node")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node ", pair.key, ", err = ", err)
		}
	}

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err == nil {
			t.Error("Expected error when creating int node but didn't get one")
		}
	}
}

func Test_Add_node_string(t *testing.T) {
	myAVL := avl.NewAvl[string]("Test Add String Node")

	for _, pair := range string_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node ", pair.key, ", err = ", err)
		}
	}

	for _, pair := range string_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err == nil {
			t.Error("Expected error when creating string node but didn't get one")
		}
	}
}

func Test_Count_Nodes(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Counting Nodes")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for count", pair.key, ", err = ", err)
		}
	}

	expected_node_count := 3
	myCount := myAVL.Count_nodes()
	if expected_node_count != myCount {
		t.Error("Error counting nodes, expected: ", expected_node_count, ", received: ", myCount)
	}
}

func Test_Delete_node(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Delete Node")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for delete", pair.key, ", err = ", err)
		}
	}

	node_to_delete := 2
	err := myAVL.Delete_node(node_to_delete)
	if err != nil {
		t.Error("Error deleting node ", node_to_delete, ", err = ", err)
	}

	node_to_delete = 10
	err = myAVL.Delete_node(node_to_delete)
	if err == nil {
		t.Error("Expected error when deleting node but didn't get one")
	}
}

func Test_Get_Payload(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Getting Payload")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node ", pair.key, ", err = ", err)
		}
	}

	test_node := 2
	expected_payload := "two"
	test_payload, err := myAVL.Get_payload(test_node)
	if (expected_payload != test_payload) || (err != nil) {
		t.Error("Error retrieving payload for node: ", test_node,
			", expected: ", expected_payload, ", received: ", test_payload, ", err = ", err)
	}

	test_node = 20
	test_payload, err = myAVL.Get_payload(test_node)
	if err == nil {
		t.Error("Expected error when getting payload but didn't get one")
	}
}

func Test_Update_Payload(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Updating Payload")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for payload update", pair.key, ", err = ", err)
		}
	}

	test_node := 1
	expected_updated_payload := "new payload"
	err := myAVL.Update_payload(test_node, expected_updated_payload)
	if err != nil {
		t.Error("Error updating payload for node: ", test_node, ", err = ", err)
	}

	test_payload, err := myAVL.Get_payload(test_node)
	if (expected_updated_payload != test_payload) || (err != nil) {
		t.Error("Error updating payload for node: ", test_node, ", err = ", err)
	}

	test_node = 20
	err = myAVL.Update_payload(test_node, expected_updated_payload)
	if err == nil {
		t.Error("Expected error when updating payload but didn't get one")
	}
}

func Test_Rotations(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test Rotations")

	// Add (50) nodes
	for _, value := range large_test {
		err := myAVL.Add_node(value, value)
		if err != nil {
			t.Error("Error creating node ", value, ", err = ", err)
		}
	}

	expected_node_count := 50
	myCount := myAVL.Count_nodes()
	if expected_node_count != myCount {
		t.Error("Error counting nodes in rotation test, expected: ", expected_node_count, ", received: ", myCount)
	}

	// Delete (50) nodes
	for _, value := range large_test {
		err := myAVL.Delete_node(value)
		if err != nil {
			t.Error("Error deleting node ", value, ", err = ", err)
		}
	}

	expected_node_count = 0
	myCount = myAVL.Count_nodes()
	if expected_node_count != myCount {
		t.Error("Error counting nodes in rotation test, expected: ", expected_node_count, ", received: ", myCount)
	}
}

func Test_Find_Min_Node(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test to Find Minimum Node Key")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for Find Min", pair.key, ", err = ", err)
		}
	}

	expected_min_key := 1
	minKey, err := myAVL.Min_node()
	if err != nil {
		t.Error("Error when trying to find minimum key")
	} else if minKey != expected_min_key {
		t.Error("Error finding minimum key, expected: ", expected_min_key, ", received: ", minKey)
	}
}

func Test_Find_Max_Node(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test to Find Maximum Node Key")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for Find Min", pair.key, ", err = ", err)
		}
	}

	expected_min_key := 3
	maxKey, err := myAVL.Max_node()
	if err != nil {
		t.Error("Error when trying to find maximum key")
	} else if maxKey != expected_min_key {
		t.Error("Error finding maximum key, expected: ", expected_min_key, ", received: ", maxKey)
	}
}

func Test_Next_Node(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test to Find Next Node")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for Find Next", pair.key, ", err = ", err)
		}
	}

	expected_next_key := 3
	this_key := 2
	nextKey, err := myAVL.Next_node(this_key)
	if err != nil {
		t.Error("Error when trying to find maximum key")
	} else if nextKey != expected_next_key {
		t.Error("Error finding next node, expected: ", expected_next_key, ", received: ", nextKey)
	}
}

func Test_Preivous_Node(t *testing.T) {
	myAVL := avl.NewAvl[int]("Test to Find Previous Node")

	for _, pair := range int_tests {
		err := myAVL.Add_node(pair.key, pair.payload)
		if err != nil {
			t.Error("Error creating node for Find Previous", pair.key, ", err = ", err)
		}
	}

	expected_previous_key := 1
	this_key := 2
	previousKey, err := myAVL.Previous_node(this_key)
	if err != nil {
		t.Error("Error when trying to find maximum key")
	} else if previousKey != expected_previous_key {
		t.Error("Error finding next node, expected: ", expected_previous_key, ", received: ", previousKey)
	}
}

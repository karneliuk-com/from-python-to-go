/* From Python to Go: Python: 004 - lists */

package main

// Import
import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// Aux functions
func getAutomationInput() string {
	/* Helper function to get the automation input */
	return os.Getenv("AUTOMATION_INPUT")
}

// Main functions
func main() {
	// Get the input
	automationInput := getAutomationInput()
	fmt.Println(automationInput)

	// Create a slice
	automationSlice := strings.Split(automationInput, ",")
	fmt.Println(automationSlice)

	// Add new element to the end of the slice
	automationSlice = append(automationSlice, "new_device")
	fmt.Println(automationSlice)

	// Add new element to the beginning of the slice
	automationSlice = slices.Insert(automationSlice, 0, "provisioning_required")
	fmt.Println(automationSlice)

	// Check if the element exist in the slice
	if slices.Contains(automationSlice, "provisioning_required") {
		fmt.Printf("provisioning_required is for the device %v\n", automationSlice[0])
	}

	// Remove elemement from the slice and by index
	deleteIndex := slices.Index(automationSlice, "new_device")
	automationSlice = append(automationSlice[:deleteIndex], automationSlice[deleteIndex+1:]...)
	fmt.Println(automationSlice)

	// Change element
	automationSlice[0] = "provisioning_done"
	fmt.Println(automationSlice)

	// Sort the slice
	slices.Sort(automationSlice)
	fmt.Println(automationSlice)

	// Reverse the slice
	slices.Reverse(automationSlice)
	fmt.Println(automationSlice)

	// Merge list into the slice
	automationString := strings.Join(automationSlice, ",")
	fmt.Println(automationString)
}

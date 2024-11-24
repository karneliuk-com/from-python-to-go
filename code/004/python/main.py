"""From Python to Go: Python: 004 - lists and slices"""

# Import part
import os


# Functions
def get_automation_input() -> str:
    """Helper function to get the automation input"""
    return os.getenv("AUTOMATION_INPUT")


# Read environmental value
automation_input = get_automation_input()
print(f"{automation_input=}")

# Create a list from string
automation_list = automation_input.split(",")
print(f"{automation_list=}")

# Add an element to the end of the list
automation_list.append("new_device")
print(f"{automation_list=}")

# Add an element to the beginning of the list
automation_list.insert(0, "provisioning_required")

# Check if that element is in the list
if "provisioning_required" in automation_list:
    print(f"provisioning_required is for the device {automation_list[1]}")

# Remove an element from the list
automation_list.remove("new_device")
print(f"{automation_list=}")

# Change the element in the list
automation_list[0] = "provisioning_done"
print(f"{automation_list=}")

# Remove the element by index
del automation_list[0]
print(f"{automation_list=}")

# Sort the list
automation_list.sort()
print(f"{automation_list=}")

# Reverse the list
automation_list.reverse()
print(f"{automation_list=}")

# Merger list into a string
automation_string = ",".join(automation_list)
print(f"{automation_string=}")

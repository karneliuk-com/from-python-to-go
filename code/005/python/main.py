"""From Python to Go: Python: 005 - Code Flow Control: Conditionals and Loops"""

# Import

# Variables
# Your initial inventory
inventory = [
    ["leaf-01", "cisco-nxos", "192.168.1.1"],
    ["leaf-02", "arista-eos", "192.168.1.2"],
    ["spine-01", "cisco-nxos", ""],
    ["spine-02", "arosta-eos", "192.168.1.12"],
]


# Functions
def get_data(ip_address: str) -> list:
    """Function that pretends to connect to network device and collect some output"""
    # Ensure there is IP address to connect
    if not ip_address:
        return (False, "There is no IP address provided")

    # Return some mock data
    return (True, "Some raw data")


def parse_data(parser: str, data: str) -> list:
    """Function that pretends to parse the output from network device depeinding on its operating system type"""
    match parser:
        case "cisco-nxos":
            return (True, f"parsed: {data}")

        case "arista-eos":
            return (True, f"parsed: {data}")

        case _:
            return (False, "There is no parser available")


# Execution
if __name__ == "__main__":
    # Loop through all network devices
    for device in inventory:
        # Collect data for each network device
        collected_data = get_data(device[2])

        parsed_data = []
        # Do parsing if data is collected
        if collected_data[0]:
            parsed_data = parse_data(device[1], collected_data[1])
        else:
            print(f"Collecting data from {device[0]} is not successful")
            continue

        # Print results
        if parsed_data[0]:
            print(f"Successfully collected and parsed data for {device[0]}")
        else:
            print(f"Successfully collected but NOT parsed data for {device[0]}")

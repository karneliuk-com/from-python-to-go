"""From Python to Go: Python: 009 - Interfaces data type"""

# Variables
inventory = [
    {
        "name":      "leaf-01",
        "os":        "cisco-nxos",
        "ip":        "192.168.1.1",
        "port":      22,
        "latitude":  51.5120898,
        "longitude": -0.0030987,
        "active":    True,
    }, {
        "name":      "leaf-02",
        "os":        "arista-eos",
        "ip":        "192.168.1.2",
        "port":      830,
        "latitude":  51.5120427,
        "longitude": -0.0044585,
        "active":    True,
    }, {
        "name":      "spine-01",
        "ip":        "192.168.1.1",
        "port":      22,
        "latitude":  51.5112179,
        "longitude": -0.0048555,
        "active":    False,
    },
]


# Execution
if __name__ == "__main__":
    # Print the entire map and data type
    print(inventory, type(inventory))

    # Print data types for each element
    for item in inventory:
        print(item, type(item))
        for key, value in item.items():
            print(key, value, type(value))

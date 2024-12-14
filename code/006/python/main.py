"""From Python to Go: Python: 006 - Dictionaries and Maps"""

# Variables
inventory = [
    {
        "name": "leaf-01",
        "os": "cisco-nxos",
        "ip": "192.168.1.1",
    },
    {
        "name": "leaf-02",
        "os": "arista-eos",
        "ip": "192.168.1.2",
    },
    {
        "name": "spine-01",
        "ip": "192.168.1.1",
    },
]


# Execution
if __name__ == "__main__":
    # Loop through all network devices
    for d in inventory:
        # Print the device data
        print(d)

        # Print the hostname
        print(f"Hostname: {d['name']}")

        # Add the OS key if it is missing
        if "os" not in d:
            d["os"] = None
        print(d)

        # Add new key-value pair
        d["location"] = "DC1"
        print(d)

        # Remove the IP key
        d.pop("ip")
        print(d)

        # Go through all keys and values
        for k, v in d.items():
            print(f"{k}: {v}")

    # Print the inventory after modifications
    print(inventory)

"""From Python to Go: Python: 008 - Object-oriented programming"""

# Import os
import os
from  typing import List


# Data models
class User:
    """Class to store user credentials"""
    def __init__(self, username: str = None, password: str = None):
        self.username: str = username
        self.password: str = password

    def get_credentials(self):
        """Function to retrieve credentials from the environment"""
        self.username, self.password = os.getenv("AUTOMATION_CREDS").split(",")


class Device:
    """Class to store device information"""
    def __init__(self, name: str, port: int, nos: str = None, ip: str = None):
        self.name: str = name
        self.port: int = port
        self.nos = nos
        self.ip = ip


class Inventory:
    """Class to store inventory information"""
    def __init__(self):
        self.devices: List[Device] = []

    def populate(self):
        """Function to retrieve inventory from the environment"""
        # Loop through the environment variables
        for key, value in os.environ.items():
            # Check if the key starts with AUTOMATION_DEVICE_
            if key.startswith('AUTOMATION_DEVICE_'):
                # Split the value by comma and create a new device object
                split_value = value.split(',')
                self.devices.append(
                    Device(
                        name=split_value[0],
                        port=int(split_value[1]),
                        nos=split_value[3],
                        ip=split_value[2],
                    )
                )


# Execution
if __name__ == "__main__":
    # Get the credentials
    user = User()
    user.get_credentials()

    # Print the credentials
    print(f"Username: {user.username}")
    print(f"Password: {user.password}")

    # Get the inventory
    inventory = Inventory()
    inventory.populate()

    # Print inventory memory address
    print(f"Memory address of inventory: {id(inventory):#x}")

    # Print the inventory
    print("Inventory:")
    for device in inventory.devices:
        print(f"Device: {device.name}")
        print(f"Port:   {device.port}")
        print(f"IP:     {device.ip}")
        print(f"NOS:    {device.nos}")
        print("\n")

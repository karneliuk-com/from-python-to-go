"""From Python to Go: Python: 007 - Classes and Structs. Using extenral modules"""

# Import os
import os
from  typing import List
from pydantic import BaseModel


# Data models
class User(BaseModel):
    """Class to store user credentials"""
    username: str
    password: str


class Device (BaseModel):
    """Class to store device information"""
    name: str
    port: int
    nos: str = None
    ip: str = None


class Inventory(BaseModel):
    """Class to store inventory information"""
    devices: List[Device]


# Functions
def get_credentials() -> User:
    """Function to retrieve credentials from the environment"""
    blocks = os.getenv("AUTOMATION_CREDS").split(",")
    return User(username=blocks[0], password=blocks[1])


def get_inventory() -> Inventory:
    """Function to retrieve inventory from the environment"""
    # Create an empty list to store devices
    result = Inventory(devices=[])

    # Loop through the environment variables
    for key, value in os.environ.items():
        # Check if the key starts with AUTOMATION_DEVICE_
        if key.startswith('AUTOMATION_DEVICE_'):
            # Split the value by comma and create a new device object
            split_value = value.split(',')
            result.devices.append(
                Device(
                    name=split_value[0],
                    port=int(split_value[1]),
                    nos=split_value[3],
                    ip=split_value[2],
                )
            )

    # Return the result
    return result


# Execution
if __name__ == "__main__":
    # Get the credentials
    user = get_credentials()

    # Print the credentials
    print(f"Username: {user.username}")
    print(f"Password: {user.password}")

    # Get the inventory
    inventory = get_inventory()

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

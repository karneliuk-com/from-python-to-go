"""From Python to Go: Python: 011 - Parsing XML/JSON/YAML files """

# Import
from dataclasses import dataclass, field
from typing import Union, List
import sys
import json
import xmltodict
import yaml


# Dataclass
@dataclass(slots=True)
class InventoryItem:
    """Device -- Inventory Item"""
    name: str
    os: Union[str, None] = None
    ip: Union[str, None] = None
    port: Union[int, None] = None
    latitude: Union[float, None] = None
    longitude: Union[float, None] = None
    active: Union[bool, None] = None


@dataclass(slots=True)
class Inventory:
    """Inventory of all devices"""
    devices: List[InventoryItem] = field(default_factory=list)


# Auxiliary functions
def load_inventory(file: str) -> Inventory:
    """Function to load inventory"""

    # Initialize result
    result = Inventory()

    # Load file
    temp_dict = {}
    with open(file, 'r', encoding="utf-8") as f:
        if file.endswith('.json'):
            temp_dict = json.load(f)
        elif file.endswith('.xml'):
            temp_dict = xmltodict.parse(f.read())["root"]
        elif file.endswith('.yaml') or file.endswith('.yml'):
            temp_dict = yaml.safe_load(f)
        else:
            raise ValueError('Unsupported file format')

    # Populate result
    for item in temp_dict['devices']:
        result.devices.append(InventoryItem(
            name=item['name'],
            os=item.get('os'),
            ip=item.get('ip'),
            port=item.get('port'),
            latitude=item.get('latitude'),
            longitude=item.get('longitude'),
            active=item.get('active')
        ))

    return result


# Main
if __name__ == "__main__":
    # Check that file is provided
    if len(sys.argv) != 2:
        print('Usage: python main.py <file>')
        sys.exit(1)

    # Load inventory
    inventory = load_inventory(sys.argv[1])

    # Print inventory
    print(inventory)

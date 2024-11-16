"""From Python to Go: Python: 003 - imports and function"""

# Import part
import os


# Functions
def get_username() -> str:
    """Helper function to get username"""
    return os.getenv("AUTOMATION_USERNAME", None)


def get_password() -> str:
    """Helper function to get password"""
    return os.getenv("AUTOMATION_PASSWORD", None)


# Execution
username = get_username()
password = get_password()

print(f"{username=}, {password=}")

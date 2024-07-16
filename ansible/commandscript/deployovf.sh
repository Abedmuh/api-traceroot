#!/bin/bash

# Function to update the credential.yml file with new values
update_credential_file() {
    local key=$1
    local value=$2
    sed -i "s#^\($key: \).*#\1$value#g" vmware-credentials.yml
}

# Check if the correct number of arguments is provided
if [ "$#" -ne 4 ]; then
    echo "Usage: $0 <esxi_name> <guest_hardware_esxi_num_cpu> <guest_hardware_esxi_memory_mb> <guest_hardware_esxi_storage>"
    exit 1
fi

# Update the credential.yml file with the provided arguments
update_credential_file "esxi_name" "$1"
update_credential_file "guest_hardware_esxi_num_cpu" "$2"
update_credential_file "guest_hardware_esxi_memory_mb" "$3"
update_credential_file "guest_hardware_esxi_storage" "$4"

# Define the inventory file and playbook file
INVENTORY_FILE="inventory.ini"
PLAYBOOK_FILE="deploy-ovf-vmdk.yml"

# Run the ansible-playbook command
ansible-playbook -i "$INVENTORY_FILE" "$PLAYBOOK_FILE"

# Capture the exit code
EXIT_CODE=$?

# Exit with the same code as the ansible-playbook command
exit $EXIT_CODE

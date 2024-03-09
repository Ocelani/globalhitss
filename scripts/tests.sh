#!/bin/bash
# This script is used to manual test the API

base_url="http://localhost:3000"
endpoint_user="/user"
endpoint_queue="/queue/user"

# Create user
echo "Create user 1"
curl -i -X POST "${base_url}${endpoint_user}" -d '{"nome":"Fulano"}'

# Get user
echo "Get user 1"
curl -i -X GET "${base_url}${endpoint_user}/1"

# Create user on queue
echo "Create user 2 on queue"
curl -i -X POST "${base_url}${endpoint_queue}" -d '{"nome":"Beltrano"}'

# Get user created by queue
echo "Get user 2 created by queue"
curl -i -X GET "${base_url}${endpoint_user}/2"

# Update user
echo "Update user 1"
curl -i -X PUT "${base_url}${endpoint_user}/1" -d '{"nome":"Ciclano"}'

# Delete user
echo "Delete user 1"
curl -i -X DELETE "${base_url}${endpoint_user}/1"
#!/bin/bash

gnome-terminal -- bash -c "swag init; go run main.go; exec bash"
gnome-terminal -- bash -c "cd frontend/ && npm run dev; exec bash"
#!/bin/bash

GREEN="\033[0;32m"
RED="\033[0;31m"
NC="\033[0m"

echo -e "${GREEN}🔍 Checking Docker daemon...${NC}"

if ! systemctl is-active --quiet docker; then
    echo -e "${RED}🚫 Docker is not running. Starting Docker...${NC}"
    sudo systemctl start docker

    if ! systemctl is-active --quiet docker; then
        echo -e "${RED}❌ Failed to start Docker. Please check Docker installation.${NC}"
        exit 1
    else
        echo -e "${GREEN}✅ Docker started successfully.${NC}"
    fi
else
    echo -e "${GREEN}✅ Docker is already running.${NC}"
fi

echo -e "${GREEN}🚀 Starting containers with Docker Compose...${NC}"

docker-compose up --build

#!/bin/bash
set -e

echo "starting ci-server workflow ...." 

REPO_DIR="$HOME/cicd-pipeline"
REPO_URL="https://github.com/btwkevin/cicd-pipeline.git"

if [ ! -d "$REPO_DIR/.git" ]; then
    echo "Directory not found. Cloning repo..."
    mkdir -p "$REPO_DIR"
    git clone "$REPO_URL" "$REPO_DIR"
else
    echo "Directory exists. Pulling latest changes..."
    cd "$REPO_DIR"
    git reset --hard
    git pull origin main || git pull
fi

cd "$REPO_DIR"
docker compose up -d

echo "complete..."
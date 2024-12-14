#!/bin/bash

# Mendapatkan timestamp (format: YYYY-MM-DD HH:MM:SS)
timestamp=$(date +"%Y-%m-%d %H:%M:%S")

# Menambahkan perubahan ke staging area
git add .

# Commit dengan pesan timestamp otomatis
git commit -m "Auto commit at $timestamp"

# Push perubahan ke branch master
git push origin master

echo "Commit and push successful!"

#!/bin/bash

echo "🚀 Initializing Git repository..."
git init
git add .
git commit -m "chore: initial commit with v0.1.0"
git branch -M main

echo "✅ Now run:"
echo "git remote add origin git@github.com:yourusername/human-presence-sim.git"
echo "git push -u origin main"
echo "git tag v0.1.0 -m 'Initial release'"
echo "git push origin v0.1.0"

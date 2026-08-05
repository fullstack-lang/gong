#!/bin/zsh

# Configuration
URL="http://localhost:8080"
INTERVAL=60

echo "🚀 Starting Chrome Restarter loop..."
echo "Target URL: $URL"
echo "Restart Interval: $INTERVAL seconds"
echo "Press [CTRL+C] to stop."
echo "----------------------------------------"

while true; do
    TIMESTAMP=$(date '+%H:%M:%S')
    echo "[$TIMESTAMP] Restarting Google Chrome..."

    # Gracefully request Chrome to quit via AppleScript
    osascript -e 'quit app "Google Chrome"' 2>/dev/null

    # Brief pause to allow processes to close cleanly
    sleep 2

    # Force kill if Chrome didn't close (e.g., dialog popups)
    killall "Google Chrome" 2>/dev/null

    # Launch Chrome to target URL
    open -a "Google Chrome" "$URL"

    # Wait 1 minute before the next restart cycle
    sleep "$INTERVAL"
done

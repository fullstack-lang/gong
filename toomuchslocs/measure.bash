#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <start> <end> <step>"
    exit 1
fi

n=$2
output_file="size_measurements.txt"
mkdir -p go/models

for ((i=$1; i<=$n; i=i+$3)); do
    # Generate Go file with i structs
    cat > go/models/model.go << EOF
package models

$(for ((j=1; j<=$i; j++)); do
echo "type Model$j struct {
    Name string

    Floatfield, Floatfield2 float64

	Intfield int

    Model$j *Model$j
    Model$js []*Model$j
 }"
done)
EOF

    # Run gongc command
    start_time=$(date +%s.%N)
    gongc -skipGoModCommands -skipNg go/models
    end_time=$(date +%s.%N)
    execution_time=$(awk -v start="$start_time" -v end="$end_time" 'BEGIN{printf "%.6f\n", end-start}') 

    # Measure file size
    size=$(du -b ./go/cmd/toomuchslocs | cut -f1)
    
    # Log size
    echo "$i $size $execution_time" >> $output_file
done
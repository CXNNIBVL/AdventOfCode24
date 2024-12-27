#! /bin/bash

# PART 1

FILE="input"

SEQ="XMAS"
RSEQ="SAMX"

numRows=$(cat "$FILE" | wc -l)
numRows=$((numRows))

numCols=$(head -n 1 "$FILE" | wc -m)
numCols=$((numCols))
numCols=$(($numCols-1)) # Tailing newline

found=0

# Find Horizontal occurrences
horizontal=$(rg -c -e "$SEQ" "$FILE")
horizontalRev=$(rg -c -e "$RSEQ" "$FILE")
found=$((found+horizontal+horizontalRev))

# Find vertical occurrences
for((i=1; i <= $numCols; i++)) do
    asRow=()
    asRow=($(cat "$FILE" | cut -b "$i"))
    # printf "%s" "${asRow[@]}" 

    vert=$(printf "%s" "${asRow[@]}" | rg -c -e "$SEQ")
    vertRev=$(printf "%s" "${asRow[@]}" | rg -c -e "$RSEQ")

    found=$((found+vert+vertRev))
done

# Find diagonal occurrences
occ=$(go run ./p1/main.go "$FILE" "$SEQ")
found=$((found+occ))
echo "Part 1: Found $found matches"
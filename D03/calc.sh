#! /bin/bash

rg -o -e "mul\(\d+,\d+\)" input | \
awk -F "mul" '{ print $2 }' | \
tr -d '(' | tr -d ')' | \
awk -F "," '{$1=$1*$2; print $1}' | \
awk '{s+=$1} END {print "Part1: " s}'

instructions=($(rg -o -e "mul\(\d+,\d+\)" -e "do\(\)" -e "don't\(\)" input |\
tr -d '(' | tr -d ')'))

filtered=()
use=true

for instruction in "${instructions[@]}"; do
    bfcount+=1
    if [[ $instruction == "do" ]]; then
        use=true
        continue
    fi

    if [[ $instruction == "don't" ]]; then
        use=false
        continue
    fi

    if $use; then
        filtered+=("$instruction")
    fi
done

printf "%s\n" "${filtered[@]}" | \
awk -F "mul" '{print $2}' | \
awk -F "," '{$1=$1*$2; print $1}' | \
awk '{s+=$1} END {print "Part2: " s}'
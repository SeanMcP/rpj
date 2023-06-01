#!/usr/bin/env bash

# create a list of files to iterate
files=(
  "rpj.bun"
  "rpj.deno"
  "rpj"
)

output="GENERATED $(date +%F_%T)\n\nExecutable file sizes:\n"

# iterate over the list of files
for file in "${files[@]}"; do
  output="$output\n$(wc -c $file)"
done

echo -e $output > sizes.txt

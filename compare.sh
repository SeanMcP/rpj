#!/usr/bin/env bash

# create a list of files to iterate
files=(
  "rpj.bun"
  "rpj.deno"
  "rpj"
)

rm -f sizes.txt

echo -e "GENERATED $(date +%F_%T)\n\nExecutable file sizes:\n" >>sizes.txt

# iterate over the list of files
for file in "${files[@]}"; do
  echo -e "$(wc -c $file)\n" >>sizes.txt
done

rm -f times.txt

echo -e "GENERATED $(date +%F_%T)\n\nExecution times:\n" >>times.txt

for file in "${files[@]}"; do
  /usr/bin/time -f "%C\n- %E real\n- %U user\n- %S sys\n" -a -o times.txt ./$file scripts &>/dev/null
done

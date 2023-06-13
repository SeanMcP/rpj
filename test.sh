#!/usr/bin/env bash

files=(
  "rpj.bun"
  "rpj.deno"
  "rpj"
)

dirs=(
  "valid"
  "invalid"
  "missing"
)

cd test/

for dir in "${dirs[@]}"; do
  cd "$dir/"
  echo -e "$dir/"

  for file in "${files[@]}"; do
    echo -e "  $file: $(../../$file n)"
  done

  cd ..
done

cd ..

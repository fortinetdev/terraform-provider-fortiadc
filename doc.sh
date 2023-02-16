#!/bin/bash
data_dir=./website/docs/d/
resource_dir=./website/docs/r/
for entry in "$resource_dir"/*
do
	filename=$(basename $entry)
	html="${filename%.*}"
	show="${html%.*}"
  echo '<li>'
  echo "<a href=\"/docs/providers/fortiadc/r/$html\">$show</a>"
  echo '</li>'
done
echo ===============================data source======================================
for entry in "$data_dir"/*
do
	filename=$(basename $entry)
	html="${filename%.*}"
	show="${html%.*}"
  echo '<li>'
  echo "<a href=\"/docs/providers/fortiadc/d/$html\">$show</a>"
  echo '</li>'
done

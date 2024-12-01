#!/bin/bash

BASE_DIR="./component/content"
SOURCE_LOCALE="en"
SOURCE_DIR="$BASE_DIR/$SOURCE_LOCALE"

# Function to check if a file is already translated in any locale
is_translated() {
    local src_file="$1"
    local relative_path="${src_file#$SOURCE_DIR/}"  # Strip SOURCE_DIR prefix

    # Check all locale directories in markdown except "en"
    for locale_dir in "$BASE_DIR"/*; do
        if [ "$locale_dir" != "$SOURCE_DIR" ] && [ -d "$locale_dir" ]; then
            if [ -f "$locale_dir/$relative_path" ]; then
                return 0  # File is already translated
            fi
        fi
    done
    return 1  # File is not translated
}

# Iterate over all .md files in the source directory
find "$SOURCE_DIR" -name '*.md' | while read -r src_file; do
    relative_path="${src_file#$SOURCE_DIR/}"  # Strip SOURCE_DIR prefix

    # Check if the file is already translated
    if is_translated "$src_file"; then
        echo "Skipping: $relative_path already translated."
        continue
    fi

    # Translate the file for each missing locale
    for locale_dir in "$BASE_DIR"/*; do
        if [ "$locale_dir" != "$SOURCE_DIR" ] && [ -d "$locale_dir" ]; then
            target_file="$locale_dir/$relative_path"
            
            # Skip if this specific locale translation exists
            if [ -f "$target_file" ]; then
                continue
            fi
            
            # Ensure the target directory exists
            mkdir -p "$(dirname "$target_file")"
            
            # Extract the target locale from the folder name
            target_locale=$(basename "$locale_dir")

            echo "Translating: $src_file -> $target_file (Locale: $target_locale)"
            trans :$target_locale -b -i "$src_file" -o "$target_file"

            if [ $? -eq 0 ]; then
                echo "Translated successfully: $target_file"
            else
                echo "Translation failed for: $src_file (Locale: $target_locale)" >&2
            fi
        fi
    done
done

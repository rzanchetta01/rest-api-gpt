while read -r dependency; do
    echo "Installing $dependency"
    pip install "$dependency"
done < python_dependencies.txt
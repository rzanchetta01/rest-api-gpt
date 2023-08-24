$dependencies = Get-Content -Path .\python_dependencies.txt

foreach ($dependency in $dependencies) {
    Write-Host "Installing $dependency"
    pip install $dependency
}
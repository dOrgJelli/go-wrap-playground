import os
import subprocess
import re

ACCEPTED_IMPORTS = {'wrap', 'env'}

def build_wam_module():
    subprocess.run(["yarn", "build"]) # this fails
    return "./build/wrap.wasm"

def convert_wasm_to_wat(wasm_file):
    wat_file = wasm_file.replace('.wasm', '.wat')
    subprocess.run(['wasm2wat', wasm_file, '-o', wat_file], check=True)
    return wat_file

def find_import_functions(wat_file):
    import_functions = []

    with open(wat_file, 'r') as f:
        content = f.read()

    pattern = r'\(import\s+"([^"]+)"\s+"([^"]+)"'
    matches = re.findall(pattern, content)

    for match in matches:
        namespace = match[0]
        if namespace not in ACCEPTED_IMPORTS:
            import_functions.append(match[1])

    return import_functions

def snip_unwanted_imports(wasm_module_path , import_functions):
    if len(import_functions) == 0:
        return
    subprocess.run(['wasm-snip', wasm_module_path, '-o', wasm_module_path, '-p', *import_functions], check=True)

def main():
    print("Building wasm module...")
    wasm_module_path = build_wam_module()
    print("Wasm module built at: ", wasm_module_path)

    print("Converting wasm to wat...")
    wat_file = convert_wasm_to_wat(wasm_module_path)
    print("Wat file generated at: ", wat_file)

    print("Finding unwanted import functions...")
    import_functions = find_import_functions(wat_file)
    print("Found unwanted import functions: ", import_functions)

    print("Snipping unwanted import functions...")
    snip_unwanted_imports(wasm_module_path, import_functions)
    print("Wasm module built at: ", wasm_module_path)

    print("Removing wat file...")
    os.remove(wat_file)
    print("Wat file removed")


if __name__ == "__main__":
    main()

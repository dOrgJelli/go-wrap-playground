use std::collections::HashMap;
use std::fs::File;
use std::io::{self, Read, Write};
use wasmparser::{Parser, ParserState, SectionCode};

// Function to reorder the import entries alphabetically
fn reorder_imports(import_section: &mut wasmparser::ImportSection) {
    import_section.entries.sort_by(|a, b| {
        let module_a = a.module().unwrap().to_string();
        let module_b = b.module().unwrap().to_string();
        module_a.cmp(&module_b)
    });
}

// Function to update call_indirect instructions with new function indices
fn update_call_indirect(
    code_section: &mut wasmparser::CodeSectionReader,
    import_indices: &HashMap<u32, u32>,
) {
    while let Ok(Some(code)) = code_section.read() {
        let mut operands = code.get_operands();
        if let wasmparser::Operator::CallIndirect(_) = code.operator {
            if let Some(import_index) = operands.pop() {
                if let Some(new_import_index) = import_indices.get(&import_index) {
                    let mut new_operands = operands.clone();
                    new_operands.push(*new_import_index);
                    let new_code = wasmparser::Operator::CallIndirect(new_operands);
                    code_section.set_operators(&[new_code]);
                }
            }
        }
    }
}

fn main() -> io::Result<()> {
    // Read the original Wasm file
    let mut file = File::open("original.wasm")?;
    let mut buffer = Vec::new();
    file.read_to_end(&mut buffer)?;

    // Parse the Wasm module
    let mut parser = Parser::new(0);
    parser.parse_all(&buffer)?;

    // Extract the import and code sections
    let mut import_section = None;
    let mut code_section = None;

    for section in parser.into_iter() {
        match section? {
            ParserState::Section(SectionCode::Import, reader) => {
                import_section = Some(reader.read()?);
            }
            ParserState::Section(SectionCode::Code, reader) => {
                code_section = Some(reader);
            }
            _ => {}
        }
    }

    // Reorder the import section
    if let Some(mut import_section) = import_section {
        reorder_imports(&mut import_section);
        
        // Create a mapping between old and new import indices
        let import_indices: HashMap<u32, u32> = import_section
            .entries
            .iter()
            .enumerate()
            .map(|(new_index, entry)| (entry.index(), new_index as u32))
            .collect();
        
        // Update the call_indirect instructions
        if let Some(mut code_section) = code_section {
            update_call_indirect(&mut code_section, &import_indices);
        }
        
        // Write the modified module to a new Wasm file
        let mut output = File::create("modified.wasm")?;
        output.write_all(&import_section.get_output())?;
    }

    Ok(())
}

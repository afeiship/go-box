# Box Package Examples

This directory contains various examples demonstrating the functionality of the `box` package.

## Examples

### 1. Basic Example (`examples/basic`)
Shows the basic usage of the box package:
- Simple text boxes
- Boxes with ANSI color codes
- Boxes with Unicode characters

```bash
cd examples/basic && go run main.go
```

### 2. Advanced Example (`examples/advanced`)
Demonstrates advanced customization options:
- Different border styles (round, single, double)
- Box padding
- Box indentation
- Combination of all options

```bash
cd examples/advanced && go run main.go
```

### 3. Interactive Example (`examples/interactive`)
Allows users to input their own text and see it rendered in boxes:
- Read input from user
- Display the text in different border styles

```bash
cd examples/interactive && go run main.go
```

### 4. Demo Example (`examples/demo`)
A comprehensive demo showing various use cases:
- Status messages
- Error messages
- Welcome messages
- Menu options
- Code examples

```bash
cd examples/demo && go run main.go
```

## Features Demonstrated

- **Basic Box Creation**: Simple ASCII boxes with text content
- **ANSI Color Support**: Automatic handling of color codes in box content
- **Unicode Support**: Proper width calculation for Unicode characters
- **Custom Border Styles**: Round, single, and double border styles
- **Padding**: Add internal spacing within boxes
- **Indentation**: Control horizontal positioning of boxes
- **Mixed Styling**: Combine multiple options for custom appearances

## Running Examples

Make sure you're in the project root directory, then run any example:

```bash
# Basic usage
cd examples/basic && go run main.go

# Advanced features
cd examples/advanced && go run main.go

# Interactive input
cd examples/interactive && go run main.go

# Complete demo
cd examples/demo && go run main.go
```

Each example showcases different aspects of the box package and can serve as a reference for your own implementations.
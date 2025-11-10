# Go Language Learning Application

## Overview

An interactive web-based learning platform designed to teach Go programming fundamentals from basic variables to object-oriented programming concepts. The application provides a clean, organized interface with progressive lessons and comprehensive code examples.

## Purpose

To help beginners learn Go programming through structured lessons with practical code examples covering:

- Variables and data types
- Constants and control structures
- Functions and methods
- Structs, interfaces, and OOP concepts

## Recent Changes

**November 09, 2025**

- Initial project setup with Go web server
- Created complete learning platform with 4 comprehensive lessons
- Implemented responsive UI with lesson navigation
- Added code examples for all major Go concepts
- **Enhanced Quiz System**:
  - Created 100-question bank covering all Go fundamentals
  - Implemented Fisher-Yates shuffle algorithm for question randomization
  - Each quiz session randomly selects 10 questions from the bank
  - Added 30-second countdown timer with auto-submit functionality
  - Visual timer display with pulsing animation when time is running low
  - Dynamic question rendering from JavaScript objects
  - Proper timer cleanup to prevent memory leaks
- Implemented percentage-based scoring system with detailed results
- Created quiz navigation button in sidebar

## Project Architecture

### Structure

```
.
├── main.go           # Go web server (serves static files on port 5000)
├── static/
│   └── index.html    # Single-page learning application with all lessons
├── .gitignore        # Go-specific gitignore
└── readme.md         # Project documentation
```

### Technology Stack

- **Backend**: Go 1.23 with standard library (net/http)
- **Frontend**: HTML5, CSS3, Vanilla JavaScript
- **Server**: Static file server running on 0.0.0.0:5000

### Key Features

1. **Interactive Navigation**: Sidebar navigation with 4 lessons
2. **Responsive Design**: Mobile-friendly layout
3. **Keyboard Controls**: Arrow keys for lesson navigation
4. **Progress Tracking**: Visual progress indicator
5. **Code Examples**: Syntax-highlighted code blocks for each topic
6. **Advanced Quiz System**:
   - 100-question bank with comprehensive Go programming topics
   - Randomized selection of 10 questions per session (Fisher-Yates shuffle)
   - 30-second countdown timer with visual countdown display
   - Auto-submit when timer expires
   - Pulsing animation alert when 10 seconds remain
   - Dynamic question rendering from JavaScript objects
   - Multiple-choice format (A/B/C/D options)
7. **Score Calculation**: Percentage-based scoring with detailed feedback
8. **Results Display**: Shows correct/incorrect answers and performance breakdown

### Lessons Content

1. **Variables & Data Types**: Declaration methods, basic types, zero values
2. **Constants & Control Flow**: Constants, if/else, switch, for loops
3. **Functions & Methods**: Parameters, returns, variadic functions, closures
4. **Structs & OOP**: Structs, methods, interfaces, composition, polymorphism

## Running the Application

The application automatically starts via the configured workflow:

```bash
go run main.go
```

Access at: http://0.0.0.0:5000

## User Preferences

- Clean, educational interface
- Progressive learning path from basics to advanced concepts
- Practical code examples for hands-on learning

# Build with Replit

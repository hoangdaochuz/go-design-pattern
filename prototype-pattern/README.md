# Prototype Pattern

## Overview
The Prototype pattern is a creational design pattern that allows you to create new objects by cloning an existing object, known as the prototype. This pattern is useful when the cost of creating an object is more expensive than copying an existing one.

## Problem It Solves
- Avoids expensive object creation
- Reduces subclassing
- Hides complexity of object creation
- Creates objects without coupling to their specific classes

## Example Use Case
In our example, we'll implement a prototype pattern for cloning different types of houses. Instead of creating new house objects from scratch each time (which could be expensive if there are many properties and calculations involved), we'll clone existing house prototypes.

## Implementation
We'll create:
1. A Prototype interface that declares the cloning method
2. Concrete prototypes (different house types) that implement the cloning interface
3. A client that creates new objects by asking a prototype to clone itself

## Code Structure

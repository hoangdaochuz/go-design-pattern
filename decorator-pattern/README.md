# Decorator Pattern

The Decorator pattern allows behavior to be added to individual objects dynamically by wrapping them in decorator objects that contain the added behavior.

## Problem It Solves
- Allows functionality to be added to objects without altering their structure
- Provides a flexible alternative to subclassing for extending functionality
- Enables adding or removing responsibilities from objects at runtime

## Example Use Case
In this example, we demonstrate a pizza ordering system where toppings can be added dynamically to a base pizza. Each topping is a decorator that wraps the pizza and adds its cost and description.

## Key Components
- **Component Interface (Pizza)**: Defines the interface for objects that can have responsibilities added
- **Concrete Component (BasePizza)**: The base object that can be decorated
- **Decorator (ToppingDecorator)**: Abstract class that implements the Component interface and has a reference to a Component object
- **Concrete Decorators (CheeseTopping, PepperoniTopping, etc)**: Add specific behaviors/responsibilities to the component

## Benefits
- Provides greater flexibility than static inheritance
- Avoids feature-laden classes high up in the hierarchy
- Allows responsibilities to be added or removed at runtime
- Supports Single Responsibility Principle by allowing functionality to be divided into classes

## When to Use
- When you need to add responsibilities to objects dynamically and transparently
- When extension by subclassing is impractical or impossible
- When you want to add or remove responsibilities at runtime
- When you want to avoid a class hierarchy bloated with every possible combination of features

## Real-World Analogies
- Adding toppings to a pizza
- Adding options to a car model
- Adding features to a software subscription

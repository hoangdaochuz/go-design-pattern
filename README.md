# Design Patterns in Go

This repository contains examples of common design patterns implemented in Go.

## Patterns Implemented

### Adapter Pattern
Allows incompatible interfaces to work together by creating a wrapper that translates one interface to another. In this example, it adapts a Lightning port charger to work with a Type-C charging interface.

Example components:
- `ChargeBattery` interface defining the target charging interface
- `TypeCCharger` implementing the standard Type-C charging
- `LightningCharger` implementing Lightning port charging
- `TypeCChargerAdapter` adapting Lightning to Type-C interface

### Factory Method Pattern
Creates objects without exposing the instantiation logic to the client.

Example components:
- Notifier factories for Slack and Mail notifications
- Common notification interface

### Strategy Pattern 
Enables selecting an algorithm's implementation at runtime.

Example components:
- Payment processor supporting multiple payment methods:
  - Credit Card
  - PayPal 
  - Cryptocurrency

### Proxy Pattern
Provides a surrogate for another object to control access to it.

Example components:
- Nginx proxy implementation with request limiting

### Facade Pattern
Provides a simplified interface to a complex subsystem.

Example components:
- Wallet facade that handles complex wallet operations through a simple interface

## Usage

See `main.go` for example usage of each pattern. Each pattern is implemented in its own package with clear separation of concerns.

## Benefits

- Clean, maintainable code through proper separation of concerns
- Flexible and extensible design
- Reduced coupling between components
- Reusable code patterns

# Proxy Pattern

## Introduction
The Proxy Pattern is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. It creates a representative object that controls access to another object, which may be remote, expensive to create, or in need of securing.

## Key Components
- **Subject Interface**: Defines common interface for RealSubject and Proxy
- **RealSubject**: The real object that the proxy represents
- **Proxy**: Maintains a reference to RealSubject and controls access to it

## Common Use Cases
1. **Virtual Proxy**: Delays creation of expensive objects until needed
2. **Protection Proxy**: Controls access rights to the object
3. **Remote Proxy**: Represents an object located remotely
4. **Logging Proxy**: Adds logging when accessing the object
5. **Caching Proxy**: Stores results of expensive operations

## Benefits
- Controls access to another object
- Can handle initialization and lifecycle of the real subject
- Adds security, logging, or other functionality
- Works even if real subject is not ready or available

## Example Scenarios
- Database connection management
- Access control systems
- Lazy loading of large images/files
- Caching frequently accessed data
- Remote service calls

## Implementation Considerations
- Ensure proxy and real subject implement same interface
- Decide whether to create real subject lazily or eagerly
- Consider thread safety if needed
- Balance added complexity vs benefits

## Code Example

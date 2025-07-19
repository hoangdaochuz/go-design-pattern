# Command Pattern

The Command pattern is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as method arguments, delay or queue a request's execution, and support undoable operations.

## Problem
Imagine you're working on a text editor application. Your current task is to create a toolbar with a bunch of buttons for various operations, such as copying, pasting, and cutting text. You created a very nice Button class that can be used for buttons in the toolbar. But how do you connect these buttons to various operations?

## Solution
The Command pattern suggests that GUI objects shouldn't send these requests directly. Instead, you should extract all of the request details, such as the object being called, the name of the method and the list of arguments into a separate command class with a single method that triggers this request.

## Structure
1. The **Command** interface declares a method for executing a command
2. **Concrete Commands** implement various kinds of requests
3. The **Invoker** asks the command to carry out the request
4. The **Receiver** knows how to perform the operations
5. The **Client** creates and configures concrete command objects

## Benefits
1. Single Responsibility Principle - decouples classes that invoke operations from classes that perform these operations
2. Open/Closed Principle - introduce new commands without breaking existing code
3. Implements undo/redo operations
4. Defers execution of commands
5. Assembles simple commands into complex ones

## Real-World Analogy
When you go to a restaurant, you (the client) write down your order (command) on a piece of paper and give it to a waiter (invoker). The waiter takes the order to the kitchen where the chef (receiver) reads the order and cooks the meal according to what's written on the order.

## Example Use Cases
1. GUI buttons and menu items
2. Macro recording
3. Multi-level undo/redo
4. Progress of operations and rollback
5. Scheduling operations

## Implementation
The implementation example will demonstrate a simple remote control system where commands can be used to control different devices like lights and music players.

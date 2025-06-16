# Adapter Pattern

## Overview
The Adapter Pattern is a structural design pattern that allows incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces by wrapping an object in an adapter to make it compatible with another class.

## Problem
Sometimes we have an existing system that needs to work with a new component, but their interfaces don't match. We can't modify the existing code, and the new component's interface is incompatible.

## Solution
The Adapter Pattern solves this by:
1. Creating an adapter class that implements the interface that the client expects
2. Internally using the new component's interface to fulfill the requests
3. Effectively translating between the two interfaces

## Example Use Case
Consider a scenario where we have:
- A legacy payment processing system that only works with USD
- A new international payment service that works with multiple currencies
- Need to make them work together without changing existing code

## Code Example

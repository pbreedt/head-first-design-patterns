# head-first-design-patterns
Implementation of design patterns covered by the book ["Head First Design Patterns"](https://www.amazon.com/_/dp/0596007124?tag=oreilly20-20) by [O'Reilly](https://www.oreilly.com/library/view/head-first-design/0596007124/)

An adaptation and continuation of work first started by [hypersport](https://github.com/hypersport/Head-First-Design-Patterns)

# Featured design patterns:

## Strategy
The Strategy Pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

## Observer
The Observer Pattern defines a one-to-many dependency between objects so that when one object changes state, all of its dependents are notified and updated automatically.

## Decorator
The Decorator Pattern attaches additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality. (A bit Java specific, but we know...)

## Factory Method
The Factory Method Pattern defines an interface for creating an object, but lets subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

Also included a simple factory example for completeness' sake - not really a design pattern.

## Design principles covered
- Identify the aspects of your application that vary and separate them from what stays the same.
- Program to an interface, not an implementation.
- Favor composition over inheritance.
- Strive for loosely coupled designs between objects that interact.
- Classes should be open for extension, but closed for modification.
- Depend upon abstractions. Do not depend upon concrete classes.

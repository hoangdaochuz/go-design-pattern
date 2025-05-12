# Template Method Pattern

## Context

We have a feature send OPT. We have 2 type of OTP need to be sent (SMS,Mail). The step of both type are the same.

- Generate n digit number OTP
- Save OTP to cache
- Get message
- Send notification to client

But there are some minor different between them at some steps like (get message, send notification). If we write each class to serve for
send OPT, it will duplicate code. And the code will be big and duplicate if we have more and more OTP method in the future.
So it is the time for `Template Method Pattern`

## Implement

`Template Method Pattern` will split the algorithms of task send OTP notification in many steps method.
And the base class (interface) will declare these steps. If there aren't different in some common step. It can implement default logic at these steps.
The subclass will implement other steps base on there business. And the algorithm to send OTP notification will a set of interface step.
And use `Dependency Injection` to inject specific subclass that implemented base class. The result we have is the algorithms of sending OTP notification
will behave differently base on the subclass we inject.

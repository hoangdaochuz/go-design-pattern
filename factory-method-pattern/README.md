# Factory method pattern

## Context

We have a feature about `Sending notification`. At the first time we only support for mail notification.
And we should not tie tightly the logic send mail into only mail. Since in the feature, maybe we also will support
some notifier like `Slack`,`Telegram`,... If we tied tightly the logic to a specific class, we must change all the codebase and
add conditions logic to handle for each case.
==> So, It's time to you `Factory Method Pattern`.

## Implement

- `Factory method pattern` is a pattern that help us work with some type or dependencies which we don't know exactly beforehand.
- By using a base class, which have a abstract method (factory method) to create a interface object (type object that we want to work)
- Subclass are factory class which implement base class -> it will create the object base on subclass
- The object which have created in above step are implement class of a base product class.

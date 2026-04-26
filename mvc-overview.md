Model-View-Controller

**Models** are about data, logic, and rules

This typically means interacting with your database, but it could also mean interacting
with data that comes from other services or APIS

Often includes validating data, normalizing it, etc.

For example, our web application is going to have user accounts, and logic for
validatiing passwords and authenticationg users will all be in the models package.


**Views** render data.

In out case, we are rendering HTML.

An API could use MVC and the views could be responsibel for generating JSON.

As little logic as possible. Only logic required to render data.

Example:
- "If next page exists, show next page link" is okay.
- logic to calculate a bunch of graphs should probably e handled elsewhere, and the passed
into a view as raw data to render.

Too much logic in views code very hard to maintain.


**Controllers** connects it all.

It won't directly render HTML, it won't directly touch the DB, but it will call
code from the models and views packages to do these things.

controllers as your[air traffic controllers]


## Walking Through WEB Request.
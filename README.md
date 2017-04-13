# LogicBoxes API tools
SDK and tools for the LogicBoxes API used by ResellerClub and other domain registrars.

## Who is it for?
This is only useful if you have a reseller account with anyone using the LogicBoxes system and its APIs, and would like to manage customers, domains and DNS from the command line.

The planned backend server will also be useful if you need to expose DNS management or other simpler management to users, either via web or custom clients.

## Features (work in progress)
- Manage customers: Look up account details, sign up new ones, reset passwords, edit details, suspend and delete, authenticate passwords (for making backend services).
- Manage domains: Look up domains per customer, list all, edit details, toggle locks, transfer, renew, change name servers.
- Manage DNS for domains: Add, remove and change all records.

## Planned features
The following features take a lower priority than getting the three major parts of the API supported (customers, domains and DNS) with a command line tool to handle everything implemented.

- Backend server: REST layer between frontends and the LB API to make it easier to set up your own custom website, rather than using LB's Supersite.
- Client tool for DNS management via the backend server.

These planned features are things I will be dogfooding. Likely to be implemented soon, depending on when the main features are done.

## Possible features
- Managing other products than just domains
- Reseller management
- Interactive mode using a fancy terminal interface.

## Dependencies
1. Go, currently tested with v1.8.x+.
2. [go-flags](https://github.com/jessevdk/go-flags)
3. [columnize](https://github.com/ryanuber/columnize)

## Platforms
Should work anywhere Go compiles to. The server tools are mainly tested on Linux but work on macOS. The command line tools and SDK should be universal, but currently needs a few lines of code for Windows support.

## Licence
MIT.

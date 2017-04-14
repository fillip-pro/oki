# Oki
Oki is a program written in Golang that will 
instantiate a clean infrastructure for the
[Fillip.pro](https://fillip.pro) projects. In
and of itself, it has no functionality other
than connecting to various services, and configuring
cloud infrastructure from either a greenfield state
or from a brownfield state. 

It is a fresh attempt at building infrastructure
with code, to reduce dependency hell, (both in the
cloud and within infrastructure development), using 
re-useable skills, and to try very hard to not
depend on shell scripting.

As a result of the greenfield and stateless approach
to infrastructure development, `Oki` also attempts 
to provide a method of cleanly destroying infrastructure
if necessary, without a fear of deleting essential data.

One instantiated, `Oki` will place itself within the
infrastructure and become the conductor, ensuring 
efficient uptime and effective scaling.

## Running
In order to run `Oki` you will need to configure your
environment variables for the cloud providers you 
intend to operate with:

 - Amazon: `AWS_TOKEN`
 - Azure: `AZURE_TOKEN`
 - Digital Ocean: `DO_TOKEN`
 - Google Compute: `GC_TOKEN`

Rather than using something like Terraform or Packer,
this project uses the SDKs provided by each of the above
to directly provision the clustered environment. Therefore
each provider-specific setup can be tuned and tailored, 
and the skills required to do that are in the `Go` 
programming language and the provider's platform. No
esoteric third-party tooling exists to complicate the 
setup.

## Development
`Oki` is almost entirely a Golang program and as such
it comes with all of the painpoints of programming in Go.
Sure, it can be clean, efficient, and effective, but its
dogmatic way of telling you how to structure workspaces
can also be a severe pain point. So, here's how it was 
done.

Firstly, the `$GOPATH` environment variable was setup 
(on macOS) to point to a root workspace folder, like so:

`export GOPATH=/myworkspace/oki`

Underneath `oki, is `bin`, `pkg`, and `src`. Underneath
`src` is the repository root (i.e. `gitlab.com`), underneath
that is the group root (i.e. `fillip`), and underneath
that is `oki`, with the real important fleshy parts. 

The reason `Oki` wasn't placed in a single workspace like
so many other developers do? Because when you're working with
a variety of organizations, on a variety of projects, and single
source of truth for `Go` can be highly inefficient and 
problematic. Opting for project specific `GOPATH`'s made more
sense and causes zero issues. 

Of course, the rest of the development tools are installed centrally,
like `godef`, `gotests`, `golint`, etc., because they're consistent
across projects. `/usr/local/go/bin` is of course added to the 
`$PATH` environment variable. Life is hard enough without managing
multiple `go` installs.

## Etymology
Named after the archipelago in the Sea of Japan.

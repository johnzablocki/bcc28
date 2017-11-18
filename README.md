# Boston Code Camp 28 Demo 

This repository contains the sample app discussed at my Going, Going, Golang presentation at Boston Code Camp 28.

## Setup

Some basic instructions for first time Go users

### Install Golang

Grab the approppriate installer at https://golang.org/dl/

Go will assume that you're working in a directory %USERPROFILE% on Windows or $HOME/go on Mac/Linux.  If you are working in another directory - say c:\dev\go\<projects> - you'll need to set your GOPATH environment variable to that path.  See https://github.com/golang/go/wiki/SettingGOPATH for more information.

### Install the Glide package manager 

Glide is one of a few package managers for Go.  You can download builds at https://github.com/Masterminds/glide/releases.  On a Mac, use `brew install glide` on Windows, copy the binary to a directory in your PATH.  I suggest copying it to the Go bin directory itself - c:\go\bin by default.

### Clone the repository

`git clone https://github.com/jzablocki/bcc28.git`

### Install the dependencies

From the root of the project (where you cloned into), run `glide install`

### Run the app

From the root (where main.go is) run `go run main.go` OR `go build` and run the bcc28 executable. 






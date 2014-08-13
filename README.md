gogel
=====

[![Build Status](https://travis-ci.org/Nevon/gogel.svg?branch=master)](https://travis-ci.org/Nevon/gogel)

The little browser engine that couldn't. A toy project for learning Go.


    $ go run main/main.go --help
        NAME:
            gogel - Parse an HTML document
    
    USAGE:
        gogel [global options] command [command options] [arguments...]
    
    VERSION:
        0.0.1
    
    COMMANDS:
        help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
        --html           path to HTML file to parse
        --help, -h           show help
        --generate-bash-completion
        --version, -v        print the version

    $ go run main/main.go --html example/html/test.html
        html
          body
            h1
              Hello friends
            div
              p
                This is a paragraph
              p
                Followed by another paragraph.

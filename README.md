# cyoa

A choose your own adventure story!

its for fun
Create your own adventure story by changing the `gopher.json` file or supplying your own json by using the `-file` flag. Adventure are in this structure

    {
        "chapter": {
            "title": "",
            "story": [
                "paragraph",
                "paragraph"
            ],
            "options": [
                {
                    "text": "",  // link text
                    "arc": ""    // chapter you can go to
                },
                {
                    "text": "",
                    "arc": ""
                }
            ]
        }
    }

    go run cmd/cyoaweb/main.go -port 3000 -file gopher.json
    open browser to localhost:3000

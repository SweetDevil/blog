# blog
golang personal blog based on revel framework, Every Thing is Experimental

Set up.

1. install revel framework
	go get github.com/revel/revel
2. install revel framework cmd tool
	go get github.com/revel/cmd/revel
	
	copy revel tool from your $gopath/bin/revel to your $PATH, I copy to $GOROOT/bin
3. get project code
	go get github.com/SweetDevil/blog
4. import db files to your database.

	see the /db_files/Readme.md
	
5. run blog code by revel cmd tool
	$ revel run github.com/SweetDevil/blog
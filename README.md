# go-prac-pro
		This is a standalone application which will invoke rest based api for
		different bus routes.
		
## Installation and Running

#### A: Install GO and set GOROOT variable  
	1: https://golang.org/doc/install		
	
#### B: Create following folder structure and then set the GOPATH variable
	1: .../workspace/src/github.com/

#### C: Clone GoPracPro/ code from the following url in specified folder:  
	1: https://github.com/kulkarnikoustubh/GoPracPro.git

### Code Package structure	
	
    local -->
			 Contains all packages which can be directly imported by any package outside local Package, they are module internal packages and are not meant to be used by any other application.
	
	local/platform:
			  Contains packages specific to logging,config,in memory storage,utils which can be used as common platform components across the module.
			
	Note:
		- No Package at same level accesses each other.
		- Packages that needs to be imported by other packages, should be either part of local package or subpackage of the package that wants to import it.			  		  	
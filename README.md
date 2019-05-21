# go-prac-pro
		This is a standalone application which will invoke rest based api for
		different bus routes.
		
## Installation and Running

#### A: Install GO and set GOROOT variable  
	1: https://golang.org/doc/install		
	
#### B: Create following folder structure and then set the GOPATH variable
	1: .../workspace/src/github.com/

#### C: Clone GoPracPro code from the following url in specified folder:  
	1: https://github.com/kulkarnikoustubh/GoPracPro.git
	
#### D: Build and Run
	1: Go to .../workspace/src/github.com/GoPracPro/src (Terminal for Linux/Mac and Command Prompt Windows)
	2: Execute command to create executable as : go build -o nextbus
	3: Run executable created with name nextbus with arguments bus route,direction and  stop name
	 ex on linux : ./nextbus "METRO Blue Line" "Bloomington Central Station" "South"  	

### Code Package structure	
	
    local -->
			 Contains all packages which can be directly imported by any package outside local Package, they are module internal packages and are not meant to be used by any other application.
	
	local/platform:
			  Contains packages specific to logging,config,in memory storage,utils which can be used as common platform components across the module.
			
	Note:
		- No Package at same level accesses each other.
		- Packages that needs to be imported by other packages, should be either part of local package or subpackage of the package that wants to import it.

### Metro Transit API Reference doc

    http://svc.metrotransit.org/					  		 	
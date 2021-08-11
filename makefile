BASE=$(shell pwd)
GOBIN := $(BASE)/bin
GOSRC := ./src
PROGS := findlinks1.exe\
		findlinks2.exe\
		findlinks3.exe\
		findlinks4.exe\
		fetch1.exe\
		fetch4.exe\
		outline.exe\
		nodecount.exe\
		printtextnodes.exe\
		wordfreq.exe
all:$(PROGS)


findlinks1.exe:$(GOSRC)/findlinks/v1/findlinks1.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

findlinks2.exe:$(GOSRC)/findlinks/v2/findlinks2.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

findlinks3.exe:$(GOSRC)/findlinks/v3/findlinks3.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

findlinks4.exe:$(GOSRC)/findlinks/v4/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

fetch1.exe:$(GOSRC)/fetch/v1/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

fetch4.exe:$(GOSRC)/fetch/v4/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

outline.exe:$(GOSRC)/outline/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

nodecount.exe:$(GOSRC)/nodecount/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^
		
printtextnodes.exe:$(GOSRC)/printtextnodes/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

wordfreq.exe:$(GOSRC)/wordfreq/main/main.go
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$@ $^

clean:
	rm -rf $(GOBIN)/

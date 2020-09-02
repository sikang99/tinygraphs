#
# Makefile for golang-example
#
.PHONY: usage edit build clean git
#----------------------------------------------------------------------------------
PROG=tinygraphs
PORT=8080
usage:
	@echo "make [edit|build]"
#----------------------------------------------------------------------------------
edit e:
	@echo "make (edit:e) [history]"
edit-go eg:
	vi main.go
edit-history eh:
	vi HISTORY.md
#----------------------------------------------------------------------------------
build b:
	go build -o $(PROG)
#----------------------------------------------------------------------------------
clean:
	rm -f bin/*
	docker system prune
#----------------------------------------------------------------------------------
run r:
	PORT=$(PORT) ./$(PROG)

open o:
	open http://localhost:$(PORT)/

kill k:
	pkill $(PROG)
#----------------------------------------------------------------------------------
git g:
	@echo "make (git:g) [update|store]"
git-update gu:
	git add .
	git commit -a -m "update contents"
	git push
git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------


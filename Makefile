.PHONY: run buildALL buildd buildp init td test gitC gitInit gitUpdateTag
default: run

program_name = github.com/medtereus/BetterConfig
GITrepo := https://github.com/medtereus/BetterConfig
GIT_tag_version:=v1.4.1

BUILD_GOOS=linux
BUILD_GOARCH=amd64

run:
	go run .

td:
	go mod tidy

test:
	go test -v .\...

init:
	find . -type f \
		-not -path './.git/*' \
		\( -name '*.go' -o -name 'go.mod' \) \
		-exec sed -i 's|ProjectTemplate|$(program_name)|g' {} +
	find . -type f \
		-not -path './.git/*' \
		-name '*.go.old_*' \
		-delete
	-go mod init $(program_name)
	go mod tidy

gitInit:
	@if [ -z "$(GITrepo)" ]; then \
		echo "GITrepo is not configured. Aborting."; exit 1; \
	fi
	@if [ -d .git ]; then \
		echo "Git repository already exists. Aborting."; exit 1; \
	fi
	@git init
	@git branch -M main
	@git add .
	@git commit -m "Initial commit"
	@git tag $(GIT_tag_version)
	@git remote add origin "$(GITrepo)"
	@git push -u origin main
	@git push origin $(GIT_tag_version)


gitC:
	git add .
	git commit -m "commit from Makefile"
	git push -u origin main

gitUpdateTag:
	@if [ -z "$(GIT_tag_version)" ]; then \
		echo "GIT_tag_version is not configured. Aborting."; exit 1; \
	fi
	@if git rev-parse "$(GIT_tag_version)" >/dev/null 2>&1; then \
		echo "Tag $(GIT_tag_version) already exists. Aborting."; exit 1; \
	fi
	@git add .
	@git commit -m "Update module version to $(GIT_tag_version)"
	@git tag $(GIT_tag_version)
	@git push origin main
	@git push origin $(GIT_tag_version)
	@echo "Version $(GIT_tag_version) published."
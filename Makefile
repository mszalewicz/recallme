.PHONY: run build dependencies

run: dependencies
	@$(call PrintSection,"Generating Go")
	@go run cmd/main.go

build: dependencies
	@$(call PrintSection,"Building Go")
	@go build -o ./bin/recallme ./cmd

dependencies:
	@$(call PrintSection,"Generating Templ")
	@templ generate

	@$(call PrintSection,"Generating Tailwind")
	@tailwindcss -i ./static/input.css -o ./static/output.css --minify

templ:
	## @templ generate -watch -proxy=http://localhost:3000
	@templ generate -watch

tailwind:
	@tailwindcss -i ./static/input.css -o ./static/output.css --minify --watch

##### Print section divider

define PrintSection

	$(eval $@SectionName = $(1))

	echo "\n"
	echo "------------------------------------------------------------------------------------------------------------------"
	echo ${$@SectionName}
	echo "------------------------------------------------------------------------------------------------------------------"
	echo "\n"
endef
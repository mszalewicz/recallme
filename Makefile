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
	@tailwindcss -i ./assets/input.css -o ./assets/output.css



##### Print section divider

define PrintSection
	$(eval $@Front = ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>     ")
	$(eval $@Back = ":     <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	$(eval $@SectionName = $(1))

    echo "\n"
	echo ${$@Front}${$@SectionName}${$@Back}
	echo "\n"
endef
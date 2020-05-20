BAZEL		= bazel

.PHONY: build
build: gazelle
	$(BAZEL) build //...

clean:
	$(BAZEL) clean --expunge

gazelle:
	$(BAZEL) run //:gazelle -- update-repos -from_file=go.mod
	$(BAZEL) run //:gazelle

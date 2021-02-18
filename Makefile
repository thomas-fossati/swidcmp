CORPUS_FILES := $(wildcard corpus/*.xml)

SWIDCMP = swidcmp

run: $(SWIDCMP) ; @for f in $(CORPUS_FILES) ; do ./$(SWIDCMP) $$f ; done

$(SWIDCMP): ; go build

clean: ; $(RM) $(SWIDCMP)

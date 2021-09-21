CORPUS_FILES := $(wildcard corpus/*.xml)
CBOR_FILES := $(CORPUS_FILES:.xml=.cbor)

SWIDCMP = swidcmp

run: $(SWIDCMP) ; @for f in $(CORPUS_FILES) ; do ./$(SWIDCMP) $$f ; done

$(SWIDCMP): main.go ; go build

clean: ; $(RM) $(SWIDCMP) $(CBOR_FILES)

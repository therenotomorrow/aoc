.PHONY: code test/smoke test/cover

code:
	@"$(CURDIR)/scripts/code.sh"

test/smoke:
	@"$(CURDIR)/scripts/test.sh" smoke

test/cover:
	@"$(CURDIR)/scripts/test.sh" cover

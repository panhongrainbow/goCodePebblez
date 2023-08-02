.PHONY: help

test:
	go test -v -run='^\QTest_Check_' ./bytez
cover:
	go test -cover -run='^\QTest_Check_' ./bytez
benchmark:
	go test -bench=. -run='^\QTest_Performance_' ./bytez
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "This Makefile provides the following targets for testing and benchmarking:"
	@echo ""
	@echo "  test       - Run unit tests"
	@echo "  cover      - Run coverage tests"
	@echo "  benchmark  - Run benchmark tests"
	@echo ""
	@echo "These targets are designed to ensure that the packages function correctly"
	@echo "and to identify any potential issues."
	@echo ""
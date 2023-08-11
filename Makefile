.PHONY: help

test_path = ./bytez ./syncPoolUtil

test:
	go test -v -run='^Test_Check_(\w+)' $(test_path)
cover:
	go test -cover -run='^Test_Check_(\w+)' $(test_path)
benchmark:
	go test -bench='^Benchmark_Performance_(\w+)' $(test_path)
Comparator:
	go test -bench='^Benchmark_Comparator_(\w+)' $(test_path)

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "This Makefile provides the following targets for testing and benchmarking:"
	@echo ""
	@echo "  test       - Run unit tests 单元测试"
	@echo "  cover      - Run coverage tests 复盖率测试"
	@echo "  benchmark  - Run benchmark tests 基准测试"
	@echo "  Comparator - Run Comparator tests 对数器测试"
	@echo ""
	@echo "These targets are designed to ensure that the packages function correctly"
	@echo "and to identify any potential issues."
	@echo ""
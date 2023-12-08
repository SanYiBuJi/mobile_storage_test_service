# 设置目标名称和文件名
TARGET := mobile_storage_test_service
FILENAME := custom_name

# 设置支持的操作系统和系统架构
OS_LIST := linux windows darwin
ARCH_LIST := amd64 arm64

# 设置编译器和编译选项
CC := go
BUILD_FLAGS :=

# 设置输出目录
OUTPUT_DIR := build

# 默认目标
all: build

# 编译目标
build: $(OS_LIST)

$(OS_LIST):
	@mkdir -p $(OUTPUT_DIR)
	$(foreach ARCH,$(ARCH_LIST), \
		$(CC) build $(BUILD_FLAGS) -o $(OUTPUT_DIR)/$(FILENAME)_$@_$(ARCH) \
	)

# 清理目标
clean:
	rm -rf $(OUTPUT_DIR)

# 声明伪目标
.PHONY: all build $(OS_LIST) clean

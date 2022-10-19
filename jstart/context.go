package jstart

import "os"

type Context interface {
	GetJdkVersion() string
	GetMemoryLimit() int64
	GetEnv(string) string
}

type ContextBuilder interface {
	JdkVersion(jdkVersion string) ContextBuilder
	MemoryLimit(memoryLimit int64) ContextBuilder
	Build() *context
}

type contextBuilder struct {
	jdkVersion  string
	memoryLimit int64
}

type context struct {
	jdkVersion  string
	memoryLimit int64
}

func (c *context) GetJdkVersion() string {
	if c != nil {
		return c.jdkVersion
	}
	return "unknown"
}

func (c *context) GetMemoryLimit() int64 {
	if c != nil {
		return c.memoryLimit
	}
	return 0
}

func (c *context) GetEnv(key string) string {
	return os.Getenv(key)
}

func NewContextBuilder() *contextBuilder {
	return &contextBuilder{}
}

func (b *contextBuilder) JdkVersion(jdkVersion string) ContextBuilder {
	b.jdkVersion = jdkVersion
	return b
}

func (b *contextBuilder) MemoryLimit(memoryLimit int64) ContextBuilder {
	b.memoryLimit = memoryLimit
	return b
}

func (b *contextBuilder) Build() *context {
	return &context{
		jdkVersion:  b.jdkVersion,
		memoryLimit: b.memoryLimit,
	}
}

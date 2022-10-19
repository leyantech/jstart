package jstart

type Context struct {
	JdkVersion  string
	MemoryLimit int64
	IsProd      bool
}

type ContextBuilder interface {
	JdkVersion(jdkVersion string) ContextBuilder
	MemoryLimit(memoryLimit int64) ContextBuilder
	IsProd(isProd bool) ContextBuilder
	Build() *Context
}

type contextBuilder struct {
	jdkVersion  string
	memoryLimit int64
	isProd      bool
}

func (c *Context) GetJdkVersion() string {
	if c != nil {
		return c.JdkVersion
	}
	return "unknown"
}

func (c *Context) GetMemoryLimit() int64 {
	if c != nil {
		return c.MemoryLimit
	}
	return 0
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

func (b *contextBuilder) IsProd(isProd bool) ContextBuilder {
	b.isProd = isProd
	return b
}

func (b *contextBuilder) Build() *Context {
	return &Context{
		JdkVersion:  b.jdkVersion,
		MemoryLimit: b.memoryLimit,
		IsProd:      b.isProd,
	}
}

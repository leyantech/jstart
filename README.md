# jstart
jstart is a scriptable version of the `java` command,
it expands java options automatically with regard to jstart rules.

Without any jstart rule, the behaviour of jstart is exactly the same with `java`.

With custom jstart rules, you can:

1. Insert a group of predefined system properties(`-D`) .
2. Choose GC algorithm automatically according to memory limit in docker containers.
3. Set Xmx with arithmetic expressions like `xmx=quota*2/3`
4. Switch remote debugging on/off according to specific environment variable.
5. Switch Native Memory Tracking on/off according to some conditions of the current running context.


With a customized jstart rules loader, you can also change the application start arguments via an external configuration service, 
instead of rewrite the start script and rebuild a container image.

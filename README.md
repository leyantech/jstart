# jstart
**Replace the boilerplate `start-java-application.sh`**

```shell
limit_in_bytes=$(cat /sys/fs/cgroup/memory/memory.limit_in_bytes)

if [ "$limit_in_bytes" -ne "9223372036854771712" ]
then
limit_in_megabytes=$(expr $limit_in_bytes \/ 1048576)
heap_size=$(expr $limit_in_megabytes \* 2 \/ 3)
export JAVA_TOOL_OPTIONS="-Xmx${heap_size}m"
fi

if [ "$ENV_PRIORITY" = "test" ]; then
jdwp_port=${JDWP_PORT:-3360}
export JAVA_TOOL_OPTIONS="-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=${jdwp_port} $JAVA_TOOL_OPTIONS"
fi

JAVA_MODULE_EXPORT_OPTIONS=""
if [ "$JAVA_VERSION" ~= "8" ]; then
  JAVA_MODULE_EXPORT_OPTIONS="--illegal-access=warn --add-opens java.base/java.nio=ALL-UNNAMED --all-opens java.base/java.lang=ALL-UNNAMED --all-opens java.base/java.lang.reflect=ALL-UNNAMED"
fi

java -DFRAMEWORK_PROP1=value1 -DFRAMEWORK_PROP2=value2 \
 $(echo $JAVA_MODULE_EXPORT_OPTIONS) -cp application.jar com.leyantech.app.Main
```

**With**:

```shell
export JSTART="xmx=quota*2/3;predefine_properties=true"
jstart -cp application.jar com.leyantech.app.Main
```

**And gain the ability to**:

> customize jvm start options via remote configuration services.
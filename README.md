# kratos-realtimemap

源自于 [Proto.Actor](https://proto.actor/) 的一个实时地图的实例 [realtimemap-go](https://github.com/asynkron/realtimemap-go.git)

它有一个在线的演示可看: <https://realtimemap.skyrise.cloud/>

它之前与网页端的实时数据传输用的是微软的SignalR,我这里则直接使用了Websocket.

## 涵盖的技术点

- 使用Kratos开发微服务
- 使用Kratos的BFF与网页端通讯
- Kratos与MQTT的融合使用
- Kratos与Websocket的融合使用

## 技术栈

- [Kratos](https://go-kratos.dev/)
- [Consul](https://www.consul.io/)
- [Jaeger](https://www.jaegertracing.io/)
- [MQTT](https://mqtt.org/)
- [Websocket](https://entgo.io/)
- [VUE](https://vuejs.org/)

## Docker部署开发服务器

### Consul

```shell
docker pull bitnami/consul:latest

docker run -itd \
    --name consul-server-standalone \
    -p 8300:8300 \
    -p 8500:8500 \
    -p 8600:8600/udp \
    -e CONSUL_BIND_INTERFACE='eth0' \
    -e CONSUL_AGENT_MODE=server \
    -e CONSUL_ENABLE_UI=true \
    -e CONSUL_BOOTSTRAP_EXPECT=1 \
    -e CONSUL_CLIENT_LAN_ADDRESS=0.0.0.0 \
    bitnami/consul:latest
```

### Jaeger

```shell
docker pull jaegertracing/all-in-one:latest

docker run -d \
    --name jaeger \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 14250:14250 \
    -p 9411:9411 \
    jaegertracing/all-in-one:latest
```

## 测试

Swagger-UI的访问地址: <http://localhost:8800/q/swagger-ui>

## 参考资料

- [GTFS Realtime Reference](https://developers.google.com/transit/gtfs-realtime/reference)
- [High-frequency positioning](https://digitransit.fi/en/developers/apis/4-realtime-api/vehicle-positions/)

FROM envoyproxy/envoy:v1.9.1

COPY ./third_party/proxy/envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

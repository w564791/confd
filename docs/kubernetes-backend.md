#### template
```
# cat /etc/confd/templates/myconfig.conf.tmpl

{{range $key := ls "/istio-system"}}
{{$subkey :=printf "/istio-system/%s" $key}}
 {{$dir := dir $subkey}}
upstream  {{base $dir }}-{{base  $subkey}} {
{{$data := json (getv $subkey)}}
    {{range $data.subsets}}
        {{$address :=  .addresses}}
        {{$port:= .ports}}
            {{range $address}}
                {{$ip:=.ip}}
                {{range $port}}
                   {{if or (eq "http" .name) (eq "http2" .name)}}
                      server {{$ip}}:{{.port}};
                   {{end}}
                {{end}}
            {{end}}
    {{end}}

}



server {
    listen 80;
    server_name {{base  $subkey}}.example.com;
    access_log /var/log/nginx/{{base  $subkey}}.access.log;
    error_log /var/log/nginx/{{base  $subkey}}.error.log;

    location / {
        index             index.html index.htm;
        proxy_pass        http://{{base $dir}}-{{base $subkey}};
        proxy_redirect    off;
        proxy_set_header  Host             $host;
        proxy_set_header  X-Real-IP        $remote_addr;
        proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}
{{end}}
```
### template conf
keys  format   "/{{namespace}}/{{endpoint}}"
```
# cat /etc/confd/conf.d/myconfig.toml
[template]
src = "myconfig.conf.tmpl"
dest = "/tmp/myconfig.conf"
keys = [
    "/istio-system/istio-egressgateway",
    "/default/reviews",
    "/default/ratings",
    "/istio-system/istio-ingressgateway",
    "/istio-system/grafana",
]
```

### confd-config
```
# cat /etc/confd/config.toml
backend = "kubernetes"
kubeconfig="/opt/kubeconfig"
log-level = "info"
interval = 2
```

### run

```
confd -config-file /etc/confd/config.toml
```

### result
```
# cat /tmp/myconfig.conf




upstream  istio-system-grafana {








                      server 172.20.113.241:3000;





}



server {
    listen 80;
    server_name grafana.example.com;
    access_log /var/log/nginx/grafana.access.log;
    error_log /var/log/nginx/grafana.error.log;

    location / {
        index             index.html index.htm;
        proxy_pass        http://istio-system-grafana;
        proxy_redirect    off;
        proxy_set_header  Host             $host;
        proxy_set_header  X-Real-IP        $remote_addr;
        proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}



upstream  istio-system-istio-egressgateway {








                      server 172.20.113.231:80;










                      server 172.20.113.244:80;









}



server {
    listen 80;
    server_name istio-egressgateway.example.com;
    access_log /var/log/nginx/istio-egressgateway.access.log;
    error_log /var/log/nginx/istio-egressgateway.error.log;

    location / {
        index             index.html index.htm;
        proxy_pass        http://istio-system-istio-egressgateway;
        proxy_redirect    off;
        proxy_set_header  Host             $host;
        proxy_set_header  X-Real-IP        $remote_addr;
        proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}



upstream  istio-system-istio-ingressgateway {












                      server 192.168.178.128:80;

















}



server {
    listen 80;
    server_name istio-ingressgateway.example.com;
    access_log /var/log/nginx/istio-ingressgateway.access.log;
    error_log /var/log/nginx/istio-ingressgateway.error.log;

    location / {
        index             index.html index.htm;
        proxy_pass        http://istio-system-istio-ingressgateway;
        proxy_redirect    off;
        proxy_set_header  Host             $host;
        proxy_set_header  X-Real-IP        $remote_addr;
        proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}


```
FROM alpine


ADD confd /opt/
RUN mkdir -p /etc/confd/{conf.d,template} && chmod +x /opt/confd
ADD config.toml /etc/confd/
ADD myconfig.conf.tmpl /etc/confd/template/
ADD myconfig.toml /etc/confd/conf.d/
ENTRYPOINT ["sh","-c", "/opt/confd -config-file /etc/confd/config.toml"]
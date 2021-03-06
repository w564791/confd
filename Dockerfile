FROM ubuntu


ADD confd /opt/
RUN mkdir -p /etc/confd/{conf.d,templates} && chmod +x /opt/confd
ADD config.toml /etc/confd/
ADD myconfig.conf.tmpl /etc/confd/templates/
ADD myconfig.toml /etc/confd/conf.d/
ENTRYPOINT ["sh","-c", "/opt/confd -config-file /etc/confd/config.toml"]
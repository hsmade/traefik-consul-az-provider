consul:
	wget -O consul.zip https://releases.hashicorp.com/consul/1.19.0/consul_1.19.0_$$(uname -o| tr '[A-Z]' '[a-z]')_$$(uname -m).zip
	unzip consul.zip consul
	rm consul.zip
traefik:
	wget -O traefik.tgz https://github.com/traefik/traefik/releases/download/v3.0.2/traefik_v3.0.2_$$(uname -o| tr '[A-Z]' '[a-z]')_$$(uname -m).tar.gz
	tar xzf traefik.tgz traefik
	rm traefik.tgz
hello:
	go build hello.go

.PHONY: test
test: consul traefik hello
	{ ./consul agent -config-dir consul.d & echo $$! > consul.pid; }
	{ ./traefik --log --log.level=DEBUG --accesslog \
		--experimental.localplugins.consulcataloglocal.modulename=github.com/hsmade/traefik-consul-az-provider \
		--providers.plugin.consulcataloglocal.localsubnet=127.0.0.1/32 \
		--providers.plugin.consulcataloglocal.consulcatalog=true \
		--providers.plugin.consulcataloglocal.consulcatalog.endpoint.address=localhost:8500 \
		--providers.plugin.consulcataloglocal.consulcatalog.exposedByDefault=true \
		--api \
		--api.insecure \
		--entryPoints.web.address=:8081 \
		--providers.consulcatalog=true \
		--providers.consulcatalog.endpoint.address=localhost:8500 \
		--providers.consulcatalog.exposedByDefault=true \
		& echo $$! > traefik.pid; }
	{ ./hello & echo $$! > hello.pid; }
	./consul services register consul.d/hello.hcl
	curl -v -HHost:hello localhost:8081

.PHONY: stop
stop:
	kill $$(cat consul.pid traefik.pid hello.pid)
	rm *.pid

clean:
	rm -rf consul.data hello traefik consul *.pid

all: test stop clean


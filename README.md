# External DNS - Adguard Home Provider (Webhook)

[![](https://img.shields.io/github/license/muhlba91/external-dns-provider-adguard?style=for-the-badge)](LICENSE.md)
[![](https://img.shields.io/github/actions/workflow/status/muhlba91/external-dns-provider-adguard/verify.yml?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/actions/workflows/verify.yml)
[![](https://img.shields.io/coverallsCoverage/github/muhlba91/external-dns-provider-adguard?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/)
[![](https://img.shields.io/github/release-date/muhlba91/external-dns-provider-adguard?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/releases)
[![](https://img.shields.io/github/all-contributors/muhlba91/external-dns-provider-adguard?color=ee8449&style=for-the-badge)](#contributors)
<a href="https://www.buymeacoffee.com/muhlba91" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="28" width="150"></a>

The Adguard Home Provider Webhook for [External DNS](https://github.com/kubernetes-sigs/external-dns) provides support for [Adguard Home filtering rules](https://github.com/AdguardTeam/AdGuardHome/wiki/Hosts-Blocklists#adblock-style).

The provider is hugely based on <https://github.com/ionos-cloud/external-dns-ionos-webhook>.

---

## Supported DNS Record Types

The following DNS record types are supported:

- `A`
- `AAAA`
- `CNAME`
- `TXT`
- `SRV`
- `NS`
- `PTR`
- `MX`

## Adguard Home Filtering Rules

The provider manages Adguard Home filtering rules following the Adguard specification, which allows this provider to - theoretically - support all kinds of DNS record types.

Each record will be added in the format `||DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET #$managed-by-external-dns`.
Examples are:

```txt
||my.domain.com^dnsrewrite=NOERROR;A;1.2.3.4 #$managed-by-external-dns
||my.domain.com^dnsrewrite=NOERROR;AAAA;1111:2222::3 #$managed-by-external-dns
```

The provider will also honour and not modify any manually created rules. The filter constraint is the `$managed-by-external-dns` comment.

---

## Configuration

See [cmd/webhook/init/configuration/configuration.go](./cmd/webhook/init/configuration/configuration.go) for all available configuration options of the webhook sidecar, and [internal/adguard/configuration.go](./internal/adguard/configuration.go) for all available configuration options of the Adguard provider.

## Kubernetes Deployment

The Adguard webhook is provided as an OCI image in [ghcr.io/muhlba91/external-dns-provider-adguard](https://ghcr.io/muhlba91/external-dns-provider-adguard).

The following example shows the deployment as a [sidecar container](https://kubernetes.io/docs/concepts/workloads/pods/#workload-resources-for-managing-pods) in the ExternalDNS pod using the [Bitnami Helm charts for ExternalDNS](https://github.com/bitnami/charts/tree/main/bitnami/external-dns).

```shell
helm repo add bitnami https://charts.bitnami.com/bitnami

# create the adguard configuration
kubectl create secret generic adguard-configuration --from-literal=url='<ADGUARD_URL>' --from-literal=user='<ADGUARD_USER>' --from-literal=password='<ADGUARD_PASSWORD>'

# create the helm values file
cat <<EOF > external-dns-adguard-values.yaml
provider: webhook

extraArgs:
  webhook-provider-url: http://localhost:8888

sidecars:
  - name: adguard-webhook
    image: ghcr.io/muhlba91/external-dns-provider-adguard:$RELEASE_VERSION
    ports:
      - containerPort: 8888
        name: http
    livenessProbe:
      httpGet:
        path: /health
        port: http
      initialDelaySeconds: 10
      timeoutSeconds: 5
    readinessProbe:
      httpGet:
        path: /health
        port: http
      initialDelaySeconds: 10
      timeoutSeconds: 5
    env:
      - name: LOG_LEVEL
        value: debug
      - name: ADGUARD_HOME
        valueFrom:
          secretKeyRef:
            name: adguard-configuration
            key: url
      - name: ADGUARD_USER
        valueFrom:
          secretKeyRef:
            name: adguard-configuration
            key: user
      - name: ADGUARD_PASSWORD
        valueFrom:
          secretKeyRef:
            name: adguard-configuration
            key: password
      - name: SERVER_HOST
        value: "0.0.0.0" 
      - name: DRY_RUN
        value: "false"  
EOF

# install external-dns with helm
helm install external-dns-adguard bitnami/external-dns -f external-dns-adguard-values.yaml
```

---

## Contributors

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://muehlbachler.io/"><img src="https://avatars.githubusercontent.com/u/653739?v=4?s=100" width="100px;" alt="Daniel Mühlbachler-Pietrzykowski"/><br /><sub><b>Daniel Mühlbachler-Pietrzykowski</b></sub></a><br /><a href="#maintenance-muhlba91" title="Maintenance">🚧</a> <a href="https://github.com/muhlba91/external-dns-provider-adguard/commits?author=muhlba91" title="Code">💻</a> <a href="https://github.com/muhlba91/external-dns-provider-adguard/commits?author=muhlba91" title="Documentation">📖</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

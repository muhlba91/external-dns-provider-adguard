# External DNS - Adguard Home Provider (Webhook)

[![](https://img.shields.io/github/license/muhlba91/external-dns-provider-adguard?style=for-the-badge)](LICENSE.md)
[![](https://img.shields.io/github/actions/workflow/status/muhlba91/external-dns-provider-adguard/verify.yml?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/actions/workflows/verify.yml)
[![](https://img.shields.io/coverallsCoverage/github/muhlba91/external-dns-provider-adguard?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/)
[![](https://api.scorecard.dev/projects/github.com/muhlba91/external-dns-provider-adguard/badge?style=for-the-badge)](https://scorecard.dev/viewer/?uri=github.com/muhlba91/external-dns-provider-adguard)
[![](https://img.shields.io/github/release-date/muhlba91/external-dns-provider-adguard?style=for-the-badge)](https://github.com/muhlba91/external-dns-provider-adguard/releases)
[![](https://img.shields.io/github/all-contributors/muhlba91/external-dns-provider-adguard?color=ee8449&style=for-the-badge)](#contributors)
<a href="https://www.buymeacoffee.com/muhlba91" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="28" width="150"></a>

The Adguard Home Provider Webhook for [External DNS](https://github.com/kubernetes-sigs/external-dns) provides support for [Adguard Home filtering rules](https://github.com/AdguardTeam/AdGuardHome/wiki/Hosts-Blocklists#adblock-style).

The provider is hugely based on <https://github.com/ionos-cloud/external-dns-ionos-webhook>.

---

> [!WARNING]
> **Please make yourself familiar with the [limitations](#limitations) before using this provider!**

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

The provider manages Adguard Home filtering rules following the [Adblock-style syntax](https://github.com/AdguardTeam/AdGuardHome/wiki/Hosts-Blocklists#adblock-style), which allows this provider to - theoretically - support all kinds of DNS record types.

Each record will be added in the format `|DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET`.
Examples are:

```txt
|my.domain.com^dnsrewrite=NOERROR;A;1.2.3.4
|my.domain.com^dnsrewrite=NOERROR;AAAA;1111:2222::3
```

## Limitations

### Rule Ownership

> [!IMPORTANT]
> This provider takes **ownership** of **all rules** matching above mentioned format!

Adguard does not support inline comments for filtering rules, making it impossible to filter out only rules set by External DNS.
If you require **manually set rules**, it is adviced to define them as **`DNSEndpoint`** objects and enable the `crd` source in External DNS.

However, rules **not matching** above format, for example, `|domain.to.block`, **will not be modified**.

---

## Migrations

> [!IMPORTANT]
> **If** an **upgrade path** between version is **listed here**, please make sure to **follow** those paths **without skipping a version**!
> Otherwise, the correct behaviour cannot be guaranteed, resulting in possible inconsistencies or errors.

### v7 to v8

`v8` introduces the `HEALTHZ_ADDRESS` (default: `0.0.0.0`) and `HEALTHZ_PORT` (default: `8080`) environment variable to introduce compatibility with the official Helm chart.

Attention: if you are using a customized Helm chart, make sure to adjust the probes accordingly.

### v5 to v6

`v6` introduces the `ADGUARD_SET_IMPORTANT_FLAG` environment variable to set the `important` flag for AdGuard rules. This is enabled by default.

To keep the previous behaviour of `v5`, set `ADGUARD_SET_IMPORTANT_FLAG` to `false`.

### v4 to v5

In `v5` removes the automated migration from the old rules syntax (`v3`) to the new syntax introduced in `v4`.

Attention: if you skip the upgrade to `v4`, old rules will be dangling and will cause issues.

### v3 to v4

In `v3` the rule format was `||DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET`.

In `v4` this was changed to `|DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET` to solve issues with handling subdomains.

`v4` also introduces an automated migration from the old syntax to the new one.
To achieve this the provider reads the old syntax when updating rules but ignores them when providing the existing rules to ExternalDNS.
In fact, ExternalDNS tries to create those rules and the provider will re-write those in AdGuard using the new syntax.

Please make sure `v4` runs for some time in your cluster to ensure the migration of all old rules.

---

## Configuration

See [cmd/webhook/init/configuration/configuration.go](./cmd/webhook/init/configuration/configuration.go) for all available configuration options of the webhook sidecar, and [internal/adguard/configuration.go](./internal/adguard/configuration.go) for all available configuration options of the Adguard provider.

---

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
      - containerPort: 8080
        name: healthz
    livenessProbe:
      httpGet:
        path: /healthz
        port: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 5
    readinessProbe:
      httpGet:
        path: /healthz
        port: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 5
    env:
      - name: LOG_LEVEL
        value: debug
      - name: ADGUARD_URL
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
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/heyzling"><img src="https://avatars.githubusercontent.com/u/30411478?v=4?s=100" width="100px;" alt="heyzling"/><br /><sub><b>heyzling</b></sub></a><br /><a href="https://github.com/muhlba91/external-dns-provider-adguard/commits?author=heyzling" title="Documentation">📖</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/gabe565"><img src="https://avatars.githubusercontent.com/u/7717888?v=4?s=100" width="100px;" alt="gabe565"/><br /><sub><b>Gabe Cook</b></sub></a><br /><a href="https://github.com/muhlba91/external-dns-provider-adguard/commits?author=gabe565" title="Ideas">🤔</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
